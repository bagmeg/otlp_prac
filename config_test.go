package grpcclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/confmap"
)

func TestValidate(t *testing.T) {

}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name      string
		configMap *confmap.Conf
		cfg       Config
		err       string
	}{
		{
			name: "traces interval",
			configMap: confmap.NewFromStringMap(map[string]interface{}{
				"traces": map[string]interface{}{
					"endpoint": "localhost:7070",
					"interval": 3,
				},
			}),
		},
	}

	f := NewFactory()
	for _, testInstance := range tests {
		t.Run(testInstance.name, func(t *testing.T) {
			cfg := f.CreateDefaultConfig().(*Config)

			err := cfg.Unmarshal(testInstance.configMap)
			if err != nil || testInstance.err != "" {
				assert.EqualError(t, err, testInstance.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
