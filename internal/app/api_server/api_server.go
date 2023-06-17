package apiserver

import (
	"context"
	"fmt"
	"net"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	icsfilemanager "github.com/nikita5637/quiz-ics-manager-api/internal/app/api_server/internal/ics_file_manager"
	logmiddleware "github.com/nikita5637/quiz-ics-manager-api/internal/app/api_server/internal/middleware/log"
	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/facade/icsfiles"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/tx"
	icsfilemanagerpb "github.com/nikita5637/quiz-ics-manager-api/pkg/pb/ics_file_manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start(ctx context.Context) error {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.ChainUnaryInterceptor(
		grpc_recovery.UnaryServerInterceptor(),
		logmiddleware.New().Log(),
	))
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

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

	icsFileManagerConfig := icsfilemanager.Config{
		ICSFilesFacade: icsFilesFacade,
	}
	icsFileManager := icsfilemanager.New(icsFileManagerConfig)

	icsfilemanagerpb.RegisterServiceServer(grpcServer, icsFileManager)

	lis, err := net.Listen("tcp", config.GetBindAddress())
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	go func() {
		err = grpcServer.Serve(lis)
		return
	}()
	if err != nil {
		return err
	}

	logger.Info(ctx, "ics-manager-api server started")

	<-ctx.Done()

	grpcServer.GracefulStop()

	logger.Info(ctx, "ics-manager-api server gracefully stopped")
	return nil
}
