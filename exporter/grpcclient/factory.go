package grpcclient

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bagmeg/otlp_prac/exporter/grpcclient/internal/metadata"
	data "github.com/bagmeg/otlp_prac/pkg/testData"
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(
			createTracesExporter,
			metadata.TracesStability,
		),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultInterval.String(),
	}
}

func createTracesExporter(_ context.Context, params exporter.CreateSettings, baseCfg component.Config) (exporter.Traces, error) {
	logger := params.Logger
	grpcClientCfg := baseCfg.(*Config)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("grpc dial failed...")
	}
	c := data.NewTraceClient()
	gCliExpt := &gCliExporter{
		logger: logger,
		config: grpcClientCfg,
		conn:   conn,
		client: c,
	}

	gCliExpt.logger.Debug("create grpc client trace exporter...")
	return gCliExpt, nil
}
