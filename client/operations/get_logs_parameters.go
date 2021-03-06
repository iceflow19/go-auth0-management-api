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

// NewGetLogsParams creates a new GetLogsParams object
// with the default values initialized.
func NewGetLogsParams() *GetLogsParams {
	var ()
	return &GetLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogsParamsWithTimeout creates a new GetLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogsParamsWithTimeout(timeout time.Duration) *GetLogsParams {
	var ()
	return &GetLogsParams{

		timeout: timeout,
	}
}

// NewGetLogsParamsWithContext creates a new GetLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogsParamsWithContext(ctx context.Context) *GetLogsParams {
	var ()
	return &GetLogsParams{

		Context: ctx,
	}
}

// NewGetLogsParamsWithHTTPClient creates a new GetLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogsParamsWithHTTPClient(client *http.Client) *GetLogsParams {
	var ()
	return &GetLogsParams{
		HTTPClient: client,
	}
}

/*GetLogsParams contains all the parameters to send to the API endpoint
for the get logs operation typically these are written to a http.Request
*/
type GetLogsParams struct {

	/*Fields
	  A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields

	*/
	Fields *string
	/*From
	  Log Event Id to start retrieving logs. You can limit the amount of logs using the <code>take</code> parameter.

	*/
	From *string
	/*IncludeFields
	  <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise. Defaults to <code>true</code>

	*/
	IncludeFields *bool
	/*IncludeTotals
	  true if a query summary must be included in the result, false otherwise. Default <code>false</code>.

	*/
	IncludeTotals *bool
	/*Page
	  The page number. Zero based

	*/
	Page *int64
	/*PerPage
	  The amount of entries per page. Default: <code>50</code>. Max value: <code>100</code>

	*/
	PerPage *int64
	/*Q
	  Query in <a target='_new' href ='http://www.lucenetutorial.com/lucene-query-syntax.html'>Lucene query string syntax</a>.

	*/
	Q *string
	/*Sort
	  The field to use for sorting. Use <code>field:order</code> where order is <code>1</code> for ascending and <code>-1</code> for descending. For example <code>date:-1</code>

	*/
	Sort *string
	/*Take
	  The total amount of entries to retrieve when using the <code>from</code> parameter. Default: <code>50</code>. Max value: <code>100</code>

	*/
	Take *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logs params
func (o *GetLogsParams) WithTimeout(timeout time.Duration) *GetLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logs params
func (o *GetLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logs params
func (o *GetLogsParams) WithContext(ctx context.Context) *GetLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logs params
func (o *GetLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) WithHTTPClient(client *http.Client) *GetLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logs params
func (o *GetLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get logs params
func (o *GetLogsParams) WithFields(fields *string) *GetLogsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get logs params
func (o *GetLogsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFrom adds the from to the get logs params
func (o *GetLogsParams) WithFrom(from *string) *GetLogsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get logs params
func (o *GetLogsParams) SetFrom(from *string) {
	o.From = from
}

// WithIncludeFields adds the includeFields to the get logs params
func (o *GetLogsParams) WithIncludeFields(includeFields *bool) *GetLogsParams {
	o.SetIncludeFields(includeFields)
	return o
}

// SetIncludeFields adds the includeFields to the get logs params
func (o *GetLogsParams) SetIncludeFields(includeFields *bool) {
	o.IncludeFields = includeFields
}

// WithIncludeTotals adds the includeTotals to the get logs params
func (o *GetLogsParams) WithIncludeTotals(includeTotals *bool) *GetLogsParams {
	o.SetIncludeTotals(includeTotals)
	return o
}

// SetIncludeTotals adds the includeTotals to the get logs params
func (o *GetLogsParams) SetIncludeTotals(includeTotals *bool) {
	o.IncludeTotals = includeTotals
}

// WithPage adds the page to the get logs params
func (o *GetLogsParams) WithPage(page *int64) *GetLogsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get logs params
func (o *GetLogsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get logs params
func (o *GetLogsParams) WithPerPage(perPage *int64) *GetLogsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get logs params
func (o *GetLogsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithQ adds the q to the get logs params
func (o *GetLogsParams) WithQ(q *string) *GetLogsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get logs params
func (o *GetLogsParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the get logs params
func (o *GetLogsParams) WithSort(sort *string) *GetLogsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get logs params
func (o *GetLogsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTake adds the take to the get logs params
func (o *GetLogsParams) WithTake(take *int64) *GetLogsParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the get logs params
func (o *GetLogsParams) SetTake(take *int64) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int64
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt64(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
