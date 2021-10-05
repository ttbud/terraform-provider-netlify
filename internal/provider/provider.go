package provider

import (
	"context"
	NetlifyOpenApiClient "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	NetlifyOpenApiClient.DefaultTimeout = 10 * time.Minute
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{
				"token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("NETLIFY_TOKEN", nil),
					Description: "The OAuth token used to authenticate with netlify",
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"netlify_site": resourceSite(),
			},
			ConfigureContextFunc: configure(version),
		}
	}
}

func configure(version string) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		token := d.Get("token").(string)
		if token == "" {
			return nil, diag.Errorf("\"token\" parameter or NETLIFY_TOKEN environment variable must be provided")
		}

		return NewMeta(ctx, version, token), nil
	}
}
