package main

import (
	"github.com/turbot/steampipe-plugin-openai/openai"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: openai.Plugin})
}
