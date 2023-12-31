package icsconsumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	generator "github.com/nikita5637/quiz-ics-manager-api/internal/app/ics_consumer/internal/generator"
	icsmessage "github.com/nikita5637/quiz-ics-manager-api/internal/app/ics_consumer/internal/handler/ics_message"
	logmiddleware "github.com/nikita5637/quiz-ics-manager-api/internal/app/ics_consumer/internal/middleware/log"
	servicenamemiddleware "github.com/nikita5637/quiz-ics-manager-api/internal/app/ics_consumer/internal/middleware/service_name"
	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/facade/icsfiles"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
	ics "github.com/nikita5637/quiz-registrator-api/pkg/ics"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	placepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Start ...
func Start(ctx context.Context) error {
	opts := grpc.WithInsecure()
	registratorAPIAddress := viper.GetString("ics_consumer.registrator_api.address")
	registratorAPIPort := viper.GetUint32("ics_consumer.registrator_api.port")
	target := fmt.Sprintf("%s:%d", registratorAPIAddress, registratorAPIPort)
	registratorAPIConn, err := grpc.DialContext(ctx, target, opts, grpc.WithChainUnaryInterceptor(
		logmiddleware.New().Log(),
		servicenamemiddleware.New().ServiceName(),
	))
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}

	leagueServiceClient, err := getLeagueServiceClient(ctx, registratorAPIConn)
	if err != nil {
		return fmt.Errorf("get league service client error: %w", err)
	}

	placeServiceClient, err := getPlaceServiceClient(ctx, registratorAPIConn)
	if err != nil {
		return fmt.Errorf("get place service client error: %w", err)
	}

	gameServiceClient, err := getGameServiceClient(ctx, registratorAPIConn)
	if err != nil {
		return fmt.Errorf("get game service client error: %w", err)
	}

	rabbitMQConn, err := amqp.Dial(config.GetRabbitMQURL())
	if err != nil {
		return fmt.Errorf("get rabbitMQ conn error: %w", err)
	}
	defer rabbitMQConn.Close()

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("get rabbitMQ channel error: %w", err)
	}
	defer rabbitMQChannel.Close()

	icsMessages, err := getICSMessages(rabbitMQChannel)
	if err != nil {
		return fmt.Errorf("get ICS messages error: %w", err)
	}

	icsGenerator := generator.New()

	driverName := viper.GetString("database.driver")
	db, err := storage.NewDB(ctx, driverName)
	if err != nil {
		return fmt.Errorf("new DB initialization error: %w", err)
	}
	defer db.Close()

	txManager := tx.NewManager(db)

	icsFileStorage := storage.NewICSFileStorage(driverName, txManager)

	icsFilesFacadeConfig := icsfiles.Config{
		ICSFileStorage: icsFileStorage,
		TxManager:      txManager,
	}
	icsFilesFacade := icsfiles.NewFacade(icsFilesFacadeConfig)

	placeStorage := storage.NewPlaceStorage(driverName, txManager)

	placesFacadeConfig := places.Config{
		PlaceStorage: placeStorage,
		TxManager:    txManager,
	}
	placesFacade := places.New(placesFacadeConfig)

	icsMessageHandlerConfig := icsmessage.Config{
		IcsFileExtension: viper.GetString("ics_consumer.ics_file_extension"),
		ICSFilesFacade:   icsFilesFacade,
		IcsFilesFolder:   viper.GetString("ics_consumer.ics_files_folder"),
		ICSGenerator:     icsGenerator,
		PlacesFacade:     placesFacade,

		GameServiceClient:   gameServiceClient,
		LeagueServiceClient: leagueServiceClient,
		PlaceServiceClient:  placeServiceClient,
	}
	icsMessageHandler := icsmessage.New(icsMessageHandlerConfig)

	go func(ctx context.Context) {
		for m := range icsMessages {
			logger.InfoKV(ctx, "accepted new message", zap.ByteString("body", m.Body))
			event := ics.Event{}
			err := json.Unmarshal(m.Body, &event)
			if err != nil {
				logger.ErrorKV(ctx, "JSON unmarshal error", zap.Error(err))
				continue
			}

			if err := icsMessageHandler.Handle(ctx, event); err != nil {
				logger.ErrorKV(ctx, "message handling error", zap.Error(err))
			}
		}
	}(ctx)

	logger.Info(ctx, "ICS messages consumer started")

	<-ctx.Done()

	logger.Info(ctx, "ICS messages consumer gracefully stopped")
	return nil
}

func getICSMessages(channel *amqp.Channel) (<-chan amqp.Delivery, error) {
	icsQueueName := viper.GetString("ics_consumer.rabbitmq.queue.name")
	if icsQueueName == "" {
		return nil, errors.New("empty rabbit MQ ICS queue name")
	}

	icsQueue, err := channel.QueueDeclare(
		icsQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("rabbitMQ queue declare error: %w", err)
	}

	icsMessages, err := channel.Consume(
		icsQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("rabbitMQ consume error: %w", err)
	}

	return icsMessages, nil
}

func getLeagueServiceClient(ctx context.Context, conn *grpc.ClientConn) (leaguepb.ServiceClient, error) {
	return leaguepb.NewServiceClient(conn), nil
}

func getPlaceServiceClient(ctx context.Context, conn *grpc.ClientConn) (placepb.ServiceClient, error) {
	return placepb.NewServiceClient(conn), nil
}

func getGameServiceClient(ctx context.Context, conn *grpc.ClientConn) (gamepb.ServiceClient, error) {
	return gamepb.NewServiceClient(conn), nil
}
