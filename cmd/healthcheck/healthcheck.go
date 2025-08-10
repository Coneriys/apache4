package healthcheck

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apache4/paerser/cli"
	"github.com/apache4/apache4/v3/pkg/config/static"
)

// NewCmd builds a new HealthCheck command.
func NewCmd(apache4Configuration *static.Configuration, loaders []cli.ResourceLoader) *cli.Command {
	return &cli.Command{
		Name:          "healthcheck",
		Description:   `Calls apache4 /ping endpoint (disabled by default) to check the health of apache4.`,
		Configuration: apache4Configuration,
		Run:           runCmd(apache4Configuration),
		Resources:     loaders,
	}
}

func runCmd(apache4Configuration *static.Configuration) func(_ []string) error {
	return func(_ []string) error {
		apache4Configuration.SetEffectiveConfiguration()

		resp, errPing := Do(*apache4Configuration)
		if resp != nil {
			resp.Body.Close()
		}
		if errPing != nil {
			fmt.Printf("Error calling healthcheck: %s\n", errPing)
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Bad healthcheck status: %s\n", resp.Status)
			os.Exit(1)
		}
		fmt.Printf("OK: %s\n", resp.Request.URL)
		os.Exit(0)
		return nil
	}
}

// Do try to do a healthcheck.
func Do(staticConfiguration static.Configuration) (*http.Response, error) {
	if staticConfiguration.Ping == nil {
		return nil, errors.New("please enable `ping` to use health check")
	}

	ep := staticConfiguration.Ping.EntryPoint
	if ep == "" {
		ep = "apache4"
	}

	pingEntryPoint, ok := staticConfiguration.EntryPoints[ep]
	if !ok {
		return nil, fmt.Errorf("ping: missing %s entry point", ep)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	protocol := "http"

	// TODO Handle TLS on ping etc...
	// if pingEntryPoint.TLS != nil {
	// 	protocol = "https"
	// 	tr := &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 	}
	// 	client.Transport = tr
	// }

	path := "/"

	return client.Head(protocol + "://" + pingEntryPoint.GetAddress() + path + "ping")
}
