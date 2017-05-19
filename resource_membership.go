// nolint
package main

import (
	"context"
	"fmt"

	okta "github.com/curtisallen/go-okta"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceMembershipCreate,
		Read:   resourceMembershipRead,
		Update: resourceMembershipUpdate,
		Delete: resourceMembershipDelete,

		Schema: map[string]*schema.Schema{
			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "User email address e.g. dr.dre@example.com",
			},
		},
	}
}

func resourceMembershipCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)
	user := d.Get("user").(string)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	err := client.CreateMembership(ctx, d.Get("group_id").(string), user)
	if err != nil {
		return fmt.Errorf("Unable to create membership: %s", err)
	}
	d.SetId(user)

	return resourceMembershipRead(d, m)
}

func resourceMembershipRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)
	user := d.Get("user").(string)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	exists, err := client.MembershipExists(ctx, d.Get("group_id").(string), user)
	if err != nil {
		return fmt.Errorf("Unable to read membership: %s", err)
	}

	if exists {
		d.SetId(user)
	}

	return nil
}

func resourceMembershipUpdate(d *schema.ResourceData, m interface{}) error {
	// not implemented
	return nil
}

func resourceMembershipDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)
	user := d.Get("user").(string)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := client.DeleteMembership(ctx, d.Get("group_id").(string), user)
	if err != nil {
		return fmt.Errorf("Unable to read membership: %s", err)
	}
	return nil
}
