package grpcclient

import (
	"errors"
	"net/url"

	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/confmap"
)

const (
	defaultInterval = 1.0
)

type TraceConfig struct {
	confignet.TCPAddr `mapstructure:",squash"`
	FlushInterval     float64 `mapstructure:"interval,omitempty"`
}

type Config struct {
	Traces TraceConfig `mapstructure:"traces"`
}

func (t TraceConfig) validate() (err error) {
	if t.Endpoint == "" {
		return errors.New("target address must not be empty")
	}
	if err = validateAddress(t.Endpoint); err != nil {
		return
	}

	return
}

func (cfg *Config) Validate() (err error) {
	if err = cfg.Traces.validate(); err != nil {
		return
	}

	return nil
}

func (cfg *Config) Unmarshal(configMap *confmap.Conf) (err error) {
	err = configMap.Unmarshal(cfg, confmap.WithErrorUnused())
	if err != nil {
		return
	}

	if !configMap.IsSet("traces::endpoint") {
		cfg.Traces.TCPAddr.Endpoint = "http://localhost:9090"
	}

	if !configMap.IsSet("traces::interval") {
		cfg.Traces.FlushInterval = defaultInterval
	}

	return
}

// validateAddress check if the given address is valid.
func validateAddress(addr string) (err error) {
	_, err = url.ParseRequestURI(addr)

	return
}
