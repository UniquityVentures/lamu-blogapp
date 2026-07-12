package main

import (
	"log"

	"github.com/NotoriousArnav/lamu-blogapp/plugins/landing"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/plugins/p_dashboard"
	"github.com/UniquityVentures/lamu/plugins/p_users"
	"github.com/UniquityVentures/lamu/registry"
)

func main() {
	// 1. Register the list of active plugins to load into the application kernel.
	plugins := []registry.Pair[string, lamu.Plugin]{
		p_dashboard.GetPlugin(),
		p_users.GetPlugin(),
		landing.GetPlugin(),
	}
	// Load database settings, server addresses, and plugin parameters from config.toml.
	config, err := lamu.LoadConfigFromFile("config.toml", plugins)
	if err != nil {
		log.Fatalf("failed loading configuration file: %v", err)
	}

	// 3. Build global registries and run the Cobra CLI bootstrapper.
	if err := lamu.Start(config, plugins); err != nil {
		log.Fatalf("failed executing application entry: %v", err)
	}
}
