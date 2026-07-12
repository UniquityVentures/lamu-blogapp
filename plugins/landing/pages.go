package landing

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/components"

	"context"
	"maragu.dev/gomponents"
)

type HomePage struct {
	components.Page
}

// rawPage wraps raw HTML string as a PageInterface so it can be a ShellBase child.
type rawPage struct {
	components.Page
	content string
}

func (r *rawPage) Build(_ context.Context) gomponents.Node {
	return gomponents.Raw(r.content)
}

func (r *rawPage) GetKey() string     { return r.Key }
func (r *rawPage) GetRoles() []string { return r.Roles }

func (e *HomePage) Build(ctx context.Context) gomponents.Node {
	return components.Render(components.ShellBase{
		Children: []components.PageInterface{
			&rawPage{content: "<h1>Harcoded, Hello World!</h1>"},
		},
	}, ctx)
}

func (e *HomePage) GetKey() string     { return e.Key }
func (e *HomePage) GetRoles() []string { return e.Roles }

func pluginPages() lamu.PluginFeatures[components.PageInterface] {
		return lamu.PluginFeatures[components.PageInterface] {
			Entries: []registry.Pair[string, components.PageInterface]{
				{ Key: "landing.Home", Value: &HomePage{} },
			},
		}
}

