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

// NewGetEnrollmentsParams creates a new GetEnrollmentsParams object
// with the default values initialized.
func NewGetEnrollmentsParams() *GetEnrollmentsParams {
	var ()
	return &GetEnrollmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnrollmentsParamsWithTimeout creates a new GetEnrollmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnrollmentsParamsWithTimeout(timeout time.Duration) *GetEnrollmentsParams {
	var ()
	return &GetEnrollmentsParams{

		timeout: timeout,
	}
}

// NewGetEnrollmentsParamsWithContext creates a new GetEnrollmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnrollmentsParamsWithContext(ctx context.Context) *GetEnrollmentsParams {
	var ()
	return &GetEnrollmentsParams{

		Context: ctx,
	}
}

// NewGetEnrollmentsParamsWithHTTPClient creates a new GetEnrollmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnrollmentsParamsWithHTTPClient(client *http.Client) *GetEnrollmentsParams {
	var ()
	return &GetEnrollmentsParams{
		HTTPClient: client,
	}
}

/*GetEnrollmentsParams contains all the parameters to send to the API endpoint
for the get enrollments operation typically these are written to a http.Request
*/
type GetEnrollmentsParams struct {

	/*ID
	  The user_id of the user to retrieve

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get enrollments params
func (o *GetEnrollmentsParams) WithTimeout(timeout time.Duration) *GetEnrollmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get enrollments params
func (o *GetEnrollmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get enrollments params
func (o *GetEnrollmentsParams) WithContext(ctx context.Context) *GetEnrollmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get enrollments params
func (o *GetEnrollmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get enrollments params
func (o *GetEnrollmentsParams) WithHTTPClient(client *http.Client) *GetEnrollmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get enrollments params
func (o *GetEnrollmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get enrollments params
func (o *GetEnrollmentsParams) WithID(id string) *GetEnrollmentsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get enrollments params
func (o *GetEnrollmentsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnrollmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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