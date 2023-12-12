package openai

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type openaiConfig struct {
	APIKey *string `hcl:"api_key"`
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
