package openai

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-openai",
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "org_id",
				Hydrate: getOrganizationId,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{ShouldIgnoreErrorFunc: isNotFoundError},
		},
		TableMap: map[string]*plugin.Table{
			"openai_completion": tableOpenAiCompletion(ctx),
			"openai_file":       tableOpenAiFile(ctx),
			"openai_model":      tableOpenAiModel(ctx),
		},
	}
	return p
}
