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

// NewGetClientGrantsParams creates a new GetClientGrantsParams object
// with the default values initialized.
func NewGetClientGrantsParams() *GetClientGrantsParams {
	var ()
	return &GetClientGrantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClientGrantsParamsWithTimeout creates a new GetClientGrantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClientGrantsParamsWithTimeout(timeout time.Duration) *GetClientGrantsParams {
	var ()
	return &GetClientGrantsParams{

		timeout: timeout,
	}
}

// NewGetClientGrantsParamsWithContext creates a new GetClientGrantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClientGrantsParamsWithContext(ctx context.Context) *GetClientGrantsParams {
	var ()
	return &GetClientGrantsParams{

		Context: ctx,
	}
}

// NewGetClientGrantsParamsWithHTTPClient creates a new GetClientGrantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClientGrantsParamsWithHTTPClient(client *http.Client) *GetClientGrantsParams {
	var ()
	return &GetClientGrantsParams{
		HTTPClient: client,
	}
}

/*GetClientGrantsParams contains all the parameters to send to the API endpoint
for the get client grants operation typically these are written to a http.Request
*/
type GetClientGrantsParams struct {

	/*Audience
	  URL Encoded audience of a client grant to filter

	*/
	Audience *string
	/*ClientID
	  id of a client to filter

	*/
	ClientID *string
	/*IncludeTotals
	  true if a query summary must be included in the result, false otherwise. Default <code>false</code>.

	*/
	IncludeTotals *bool
	/*Page
	  The page number. Zero based

	*/
	Page *int64
	/*PerPage
	  The amount of entries per page.

	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get client grants params
func (o *GetClientGrantsParams) WithTimeout(timeout time.Duration) *GetClientGrantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get client grants params
func (o *GetClientGrantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get client grants params
func (o *GetClientGrantsParams) WithContext(ctx context.Context) *GetClientGrantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get client grants params
func (o *GetClientGrantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get client grants params
func (o *GetClientGrantsParams) WithHTTPClient(client *http.Client) *GetClientGrantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get client grants params
func (o *GetClientGrantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudience adds the audience to the get client grants params
func (o *GetClientGrantsParams) WithAudience(audience *string) *GetClientGrantsParams {
	o.SetAudience(audience)
	return o
}

// SetAudience adds the audience to the get client grants params
func (o *GetClientGrantsParams) SetAudience(audience *string) {
	o.Audience = audience
}

// WithClientID adds the clientID to the get client grants params
func (o *GetClientGrantsParams) WithClientID(clientID *string) *GetClientGrantsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get client grants params
func (o *GetClientGrantsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithIncludeTotals adds the includeTotals to the get client grants params
func (o *GetClientGrantsParams) WithIncludeTotals(includeTotals *bool) *GetClientGrantsParams {
	o.SetIncludeTotals(includeTotals)
	return o
}

// SetIncludeTotals adds the includeTotals to the get client grants params
func (o *GetClientGrantsParams) SetIncludeTotals(includeTotals *bool) {
	o.IncludeTotals = includeTotals
}

// WithPage adds the page to the get client grants params
func (o *GetClientGrantsParams) WithPage(page *int64) *GetClientGrantsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get client grants params
func (o *GetClientGrantsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get client grants params
func (o *GetClientGrantsParams) WithPerPage(perPage *int64) *GetClientGrantsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get client grants params
func (o *GetClientGrantsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetClientGrantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audience != nil {

		// query param audience
		var qrAudience string
		if o.Audience != nil {
			qrAudience = *o.Audience
		}
		qAudience := qrAudience
		if qAudience != "" {
			if err := r.SetQueryParam("audience", qAudience); err != nil {
				return err
			}
		}

	}

	if o.ClientID != nil {

		// query param client_id
		var qrClientID string
		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {
			if err := r.SetQueryParam("client_id", qClientID); err != nil {
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
