package grpcclient

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/bagmeg/otlp_prac/pkg/data"
)

type gCliExporter struct {
	host   component.Host
	cancel context.CancelFunc
	logger *zap.Logger
	config *Config
	client pb.TraceClient
	conn   *grpc.ClientConn
}

func (g *gCliExporter) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (g *gCliExporter) Shutdown(ctx context.Context) error {
	return nil
}

func (g *gCliExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{}
}

func (g *gCliExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	return nil
}
