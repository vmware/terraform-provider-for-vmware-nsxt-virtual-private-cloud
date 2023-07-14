/***************************************************************************
 * ========================================================================
 * Copyright 2022-2023 VMware, Inc.  All rights reserved. VMware Confidential
 * SPDX-License-Identifier: MPL-2.0
 * ========================================================================
 */

// Auto generated code. DO NOT EDIT.

// nolint
package nsxt

import (
	nsxtclient "github.com/vmware/terraform-provider-for-vmware-nsxt-virtual-private-cloud/nsxt/clients"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpcIpAddressAllocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"allocation_ip": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"_revision": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"ip_address_block_visibility": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"EXTERNAL", "PRIVATE"}, false),
			Default:      "EXTERNAL",
		},
		"ip_address_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"IPV4", "IPV6"}, false),
			Default:      "IPV4",
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 30,
			Elem:     resourceTagSchema(),
		},
		"nsx_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func resourceNsxtVpcIpAddressAllocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtVpcIpAddressAllocationCreate,
		Read:   resourceNsxtVpcIpAddressAllocationRead,
		Update: resourceNsxtVpcIpAddressAllocationUpdate,
		Delete: resourceNsxtVpcIpAddressAllocationDelete,
		Schema: resourceVpcIpAddressAllocationSchema(),
		Importer: &schema.ResourceImporter{
			State: resourceVpcIpAddressAllocationImporter,
		},
	}
}

func resourceVpcIpAddressAllocationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := resourceVpcIpAddressAllocationSchema()
	return ResourceImporter(d, m, "VpcIpAddressAllocation", s, d.Id())
}

func resourceNsxtVpcIpAddressAllocationRead(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcIpAddressAllocationSchema()
	err := APIRead(d, meta, "VpcIpAddressAllocation", s)
	// if 404 not found error occurs, terraform should swallow it and not fail read on object
	if err != nil && strings.Contains(err.Error(), "404") {
		log.Printf("[WARNING] Failed to read object VpcIpAddressAllocation %v\n", err)
		return nil
	}
	return err
}

func resourceNsxtVpcIpAddressAllocationCreate(d *schema.ResourceData, meta interface{}) error {
	s := resourceVpcIpAddressAllocationSchema()
	err := APICreateOrUpdate(d, meta, "VpcIpAddressAllocation", s)
	if err == nil {
		err = resourceNsxtVpcIpAddressAllocationRead(d, meta)
	}
	return err
}

func resourceNsxtVpcIpAddressAllocationUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceNsxtVpcIpAddressAllocationDelete(d *schema.ResourceData, meta interface{}) error {
	nsxtClient := meta.(*nsxtclient.NsxtClient)
	resourceID := d.Id()
	if resourceID != "" {
		path := nsxtClient.Config.BasePath + d.Get("path").(string)
		err := nsxtClient.NsxtSession.Delete(path)
		// if 'object not found' or 'forbidden' or 'success with no response' response occurs, terraform should swallow it and not fail apply on object, else throw error and fail
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Printf("[INFO] Error occurred in Delete for resource VpcIpAddressAllocation \n")
			return err
		}
		d.SetId("")
	}
	return nil
}
