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

// NewPostRotateSecretParams creates a new PostRotateSecretParams object
// with the default values initialized.
func NewPostRotateSecretParams() *PostRotateSecretParams {
	var ()
	return &PostRotateSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRotateSecretParamsWithTimeout creates a new PostRotateSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRotateSecretParamsWithTimeout(timeout time.Duration) *PostRotateSecretParams {
	var ()
	return &PostRotateSecretParams{

		timeout: timeout,
	}
}

// NewPostRotateSecretParamsWithContext creates a new PostRotateSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRotateSecretParamsWithContext(ctx context.Context) *PostRotateSecretParams {
	var ()
	return &PostRotateSecretParams{

		Context: ctx,
	}
}

// NewPostRotateSecretParamsWithHTTPClient creates a new PostRotateSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRotateSecretParamsWithHTTPClient(client *http.Client) *PostRotateSecretParams {
	var ()
	return &PostRotateSecretParams{
		HTTPClient: client,
	}
}

/*PostRotateSecretParams contains all the parameters to send to the API endpoint
for the post rotate secret operation typically these are written to a http.Request
*/
type PostRotateSecretParams struct {

	/*ID
	  The id of the client to rotate secrets

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post rotate secret params
func (o *PostRotateSecretParams) WithTimeout(timeout time.Duration) *PostRotateSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post rotate secret params
func (o *PostRotateSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post rotate secret params
func (o *PostRotateSecretParams) WithContext(ctx context.Context) *PostRotateSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post rotate secret params
func (o *PostRotateSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post rotate secret params
func (o *PostRotateSecretParams) WithHTTPClient(client *http.Client) *PostRotateSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post rotate secret params
func (o *PostRotateSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post rotate secret params
func (o *PostRotateSecretParams) WithID(id string) *PostRotateSecretParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post rotate secret params
func (o *PostRotateSecretParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRotateSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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