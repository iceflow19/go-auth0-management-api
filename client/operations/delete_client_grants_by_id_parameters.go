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

// NewDeleteClientGrantsByIDParams creates a new DeleteClientGrantsByIDParams object
// with the default values initialized.
func NewDeleteClientGrantsByIDParams() *DeleteClientGrantsByIDParams {
	var ()
	return &DeleteClientGrantsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClientGrantsByIDParamsWithTimeout creates a new DeleteClientGrantsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClientGrantsByIDParamsWithTimeout(timeout time.Duration) *DeleteClientGrantsByIDParams {
	var ()
	return &DeleteClientGrantsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteClientGrantsByIDParamsWithContext creates a new DeleteClientGrantsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClientGrantsByIDParamsWithContext(ctx context.Context) *DeleteClientGrantsByIDParams {
	var ()
	return &DeleteClientGrantsByIDParams{

		Context: ctx,
	}
}

// NewDeleteClientGrantsByIDParamsWithHTTPClient creates a new DeleteClientGrantsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClientGrantsByIDParamsWithHTTPClient(client *http.Client) *DeleteClientGrantsByIDParams {
	var ()
	return &DeleteClientGrantsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteClientGrantsByIDParams contains all the parameters to send to the API endpoint
for the delete client grants by id operation typically these are written to a http.Request
*/
type DeleteClientGrantsByIDParams struct {

	/*ID
	  The id of the client grant to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) WithTimeout(timeout time.Duration) *DeleteClientGrantsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) WithContext(ctx context.Context) *DeleteClientGrantsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) WithHTTPClient(client *http.Client) *DeleteClientGrantsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) WithID(id string) *DeleteClientGrantsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete client grants by id params
func (o *DeleteClientGrantsByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClientGrantsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
