provider "okta" {
  organization = "dev-111111"
  token = "api token"
  preview = true
}

resource "okta_group" "my-group" {
  name = "test group terraform"
  description = "holla"
}

resource "okta_membership" "curtis" {
  group_id = "${okta_group.my-group.id}"
  user = "dr.dre@example.com"
}

resource "okta_membership" "jim" {
  group_id = "${okta_group.my-group.id}"
  user = "easy.e@example.com"
}
