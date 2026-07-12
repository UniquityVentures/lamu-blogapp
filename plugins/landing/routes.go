package landing

import (
	"log"
	"net/http"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func LandingRoutePatch (old lamu.Route) lamu.Route {
	old.Handler = lamu.NewDynamicView("landing.HomeView")
	return old
}

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Patches: []registry.Pair[string, func(lamu.Route) lamu.Route]{
			Key: "landing.DefaultRoute",
			Value: LandingRoutePatch
		}
	}
}
