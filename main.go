package main

import (
	"github.com/BESTSELLER/terraform-provider-luis/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider})
}
