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

// NewGetDailyParams creates a new GetDailyParams object
// with the default values initialized.
func NewGetDailyParams() *GetDailyParams {
	var ()
	return &GetDailyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDailyParamsWithTimeout creates a new GetDailyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDailyParamsWithTimeout(timeout time.Duration) *GetDailyParams {
	var ()
	return &GetDailyParams{

		timeout: timeout,
	}
}

// NewGetDailyParamsWithContext creates a new GetDailyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDailyParamsWithContext(ctx context.Context) *GetDailyParams {
	var ()
	return &GetDailyParams{

		Context: ctx,
	}
}

// NewGetDailyParamsWithHTTPClient creates a new GetDailyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDailyParamsWithHTTPClient(client *http.Client) *GetDailyParams {
	var ()
	return &GetDailyParams{
		HTTPClient: client,
	}
}

/*GetDailyParams contains all the parameters to send to the API endpoint
for the get daily operation typically these are written to a http.Request
*/
type GetDailyParams struct {

	/*From
	  The first day of the period (inclusive) in YYYYMMDD format.

	*/
	From *string
	/*To
	  The last day of the period (inclusive) in YYYYMMDD format.

	*/
	To *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get daily params
func (o *GetDailyParams) WithTimeout(timeout time.Duration) *GetDailyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get daily params
func (o *GetDailyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get daily params
func (o *GetDailyParams) WithContext(ctx context.Context) *GetDailyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get daily params
func (o *GetDailyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get daily params
func (o *GetDailyParams) WithHTTPClient(client *http.Client) *GetDailyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get daily params
func (o *GetDailyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get daily params
func (o *GetDailyParams) WithFrom(from *string) *GetDailyParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get daily params
func (o *GetDailyParams) SetFrom(from *string) {
	o.From = from
}

// WithTo adds the to to the get daily params
func (o *GetDailyParams) WithTo(to *string) *GetDailyParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get daily params
func (o *GetDailyParams) SetTo(to *string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetDailyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.To != nil {

		// query param to
		var qrTo string
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
