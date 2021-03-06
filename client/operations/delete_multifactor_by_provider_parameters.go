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

// NewDeleteMultifactorByProviderParams creates a new DeleteMultifactorByProviderParams object
// with the default values initialized.
func NewDeleteMultifactorByProviderParams() *DeleteMultifactorByProviderParams {
	var ()
	return &DeleteMultifactorByProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMultifactorByProviderParamsWithTimeout creates a new DeleteMultifactorByProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMultifactorByProviderParamsWithTimeout(timeout time.Duration) *DeleteMultifactorByProviderParams {
	var ()
	return &DeleteMultifactorByProviderParams{

		timeout: timeout,
	}
}

// NewDeleteMultifactorByProviderParamsWithContext creates a new DeleteMultifactorByProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMultifactorByProviderParamsWithContext(ctx context.Context) *DeleteMultifactorByProviderParams {
	var ()
	return &DeleteMultifactorByProviderParams{

		Context: ctx,
	}
}

// NewDeleteMultifactorByProviderParamsWithHTTPClient creates a new DeleteMultifactorByProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMultifactorByProviderParamsWithHTTPClient(client *http.Client) *DeleteMultifactorByProviderParams {
	var ()
	return &DeleteMultifactorByProviderParams{
		HTTPClient: client,
	}
}

/*DeleteMultifactorByProviderParams contains all the parameters to send to the API endpoint
for the delete multifactor by provider operation typically these are written to a http.Request
*/
type DeleteMultifactorByProviderParams struct {

	/*ID
	  The user_id of the user to delete

	*/
	ID string
	/*Provider
	  The multifactor provider. Supported values 'duo' or 'google-authenticator'

	*/
	Provider string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) WithTimeout(timeout time.Duration) *DeleteMultifactorByProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) WithContext(ctx context.Context) *DeleteMultifactorByProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) WithHTTPClient(client *http.Client) *DeleteMultifactorByProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) WithID(id string) *DeleteMultifactorByProviderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) SetID(id string) {
	o.ID = id
}

// WithProvider adds the provider to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) WithProvider(provider string) *DeleteMultifactorByProviderParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the delete multifactor by provider params
func (o *DeleteMultifactorByProviderParams) SetProvider(provider string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMultifactorByProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
