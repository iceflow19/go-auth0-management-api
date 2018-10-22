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

// NewPostEmailTemplatesParams creates a new PostEmailTemplatesParams object
// with the default values initialized.
func NewPostEmailTemplatesParams() *PostEmailTemplatesParams {
	var ()
	return &PostEmailTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEmailTemplatesParamsWithTimeout creates a new PostEmailTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEmailTemplatesParamsWithTimeout(timeout time.Duration) *PostEmailTemplatesParams {
	var ()
	return &PostEmailTemplatesParams{

		timeout: timeout,
	}
}

// NewPostEmailTemplatesParamsWithContext creates a new PostEmailTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEmailTemplatesParamsWithContext(ctx context.Context) *PostEmailTemplatesParams {
	var ()
	return &PostEmailTemplatesParams{

		Context: ctx,
	}
}

// NewPostEmailTemplatesParamsWithHTTPClient creates a new PostEmailTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEmailTemplatesParamsWithHTTPClient(client *http.Client) *PostEmailTemplatesParams {
	var ()
	return &PostEmailTemplatesParams{
		HTTPClient: client,
	}
}

/*PostEmailTemplatesParams contains all the parameters to send to the API endpoint
for the post email templates operation typically these are written to a http.Request
*/
type PostEmailTemplatesParams struct {

	/*Body*/
	Body *models.PostEmailTemplatesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post email templates params
func (o *PostEmailTemplatesParams) WithTimeout(timeout time.Duration) *PostEmailTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post email templates params
func (o *PostEmailTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post email templates params
func (o *PostEmailTemplatesParams) WithContext(ctx context.Context) *PostEmailTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post email templates params
func (o *PostEmailTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post email templates params
func (o *PostEmailTemplatesParams) WithHTTPClient(client *http.Client) *PostEmailTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post email templates params
func (o *PostEmailTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post email templates params
func (o *PostEmailTemplatesParams) WithBody(body *models.PostEmailTemplatesBody) *PostEmailTemplatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post email templates params
func (o *PostEmailTemplatesParams) SetBody(body *models.PostEmailTemplatesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEmailTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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