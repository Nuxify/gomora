package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
)

// Config handles the hystrix configurations
type Config struct{}

// Settings returns the hystrix command config
func (c Config) Settings() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout: 3000,
	}
}
