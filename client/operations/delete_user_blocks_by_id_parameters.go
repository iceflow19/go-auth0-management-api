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

// NewDeleteUserBlocksByIDParams creates a new DeleteUserBlocksByIDParams object
// with the default values initialized.
func NewDeleteUserBlocksByIDParams() *DeleteUserBlocksByIDParams {
	var ()
	return &DeleteUserBlocksByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserBlocksByIDParamsWithTimeout creates a new DeleteUserBlocksByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserBlocksByIDParamsWithTimeout(timeout time.Duration) *DeleteUserBlocksByIDParams {
	var ()
	return &DeleteUserBlocksByIDParams{

		timeout: timeout,
	}
}

// NewDeleteUserBlocksByIDParamsWithContext creates a new DeleteUserBlocksByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserBlocksByIDParamsWithContext(ctx context.Context) *DeleteUserBlocksByIDParams {
	var ()
	return &DeleteUserBlocksByIDParams{

		Context: ctx,
	}
}

// NewDeleteUserBlocksByIDParamsWithHTTPClient creates a new DeleteUserBlocksByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserBlocksByIDParamsWithHTTPClient(client *http.Client) *DeleteUserBlocksByIDParams {
	var ()
	return &DeleteUserBlocksByIDParams{
		HTTPClient: client,
	}
}

/*DeleteUserBlocksByIDParams contains all the parameters to send to the API endpoint
for the delete user blocks by id operation typically these are written to a http.Request
*/
type DeleteUserBlocksByIDParams struct {

	/*ID
	  The user_id of the user to update.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) WithTimeout(timeout time.Duration) *DeleteUserBlocksByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) WithContext(ctx context.Context) *DeleteUserBlocksByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) WithHTTPClient(client *http.Client) *DeleteUserBlocksByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) WithID(id string) *DeleteUserBlocksByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete user blocks by id params
func (o *DeleteUserBlocksByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserBlocksByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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