package openai

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type openaiConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &openaiConfig{}
}

func GetConfig(connection *plugin.Connection) openaiConfig {
	if connection == nil || connection.GetConfig() == nil {
		return openaiConfig{}
	}
	config, _ := connection.GetConfig().(openaiConfig)
	return config
}
