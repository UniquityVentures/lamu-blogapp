package landing

import (
	// "log"
	// "net/url"

	"github.com/lariv-in/lago/lago"
	"github.com/lariv-in/lago/registry"
)

// Main entry point for the plugin, everything that this plugin does ends up here.
func GetPlugin() registry.Pair[string, lago.Plugin] {
	return registry.Pair[string, lago.Plugin]{
		Key: "landing",
		Value: lago.Plugin{
			// Type is Addon because we don't want this to show up in the dashboard.
			Type: lago.PluginTypeAddon,
			// URL:         "/",
			// VerboseName: "Public Pages",
			Views:  lago.PluginStages(pluginViews),
			Pages:  lago.PluginStages(pluginPages),
			Routes: lago.PluginStages(pluginRoutes),
		},
	}
}
