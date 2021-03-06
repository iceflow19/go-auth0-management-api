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

// NewGetActiveUsersParams creates a new GetActiveUsersParams object
// with the default values initialized.
func NewGetActiveUsersParams() *GetActiveUsersParams {

	return &GetActiveUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetActiveUsersParamsWithTimeout creates a new GetActiveUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetActiveUsersParamsWithTimeout(timeout time.Duration) *GetActiveUsersParams {

	return &GetActiveUsersParams{

		timeout: timeout,
	}
}

// NewGetActiveUsersParamsWithContext creates a new GetActiveUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetActiveUsersParamsWithContext(ctx context.Context) *GetActiveUsersParams {

	return &GetActiveUsersParams{

		Context: ctx,
	}
}

// NewGetActiveUsersParamsWithHTTPClient creates a new GetActiveUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetActiveUsersParamsWithHTTPClient(client *http.Client) *GetActiveUsersParams {

	return &GetActiveUsersParams{
		HTTPClient: client,
	}
}

/*GetActiveUsersParams contains all the parameters to send to the API endpoint
for the get active users operation typically these are written to a http.Request
*/
type GetActiveUsersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get active users params
func (o *GetActiveUsersParams) WithTimeout(timeout time.Duration) *GetActiveUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get active users params
func (o *GetActiveUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get active users params
func (o *GetActiveUsersParams) WithContext(ctx context.Context) *GetActiveUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get active users params
func (o *GetActiveUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get active users params
func (o *GetActiveUsersParams) WithHTTPClient(client *http.Client) *GetActiveUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get active users params
func (o *GetActiveUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetActiveUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
