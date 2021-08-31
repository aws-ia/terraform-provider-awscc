// Code generated by generators/resource/main.go; DO NOT EDIT.

package emr

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_emr_studio_session_mapping", studioSessionMappingResourceType)
}

// studioSessionMappingResourceType returns the Terraform awscc_emr_studio_session_mapping resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EMR::StudioSessionMapping resource type.
func studioSessionMappingResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"identity_name": {
			// Property: IdentityName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.",
			//   "type": "string"
			// }
			Description: "The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.",
			Type:        types.StringType,
			Required:    true,
			// IdentityName is a force-new attribute.
		},
		"identity_type": {
			// Property: IdentityType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the identity to map to the Studio is a user or a group.",
			//   "enum": [
			//     "USER",
			//     "GROUP"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies whether the identity to map to the Studio is a user or a group.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"USER",
					"GROUP",
				}),
			},
			// IdentityType is a force-new attribute.
		},
		"session_policy_arn": {
			// Property: SessionPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the Amazon EMR Studio to which the user or group will be mapped.",
			//   "maxLength": 256,
			//   "minLength": 4,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the Amazon EMR Studio to which the user or group will be mapped.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(4, 256),
			},
			// StudioId is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMR::StudioSessionMapping").WithTerraformTypeName("awscc_emr_studio_session_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"identity_name":      "IdentityName",
		"identity_type":      "IdentityType",
		"session_policy_arn": "SessionPolicyArn",
		"studio_id":          "StudioId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_emr_studio_session_mapping", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
