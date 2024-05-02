package openai

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiModel(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_model",
		Description: "Models available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: listModel,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the model, e.g. davinci."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(transform.UnixToTimestamp), Description: "Timestamp of when the model was created."},
			{Name: "owned_by", Type: proto.ColumnType_STRING, Description: "Organization that owns the model, e.g. openai."},
			// Other columns
			{Name: "object", Type: proto.ColumnType_STRING, Description: "Type of the object, e.g. model."},
			{Name: "root", Type: proto.ColumnType_STRING, Description: "Root of this model."},
			{Name: "permission", Type: proto.ColumnType_JSON, Description: "Permissions for the model."},
			// Always null in testing? {Name: "parent", Type: proto.ColumnType_STRING, Description: ""},
		}),
	}
}

func listModel(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("openai_model.listModel", "connection_error", err)
		return nil, err
	}
	resp, err := conn.ListModels(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("openai_model.listModel", "query_error", err)
		return nil, err
	}
	for _, i := range resp.Models {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
