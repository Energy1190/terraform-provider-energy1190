package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"base_url": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: func() (interface{}, error) {
						return nil, nil
					},
				},
				"username": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: func() (interface{}, error) {
						return nil, nil
					},
				},
				"password": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: func() (interface{}, error) {
						return nil, nil
					},
				},
			},

			DataSourcesMap: map[string]*schema.Resource{
				"superset_test_resource": dataSourceTest(),
			},
		}

		return p
	}
}

func dataSourceTest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTestRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
