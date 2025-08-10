package cmd

import (
	"time"

	ptypes "github.com/apache4/paerser/types"
	"github.com/apache4/apache4/v3/pkg/config/static"
)

// apache4CmdConfiguration wraps the static configuration and extra parameters.
type apache4CmdConfiguration struct {
	static.Configuration `export:"true"`
	// ConfigFile is the path to the configuration file.
	ConfigFile string `description:"Configuration file to use. If specified all other flags are ignored." export:"true"`
}

// Newapache4Configuration creates a apache4CmdConfiguration with default values.
func Newapache4Configuration() *apache4CmdConfiguration {
	return &apache4CmdConfiguration{
		Configuration: static.Configuration{
			Global: &static.Global{
				CheckNewVersion: true,
			},
			EntryPoints: make(static.EntryPoints),
			Providers: &static.Providers{
				ProvidersThrottleDuration: ptypes.Duration(2 * time.Second),
			},
			ServersTransport: &static.ServersTransport{
				MaxIdleConnsPerHost: 200,
			},
			TCPServersTransport: &static.TCPServersTransport{
				DialTimeout:   ptypes.Duration(30 * time.Second),
				DialKeepAlive: ptypes.Duration(15 * time.Second),
			},
		},
		ConfigFile: "",
	}
}
