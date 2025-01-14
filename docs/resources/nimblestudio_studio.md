---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_nimblestudio_studio Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a studio that contains other Nimble Studio resources
---

# awscc_nimblestudio_studio (Resource)

Represents a studio that contains other Nimble Studio resources



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_role_arn` (String) <p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>
- `display_name` (String) <p>A friendly name for the studio.</p>
- `studio_name` (String) <p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>
- `user_role_arn` (String) <p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>

### Optional

- `studio_encryption_configuration` (Attributes) <p>Configuration of the encryption method that is used for the studio.</p> (see [below for nested schema](#nestedatt--studio_encryption_configuration))
- `tags` (Map of String)

### Read-Only

- `home_region` (String) <p>The Amazon Web Services Region where the studio resource is located.</p>
- `id` (String) Uniquely identifies the resource.
- `sso_client_id` (String) <p>The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.</p>
- `studio_id` (String)
- `studio_url` (String) <p>The address of the web page for the studio.</p>

<a id="nestedatt--studio_encryption_configuration"></a>
### Nested Schema for `studio_encryption_configuration`

Optional:

- `key_arn` (String) <p>The ARN for a KMS key that is used to encrypt studio data.</p>
- `key_type` (String) <p>The type of KMS key that is used to encrypt studio data.</p>

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_nimblestudio_studio.example <resource ID>
```
