package provider

import (
	"github.com/BESTSELLER/terraform-provider-luis/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_ENDPOINT", ""),
			},
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_APP_ID", ""),
			},
			"luis_version": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_VERSION", ""),
			},
			"subscription_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_SUBSCRIPTION_KEY", ""),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"luis_closed_list_sublist": resourceClosedListSublist(),
		},
		DataSourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return client.New(client.Client{
		Endpoint:        d.Get("endpoint").(string),
		AppID:           d.Get("app_id").(string),
		Version:         d.Get("luis_version").(string),
		SubscriptionKey: d.Get("subscription_key").(string),
	}), nil
}
