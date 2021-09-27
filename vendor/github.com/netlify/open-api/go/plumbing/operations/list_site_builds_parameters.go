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
)

// NewListSiteBuildsParams creates a new ListSiteBuildsParams object
// with the default values initialized.
func NewListSiteBuildsParams() *ListSiteBuildsParams {
	var ()
	return &ListSiteBuildsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSiteBuildsParamsWithTimeout creates a new ListSiteBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSiteBuildsParamsWithTimeout(timeout time.Duration) *ListSiteBuildsParams {
	var ()
	return &ListSiteBuildsParams{

		timeout: timeout,
	}
}

// NewListSiteBuildsParamsWithContext creates a new ListSiteBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSiteBuildsParamsWithContext(ctx context.Context) *ListSiteBuildsParams {
	var ()
	return &ListSiteBuildsParams{

		Context: ctx,
	}
}

// NewListSiteBuildsParamsWithHTTPClient creates a new ListSiteBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSiteBuildsParamsWithHTTPClient(client *http.Client) *ListSiteBuildsParams {
	var ()
	return &ListSiteBuildsParams{
		HTTPClient: client,
	}
}

/*ListSiteBuildsParams contains all the parameters to send to the API endpoint
for the list site builds operation typically these are written to a http.Request
*/
type ListSiteBuildsParams struct {

	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list site builds params
func (o *ListSiteBuildsParams) WithTimeout(timeout time.Duration) *ListSiteBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list site builds params
func (o *ListSiteBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list site builds params
func (o *ListSiteBuildsParams) WithContext(ctx context.Context) *ListSiteBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list site builds params
func (o *ListSiteBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list site builds params
func (o *ListSiteBuildsParams) WithHTTPClient(client *http.Client) *ListSiteBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list site builds params
func (o *ListSiteBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the list site builds params
func (o *ListSiteBuildsParams) WithPage(page *int32) *ListSiteBuildsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list site builds params
func (o *ListSiteBuildsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the list site builds params
func (o *ListSiteBuildsParams) WithPerPage(perPage *int32) *ListSiteBuildsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the list site builds params
func (o *ListSiteBuildsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithSiteID adds the siteID to the list site builds params
func (o *ListSiteBuildsParams) WithSiteID(siteID string) *ListSiteBuildsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the list site builds params
func (o *ListSiteBuildsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
