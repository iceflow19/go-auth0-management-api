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

// NewGetRulesParams creates a new GetRulesParams object
// with the default values initialized.
func NewGetRulesParams() *GetRulesParams {
	var ()
	return &GetRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRulesParamsWithTimeout creates a new GetRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRulesParamsWithTimeout(timeout time.Duration) *GetRulesParams {
	var ()
	return &GetRulesParams{

		timeout: timeout,
	}
}

// NewGetRulesParamsWithContext creates a new GetRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRulesParamsWithContext(ctx context.Context) *GetRulesParams {
	var ()
	return &GetRulesParams{

		Context: ctx,
	}
}

// NewGetRulesParamsWithHTTPClient creates a new GetRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRulesParamsWithHTTPClient(client *http.Client) *GetRulesParams {
	var ()
	return &GetRulesParams{
		HTTPClient: client,
	}
}

/*GetRulesParams contains all the parameters to send to the API endpoint
for the get rules operation typically these are written to a http.Request
*/
type GetRulesParams struct {

	/*Enabled
	  If provided retrieves rules that match the value, otherwise all rules are retrieved

	*/
	Enabled *bool
	/*Fields
	  A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields

	*/
	Fields *string
	/*IncludeFields
	  <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise (defaults to <code>true</code>)

	*/
	IncludeFields *bool
	/*IncludeTotals
	  true if a query summary must be included in the result

	*/
	IncludeTotals *bool
	/*Page
	  The page number. Zero based.

	*/
	Page *int64
	/*PerPage
	  The amount of entries per page. Default: <code>50</code>. Max value: <code>100</code>. If not present, pagination will be disabled

	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rules params
func (o *GetRulesParams) WithTimeout(timeout time.Duration) *GetRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rules params
func (o *GetRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rules params
func (o *GetRulesParams) WithContext(ctx context.Context) *GetRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rules params
func (o *GetRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rules params
func (o *GetRulesParams) WithHTTPClient(client *http.Client) *GetRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rules params
func (o *GetRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the get rules params
func (o *GetRulesParams) WithEnabled(enabled *bool) *GetRulesParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the get rules params
func (o *GetRulesParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithFields adds the fields to the get rules params
func (o *GetRulesParams) WithFields(fields *string) *GetRulesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get rules params
func (o *GetRulesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeFields adds the includeFields to the get rules params
func (o *GetRulesParams) WithIncludeFields(includeFields *bool) *GetRulesParams {
	o.SetIncludeFields(includeFields)
	return o
}

// SetIncludeFields adds the includeFields to the get rules params
func (o *GetRulesParams) SetIncludeFields(includeFields *bool) {
	o.IncludeFields = includeFields
}

// WithIncludeTotals adds the includeTotals to the get rules params
func (o *GetRulesParams) WithIncludeTotals(includeTotals *bool) *GetRulesParams {
	o.SetIncludeTotals(includeTotals)
	return o
}

// SetIncludeTotals adds the includeTotals to the get rules params
func (o *GetRulesParams) SetIncludeTotals(includeTotals *bool) {
	o.IncludeTotals = includeTotals
}

// WithPage adds the page to the get rules params
func (o *GetRulesParams) WithPage(page *int64) *GetRulesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get rules params
func (o *GetRulesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get rules params
func (o *GetRulesParams) WithPerPage(perPage *int64) *GetRulesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get rules params
func (o *GetRulesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool
		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {
			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}

	}

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

	if o.IncludeTotals != nil {

		// query param include_totals
		var qrIncludeTotals bool
		if o.IncludeTotals != nil {
			qrIncludeTotals = *o.IncludeTotals
		}
		qIncludeTotals := swag.FormatBool(qrIncludeTotals)
		if qIncludeTotals != "" {
			if err := r.SetQueryParam("include_totals", qIncludeTotals); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
