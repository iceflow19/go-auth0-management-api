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

// NewPatchSettingsParams creates a new PatchSettingsParams object
// with the default values initialized.
func NewPatchSettingsParams() *PatchSettingsParams {
	var ()
	return &PatchSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSettingsParamsWithTimeout creates a new PatchSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSettingsParamsWithTimeout(timeout time.Duration) *PatchSettingsParams {
	var ()
	return &PatchSettingsParams{

		timeout: timeout,
	}
}

// NewPatchSettingsParamsWithContext creates a new PatchSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSettingsParamsWithContext(ctx context.Context) *PatchSettingsParams {
	var ()
	return &PatchSettingsParams{

		Context: ctx,
	}
}

// NewPatchSettingsParamsWithHTTPClient creates a new PatchSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSettingsParamsWithHTTPClient(client *http.Client) *PatchSettingsParams {
	var ()
	return &PatchSettingsParams{
		HTTPClient: client,
	}
}

/*PatchSettingsParams contains all the parameters to send to the API endpoint
for the patch settings operation typically these are written to a http.Request
*/
type PatchSettingsParams struct {

	/*Body*/
	Body *models.PatchSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch settings params
func (o *PatchSettingsParams) WithTimeout(timeout time.Duration) *PatchSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch settings params
func (o *PatchSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch settings params
func (o *PatchSettingsParams) WithContext(ctx context.Context) *PatchSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch settings params
func (o *PatchSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch settings params
func (o *PatchSettingsParams) WithHTTPClient(client *http.Client) *PatchSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch settings params
func (o *PatchSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch settings params
func (o *PatchSettingsParams) WithBody(body *models.PatchSettingsBody) *PatchSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch settings params
func (o *PatchSettingsParams) SetBody(body *models.PatchSettingsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
