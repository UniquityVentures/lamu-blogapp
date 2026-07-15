package landing

import (
	// "log"
	// "net/http"

	"github.com/lariv-in/lago/lago"
	"github.com/lariv-in/lago/registry"
)

func LandingRoutePatch(old lago.Route) lago.Route {
	old.Handler = lago.NewDynamicView("landing.HomeView")
	return old
}

// Patching the "/" route to make it render html instead of redirecting to login page/dashboard
func pluginRoutes() lago.PluginFeatures[lago.Route] {
	return lago.PluginFeatures[lago.Route]{
		Patches: []registry.Pair[string, func(lago.Route) lago.Route]{
			{Key: "core.HomeRoute", Value: LandingRoutePatch},
		},
	}
}
