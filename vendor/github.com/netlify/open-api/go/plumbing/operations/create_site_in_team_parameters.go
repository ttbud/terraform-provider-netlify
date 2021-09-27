// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netlify/open-api/go/models"
)

// NewCreateSiteInTeamParams creates a new CreateSiteInTeamParams object
// with the default values initialized.
func NewCreateSiteInTeamParams() *CreateSiteInTeamParams {
	var ()
	return &CreateSiteInTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSiteInTeamParamsWithTimeout creates a new CreateSiteInTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSiteInTeamParamsWithTimeout(timeout time.Duration) *CreateSiteInTeamParams {
	var ()
	return &CreateSiteInTeamParams{

		timeout: timeout,
	}
}

// NewCreateSiteInTeamParamsWithContext creates a new CreateSiteInTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSiteInTeamParamsWithContext(ctx context.Context) *CreateSiteInTeamParams {
	var ()
	return &CreateSiteInTeamParams{

		Context: ctx,
	}
}

// NewCreateSiteInTeamParamsWithHTTPClient creates a new CreateSiteInTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSiteInTeamParamsWithHTTPClient(client *http.Client) *CreateSiteInTeamParams {
	var ()
	return &CreateSiteInTeamParams{
		HTTPClient: client,
	}
}

/*CreateSiteInTeamParams contains all the parameters to send to the API endpoint
for the create site in team operation typically these are written to a http.Request
*/
type CreateSiteInTeamParams struct {

	/*AccountSlug*/
	AccountSlug string
	/*ConfigureDNS*/
	ConfigureDNS *bool
	/*Site*/
	Site *models.SiteSetup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create site in team params
func (o *CreateSiteInTeamParams) WithTimeout(timeout time.Duration) *CreateSiteInTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create site in team params
func (o *CreateSiteInTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create site in team params
func (o *CreateSiteInTeamParams) WithContext(ctx context.Context) *CreateSiteInTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create site in team params
func (o *CreateSiteInTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create site in team params
func (o *CreateSiteInTeamParams) WithHTTPClient(client *http.Client) *CreateSiteInTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create site in team params
func (o *CreateSiteInTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountSlug adds the accountSlug to the create site in team params
func (o *CreateSiteInTeamParams) WithAccountSlug(accountSlug string) *CreateSiteInTeamParams {
	o.SetAccountSlug(accountSlug)
	return o
}

// SetAccountSlug adds the accountSlug to the create site in team params
func (o *CreateSiteInTeamParams) SetAccountSlug(accountSlug string) {
	o.AccountSlug = accountSlug
}

// WithConfigureDNS adds the configureDNS to the create site in team params
func (o *CreateSiteInTeamParams) WithConfigureDNS(configureDNS *bool) *CreateSiteInTeamParams {
	o.SetConfigureDNS(configureDNS)
	return o
}

// SetConfigureDNS adds the configureDns to the create site in team params
func (o *CreateSiteInTeamParams) SetConfigureDNS(configureDNS *bool) {
	o.ConfigureDNS = configureDNS
}

// WithSite adds the site to the create site in team params
func (o *CreateSiteInTeamParams) WithSite(site *models.SiteSetup) *CreateSiteInTeamParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the create site in team params
func (o *CreateSiteInTeamParams) SetSite(site *models.SiteSetup) {
	o.Site = site
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSiteInTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_slug
	if err := r.SetPathParam("account_slug", o.AccountSlug); err != nil {
		return err
	}

	if o.ConfigureDNS != nil {

		// query param configure_dns
		var qrConfigureDNS bool
		if o.ConfigureDNS != nil {
			qrConfigureDNS = *o.ConfigureDNS
		}
		qConfigureDNS := swag.FormatBool(qrConfigureDNS)
		if qConfigureDNS != "" {
			if err := r.SetQueryParam("configure_dns", qConfigureDNS); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {
		if err := r.SetBodyParam(o.Site); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
