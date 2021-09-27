---
page_title: "Netlify Provider"
subcategory: ""
description: |-
  
---

# netlify Provider



## Example Usage

```terraform
provider "netlify" {
  token: var.netlify_token
}
```

## Schema

### Optional

- **token** (String, Optional) The access token to use to deploy to your site, must be specified if the NETLIFY_TOKEN
environment variable is not set
