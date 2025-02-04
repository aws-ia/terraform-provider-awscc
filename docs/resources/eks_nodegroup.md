---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_eks_nodegroup Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::EKS::Nodegroup
---

# awscc_eks_nodegroup (Resource)

Resource schema for AWS::EKS::Nodegroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) Name of the cluster to create the node group in.
- `node_role` (String) The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
- `subnets` (List of String) The subnets to use for the Auto Scaling group that is created for your node group.

### Optional

- `ami_type` (String) The AMI type for your node group.
- `capacity_type` (String) The capacity type of your managed node group.
- `disk_size` (Number) The root device disk size (in GiB) for your node group instances.
- `force_update_enabled` (Boolean) Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
- `instance_types` (List of String) Specify the instance types for a node group.
- `labels` (Map of String) The Kubernetes labels to be applied to the nodes in the node group when they are created.
- `launch_template` (Attributes) An object representing a node group's launch template specification. (see [below for nested schema](#nestedatt--launch_template))
- `nodegroup_name` (String) The unique name to give your node group.
- `release_version` (String) The AMI version of the Amazon EKS-optimized AMI to use with your node group.
- `remote_access` (Attributes) The remote access (SSH) configuration to use with your node group. (see [below for nested schema](#nestedatt--remote_access))
- `scaling_config` (Attributes) The scaling configuration details for the Auto Scaling group that is created for your node group. (see [below for nested schema](#nestedatt--scaling_config))
- `tags` (Map of String) The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
- `taints` (Attributes List) The Kubernetes taints to be applied to the nodes in the node group when they are created. (see [below for nested schema](#nestedatt--taints))
- `update_config` (Attributes) The node group update configuration. (see [below for nested schema](#nestedatt--update_config))
- `version` (String) The Kubernetes version to use for your managed nodes.

### Read-Only

- `arn` (String)
- `id` (String) The ID of this resource.

<a id="nestedatt--launch_template"></a>
### Nested Schema for `launch_template`

Optional:

- `id` (String) The ID of this resource.
- `name` (String)
- `version` (String)


<a id="nestedatt--remote_access"></a>
### Nested Schema for `remote_access`

Optional:

- `ec_2_ssh_key` (String)
- `source_security_groups` (List of String)


<a id="nestedatt--scaling_config"></a>
### Nested Schema for `scaling_config`

Optional:

- `desired_size` (Number)
- `max_size` (Number)
- `min_size` (Number)


<a id="nestedatt--taints"></a>
### Nested Schema for `taints`

Optional:

- `effect` (String)
- `key` (String)
- `value` (String)


<a id="nestedatt--update_config"></a>
### Nested Schema for `update_config`

Optional:

- `max_unavailable` (Number) The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100.
- `max_unavailable_percentage` (Number) The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_eks_nodegroup.example <resource ID>
```
