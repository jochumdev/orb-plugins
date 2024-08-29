package natjs

import (
	"crypto/tls"
	"fmt"

	"github.com/go-orb/go-orb/config"
	"github.com/go-orb/go-orb/config/source/cli"
	"github.com/go-orb/go-orb/event"
	"github.com/go-orb/go-orb/types"
)

// Name provides the name of this event client.
const Name = "natsjs"

// Defaults.
//
//nolint:gochecknoglobals
var (
	DefaultAddresses = []string{"nats://localhost:4222"}
	DefaultCodec     = "application/json"
)

func init() {
	//nolint:errcheck
	_ = cli.Flags.Add(cli.NewFlag(
		event.ComponentType+"_addresses",
		DefaultAddresses,
		cli.ConfigPathSlice([]string{event.ComponentType, "addresses"}),
		cli.Usage("Events addresses."),
	))

	event.Register(Name, Provide)
}

// Config provides configuration for the NATS registry.
type Config struct {
	event.Config `yaml:",inline"`

	Addresses []string    `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	TLSConfig *tls.Config `json:"-"                   yaml:"-"`
	Codec     string      `json:"codec"               yaml:"codec"`
}

// ApplyOptions applies a set of options to the config.
func (c *Config) ApplyOptions(opts ...event.Option) {
	for _, o := range opts {
		o(c)
	}
}

// WithAddresses sets the NATS server addresses.
func WithAddresses(n ...string) event.Option {
	return func(c event.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.Addresses = n
		}
	}
}

// WithTLSConfig defines the TLS config to use for the secure connection.
func WithTLSConfig(n *tls.Config) event.Option {
	return func(c event.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.TLSConfig = n
		}
	}
}

// NewConfig creates a new config object.
func NewConfig(
	serviceName types.ServiceName,
	datas types.ConfigData,
	opts ...event.Option,
) (Config, error) {
	cfg := Config{
		Config:    event.NewConfig(),
		Addresses: DefaultAddresses,
		Codec:     DefaultCodec,
	}

	cfg.ApplyOptions(opts...)

	sections := types.SplitServiceName(serviceName)
	if err := config.Parse(append(sections, event.ComponentType), datas, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}