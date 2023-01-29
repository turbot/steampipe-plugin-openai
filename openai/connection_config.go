package openai

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type openaiConfig struct {
	APIKey *string `cty:"api_key" hcl:"api_key"`

	DataSets []DataSet `cty:"data_sets" hcl:"data_sets"`
}

type DataSet struct {
	Name        string   `cty:"name" hcl:"name"`
	Description string   `cty:"description" hcl:"description"`
	Columns     []string `cty:"columns" hcl:"columns"`
}

func ConfigInstance() interface{} {
	return &openaiConfig{}
}

func GetConfig(connection *plugin.Connection) openaiConfig {
	if connection == nil || connection.Config == nil {
		return openaiConfig{}
	}
	config, _ := connection.Config.(openaiConfig)
	return config
}
