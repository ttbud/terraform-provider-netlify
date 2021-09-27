---
page_title: "netlify_site Resource - terraform-provider-netlify"
subcategory: ""
description: |-
  Resource to deploy sites to Netlify
---

# Resource `netlify_site`

Resource to deploy sites to Netlify

## Example Usage

```terraform
resource "netlify_site" "example" {
  name = "foo"
  custom_domain = "example.com"
  source_url = "https://example.com/path/to/tarball"
}
```

## Schema

- **name** (String) The name of the site.
- **source_url** (String) A URL that contains a tarball of the source to be deployed.
- 
### Optional

- **custom_domain** (String, Optional) The custom domain to deploy this site to
