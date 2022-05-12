package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

		},
		ResourcesMap: map[string]*schema.Resource{
			"mlin_test": resourceMlinTest(),
		},
		ConfigureContextFunc: mlinConfigure,
	}
}

func mlinConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, nil
}