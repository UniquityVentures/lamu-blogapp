package landing

import (
	//	"log"
	// "net/http"

	"github.com/lariv-in/lago/lago"
	"github.com/lariv-in/lago/registry"
	"github.com/lariv-in/lago/views"
)

// Returns a single view that will be used to display the Landing home page
func pluginViews() lago.PluginFeatures[*views.View] {
	return lago.PluginFeatures[*views.View]{
		Entries: []registry.Pair[string, *views.View]{
			{Key: "landing.HomeView", Value: lago.GetPageView("landing.Home")},
		},
	}
}
