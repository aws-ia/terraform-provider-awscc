// Code generated by generators/resource/main.go; DO NOT EDIT.

package location

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_location_map", mapResourceType)
}

// mapResourceType returns the Terraform aws_location_map resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Location::Map resource type.
func mapResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1600,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Style": {
			         "maxLength": 100,
			         "minLength": 1,
			         "pattern": "",
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/MapConfiguration",
			     "required": [
			       "Style"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"style": {
						// Property: Style
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 100,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
			// Configuration is a force-new attribute.
		},
		"create_time": {
			// Property: CreateTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			     "pattern": "",
			     "$ref": "#/definitions/iso8601UTC",
			     "type": "string"
			   }
			*/
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
		"data_source": {
			// Property: DataSource
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1000,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// Description is a force-new attribute.
		},
		"map_arn": {
			// Property: MapArn
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1600,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"map_name": {
			// Property: MapName
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 100,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// MapName is a force-new attribute.
		},
		"pricing_plan": {
			// Property: PricingPlan
			// CloudFormation resource type schema:
			/*
			   {
			     "enum": [
			       "RequestBasedUsage",
			       "MobileAssetTracking",
			       "MobileAssetManagement"
			     ],
			     "$ref": "#/definitions/PricingPlan",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// PricingPlan is a force-new attribute.
		},
		"update_time": {
			// Property: UpdateTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			     "pattern": "",
			     "$ref": "#/definitions/iso8601UTC",
			     "type": "string"
			   }
			*/
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Definition of AWS::Location::Map Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::Map").WithTerraformTypeName("aws_location_map").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_location_map", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}