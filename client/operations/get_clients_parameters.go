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

// NewGetClientsParams creates a new GetClientsParams object
// with the default values initialized.
func NewGetClientsParams() *GetClientsParams {
	var ()
	return &GetClientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClientsParamsWithTimeout creates a new GetClientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClientsParamsWithTimeout(timeout time.Duration) *GetClientsParams {
	var ()
	return &GetClientsParams{

		timeout: timeout,
	}
}

// NewGetClientsParamsWithContext creates a new GetClientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClientsParamsWithContext(ctx context.Context) *GetClientsParams {
	var ()
	return &GetClientsParams{

		Context: ctx,
	}
}

// NewGetClientsParamsWithHTTPClient creates a new GetClientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClientsParamsWithHTTPClient(client *http.Client) *GetClientsParams {
	var ()
	return &GetClientsParams{
		HTTPClient: client,
	}
}

/*GetClientsParams contains all the parameters to send to the API endpoint
for the get clients operation typically these are written to a http.Request
*/
type GetClientsParams struct {

	/*AppType
	  A comma separated list of application types used to filter the returned clients (native, spa, regular_web, non_interactive, rms, box, cloudbees, concur, dropbox, mscrm, echosign, egnyte, newrelic, office365, salesforce, sentry, sharepoint, slack, springcm, zendesk, zoom)

	*/
	AppType *string
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
	/*IsFirstParty
	  Filter on whether or not a client is a first party client

	*/
	IsFirstParty *bool
	/*IsGlobal
	  Optionally filter on the global client parameter

	*/
	IsGlobal *bool
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

// WithTimeout adds the timeout to the get clients params
func (o *GetClientsParams) WithTimeout(timeout time.Duration) *GetClientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clients params
func (o *GetClientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clients params
func (o *GetClientsParams) WithContext(ctx context.Context) *GetClientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clients params
func (o *GetClientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clients params
func (o *GetClientsParams) WithHTTPClient(client *http.Client) *GetClientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clients params
func (o *GetClientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppType adds the appType to the get clients params
func (o *GetClientsParams) WithAppType(appType *string) *GetClientsParams {
	o.SetAppType(appType)
	return o
}

// SetAppType adds the appType to the get clients params
func (o *GetClientsParams) SetAppType(appType *string) {
	o.AppType = appType
}

// WithFields adds the fields to the get clients params
func (o *GetClientsParams) WithFields(fields *string) *GetClientsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get clients params
func (o *GetClientsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeFields adds the includeFields to the get clients params
func (o *GetClientsParams) WithIncludeFields(includeFields *bool) *GetClientsParams {
	o.SetIncludeFields(includeFields)
	return o
}

// SetIncludeFields adds the includeFields to the get clients params
func (o *GetClientsParams) SetIncludeFields(includeFields *bool) {
	o.IncludeFields = includeFields
}

// WithIncludeTotals adds the includeTotals to the get clients params
func (o *GetClientsParams) WithIncludeTotals(includeTotals *bool) *GetClientsParams {
	o.SetIncludeTotals(includeTotals)
	return o
}

// SetIncludeTotals adds the includeTotals to the get clients params
func (o *GetClientsParams) SetIncludeTotals(includeTotals *bool) {
	o.IncludeTotals = includeTotals
}

// WithIsFirstParty adds the isFirstParty to the get clients params
func (o *GetClientsParams) WithIsFirstParty(isFirstParty *bool) *GetClientsParams {
	o.SetIsFirstParty(isFirstParty)
	return o
}

// SetIsFirstParty adds the isFirstParty to the get clients params
func (o *GetClientsParams) SetIsFirstParty(isFirstParty *bool) {
	o.IsFirstParty = isFirstParty
}

// WithIsGlobal adds the isGlobal to the get clients params
func (o *GetClientsParams) WithIsGlobal(isGlobal *bool) *GetClientsParams {
	o.SetIsGlobal(isGlobal)
	return o
}

// SetIsGlobal adds the isGlobal to the get clients params
func (o *GetClientsParams) SetIsGlobal(isGlobal *bool) {
	o.IsGlobal = isGlobal
}

// WithPage adds the page to the get clients params
func (o *GetClientsParams) WithPage(page *int64) *GetClientsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get clients params
func (o *GetClientsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get clients params
func (o *GetClientsParams) WithPerPage(perPage *int64) *GetClientsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get clients params
func (o *GetClientsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetClientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppType != nil {

		// query param app_type
		var qrAppType string
		if o.AppType != nil {
			qrAppType = *o.AppType
		}
		qAppType := qrAppType
		if qAppType != "" {
			if err := r.SetQueryParam("app_type", qAppType); err != nil {
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

	if o.IsFirstParty != nil {

		// query param is_first_party
		var qrIsFirstParty bool
		if o.IsFirstParty != nil {
			qrIsFirstParty = *o.IsFirstParty
		}
		qIsFirstParty := swag.FormatBool(qrIsFirstParty)
		if qIsFirstParty != "" {
			if err := r.SetQueryParam("is_first_party", qIsFirstParty); err != nil {
				return err
			}
		}

	}

	if o.IsGlobal != nil {

		// query param is_global
		var qrIsGlobal bool
		if o.IsGlobal != nil {
			qrIsGlobal = *o.IsGlobal
		}
		qIsGlobal := swag.FormatBool(qrIsGlobal)
		if qIsGlobal != "" {
			if err := r.SetQueryParam("is_global", qIsGlobal); err != nil {
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