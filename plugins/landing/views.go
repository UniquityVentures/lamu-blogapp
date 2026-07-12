package landing

import (
//	"log"
	// "net/http"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/views"
)

func pluginViews() lamu.PluginFeatures[*views.View] {
	return lamu.PluginFeatures[*views.View] {
		Entries: []registry.Pair[string, *views.View]{
			{ Key: "landing.HomeView", Value: lamu.GetPageView("landing.Home") },
		},
	}
}