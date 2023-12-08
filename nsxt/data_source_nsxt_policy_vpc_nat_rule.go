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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceNsxtPolicyVpcNatRule() *schema.Resource {
	additionalSchemaItems := map[string]*schema.Schema{} // Define any additional schema items specific to the data source

	additionalSchemaItems["parent_path"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	}

	return &schema.Resource{
		Read:   dataSourceNsxtPolicyVpcNatRuleRead,
		Schema: getDataSourceCommonSchema(additionalSchemaItems),
	}
}

func dataSourceNsxtPolicyVpcNatRuleRead(d *schema.ResourceData, meta interface{}) error {
	s := dataSourceNsxtPolicyVpcNatRule()
	err := DatasourceRead(d, meta, "PolicyVpcNatRule", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}
