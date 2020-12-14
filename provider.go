package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		// terraformのresourceとの関連づけを行っている
		ResourcesMap: map[string]*schema.Resource{
			"example_server": resourceServer(),
		},
	}
}
