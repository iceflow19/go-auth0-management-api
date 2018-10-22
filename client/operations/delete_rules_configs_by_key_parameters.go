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

// NewDeleteRulesConfigsByKeyParams creates a new DeleteRulesConfigsByKeyParams object
// with the default values initialized.
func NewDeleteRulesConfigsByKeyParams() *DeleteRulesConfigsByKeyParams {
	var ()
	return &DeleteRulesConfigsByKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRulesConfigsByKeyParamsWithTimeout creates a new DeleteRulesConfigsByKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRulesConfigsByKeyParamsWithTimeout(timeout time.Duration) *DeleteRulesConfigsByKeyParams {
	var ()
	return &DeleteRulesConfigsByKeyParams{

		timeout: timeout,
	}
}

// NewDeleteRulesConfigsByKeyParamsWithContext creates a new DeleteRulesConfigsByKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRulesConfigsByKeyParamsWithContext(ctx context.Context) *DeleteRulesConfigsByKeyParams {
	var ()
	return &DeleteRulesConfigsByKeyParams{

		Context: ctx,
	}
}

// NewDeleteRulesConfigsByKeyParamsWithHTTPClient creates a new DeleteRulesConfigsByKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRulesConfigsByKeyParamsWithHTTPClient(client *http.Client) *DeleteRulesConfigsByKeyParams {
	var ()
	return &DeleteRulesConfigsByKeyParams{
		HTTPClient: client,
	}
}

/*DeleteRulesConfigsByKeyParams contains all the parameters to send to the API endpoint
for the delete rules configs by key operation typically these are written to a http.Request
*/
type DeleteRulesConfigsByKeyParams struct {

	/*Key
	  The key of the rules config to remove (Max length: <code>127</code>)

	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) WithTimeout(timeout time.Duration) *DeleteRulesConfigsByKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) WithContext(ctx context.Context) *DeleteRulesConfigsByKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) WithHTTPClient(client *http.Client) *DeleteRulesConfigsByKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) WithKey(key string) *DeleteRulesConfigsByKeyParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete rules configs by key params
func (o *DeleteRulesConfigsByKeyParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRulesConfigsByKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
