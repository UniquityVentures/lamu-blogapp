package landing

import (
	"log"
	"net/http"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginViews() lamu.PluginFeatures[*lamu.View] {
	return lamu.PluginFeatures[*lamu.View] {
		Entries: []registry.Pair[string, *lamu.View]{
			{
				Key: "landing.HomeView",
				Value: lamu.GetPageView("landing.HomePage")
			}
	}
}