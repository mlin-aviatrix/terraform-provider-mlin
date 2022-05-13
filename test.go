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
			"attr": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
				Deprecated: "Use 'attribute' instead of 'attr'",
				Description: "attr.",
			},
			"attribute": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
				ConflictsWith: []string{"attr"},
				Description: "Attribute",
			},
		},
	}
}

func resourceMlinTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if _, ok := d.GetOk("attr"); ok {
		attr := d.Get("attr").(string)
		d.SetId(attr)
	} else if _, ok := d.GetOk("attribute"); ok {
		attribute := d.Get("attribute").(string)
		d.SetId(attribute)
	} else {
		return diag.Errorf("one of %q or %q must be set", "attr", "attribute")
	}

	return nil
}

func resourceMlinTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceMlinTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}