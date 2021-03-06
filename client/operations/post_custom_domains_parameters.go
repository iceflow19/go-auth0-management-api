// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	models "bitbucket.org/troyko/go-auth0-management-api/models"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewPostCustomDomainsParams creates a new PostCustomDomainsParams object
// with the default values initialized.
func NewPostCustomDomainsParams() *PostCustomDomainsParams {
	var ()
	return &PostCustomDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomDomainsParamsWithTimeout creates a new PostCustomDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomDomainsParamsWithTimeout(timeout time.Duration) *PostCustomDomainsParams {
	var ()
	return &PostCustomDomainsParams{

		timeout: timeout,
	}
}

// NewPostCustomDomainsParamsWithContext creates a new PostCustomDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomDomainsParamsWithContext(ctx context.Context) *PostCustomDomainsParams {
	var ()
	return &PostCustomDomainsParams{

		Context: ctx,
	}
}

// NewPostCustomDomainsParamsWithHTTPClient creates a new PostCustomDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomDomainsParamsWithHTTPClient(client *http.Client) *PostCustomDomainsParams {
	var ()
	return &PostCustomDomainsParams{
		HTTPClient: client,
	}
}

/*PostCustomDomainsParams contains all the parameters to send to the API endpoint
for the post custom domains operation typically these are written to a http.Request
*/
type PostCustomDomainsParams struct {

	/*Body*/
	Body *models.PostCustomDomainsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post custom domains params
func (o *PostCustomDomainsParams) WithTimeout(timeout time.Duration) *PostCustomDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post custom domains params
func (o *PostCustomDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post custom domains params
func (o *PostCustomDomainsParams) WithContext(ctx context.Context) *PostCustomDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post custom domains params
func (o *PostCustomDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post custom domains params
func (o *PostCustomDomainsParams) WithHTTPClient(client *http.Client) *PostCustomDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post custom domains params
func (o *PostCustomDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post custom domains params
func (o *PostCustomDomainsParams) WithBody(body *models.PostCustomDomainsBody) *PostCustomDomainsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post custom domains params
func (o *PostCustomDomainsParams) SetBody(body *models.PostCustomDomainsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
