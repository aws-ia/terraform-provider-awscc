---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_inspectorv2_filter Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Inspector Filter resource schema
---

# awscc_inspectorv2_filter (Resource)

Inspector Filter resource schema



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter_action` (String) Findings filter action.
- `filter_criteria` (Attributes) Findings filter criteria. (see [below for nested schema](#nestedatt--filter_criteria))
- `name` (String) Findings filter name.

### Optional

- `description` (String) Findings filter description.

### Read-Only

- `arn` (String) Findings filter ARN.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--filter_criteria"></a>
### Nested Schema for `filter_criteria`

Required:

- `aws_account_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--aws_account_id))
- `component_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--component_id))
- `component_type` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--component_type))
- `ec_2_instance_image_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ec_2_instance_image_id))
- `ec_2_instance_subnet_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ec_2_instance_subnet_id))
- `ec_2_instance_vpc_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ec_2_instance_vpc_id))
- `ecr_image_architecture` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ecr_image_architecture))
- `ecr_image_hash` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ecr_image_hash))
- `ecr_image_pushed_at` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ecr_image_pushed_at))
- `ecr_image_registry` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ecr_image_registry))
- `ecr_image_repository_name` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ecr_image_repository_name))
- `ecr_image_tags` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--ecr_image_tags))
- `finding_arn` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--finding_arn))
- `finding_status` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--finding_status))
- `finding_type` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--finding_type))
- `first_observed_at` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--first_observed_at))
- `inspector_score` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--inspector_score))
- `last_observed_at` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--last_observed_at))
- `network_protocol` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--network_protocol))
- `port_range` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--port_range))
- `related_vulnerabilities` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--related_vulnerabilities))
- `resource_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--resource_id))
- `resource_tags` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--resource_tags))
- `resource_type` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--resource_type))
- `severity` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--severity))
- `title` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--title))
- `updated_at` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--updated_at))
- `vendor_severity` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--vendor_severity))
- `vulnerability_id` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--vulnerability_id))
- `vulnerability_source` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--vulnerability_source))
- `vulnerable_packages` (Attributes List) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages))

<a id="nestedatt--filter_criteria--aws_account_id"></a>
### Nested Schema for `filter_criteria.aws_account_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--component_id"></a>
### Nested Schema for `filter_criteria.component_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--component_type"></a>
### Nested Schema for `filter_criteria.component_type`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ec_2_instance_image_id"></a>
### Nested Schema for `filter_criteria.ec_2_instance_image_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ec_2_instance_subnet_id"></a>
### Nested Schema for `filter_criteria.ec_2_instance_subnet_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ec_2_instance_vpc_id"></a>
### Nested Schema for `filter_criteria.ec_2_instance_vpc_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ecr_image_architecture"></a>
### Nested Schema for `filter_criteria.ecr_image_architecture`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ecr_image_hash"></a>
### Nested Schema for `filter_criteria.ecr_image_hash`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ecr_image_pushed_at"></a>
### Nested Schema for `filter_criteria.ecr_image_pushed_at`

Required:

- `end_inclusive` (Number)
- `start_inclusive` (Number)


<a id="nestedatt--filter_criteria--ecr_image_registry"></a>
### Nested Schema for `filter_criteria.ecr_image_registry`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ecr_image_repository_name"></a>
### Nested Schema for `filter_criteria.ecr_image_repository_name`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--ecr_image_tags"></a>
### Nested Schema for `filter_criteria.ecr_image_tags`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--finding_arn"></a>
### Nested Schema for `filter_criteria.finding_arn`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--finding_status"></a>
### Nested Schema for `filter_criteria.finding_status`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--finding_type"></a>
### Nested Schema for `filter_criteria.finding_type`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--first_observed_at"></a>
### Nested Schema for `filter_criteria.first_observed_at`

Required:

- `end_inclusive` (Number)
- `start_inclusive` (Number)


<a id="nestedatt--filter_criteria--inspector_score"></a>
### Nested Schema for `filter_criteria.inspector_score`

Required:

- `lower_inclusive` (Number)
- `upper_inclusive` (Number)


<a id="nestedatt--filter_criteria--last_observed_at"></a>
### Nested Schema for `filter_criteria.last_observed_at`

Required:

- `end_inclusive` (Number)
- `start_inclusive` (Number)


<a id="nestedatt--filter_criteria--network_protocol"></a>
### Nested Schema for `filter_criteria.network_protocol`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--port_range"></a>
### Nested Schema for `filter_criteria.port_range`

Required:

- `begin_inclusive` (Number)
- `end_inclusive` (Number)


<a id="nestedatt--filter_criteria--related_vulnerabilities"></a>
### Nested Schema for `filter_criteria.related_vulnerabilities`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--resource_id"></a>
### Nested Schema for `filter_criteria.resource_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--resource_tags"></a>
### Nested Schema for `filter_criteria.resource_tags`

Required:

- `comparison` (String)
- `key` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--resource_type"></a>
### Nested Schema for `filter_criteria.resource_type`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--severity"></a>
### Nested Schema for `filter_criteria.severity`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--title"></a>
### Nested Schema for `filter_criteria.title`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--updated_at"></a>
### Nested Schema for `filter_criteria.updated_at`

Required:

- `end_inclusive` (Number)
- `start_inclusive` (Number)


<a id="nestedatt--filter_criteria--vendor_severity"></a>
### Nested Schema for `filter_criteria.vendor_severity`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerability_id"></a>
### Nested Schema for `filter_criteria.vulnerability_id`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerability_source"></a>
### Nested Schema for `filter_criteria.vulnerability_source`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerable_packages"></a>
### Nested Schema for `filter_criteria.vulnerable_packages`

Required:

- `architecture` (Attributes) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages--architecture))
- `epoch` (Attributes) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages--epoch))
- `name` (Attributes) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages--name))
- `release` (Attributes) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages--release))
- `source_layer_hash` (Attributes) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages--source_layer_hash))
- `version` (Attributes) (see [below for nested schema](#nestedatt--filter_criteria--vulnerable_packages--version))

<a id="nestedatt--filter_criteria--vulnerable_packages--architecture"></a>
### Nested Schema for `filter_criteria.vulnerable_packages.architecture`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerable_packages--epoch"></a>
### Nested Schema for `filter_criteria.vulnerable_packages.epoch`

Required:

- `lower_inclusive` (Number)
- `upper_inclusive` (Number)


<a id="nestedatt--filter_criteria--vulnerable_packages--name"></a>
### Nested Schema for `filter_criteria.vulnerable_packages.name`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerable_packages--release"></a>
### Nested Schema for `filter_criteria.vulnerable_packages.release`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerable_packages--source_layer_hash"></a>
### Nested Schema for `filter_criteria.vulnerable_packages.source_layer_hash`

Required:

- `comparison` (String)
- `value` (String)


<a id="nestedatt--filter_criteria--vulnerable_packages--version"></a>
### Nested Schema for `filter_criteria.vulnerable_packages.version`

Required:

- `comparison` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_inspectorv2_filter.example <resource ID>
```
