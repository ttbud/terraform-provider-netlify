package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/porcelain"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{
		Description: "Netlify Site Resource",

		CreateContext: resourceSiteCreate,
		ReadContext:   resourceSiteRead,
		UpdateContext: resourceSiteUpdate,
		DeleteContext: resourceSiteDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_url": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "URL to a targz of the source to deploy." +
					" Destination contents must be immutable, only changes to the URL will trigger a re-deploy",
			},
		},
	}
}

func resourceSiteCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	meta := m.(*Meta)

	siteSetup := &models.SiteSetup{
		Site: models.Site{
			Name:         d.Get("name").(string),
			CustomDomain: d.Get("custom_domain").(string),
		},
	}

	site, err := meta.client.CreateSite(meta.netlifyCtx, siteSetup, false)
	if err != nil {
		return diag.FromErr(err)
	}

	sourceURL, ok := d.GetOk("source_url")
	if ok {
		if err = deploySource(meta, site.ID, sourceURL.(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	if err = applySiteProperties(d, site); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func resourceSiteRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	meta := m.(*Meta)
	siteID := d.Get("site_id").(string)

	site, err := meta.client.GetSite(meta.netlifyCtx, siteID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = applySiteProperties(d, site); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}

func resourceSiteUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	meta := m.(*Meta)
	if d.HasChanges("name", "custom_domain") {
		siteSetup := &models.SiteSetup{
			Site: models.Site{
				ID:           d.Get("site_id").(string),
				Name:         d.Get("name").(string),
				CustomDomain: d.Get("custom_domain").(string),
			},
		}

		site, err := meta.client.UpdateSite(meta.netlifyCtx, siteSetup)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = applySiteProperties(d, site); err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("source_url") {
		if err := deploySource(meta, d.Id(), d.Get("source_url").(string)); err != nil {
			return diag.FromErr(err)
		}
	}

	return diag.Diagnostics{}
}

func resourceSiteDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	meta := m.(*Meta)
	siteID := d.Get("site_id").(string)
	if err := meta.client.DeleteSite(meta.netlifyCtx, siteID); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

func applySiteProperties(d *schema.ResourceData, site *models.Site) error {
	d.SetId(site.ID)
	if err := d.Set("site_id", site.ID); err != nil {
		return err
	}
	if err := d.Set("name", site.Name); err != nil {
		return err
	}
	if err := d.Set("custom_domain", site.CustomDomain); err != nil {
		return err
	}

	return nil
}

func deploySource(meta *Meta, siteID string, sourceURL string) error {
	// Unpack remote tarball to filesystem
	resp, err := http.Get(sourceURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	unpackDir, err := os.MkdirTemp("", "netlify-terraform-deploy-source-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(unpackDir)
	if err = Untar(unpackDir, resp.Body); err != nil {
		return err
	}

	deployOptions := porcelain.DeployOptions{SiteID: siteID, Dir: unpackDir}
	deploy, err := meta.client.DeploySite(meta.netlifyCtx, deployOptions)
	if err != nil {
		return err
	}

	_, err = meta.client.WaitUntilDeployLive(meta.netlifyCtx, deploy, time.Minute*10)
	if err != nil {
		return err
	}

	return nil
}
