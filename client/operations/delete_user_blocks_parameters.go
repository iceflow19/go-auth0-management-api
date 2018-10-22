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

// NewDeleteUserBlocksParams creates a new DeleteUserBlocksParams object
// with the default values initialized.
func NewDeleteUserBlocksParams() *DeleteUserBlocksParams {
	var ()
	return &DeleteUserBlocksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserBlocksParamsWithTimeout creates a new DeleteUserBlocksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserBlocksParamsWithTimeout(timeout time.Duration) *DeleteUserBlocksParams {
	var ()
	return &DeleteUserBlocksParams{

		timeout: timeout,
	}
}

// NewDeleteUserBlocksParamsWithContext creates a new DeleteUserBlocksParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserBlocksParamsWithContext(ctx context.Context) *DeleteUserBlocksParams {
	var ()
	return &DeleteUserBlocksParams{

		Context: ctx,
	}
}

// NewDeleteUserBlocksParamsWithHTTPClient creates a new DeleteUserBlocksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserBlocksParamsWithHTTPClient(client *http.Client) *DeleteUserBlocksParams {
	var ()
	return &DeleteUserBlocksParams{
		HTTPClient: client,
	}
}

/*DeleteUserBlocksParams contains all the parameters to send to the API endpoint
for the delete user blocks operation typically these are written to a http.Request
*/
type DeleteUserBlocksParams struct {

	/*Identifier
	  Should be any of: username, phone_number, email.

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user blocks params
func (o *DeleteUserBlocksParams) WithTimeout(timeout time.Duration) *DeleteUserBlocksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user blocks params
func (o *DeleteUserBlocksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user blocks params
func (o *DeleteUserBlocksParams) WithContext(ctx context.Context) *DeleteUserBlocksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user blocks params
func (o *DeleteUserBlocksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user blocks params
func (o *DeleteUserBlocksParams) WithHTTPClient(client *http.Client) *DeleteUserBlocksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user blocks params
func (o *DeleteUserBlocksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the delete user blocks params
func (o *DeleteUserBlocksParams) WithIdentifier(identifier string) *DeleteUserBlocksParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the delete user blocks params
func (o *DeleteUserBlocksParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserBlocksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param identifier
	qrIdentifier := o.Identifier
	qIdentifier := qrIdentifier
	if qIdentifier != "" {
		if err := r.SetQueryParam("identifier", qIdentifier); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}