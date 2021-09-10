---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53resolver_resolver_query_logging_config Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Route53Resolver::ResolverQueryLoggingConfig
---

# awscc_route53resolver_resolver_query_logging_config (Data Source)

Data Source schema for AWS::Route53Resolver::ResolverQueryLoggingConfig



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) Arn
- **association_count** (Number) Count
- **creation_time** (String) Rfc3339TimeString
- **creator_request_id** (String) The id of the creator request.
- **destination_arn** (String) destination arn
- **name** (String) ResolverQueryLogConfigName
- **owner_id** (String) AccountId
- **share_status** (String) ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
- **status** (String) ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.

