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

// NewPostRulesParams creates a new PostRulesParams object
// with the default values initialized.
func NewPostRulesParams() *PostRulesParams {
	var ()
	return &PostRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRulesParamsWithTimeout creates a new PostRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRulesParamsWithTimeout(timeout time.Duration) *PostRulesParams {
	var ()
	return &PostRulesParams{

		timeout: timeout,
	}
}

// NewPostRulesParamsWithContext creates a new PostRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRulesParamsWithContext(ctx context.Context) *PostRulesParams {
	var ()
	return &PostRulesParams{

		Context: ctx,
	}
}

// NewPostRulesParamsWithHTTPClient creates a new PostRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRulesParamsWithHTTPClient(client *http.Client) *PostRulesParams {
	var ()
	return &PostRulesParams{
		HTTPClient: client,
	}
}

/*PostRulesParams contains all the parameters to send to the API endpoint
for the post rules operation typically these are written to a http.Request
*/
type PostRulesParams struct {

	/*Body*/
	Body *models.PostRulesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post rules params
func (o *PostRulesParams) WithTimeout(timeout time.Duration) *PostRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post rules params
func (o *PostRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post rules params
func (o *PostRulesParams) WithContext(ctx context.Context) *PostRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post rules params
func (o *PostRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post rules params
func (o *PostRulesParams) WithHTTPClient(client *http.Client) *PostRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post rules params
func (o *PostRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post rules params
func (o *PostRulesParams) WithBody(body *models.PostRulesBody) *PostRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post rules params
func (o *PostRulesParams) SetBody(body *models.PostRulesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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