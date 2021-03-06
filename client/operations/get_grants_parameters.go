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

// NewGetGrantsParams creates a new GetGrantsParams object
// with the default values initialized.
func NewGetGrantsParams() *GetGrantsParams {
	var ()
	return &GetGrantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGrantsParamsWithTimeout creates a new GetGrantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGrantsParamsWithTimeout(timeout time.Duration) *GetGrantsParams {
	var ()
	return &GetGrantsParams{

		timeout: timeout,
	}
}

// NewGetGrantsParamsWithContext creates a new GetGrantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGrantsParamsWithContext(ctx context.Context) *GetGrantsParams {
	var ()
	return &GetGrantsParams{

		Context: ctx,
	}
}

// NewGetGrantsParamsWithHTTPClient creates a new GetGrantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGrantsParamsWithHTTPClient(client *http.Client) *GetGrantsParams {
	var ()
	return &GetGrantsParams{
		HTTPClient: client,
	}
}

/*GetGrantsParams contains all the parameters to send to the API endpoint
for the get grants operation typically these are written to a http.Request
*/
type GetGrantsParams struct {

	/*Audience
	  The audience of the grants to retrieve

	*/
	Audience *string
	/*ClientID
	  The client_id of the grants to retrieve

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
	  The amount of entries per page. Default: no paging is used, all grants are returned

	*/
	PerPage *int64
	/*UserID
	  The user_id of the grants to retrieve

	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get grants params
func (o *GetGrantsParams) WithTimeout(timeout time.Duration) *GetGrantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get grants params
func (o *GetGrantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get grants params
func (o *GetGrantsParams) WithContext(ctx context.Context) *GetGrantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get grants params
func (o *GetGrantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get grants params
func (o *GetGrantsParams) WithHTTPClient(client *http.Client) *GetGrantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get grants params
func (o *GetGrantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudience adds the audience to the get grants params
func (o *GetGrantsParams) WithAudience(audience *string) *GetGrantsParams {
	o.SetAudience(audience)
	return o
}

// SetAudience adds the audience to the get grants params
func (o *GetGrantsParams) SetAudience(audience *string) {
	o.Audience = audience
}

// WithClientID adds the clientID to the get grants params
func (o *GetGrantsParams) WithClientID(clientID *string) *GetGrantsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get grants params
func (o *GetGrantsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithIncludeTotals adds the includeTotals to the get grants params
func (o *GetGrantsParams) WithIncludeTotals(includeTotals *bool) *GetGrantsParams {
	o.SetIncludeTotals(includeTotals)
	return o
}

// SetIncludeTotals adds the includeTotals to the get grants params
func (o *GetGrantsParams) SetIncludeTotals(includeTotals *bool) {
	o.IncludeTotals = includeTotals
}

// WithPage adds the page to the get grants params
func (o *GetGrantsParams) WithPage(page *int64) *GetGrantsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get grants params
func (o *GetGrantsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get grants params
func (o *GetGrantsParams) WithPerPage(perPage *int64) *GetGrantsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get grants params
func (o *GetGrantsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithUserID adds the userID to the get grants params
func (o *GetGrantsParams) WithUserID(userID *string) *GetGrantsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get grants params
func (o *GetGrantsParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGrantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
