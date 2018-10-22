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

// NewGetTemplatesParams creates a new GetTemplatesParams object
// with the default values initialized.
func NewGetTemplatesParams() *GetTemplatesParams {

	return &GetTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplatesParamsWithTimeout creates a new GetTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTemplatesParamsWithTimeout(timeout time.Duration) *GetTemplatesParams {

	return &GetTemplatesParams{

		timeout: timeout,
	}
}

// NewGetTemplatesParamsWithContext creates a new GetTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTemplatesParamsWithContext(ctx context.Context) *GetTemplatesParams {

	return &GetTemplatesParams{

		Context: ctx,
	}
}

// NewGetTemplatesParamsWithHTTPClient creates a new GetTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTemplatesParamsWithHTTPClient(client *http.Client) *GetTemplatesParams {

	return &GetTemplatesParams{
		HTTPClient: client,
	}
}

/*GetTemplatesParams contains all the parameters to send to the API endpoint
for the get templates operation typically these are written to a http.Request
*/
type GetTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get templates params
func (o *GetTemplatesParams) WithTimeout(timeout time.Duration) *GetTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get templates params
func (o *GetTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get templates params
func (o *GetTemplatesParams) WithContext(ctx context.Context) *GetTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get templates params
func (o *GetTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get templates params
func (o *GetTemplatesParams) WithHTTPClient(client *http.Client) *GetTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get templates params
func (o *GetTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}