package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// schema.ResourceData APIを利用している
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	// SetIdを利用することで、resouceのaddressを変更を検知できるようになる
	d.SetId(address)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("address") {
		// Try updating the address
		if err := updateAddress(d, meta); err != nil {
			return err
		}

		d.SetPartial("address")
	}

	// If we were to return here, before disabling partial mode below,
	// then only the "address" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// save all fields again.
	d.Partial(false)

	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			// パラメータの定義
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
