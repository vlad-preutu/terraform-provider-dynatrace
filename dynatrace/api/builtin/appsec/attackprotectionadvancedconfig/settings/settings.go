/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package attackprotectionadvancedconfig

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	AttackHandling              *AttackHandling             `json:"attackHandling"`                        // Step 1: Select attack protection behavior
	Criteria                    *Criteria                   `json:"criteria"`                              // Step 2: Select attack type
	Enabled                     bool                        `json:"enabled"`                               // This setting is enabled (`true`) or disabled (`false`)
	Metadata                    *Metadata                   `json:"metadata"`                              // Step 4: Leave comment (optional)
	ResourceAttributeConditions ResourceAttributeConditions `json:"resourceAttributeConditions,omitempty"` // If you add more than one condition, note that all conditions must be true simultaneously for the rule to apply.\n\nWe provide suggestions for resource attribute keys and values based on what we currently see in your environment. You can also enter any value not currently seen in the list. Resource attributes come out of the box from the OneAgent, and you can set them up from [data enrichment](https://docs.dynatrace.com/docs/extend-dynatrace/extend-data).
	RuleName                    *string                     `json:"ruleName,omitempty"`                    // Rule name
	InsertAfter                 string                      `json:"-"`
}

func (me *Settings) Name() string {
	return uuid.NewString()
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attack_handling": {
			Type:        schema.TypeList,
			Description: "Step 1: Select attack protection behavior",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(AttackHandling).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"criteria": {
			Type:        schema.TypeList,
			Description: "Step 2: Select attack type",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Criteria).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        schema.TypeBool,
			Description: "This setting is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"metadata": {
			Type:        schema.TypeList,
			Description: "Step 4: Leave comment (optional)",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Metadata).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"resource_attribute_conditions": {
			Type:        schema.TypeList,
			Description: "If you add more than one condition, note that all conditions must be true simultaneously for the rule to apply.\n\nWe provide suggestions for resource attribute keys and values based on what we currently see in your environment. You can also enter any value not currently seen in the list. Resource attributes come out of the box from the OneAgent, and you can set them up from [data enrichment](https://docs.dynatrace.com/docs/extend-dynatrace/extend-data).",
			Optional:    true, // minobjects == 0
			Elem:        &schema.Resource{Schema: new(ResourceAttributeConditions).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"rule_name": {
			Type:        schema.TypeString,
			Description: "Rule name",
			Optional:    true, // nullable
		},
		"insert_after": {
			Type:        schema.TypeString,
			Description: "Because this resource allows for ordering you may specify the ID of the resource instance that comes before this instance regarding order. If not specified when creating the setting will be added to the end of the list. If not specified during update the order will remain untouched",
			Optional:    true,
			Computed:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"attack_handling":               me.AttackHandling,
		"criteria":                      me.Criteria,
		"enabled":                       me.Enabled,
		"metadata":                      me.Metadata,
		"resource_attribute_conditions": me.ResourceAttributeConditions,
		"rule_name":                     me.RuleName,
		"insert_after":                  me.InsertAfter,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"attack_handling":               &me.AttackHandling,
		"criteria":                      &me.Criteria,
		"enabled":                       &me.Enabled,
		"metadata":                      &me.Metadata,
		"resource_attribute_conditions": &me.ResourceAttributeConditions,
		"rule_name":                     &me.RuleName,
		"insert_after":                  &me.InsertAfter,
	})
}
