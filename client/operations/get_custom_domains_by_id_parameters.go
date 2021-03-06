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

// NewGetCustomDomainsByIDParams creates a new GetCustomDomainsByIDParams object
// with the default values initialized.
func NewGetCustomDomainsByIDParams() *GetCustomDomainsByIDParams {
	var ()
	return &GetCustomDomainsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomDomainsByIDParamsWithTimeout creates a new GetCustomDomainsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomDomainsByIDParamsWithTimeout(timeout time.Duration) *GetCustomDomainsByIDParams {
	var ()
	return &GetCustomDomainsByIDParams{

		timeout: timeout,
	}
}

// NewGetCustomDomainsByIDParamsWithContext creates a new GetCustomDomainsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomDomainsByIDParamsWithContext(ctx context.Context) *GetCustomDomainsByIDParams {
	var ()
	return &GetCustomDomainsByIDParams{

		Context: ctx,
	}
}

// NewGetCustomDomainsByIDParamsWithHTTPClient creates a new GetCustomDomainsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomDomainsByIDParamsWithHTTPClient(client *http.Client) *GetCustomDomainsByIDParams {
	var ()
	return &GetCustomDomainsByIDParams{
		HTTPClient: client,
	}
}

/*GetCustomDomainsByIDParams contains all the parameters to send to the API endpoint
for the get custom domains by id operation typically these are written to a http.Request
*/
type GetCustomDomainsByIDParams struct {

	/*ID
	  The id of the custom domain to retrieve

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) WithTimeout(timeout time.Duration) *GetCustomDomainsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) WithContext(ctx context.Context) *GetCustomDomainsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) WithHTTPClient(client *http.Client) *GetCustomDomainsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) WithID(id string) *GetCustomDomainsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get custom domains by id params
func (o *GetCustomDomainsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomDomainsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
