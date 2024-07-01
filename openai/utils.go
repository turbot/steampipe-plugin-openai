package openai

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*openai.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*openai.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	var conn *openai.Client

	plugin.Logger(ctx).Debug("openai connectUncached")

	// Default to the env var settings
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Prefer config settings
	openaiConfig := GetConfig(d.Connection)
	if openaiConfig.APIKey != nil {
		apiKey = *openaiConfig.APIKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return conn, errors.New("api_key must be configured")
	}

	conn = openai.NewClient(apiKey)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "status code: 404")
}

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{Name: "org_id", Type: proto.ColumnType_STRING, Description: "The ID of the organization.", Hydrate: getOrganizationId, Transform: transform.FromValue()},
	}, c...)
}

// Profile represents the structure of our profile data.
type Profile struct {
	UserID    string `json:"id"`
	FirstName string `json:"name"`
	Email     string `json:"email"`
	Orgs      struct {
		Data []OrganizationInfo `json:"data"`
	} `json:"orgs"`
}

type OrganizationInfo struct {
	Object      string `json:"object"`
	OrgId       string `json:"id"`
	Created     int64  `json:"created"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Personal    bool   `json:"personal"`
	ParentOrgId string `json:"parent_org_id"`
}

var getOrganizationIdMemoize = plugin.HydrateFunc(getOrganizationIdUncached).Memoize(memoize.WithCacheKeyFunction(getOrganizationIdCacheKey))

func getOrganizationIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getOrganizationId"
	return cacheKey, nil
}

func getOrganizationId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	config, err := getOrganizationIdMemoize(ctx, d, h)
	if err != nil {
		return nil, err
	}

	c := config.(string)

	return c, nil
}

func getOrganizationIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	cacheKey := "getOrganizationId"

	var orgId string
	var profile Profile

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		orgId = cachedData.(string)
	} else {
		url := "https://api.openai.com/v1/me"

		// Default to the env var settings
		apiKey := os.Getenv("OPENAI_API_KEY")

		// Prefer config settings
		openaiConfig := GetConfig(d.Connection)
		if openaiConfig.APIKey != nil {
			apiKey = *openaiConfig.APIKey
		}
		client := &http.Client{}

		// Create a new GET request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+apiKey)

		// Perform the request
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(body, &profile); err != nil {
			return nil, err
		}

		o := profile.Orgs.Data[0]
		orgId = o.OrgId

		d.ConnectionManager.Cache.Set(cacheKey, orgId)
	}

	return orgId, nil
}
