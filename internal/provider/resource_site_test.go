package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/netlify/open-api/go/plumbing/operations"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("netlify_site", &resource.Sweeper{
		Name: "netlify_site",
		F: func(region string) error {
			meta := NewTestMeta()
			filter := "terraform-test-"
			var page int32
			for page = 0; ; page++ {
				params := &operations.ListSitesParams{
					Filter:  &filter,
					Page:    &page,
					Context: meta.netlifyCtx,
				}
				sites, err := meta.client.ListSites(meta.netlifyCtx, params)
				if err != nil {
					return err
				}
				if len(sites) < 1 {
					break
				}
				for _, site := range sites {
					log.Printf("[INFO] Deleting site %s (%s)\n", site.Name, site.ID)
					if err := meta.client.DeleteSite(meta.netlifyCtx, site.ID); err != nil {
						return err
					}
				}
			}
			return nil
		},
	})
}

func TestAccResourceSite_basic(t *testing.T) {
	meta := NewTestMeta()
	siteName := "terraform-test-" + acctest.RandStringFromCharSet(15, acctest.CharSetAlphaNum)
	newSiteName := "terraform-test-" + acctest.RandStringFromCharSet(15, acctest.CharSetAlphaNum)

	// go test uses a different working directory depending on whether you use a directory or package specification when
	// choosing which tests to run, so we have to find the testdata directory based on our current file
	_, filename, _, _ := runtime.Caller(0)
	serveDir := path.Join(path.Dir(filename), "testdata", "serve")
	server := httptest.NewServer(http.FileServer(http.Dir(serveDir)))
	defer server.Close()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      verifySitesDestroyed(meta),
		Steps: []resource.TestStep{
			{
				Config: netlifySiteConfig(siteName, server.URL+"/initial_site.tar.gz"),
				Check:  verifySiteContentExists(siteName, "Initial Site", meta),
			},
			{
				Config: netlifySiteConfig(newSiteName, server.URL+"/updated_site.tar.gz"),
				Check:  verifySiteContentExists(newSiteName, "Updated Site", meta),
			},
		},
	})
}

func TestAccResourceSite_no_source_url(t *testing.T) {
	meta := NewTestMeta()
	siteName := "terraform-test-" + acctest.RandStringFromCharSet(15, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      verifySitesDestroyed(meta),
		Steps: []resource.TestStep{
			{
				Config: netlifySiteConfigSansURL(siteName),
			},
			{
				ResourceName:      "netlify_site.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func verifySiteContentExists(siteName string, expectedContent string, meta *Meta) func(state *terraform.State) error {

	return func(s *terraform.State) error {
		for _, res := range s.RootModule().Resources {
			if res.Type != "netlify_site" || res.Primary.Attributes["name"] != siteName {
				continue
			}

			site, err := meta.client.GetSite(meta.netlifyCtx, res.Primary.ID)
			if err != nil {
				return fmt.Errorf("site %s does not exist in netlify", res.Primary.ID)
			}
			if site.Name != siteName {
				return fmt.Errorf("expected site name %s in netlify, found site name %s", siteName, site.Name)
			}

			resp, err := http.Get(site.URL)
			if err != nil {
				return err
			}
			bytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			text := string(bytes)
			if strings.Contains(text, expectedContent) {
				resp.Body.Close()
				return nil
			} else {
				resp.Body.Close()
				return fmt.Errorf("could not find %s in response:\n%s\n", expectedContent, text)
			}
		}

		return fmt.Errorf("site %s does not exist in terraform state", siteName)
	}
}

func verifySitesDestroyed(meta *Meta) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for _, res := range s.RootModule().Resources {
			if res.Type != "netlify_site" {
				continue
			}

			if _, err := meta.client.GetSite(meta.netlifyCtx, res.Primary.ID); err == nil {
				return fmt.Errorf("site %s still exists", res.Primary.ID)
			}
		}

		return nil
	}
}

func netlifySiteConfigSansURL(siteName string) string {
	return fmt.Sprintf(`
resource "netlify_site" "example" {
	name = "%s"
}
`, siteName)
}

func netlifySiteConfig(siteName string, sourceURL string) string {
	return fmt.Sprintf(`
resource "netlify_site" "example" {
	name = "%s"
	source_url = "%s"
}
`, siteName, sourceURL)
}
