package main

import (
	"github.com/tnelson-doghouse/steampipe-plugin-hubspot/hubspot"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: hubspot.Plugin})
}
