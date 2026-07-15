package main

import (
	"log"

	// Custom plugins
	"github.com/NotoriousArnav/lago-blogapp/plugins/landing"

	// Includes the Core Plugin
	"github.com/lariv-in/lago/lago"

	// Bundled plugins, for common functionality
	"github.com/lariv-in/lago/plugins/p_dashboard"
	"github.com/lariv-in/lago/plugins/p_users"

	"github.com/lariv-in/lago/registry"
)

func main() {
	// Register the list of active plugins to load into the application kernel.
	// Order matters here, especially when plugins patch other packages
	plugins := []registry.Pair[string, lago.Plugin]{
		// Provides the main dashboard, with the apps grid
		p_dashboard.GetPlugin(),
		// Provides auth, login page, signup page, and authorization helpers
		p_users.GetPlugin(),
		// Custom plugin, patches the HomeRoute to prevent redirecting to login, instead displayes a landing page.
		landing.GetPlugin(),
	}
	// Load database settings, server addresses, and plugin parameters from config.toml.
	config, err := lago.LoadConfigFromFile("config.toml", plugins)
	if err != nil {
		log.Fatalf("failed loading configuration file: %v", err)
	}

	// Build global registries and run the Cobra CLI bootstrapper.
	if err := lago.Start(config, plugins); err != nil {
		log.Fatalf("failed executing application entry: %v", err)
	}
}
