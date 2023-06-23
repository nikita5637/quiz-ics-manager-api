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
	registratorpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

// Start ...
func Start(ctx context.Context) error {
	registratorServiceClient, err := getRegistratorServiceClient(ctx)
	if err != nil {
		return fmt.Errorf("get registrator service client error: %w", err)
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

	db, err := storage.NewDB()
	if err != nil {
		logger.Fatalf(ctx, "new DB initialization error: %s", err.Error())
	}
	defer db.Close()

	txManager := tx.NewManager(db)

	icsFileStorage := storage.NewICSFileStorage(txManager)

	icsFilesFacadeConfig := icsfiles.Config{
		ICSFileStorage: icsFileStorage,
		TxManager:      txManager,
	}
	icsFilesFacade := icsfiles.NewFacade(icsFilesFacadeConfig)

	placeStorage := storage.NewPlaceStorage(txManager)

	placesFacadeConfig := places.Config{
		PlaceStorage: placeStorage,
		TxManager:    txManager,
	}
	placesFacade := places.New(placesFacadeConfig)

	icsMessageHandlerConfig := icsmessage.Config{
		IcsFileExtension:         config.GetValue("ICSFileExtension").String(),
		ICSFilesFacade:           icsFilesFacade,
		IcsFilesFolder:           config.GetValue("ICSFilesFolder").String(),
		ICSGenerator:             icsGenerator,
		PlacesFacade:             placesFacade,
		RegistratorServiceClient: registratorServiceClient,
	}
	icsMessageHandler := icsmessage.New(icsMessageHandlerConfig)

	go func(ctx context.Context) {
		for m := range icsMessages {
			logger.InfoKV(ctx, "accepted new message", "body", m.Body)
			event := ics.Event{}
			err := json.Unmarshal(m.Body, &event)
			if err != nil {
				logger.Errorf(ctx, "JSON unmarshal error: %s", err.Error())
				continue
			}

			if err := icsMessageHandler.Handle(ctx, event); err != nil {
				logger.Errorf(ctx, "message handling error: %s", err.Error())
			}
		}
	}(ctx)

	logger.Info(ctx, "ICS messages consumer started")

	<-ctx.Done()

	logger.Info(ctx, "ICS messages consumer gracefully stopped")
	return nil
}

func getICSMessages(channel *amqp.Channel) (<-chan amqp.Delivery, error) {
	icsQueueName := config.GetValue("RabbitMQICSQueueName").String()
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

func getRegistratorServiceClient(ctx context.Context) (registratorpb.RegistratorServiceClient, error) {
	opts := grpc.WithInsecure()
	registratorAPIAddress := config.GetValue("RegistratorAPIAddress").String()
	registratorAPIPort := config.GetValue("RegistratorAPIPort").Uint16()
	target := fmt.Sprintf("%s:%d", registratorAPIAddress, registratorAPIPort)
	cc, err := grpc.DialContext(ctx, target, opts, grpc.WithChainUnaryInterceptor(
		logmiddleware.New().Log(),
		servicenamemiddleware.New().ServiceName(),
	))
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}

	return registratorpb.NewRegistratorServiceClient(cc), nil
}
