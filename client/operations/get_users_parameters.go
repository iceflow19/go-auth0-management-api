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

// NewGetUsersParams creates a new GetUsersParams object
// with the default values initialized.
func NewGetUsersParams() *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	var ()
	return &GetUsersParams{

		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	var ()
	return &GetUsersParams{

		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	var ()
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/*GetUsersParams contains all the parameters to send to the API endpoint
for the get users operation typically these are written to a http.Request
*/
type GetUsersParams struct {

	/*Connection
	  Connection filter. Only applies when using <code>search_engine=v1</code>. To filter by connection with <code>search_engine=v2|v3</code>, use <code>q=identities.connection:"connection_name"</code>

	*/
	Connection *string
	/*Fields
	  A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields

	*/
	Fields *string
	/*IncludeFields
	  <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise. Defaults to <code>true</code>

	*/
	IncludeFields *bool
	/*IncludeTotals
	  true if a query summary must be included in the result

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
	  Query in <a target='_new' href ='http://www.lucenetutorial.com/lucene-query-syntax.html'>Lucene query string syntax</a>. Not all metadata fields are searchable when using <code>search_engine=v2</code>, for details see <a href='https://auth0.com/docs/users/search/v2/query-syntax#searchable-fields'>Searchable Fields (v2)</a>. When using <code>search_engine=v3</code>, some query types cannot be used on metadata fields, for details see <a href='https://auth0.com/docs/users/search/v3/query-syntax#searchable-fields'>Searchable Fields (v3)</a>.

	*/
	Q *string
	/*SearchEngine
	  The version of the search engine

	*/
	SearchEngine *string
	/*Sort
	  The field to use for sorting. Use <code>field:order</code> where order is <code>1</code> for ascending and <code>-1</code> for descending. For example <code>created_at:1</code>

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnection adds the connection to the get users params
func (o *GetUsersParams) WithConnection(connection *string) *GetUsersParams {
	o.SetConnection(connection)
	return o
}

// SetConnection adds the connection to the get users params
func (o *GetUsersParams) SetConnection(connection *string) {
	o.Connection = connection
}

// WithFields adds the fields to the get users params
func (o *GetUsersParams) WithFields(fields *string) *GetUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get users params
func (o *GetUsersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeFields adds the includeFields to the get users params
func (o *GetUsersParams) WithIncludeFields(includeFields *bool) *GetUsersParams {
	o.SetIncludeFields(includeFields)
	return o
}

// SetIncludeFields adds the includeFields to the get users params
func (o *GetUsersParams) SetIncludeFields(includeFields *bool) {
	o.IncludeFields = includeFields
}

// WithIncludeTotals adds the includeTotals to the get users params
func (o *GetUsersParams) WithIncludeTotals(includeTotals *bool) *GetUsersParams {
	o.SetIncludeTotals(includeTotals)
	return o
}

// SetIncludeTotals adds the includeTotals to the get users params
func (o *GetUsersParams) SetIncludeTotals(includeTotals *bool) {
	o.IncludeTotals = includeTotals
}

// WithPage adds the page to the get users params
func (o *GetUsersParams) WithPage(page *int64) *GetUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get users params
func (o *GetUsersParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get users params
func (o *GetUsersParams) WithPerPage(perPage *int64) *GetUsersParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get users params
func (o *GetUsersParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithQ adds the q to the get users params
func (o *GetUsersParams) WithQ(q *string) *GetUsersParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get users params
func (o *GetUsersParams) SetQ(q *string) {
	o.Q = q
}

// WithSearchEngine adds the searchEngine to the get users params
func (o *GetUsersParams) WithSearchEngine(searchEngine *string) *GetUsersParams {
	o.SetSearchEngine(searchEngine)
	return o
}

// SetSearchEngine adds the searchEngine to the get users params
func (o *GetUsersParams) SetSearchEngine(searchEngine *string) {
	o.SearchEngine = searchEngine
}

// WithSort adds the sort to the get users params
func (o *GetUsersParams) WithSort(sort *string) *GetUsersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get users params
func (o *GetUsersParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Connection != nil {

		// query param connection
		var qrConnection string
		if o.Connection != nil {
			qrConnection = *o.Connection
		}
		qConnection := qrConnection
		if qConnection != "" {
			if err := r.SetQueryParam("connection", qConnection); err != nil {
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

	if o.SearchEngine != nil {

		// query param search_engine
		var qrSearchEngine string
		if o.SearchEngine != nil {
			qrSearchEngine = *o.SearchEngine
		}
		qSearchEngine := qrSearchEngine
		if qSearchEngine != "" {
			if err := r.SetQueryParam("search_engine", qSearchEngine); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
