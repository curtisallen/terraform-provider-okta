// nolint
package main

import (
	"context"
	"fmt"
	"time"

	okta "github.com/curtisallen/go-okta"
	"github.com/hashicorp/terraform/helper/schema"
)

const timeout = 5 * time.Second

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)

	// Create the new group
	group := okta.Group{
		Profile: okta.GroupProfile{
			Name: d.Get("name").(string),
		},
	}

	if desc, ok := d.GetOk("description"); ok {
		group.Profile.Description = desc.(string)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resp, err := client.CreateGroup(ctx, group)
	if err != nil {
		return fmt.Errorf("Failed to create Okta group: %s", err)
	}

	d.SetId(resp.ID)
	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	group, err := client.Group(ctx, d.Id())
	if err != nil {
		return fmt.Errorf("Couldn't get group: %s", err)
	}

	d.Set("name", group.Profile.Name)
	d.Set("description", group.Profile.Description)

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)

	updateGroup := okta.Group{ID: d.Id()}

	if attr, ok := d.GetOk("name"); ok {
		updateGroup.Profile.Name = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		updateGroup.Profile.Description = attr.(string)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := client.UpdateGroup(ctx, updateGroup)
	if err != nil {
		return fmt.Errorf("Unable to update group : %s", err)
	}

	return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*okta.Client)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := client.DeleteGroup(ctx, d.Id())
	if err != nil {
		return fmt.Errorf("Unable to delete group: %s", err)
	}

	return nil
}
