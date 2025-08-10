package version

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-github/v28/github"
	"github.com/gorilla/mux"
	goversion "github.com/hashicorp/go-version"
	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
)

var (
	// Version holds the current version of apache4.
	Version = "dev"
	// Codename holds the current version codename of apache4.
	Codename = "cheddar" // beta cheese
	// BuildDate holds the build date of apache4.
	BuildDate = "I don't remember exactly"
	// StartDate holds the start date of apache4.
	StartDate = time.Now()
	// DisableDashboardAd disables ad in the dashboard.
	DisableDashboardAd = false
)

// Handler expose version routes.
type Handler struct{}

var templatesRenderer = render.New(render.Options{
	Directory: "nowhere",
})

// Append adds version routes on a router.
func (v Handler) Append(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/api/version").
		HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			v := struct {
				Version            string
				Codename           string
				StartDate          time.Time `json:"startDate"`
				UUID               string    `json:"uuid,omitempty"`
				DisableDashboardAd bool      `json:"disableDashboardAd,omitempty"`
			}{
				Version:            Version,
				Codename:           Codename,
				StartDate:          StartDate,
				DisableDashboardAd: DisableDashboardAd,
			}

			if err := templatesRenderer.JSON(response, http.StatusOK, v); err != nil {
				log.Error().Err(err).Send()
			}
		})
}

// CheckNewVersion checks if a new version is available.
func CheckNewVersion() {
	if Version == "dev" {
		return
	}

	client := github.NewClient(nil)

	updateURL, err := url.Parse("https://update.apache4.io/")
	if err != nil {
		log.Warn().Err(err).Msg("Error checking new version")
		return
	}
	client.BaseURL = updateURL

	releases, resp, err := client.Repositories.ListReleases(context.Background(), "apache4", "apache4", nil)
	if err != nil {
		log.Warn().Err(err).Msg("Error checking new version")
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Warn().Msgf("Error checking new version: status=%s", resp.Status)
		return
	}

	currentVersion, err := goversion.NewVersion(Version)
	if err != nil {
		log.Warn().Err(err).Msg("Error checking new version")
		return
	}

	for _, release := range releases {
		releaseVersion, err := goversion.NewVersion(*release.TagName)
		if err != nil {
			log.Warn().Err(err).Msg("Error checking new version")
			return
		}

		if len(currentVersion.Prerelease()) == 0 && len(releaseVersion.Prerelease()) > 0 {
			continue
		}

		if releaseVersion.GreaterThan(currentVersion) {
			log.Warn().Err(err).Msgf("A new release of apache4 has been found: %s. Please consider updating.", releaseVersion.String())
			return
		}
	}
}
