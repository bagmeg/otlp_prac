package grpcclient

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/exporter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bagmeg/otlp_prac/internal/metadata"

	pb "github.com/bagmeg/otlp_prac/data"
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
		Traces: TraceConfig{
			FlushInterval: defaultInterval,
			TCPAddr: confignet.TCPAddr{
				Endpoint: "localhost:9090",
			},
		},
	}
}

func createTracesExporter(_ context.Context, params exporter.CreateSettings, baseCfg component.Config) (exporter.Traces, error) {
	logger := params.Logger
	grpcClientCfg := baseCfg.(*Config)

	logger.Debug(fmt.Sprintf("GRPC Addr: %s", grpcClientCfg.Traces.Endpoint))
	conn, err := grpc.Dial(grpcClientCfg.Traces.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("grpc dial failed...")
	}
	c := pb.NewTraceClient(conn)
	gCliExpt := &gCliExporter{
		logger: logger,
		config: grpcClientCfg,
		conn:   conn,
		client: c,
	}

	gCliExpt.logger.Debug("create grpc client trace exporter...")
	return gCliExpt, nil
}
