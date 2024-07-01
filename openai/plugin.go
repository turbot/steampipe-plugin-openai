package openai

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"openai_completion": tableOpenAiCompletion(ctx),
			"openai_file":       tableOpenAiFile(ctx),
			"openai_model":      tableOpenAiModel(ctx),
		},
	}
	return p
}
