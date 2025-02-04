---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_billingconductor_pricing_plan Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.
---

# awscc_billingconductor_pricing_plan (Resource)

Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `description` (String)
- `pricing_rule_arns` (List of String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Pricing Plan ARN
- `creation_time` (Number) Creation timestamp in UNIX epoch time format
- `id` (String) Uniquely identifies the resource.
- `last_modified_time` (Number) Latest modified timestamp in UNIX epoch time format
- `size` (Number) Number of associated pricing rules

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_billingconductor_pricing_plan.example <resource ID>
```
