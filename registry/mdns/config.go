package mdns

import (
	"fmt"

	"github.com/go-orb/go-orb/config"
	"github.com/go-orb/go-orb/config/source/cli"
	"github.com/go-orb/go-orb/registry"
	"github.com/go-orb/go-orb/types"
)

<<<<<<< Updated upstream
// metaSchemeKey is the key to use to store the scheme in metadata.
const metaSchemeKey = "_md_scheme_"

=======
>>>>>>> Stashed changes
// Name provides the name of this registry.
const Name = "mdns"

// Defaults.
//
//nolint:gochecknoglobals
var (
	DefaultDomain = "micro"
)

func init() {
	//nolint:errcheck
	_ = cli.Flags.Add(cli.NewFlag(
		"registry_domain",
		DefaultDomain,
		cli.ConfigPathSlice([]string{"registry", "domain"}),
		cli.Usage("Registry domain."),
	))

	if err := registry.Plugins.Add(Name, registry.ProviderFunc(ProvideRegistryMDNS)); err != nil {
		panic(err)
	}
}

// Config provides configuration for the mDNS registry.
type Config struct {
	registry.Config `yaml:",inline"`

	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}

// NewConfig creates a new config object.
func NewConfig(
	serviceName types.ServiceName,
	datas types.ConfigData,
	opts ...registry.Option,
) (Config, error) {
	cfg := Config{
		Config: registry.NewConfig(),
	}

	cfg.Config.Timeout = 500

	cfg.ApplyOptions(opts...)

	sections := types.SplitServiceName(serviceName)
	if err := config.Parse(append(sections, registry.ComponentType), datas, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}

// WithDomain sets the mDNS domain.
func WithDomain(domain string) registry.Option {
	return func(c registry.ConfigType) {
		cfg, ok := c.(*Config)
		if ok {
			cfg.Domain = domain
		}
	}
}
