package grpcclient

import (
	"context"
	"fmt"

	pb "github.com/bagmeg/otlp_prac/data"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	g.logger.Info("Starting grpc client..... ❤️❤️❤️❤️❤️❤️")

	return nil
}

func (g *gCliExporter) Shutdown(ctx context.Context) error {
	g.logger.Debug("Shutting down grpc client.... 💕💕💕💕💕💕💕")

	return nil
}

func (g *gCliExporter) Capabilities() consumer.Capabilities {
	g.logger.Debug("What is wrong with this capatibilities.... .🤣🤣🤣🤣")

	return consumer.Capabilities{}
}

func (g *gCliExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	g.logger.Debug("Consuming Traces.... 🎶🎶🎶🎶🎶🎶")

	traceData := pb.TraceData{
		Data: "Some data...",
	}

	reply, err := g.client.Consume(ctx, &traceData)
	if err != nil {
		g.logger.Warn("grpc error")
		g.logger.Warn(fmt.Sprintf("err: %v", err))
	}
	g.logger.Info(reply.GetMessage())

	return nil
}
