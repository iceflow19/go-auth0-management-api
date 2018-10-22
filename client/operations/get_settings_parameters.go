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

// NewGetSettingsParams creates a new GetSettingsParams object
// with the default values initialized.
func NewGetSettingsParams() *GetSettingsParams {
	var ()
	return &GetSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsParamsWithTimeout creates a new GetSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSettingsParamsWithTimeout(timeout time.Duration) *GetSettingsParams {
	var ()
	return &GetSettingsParams{

		timeout: timeout,
	}
}

// NewGetSettingsParamsWithContext creates a new GetSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSettingsParamsWithContext(ctx context.Context) *GetSettingsParams {
	var ()
	return &GetSettingsParams{

		Context: ctx,
	}
}

// NewGetSettingsParamsWithHTTPClient creates a new GetSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSettingsParamsWithHTTPClient(client *http.Client) *GetSettingsParams {
	var ()
	return &GetSettingsParams{
		HTTPClient: client,
	}
}

/*GetSettingsParams contains all the parameters to send to the API endpoint
for the get settings operation typically these are written to a http.Request
*/
type GetSettingsParams struct {

	/*Fields
	  A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields

	*/
	Fields *string
	/*IncludeFields
	  <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise (defaults to <code>true</code>)

	*/
	IncludeFields *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get settings params
func (o *GetSettingsParams) WithTimeout(timeout time.Duration) *GetSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings params
func (o *GetSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings params
func (o *GetSettingsParams) WithContext(ctx context.Context) *GetSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings params
func (o *GetSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings params
func (o *GetSettingsParams) WithHTTPClient(client *http.Client) *GetSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings params
func (o *GetSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get settings params
func (o *GetSettingsParams) WithFields(fields *string) *GetSettingsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get settings params
func (o *GetSettingsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeFields adds the includeFields to the get settings params
func (o *GetSettingsParams) WithIncludeFields(includeFields *bool) *GetSettingsParams {
	o.SetIncludeFields(includeFields)
	return o
}

// SetIncludeFields adds the includeFields to the get settings params
func (o *GetSettingsParams) SetIncludeFields(includeFields *bool) {
	o.IncludeFields = includeFields
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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