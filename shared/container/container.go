package container

import (
	"github.com/fgrosse/goldi"
	"github.com/rohanchauhan02/automation-engine/shared/config"
	"github.com/rohanchauhan02/automation-engine/shared/postgres"
)

func DefaultContainer() *goldi.Container {
	registry := goldi.NewTypeRegistry()
	conf := make(map[string]interface{})
	container := goldi.NewContainer(registry, conf)
	container.RegisterType("shared.config", config.NewImmutableConfig)
	container.RegisterType("shared.database", postgres.NewPostgres, "@shared.config")
	return container
}
