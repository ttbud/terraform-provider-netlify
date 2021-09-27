resource "netlify_site" "example" {
  name = "foo"
  custom_domain = "example.com"
  source_url = "https://example.com/path/to/tarball"
}