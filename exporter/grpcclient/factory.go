package grpcclient

import "go.opentelemetry.io/collector/exporter"

func NewFactory() exporter.Factory {
	return &exporter.Factory{}
}
