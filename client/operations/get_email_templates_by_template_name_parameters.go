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

// NewGetEmailTemplatesByTemplateNameParams creates a new GetEmailTemplatesByTemplateNameParams object
// with the default values initialized.
func NewGetEmailTemplatesByTemplateNameParams() *GetEmailTemplatesByTemplateNameParams {
	var ()
	return &GetEmailTemplatesByTemplateNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailTemplatesByTemplateNameParamsWithTimeout creates a new GetEmailTemplatesByTemplateNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailTemplatesByTemplateNameParamsWithTimeout(timeout time.Duration) *GetEmailTemplatesByTemplateNameParams {
	var ()
	return &GetEmailTemplatesByTemplateNameParams{

		timeout: timeout,
	}
}

// NewGetEmailTemplatesByTemplateNameParamsWithContext creates a new GetEmailTemplatesByTemplateNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailTemplatesByTemplateNameParamsWithContext(ctx context.Context) *GetEmailTemplatesByTemplateNameParams {
	var ()
	return &GetEmailTemplatesByTemplateNameParams{

		Context: ctx,
	}
}

// NewGetEmailTemplatesByTemplateNameParamsWithHTTPClient creates a new GetEmailTemplatesByTemplateNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailTemplatesByTemplateNameParamsWithHTTPClient(client *http.Client) *GetEmailTemplatesByTemplateNameParams {
	var ()
	return &GetEmailTemplatesByTemplateNameParams{
		HTTPClient: client,
	}
}

/*GetEmailTemplatesByTemplateNameParams contains all the parameters to send to the API endpoint
for the get email templates by template name operation typically these are written to a http.Request
*/
type GetEmailTemplatesByTemplateNameParams struct {

	/*TemplateName
	  The template name (verify_email, reset_email, welcome_email, blocked_account, stolen_credentials, enrollment_email, mfa_oob_code, <s>change_password</s> (legacy), <s>password_reset</s> (legacy)).

	*/
	TemplateName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) WithTimeout(timeout time.Duration) *GetEmailTemplatesByTemplateNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) WithContext(ctx context.Context) *GetEmailTemplatesByTemplateNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) WithHTTPClient(client *http.Client) *GetEmailTemplatesByTemplateNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTemplateName adds the templateName to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) WithTemplateName(templateName string) *GetEmailTemplatesByTemplateNameParams {
	o.SetTemplateName(templateName)
	return o
}

// SetTemplateName adds the templateName to the get email templates by template name params
func (o *GetEmailTemplatesByTemplateNameParams) SetTemplateName(templateName string) {
	o.TemplateName = templateName
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailTemplatesByTemplateNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param templateName
	if err := r.SetPathParam("templateName", o.TemplateName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
