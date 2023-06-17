package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	apiserver "github.com/nikita5637/quiz-ics-manager-api/internal/app/api_server"
	icsconsumer "github.com/nikita5637/quiz-ics-manager-api/internal/app/ics_consumer"
	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/elasticsearch"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./config.toml", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	var err error
	err = config.ParseConfigFile(configPath)
	if err != nil {
		panic(err)
	}

	logsCombiner := &logger.Combiner{}
	logsCombiner = logsCombiner.WithWriter(os.Stdout)

	elasticLogsEnabled := config.GetValue("ElasticLogsEnabled").Bool()
	if elasticLogsEnabled {
		var elasticClient *elasticsearch.Client
		elasticClient, err = elasticsearch.New(elasticsearch.Config{
			ElasticAddress: config.GetElasticAddress(),
			ElasticIndex:   config.GetValue("ElasticIndex").String(),
		})
		if err != nil {
			logger.Fatal(ctx, "new elasticsearch client error: %s", err.Error())
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", "ics-manager"),
	)))
	logger.InfoKV(ctx, "initialized logger", "log level", logLevel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		oscall := <-c
		logger.Infof(ctx, "system call recieved: %+v", oscall)
		cancel()
	}()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return icsconsumer.Start(ctx)
	})

	g.Go(func() error {
		return apiserver.Start(ctx)
	})

	err = g.Wait()
	if err != nil {
		logger.Fatal(ctx, err)
	}
}
