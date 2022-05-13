package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMlinTest() *schema.Resource {
	return &schema.Resource{
		CreateContext:resourceMlinTestCreate,
		ReadContext: resourceMlinTestRead,
		DeleteContext: resourceMlinTestDelete,

		Schema: map[string]*schema.Schema{
			"attribute": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMlinTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	attr := d.Get("attribute").(string)
	d.SetId(attr)
	return nil
}

func resourceMlinTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceMlinTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}