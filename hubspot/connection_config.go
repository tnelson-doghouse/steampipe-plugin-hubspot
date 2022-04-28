package hubspot

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type hubspotConfig struct {
	BaseUrl  *string `cty:"base_url"`
	APIKey   *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"base_url": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &hubspotConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) hubspotConfig {
	if connection == nil || connection.Config == nil {
		return hubspotConfig{}
	}
	config, _ := connection.Config.(hubspotConfig)
	return config
}
