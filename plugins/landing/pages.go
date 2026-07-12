package landing

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/components"

)

type HomePage struct {
	lamu.Page
}

func (e *HomePage) Build(ctx context.Context) gomponents.Node {
	return components.Render(components.ShellBase{
		Children: []components.PageInterface{
			&rawPage{content: "<h1>Harcoded, Hello World!</h1>"},
		},
	}, ctx)
}

func (e *HomePage) GetKey() string     { return e.Key }
func (e *HomePage) GetRoles() []string { return e.Roles }

func pluginPages() lamu.PluginFeatures[lamu.Page] {
		return lamu.PluginFeatures[lamu.Page] {
			Entries: []registry.Pair[string, lamu.Page]{
				{
					Key: "landing.HomePage",
					Value: &HomePage{}
				}
			}
		}
}

