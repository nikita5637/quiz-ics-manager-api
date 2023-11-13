package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-ics-manager-api/internal/app/apiserver"
	icsconsumer "github.com/nikita5637/quiz-ics-manager-api/internal/app/ics_consumer"
	"github.com/nikita5637/quiz-ics-manager-api/internal/config"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/elasticsearch"
	"github.com/nikita5637/quiz-ics-manager-api/internal/pkg/logger"
	"github.com/posener/ctxutil"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func init() {
	pflag.StringP("config", "c", "", "path to config file")
	_ = viper.BindPFlag("config", pflag.Lookup("config"))
}

func main() {
	ctx := ctxutil.Interrupt()

	pflag.Parse()

	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	logsCombiner := &logger.Combiner{}
	logsCombiner = logsCombiner.WithWriter(os.Stdout)

	elasticLogsEnabled := viper.GetBool("log.elastic.enabled")
	if elasticLogsEnabled {
		var elasticClient *elasticsearch.Client
		elasticClient, err := elasticsearch.New(elasticsearch.Config{
			ElasticAddress: config.GetElasticAddress(),
			ElasticIndex:   viper.GetString("log.elastic.index"),
		})
		if err != nil {
			logger.Fatal(ctx, "new elasticsearch client error: %s", err.Error())
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", viper.GetString("log.module_name")),
	)))
	logger.InfoKV(ctx, "initialized logger", "log level", logLevel)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return icsconsumer.Start(ctx)
	})

	g.Go(func() error {
		return apiserver.Start(ctx)
	})

	if err := g.Wait(); err != nil {
		logger.Fatal(ctx, err)
	}
}
