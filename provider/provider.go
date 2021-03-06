package provider

import (
	"fmt"
	"strings"

	"github.com/BESTSELLER/terraform-provider-luis/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		if s.Deprecated != "" {
			desc += " " + s.Deprecated
		}
		return strings.TrimSpace(desc)
	}
}

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_ENDPOINT", ""),
				Description: "LUIS API Endpoint, e.g West Europe should be \"westeurope.api.cognitive.microsoft.com\"",
				Default:     "westeurope.api.cognitive.microsoft.com",
			},
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_APP_ID", ""),
				Description: "LUIS Application ID",
			},
			"luis_version": {
				Type:        schema.TypeString,
				Required:    true,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_VERSION", ""),
				Description: "Version of you LUIS application",
			},
			"subscription_key": {
				Type:        schema.TypeString,
				Required:    true,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("LUIS_SUBSCRIPTION_KEY", ""),
				Description: "Subscription key which provides access to this API. Found in your Cognitive Services accounts.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"luis_closed_list_sublist": resourceClosedListSublist(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
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
