package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/netlify/open-api/go/plumbing/operations"
	"io"
	"log"
	"net/http"
	"os"
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
			for page = 0;; page++ {
				params := &operations.ListSitesParams{
					Filter: &filter,
					Page: &page,
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
					log.Printf("Deleting site %s (%s)\n", site.Name, site.ID)
					if err := meta.client.DeleteSite(meta.netlifyCtx, site.ID); err != nil {
						return err
					}
				}
			}
			return nil
		},
	})
}

func TestAccResourceSite(t *testing.T) {
	meta := NewMeta(context.Background(), "test", os.Getenv("NETLIFY_TOKEN"))
	siteName := "terraform-test-" + acctest.RandStringFromCharSet(15, acctest.CharSetAlphaNum)
	resource.Test(t, resource.TestCase{
		PreCheck: testAccPreCheck(t),
		ProviderFactories: providerFactories,
		CheckDestroy:      testVerifySitesDestroyed(meta),
		Steps: []resource.TestStep{
			{
				Config: netlifySiteConfig(siteName),
				Check:  testVerifySiteExists(siteName, meta),
			},
		},
	})
}

func testVerifySiteExists(siteName string, meta *Meta) func(state *terraform.State) error {

	return func(s *terraform.State) error {
		for _, res := range s.RootModule().Resources {
			if res.Type != "netlify_site" || res.Primary.Attributes["name"] != siteName {
				continue
			}

			site, err := meta.client.GetSite(meta.netlifyCtx, res.Primary.ID)
			if err != nil {
				return fmt.Errorf("site %s does not exist in netlify", res.Primary.ID)
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
			if strings.Contains(text, "You need to enable Javascript to run this app") {
				resp.Body.Close()
				return nil
			} else {
				resp.Body.Close()
				return fmt.Errorf("site contents were not deployed")
			}
		}

		return fmt.Errorf("site %s does not exist in terraform state", siteName)
	}
}

func testVerifySitesDestroyed(meta *Meta) func(s *terraform.State) error {
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

func netlifySiteConfig(siteName string) string {
	return fmt.Sprintf(`
resource "netlify_site" "example" {
	name = "%s"
	source_url = "https://github.com/ttbud/ttbud/tarball/0fd1df2188e120e00528321cdfb60f52c1a3a683"
}
`, siteName)
}
