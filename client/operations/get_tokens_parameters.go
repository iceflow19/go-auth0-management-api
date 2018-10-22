// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewGetTokensParams creates a new GetTokensParams object
// with the default values initialized.
func NewGetTokensParams() *GetTokensParams {
	var ()
	return &GetTokensParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTokensParamsWithTimeout creates a new GetTokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTokensParamsWithTimeout(timeout time.Duration) *GetTokensParams {
	var ()
	return &GetTokensParams{

		timeout: timeout,
	}
}

// NewGetTokensParamsWithContext creates a new GetTokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTokensParamsWithContext(ctx context.Context) *GetTokensParams {
	var ()
	return &GetTokensParams{

		Context: ctx,
	}
}

// NewGetTokensParamsWithHTTPClient creates a new GetTokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTokensParamsWithHTTPClient(client *http.Client) *GetTokensParams {
	var ()
	return &GetTokensParams{
		HTTPClient: client,
	}
}

/*GetTokensParams contains all the parameters to send to the API endpoint
for the get tokens operation typically these are written to a http.Request
*/
type GetTokensParams struct {

	/*Aud
	  The JWT's aud claim. The client_id of the client for which it was issued

	*/
	Aud *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tokens params
func (o *GetTokensParams) WithTimeout(timeout time.Duration) *GetTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tokens params
func (o *GetTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tokens params
func (o *GetTokensParams) WithContext(ctx context.Context) *GetTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tokens params
func (o *GetTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tokens params
func (o *GetTokensParams) WithHTTPClient(client *http.Client) *GetTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tokens params
func (o *GetTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAud adds the aud to the get tokens params
func (o *GetTokensParams) WithAud(aud *string) *GetTokensParams {
	o.SetAud(aud)
	return o
}

// SetAud adds the aud to the get tokens params
func (o *GetTokensParams) SetAud(aud *string) {
	o.Aud = aud
}

// WriteToRequest writes these params to a swagger request
func (o *GetTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Aud != nil {

		// query param aud
		var qrAud string
		if o.Aud != nil {
			qrAud = *o.Aud
		}
		qAud := qrAud
		if qAud != "" {
			if err := r.SetQueryParam("aud", qAud); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
