package main

import (
	"terraform-provider-superset/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	version string = "dev"
)

func main() {
	opts := &plugin.ServeOpts{ProviderFunc: provider.New(version)}
	plugin.Serve(opts)
}
