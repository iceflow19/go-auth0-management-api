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

// NewPostIdentitiesParams creates a new PostIdentitiesParams object
// with the default values initialized.
func NewPostIdentitiesParams() *PostIdentitiesParams {
	var ()
	return &PostIdentitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIdentitiesParamsWithTimeout creates a new PostIdentitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIdentitiesParamsWithTimeout(timeout time.Duration) *PostIdentitiesParams {
	var ()
	return &PostIdentitiesParams{

		timeout: timeout,
	}
}

// NewPostIdentitiesParamsWithContext creates a new PostIdentitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIdentitiesParamsWithContext(ctx context.Context) *PostIdentitiesParams {
	var ()
	return &PostIdentitiesParams{

		Context: ctx,
	}
}

// NewPostIdentitiesParamsWithHTTPClient creates a new PostIdentitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIdentitiesParamsWithHTTPClient(client *http.Client) *PostIdentitiesParams {
	var ()
	return &PostIdentitiesParams{
		HTTPClient: client,
	}
}

/*PostIdentitiesParams contains all the parameters to send to the API endpoint
for the post identities operation typically these are written to a http.Request
*/
type PostIdentitiesParams struct {

	/*Body*/
	Body *models.PostIdentitiesBody
	/*ID
	  The user_id of the primary identity where you are linking the secondary account to.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post identities params
func (o *PostIdentitiesParams) WithTimeout(timeout time.Duration) *PostIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post identities params
func (o *PostIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post identities params
func (o *PostIdentitiesParams) WithContext(ctx context.Context) *PostIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post identities params
func (o *PostIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post identities params
func (o *PostIdentitiesParams) WithHTTPClient(client *http.Client) *PostIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post identities params
func (o *PostIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post identities params
func (o *PostIdentitiesParams) WithBody(body *models.PostIdentitiesBody) *PostIdentitiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post identities params
func (o *PostIdentitiesParams) SetBody(body *models.PostIdentitiesBody) {
	o.Body = body
}

// WithID adds the id to the post identities params
func (o *PostIdentitiesParams) WithID(id string) *PostIdentitiesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post identities params
func (o *PostIdentitiesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}