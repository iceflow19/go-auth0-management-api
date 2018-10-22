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

// NewPostUsersExportsParams creates a new PostUsersExportsParams object
// with the default values initialized.
func NewPostUsersExportsParams() *PostUsersExportsParams {
	var ()
	return &PostUsersExportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersExportsParamsWithTimeout creates a new PostUsersExportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsersExportsParamsWithTimeout(timeout time.Duration) *PostUsersExportsParams {
	var ()
	return &PostUsersExportsParams{

		timeout: timeout,
	}
}

// NewPostUsersExportsParamsWithContext creates a new PostUsersExportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsersExportsParamsWithContext(ctx context.Context) *PostUsersExportsParams {
	var ()
	return &PostUsersExportsParams{

		Context: ctx,
	}
}

// NewPostUsersExportsParamsWithHTTPClient creates a new PostUsersExportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsersExportsParamsWithHTTPClient(client *http.Client) *PostUsersExportsParams {
	var ()
	return &PostUsersExportsParams{
		HTTPClient: client,
	}
}

/*PostUsersExportsParams contains all the parameters to send to the API endpoint
for the post users exports operation typically these are written to a http.Request
*/
type PostUsersExportsParams struct {

	/*Body*/
	Body *models.PostUsersExportsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post users exports params
func (o *PostUsersExportsParams) WithTimeout(timeout time.Duration) *PostUsersExportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users exports params
func (o *PostUsersExportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users exports params
func (o *PostUsersExportsParams) WithContext(ctx context.Context) *PostUsersExportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users exports params
func (o *PostUsersExportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users exports params
func (o *PostUsersExportsParams) WithHTTPClient(client *http.Client) *PostUsersExportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users exports params
func (o *PostUsersExportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post users exports params
func (o *PostUsersExportsParams) WithBody(body *models.PostUsersExportsBody) *PostUsersExportsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post users exports params
func (o *PostUsersExportsParams) SetBody(body *models.PostUsersExportsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersExportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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