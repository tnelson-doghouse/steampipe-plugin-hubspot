package hubspot

import (
	"context"
	"errors"
//	"fmt"

	"github.com/tnelson-doghouse/hubspot"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*hubspot.Client, error) {
//	logger := plugin.Logger(ctx)

//	logger.Error("t1")
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "hubspot"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*hubspot.Client), nil
	}

	// Start with an empty default config
	var baseUrl, APIKey string

	// Prefer config options given in Steampipe
	hubspotConfig := GetConfig(d.Connection)

	if hubspotConfig.BaseUrl != nil {
		baseUrl = *hubspotConfig.BaseUrl
	}
	if hubspotConfig.APIKey != nil {
		APIKey = *hubspotConfig.APIKey
	}

	if baseUrl == "" {
		return nil, errors.New("'base_url' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	if APIKey == "" {
		return nil, errors.New("'api_token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// Create the client
	var clientConfig = hubspot.NewClientConfig()
	clientConfig.APIHost = baseUrl
	clientConfig.APIKey = APIKey
	client := hubspot.NewClient(clientConfig)
//	client := hubspot.NewClient(hubspot.NewClientConfig())
/*	hubspot.NewClient(hubspot.ClientConfig{
		APIHost: baseUrl,
		APIKey: APIKey,
	}) */
//	logger.Error("client/hubspot", client)

/*	if err != nil {
		return nil, fmt.Errorf("error creating Hubspot client: %s", err.Error())
	} */

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, &client)

	// Done
	return &client, nil
}
