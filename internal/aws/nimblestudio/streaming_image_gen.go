// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

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
	registry.AddResourceTypeFactory("aws_nimblestudio_streaming_image", streamingImageResourceType)
}

// streamingImageResourceType returns the Terraform aws_nimblestudio_streaming_image resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NimbleStudio::StreamingImage resource type.
func streamingImageResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"ec_2_image_id": {
			// Property: Ec2ImageId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// Ec2ImageId is a force-new attribute.
		},
		"encryption_configuration": {
			// Property: EncryptionConfiguration
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "KeyArn": {
			         "type": "string"
			       },
			       "KeyType": {
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/StreamingImageEncryptionConfiguration",
			     "required": [
			       "KeyType"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"key_arn": {
						// Property: KeyArn
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"key_type": {
						// Property: KeyType
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Computed: true,
		},
		"eula_ids": {
			// Property: EulaIds
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "type": "string"
			     },
			     "type": "array"
			   }
			*/
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
		},
		"owner": {
			// Property: Owner
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"platform": {
			// Property: Platform
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"streaming_image_id": {
			// Property: StreamingImageId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// StudioId is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "patternProperties": {
			       "": {
			         "type": "string"
			       }
			     },
			     "type": "object"
			   }
			*/
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::NimbleStudio::StreamingImage.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StreamingImage").WithTerraformTypeName("aws_nimblestudio_streaming_image").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_nimblestudio_streaming_image", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}