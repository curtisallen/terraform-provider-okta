# Terraform Okta provider

A minimal terraform provider that manages group memberships in Okta.

## Usage

To install the plugin

    go get -u github.com/curtisallen/terraform-provider-okta

## Build

    make init
    make build

## Configuration

To build this plugin you'll need a [golang development environment setup](https://golang.org/doc/install) and [terraform](https://github.com/hashicorp/terraform) cloned.

The Okta API token can be set with `OKTA_TOKEN` environment variable.

## Example Terraform

```hcl
provider "okta" {
  organization = "org id"
  // can also use OKTA_TOKEN environment variable
  token = "api token"
  // set to true if targeting preview environment
  preview = true
}

resource "okta_group" "my-group" {
  name = "test group terraform"
  description = "holla"
}

resource "okta_membership" "curtis" {
  group_id = "${okta_group.my-group.id}"
  user = "easy.e@example.com"
}

resource "okta_membership" "jim" {
  group_id = "${okta_group.my-group.id}"
  user = "dr.dre@example.com"
}
```
