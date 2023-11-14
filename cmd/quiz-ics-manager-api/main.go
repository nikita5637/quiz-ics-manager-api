package main

import (
	"os"
	"runtime/debug"

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
			logger.FatalKV(ctx, "new elasticsearch client error", zap.Error(err))
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", viper.GetString("log.module_name")),
	)))
	logger.InfoKV(ctx, "initialized logger", zap.String("log_level", logLevel.String()))

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return icsconsumer.Start(ctx)
	})

	g.Go(func() error {
		return apiserver.Start(ctx)
	})

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range buildInfo.Settings {
			if setting.Key == "vcs.revision" {
				logger.InfoKV(ctx, "application started", zap.String("vcs.revision", setting.Value))
			}
		}
	}

	if err := g.Wait(); err != nil {
		logger.Fatal(ctx, err)
	}
}
