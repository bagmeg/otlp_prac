package grpcclient

import (
	"fmt"
	"net/url"
	"time"
)

type Config struct {
	Interval string `mapstructure:"interval"`
	Address  string `mapstructure:"address"`
}

const (
	defaultInterval = 1 * time.Minute
)

func (cfg *Config) Validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Seconds() < 1 {
		return fmt.Errorf("when defined, the interval has to be set to at least 1 second (1s)")
	}

	if _, err := url.ParseRequestURI(cfg.Address); err != nil {
		return err
	}
	return nil
}
