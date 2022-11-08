package test

import (
	"net/url"
	"os"
	"testing"

	_ "github.com/go-micro/plugins/codecs/json"
	_ "github.com/go-micro/plugins/codecs/yaml"
	_ "github.com/go-micro/plugins/config/source/cli/urfave"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/cli"
)

func testSections(t *testing.T, sections []string) {
	t.Helper()

	// Clear flags
	cli.Flags.Clear()

	err := cli.Flags.Add(cli.NewFlag(
		"registry",
		"mdns",
		cli.ConfigPath("registry.plugin"),
		cli.Usage("string flag usage"),
	))
	require.NoError(t, err)

	err = cli.Flags.Add(cli.NewFlag(
		"registry_ttl",
		300,
		cli.CPSlice([]string{"registry", "ttl"}),
		cli.Usage("int flag usage"),
	))
	require.NoError(t, err)

	err = cli.Flags.Add(cli.NewFlag(
		"nats-address",
		[]string{},
		cli.CPSlice([]string{"registry", "addresses"}),
		cli.Usage("NATS Address"),
	))
	require.NoError(t, err)

	// Setup os.Args
	os.Args = []string{
		"testapp",
		"--registry",
		"nats",
		"--registry_ttl",
		"600",
		"--nats-address",
		"nats://localhost:4222",
	}

	// Setup the urls.
	u1, err := url.Parse("cli://urfave")
	require.NoError(t, err)

	// Read the urls.
	datas, err := config.Read([]*url.URL{u1}, sections)
	require.NoError(t, err)

	// Merge all data from the URL's.
	cfg := newRegistryNatsConfig()
	err = config.Parse(append(sections, "registry"), datas, cfg)
	require.NoError(t, err)

	// Check if it merges right.
	assert.Equal(t, true, cfg.Enabled, "Enabled by default")
	assert.Equal(t, "nats", cfg.Plugin, "Plugin")
	assert.Equal(t, 600, cfg.Timeout, "Timeout")
	assert.Equal(t, true, cfg.Secure, "Secure by default")
	assert.EqualValues(t, []string{"nats://localhost:4222"}, cfg.Addresses, "Addresses")
}

func TestCliSingleSection(t *testing.T) {
	testSections(t, []string{"app"})
}

func TestCliNoSection(t *testing.T) {
	testSections(t, []string{})
}

func TestCliMultiSection(t *testing.T) {
	testSections(t, []string{"com", "example", "abc"})
}