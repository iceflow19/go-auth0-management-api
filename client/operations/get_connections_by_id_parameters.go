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
	"github.com/go-openapi/swag"
	"golang.org/x/net/context"
)

// NewGetConnectionsByIDParams creates a new GetConnectionsByIDParams object
// with the default values initialized.
func NewGetConnectionsByIDParams() *GetConnectionsByIDParams {
	var ()
	return &GetConnectionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConnectionsByIDParamsWithTimeout creates a new GetConnectionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConnectionsByIDParamsWithTimeout(timeout time.Duration) *GetConnectionsByIDParams {
	var ()
	return &GetConnectionsByIDParams{

		timeout: timeout,
	}
}

// NewGetConnectionsByIDParamsWithContext creates a new GetConnectionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConnectionsByIDParamsWithContext(ctx context.Context) *GetConnectionsByIDParams {
	var ()
	return &GetConnectionsByIDParams{

		Context: ctx,
	}
}

// NewGetConnectionsByIDParamsWithHTTPClient creates a new GetConnectionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConnectionsByIDParamsWithHTTPClient(client *http.Client) *GetConnectionsByIDParams {
	var ()
	return &GetConnectionsByIDParams{
		HTTPClient: client,
	}
}

/*GetConnectionsByIDParams contains all the parameters to send to the API endpoint
for the get connections by id operation typically these are written to a http.Request
*/
type GetConnectionsByIDParams struct {

	/*Fields
	  A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields

	*/
	Fields *string
	/*ID
	  The id of the connection to retrieve

	*/
	ID string
	/*IncludeFields
	  <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise (defaults to <code>true</code>)

	*/
	IncludeFields *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get connections by id params
func (o *GetConnectionsByIDParams) WithTimeout(timeout time.Duration) *GetConnectionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get connections by id params
func (o *GetConnectionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get connections by id params
func (o *GetConnectionsByIDParams) WithContext(ctx context.Context) *GetConnectionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get connections by id params
func (o *GetConnectionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get connections by id params
func (o *GetConnectionsByIDParams) WithHTTPClient(client *http.Client) *GetConnectionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get connections by id params
func (o *GetConnectionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get connections by id params
func (o *GetConnectionsByIDParams) WithFields(fields *string) *GetConnectionsByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get connections by id params
func (o *GetConnectionsByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get connections by id params
func (o *GetConnectionsByIDParams) WithID(id string) *GetConnectionsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get connections by id params
func (o *GetConnectionsByIDParams) SetID(id string) {
	o.ID = id
}

// WithIncludeFields adds the includeFields to the get connections by id params
func (o *GetConnectionsByIDParams) WithIncludeFields(includeFields *bool) *GetConnectionsByIDParams {
	o.SetIncludeFields(includeFields)
	return o
}

// SetIncludeFields adds the includeFields to the get connections by id params
func (o *GetConnectionsByIDParams) SetIncludeFields(includeFields *bool) {
	o.IncludeFields = includeFields
}

// WriteToRequest writes these params to a swagger request
func (o *GetConnectionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IncludeFields != nil {

		// query param include_fields
		var qrIncludeFields bool
		if o.IncludeFields != nil {
			qrIncludeFields = *o.IncludeFields
		}
		qIncludeFields := swag.FormatBool(qrIncludeFields)
		if qIncludeFields != "" {
			if err := r.SetQueryParam("include_fields", qIncludeFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}