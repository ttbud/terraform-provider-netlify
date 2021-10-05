package provider

import (
	"context"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/netlify/open-api/go/porcelain"
	NetlifyContext "github.com/netlify/open-api/go/porcelain/context"
)

type Meta struct {
	client     *porcelain.Netlify
	netlifyCtx context.Context
}

func NewMeta(ctx context.Context, version string, token string) *Meta {
	authInfo := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("User-Agent", "github.com/ttbud/terraform-netlify-provider "+version); err != nil {
			return err
		}
		if err := r.SetHeaderParam("Authorization", "Bearer "+token); err != nil {
			return err
		}
		return nil
	})

	client := porcelain.NewRetryableHTTPClient(strfmt.Default, porcelain.DefaultRetryAttempts)
	netlifyCtx := NetlifyContext.WithAuthInfo(ctx, authInfo)

	return &Meta{client: client, netlifyCtx: netlifyCtx}
}
