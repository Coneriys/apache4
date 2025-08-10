package provider

import (
	"github.com/apache4/apache4/v3/pkg/config/dynamic"
	"github.com/apache4/apache4/v3/pkg/safe"
)

// Provider defines methods of a provider.
type Provider interface {
	// Provide allows the provider to provide configurations to apache4
	// using the given configuration channel.
	Provide(configurationChan chan<- dynamic.Message, pool *safe.Pool) error
	Init() error
}
