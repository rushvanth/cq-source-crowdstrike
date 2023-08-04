package plugin

import (
	"github.com/justmiles/cq-source-crowdstrike/client"
	"github.com/justmiles/cq-source-crowdstrike/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"rushvanth-crowdstrike",
		Version,
		schema.Tables{
			resources.Vulnerabilities(),
		},
		client.New,
	)
}
