// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	models "bitbucket.org/troyko/go-auth0-management-api/models"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// PatchEmailTemplatesByTemplateNameReader is a Reader for the PatchEmailTemplatesByTemplateName structure.
type PatchEmailTemplatesByTemplateNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEmailTemplatesByTemplateNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchEmailTemplatesByTemplateNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchEmailTemplatesByTemplateNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchEmailTemplatesByTemplateNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchEmailTemplatesByTemplateNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchEmailTemplatesByTemplateNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPatchEmailTemplatesByTemplateNameTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchEmailTemplatesByTemplateNameOK creates a PatchEmailTemplatesByTemplateNameOK with default headers values
func NewPatchEmailTemplatesByTemplateNameOK() *PatchEmailTemplatesByTemplateNameOK {
	return &PatchEmailTemplatesByTemplateNameOK{}
}

/*PatchEmailTemplatesByTemplateNameOK handles this case with default header values.

The template was updated. See <strong>Response Class</strong> below for schema.
*/
type PatchEmailTemplatesByTemplateNameOK struct {
	Payload *models.PatchEmailTemplatesByTemplateNameResponse
}

func (o *PatchEmailTemplatesByTemplateNameOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/email-templates/{templateName}][%d] patchEmailTemplatesByTemplateNameOK  %+v", 200, o.Payload)
}

func (o *PatchEmailTemplatesByTemplateNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchEmailTemplatesByTemplateNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEmailTemplatesByTemplateNameBadRequest creates a PatchEmailTemplatesByTemplateNameBadRequest with default headers values
func NewPatchEmailTemplatesByTemplateNameBadRequest() *PatchEmailTemplatesByTemplateNameBadRequest {
	return &PatchEmailTemplatesByTemplateNameBadRequest{}
}

/*PatchEmailTemplatesByTemplateNameBadRequest handles this case with default header values.

Invalid request body. The message will vary depending on the cause.
*/
type PatchEmailTemplatesByTemplateNameBadRequest struct {
}

func (o *PatchEmailTemplatesByTemplateNameBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/email-templates/{templateName}][%d] patchEmailTemplatesByTemplateNameBadRequest ", 400)
}

func (o *PatchEmailTemplatesByTemplateNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailTemplatesByTemplateNameUnauthorized creates a PatchEmailTemplatesByTemplateNameUnauthorized with default headers values
func NewPatchEmailTemplatesByTemplateNameUnauthorized() *PatchEmailTemplatesByTemplateNameUnauthorized {
	return &PatchEmailTemplatesByTemplateNameUnauthorized{}
}

/*PatchEmailTemplatesByTemplateNameUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation.
*/
type PatchEmailTemplatesByTemplateNameUnauthorized struct {
}

func (o *PatchEmailTemplatesByTemplateNameUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/email-templates/{templateName}][%d] patchEmailTemplatesByTemplateNameUnauthorized ", 401)
}

func (o *PatchEmailTemplatesByTemplateNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailTemplatesByTemplateNameForbidden creates a PatchEmailTemplatesByTemplateNameForbidden with default headers values
func NewPatchEmailTemplatesByTemplateNameForbidden() *PatchEmailTemplatesByTemplateNameForbidden {
	return &PatchEmailTemplatesByTemplateNameForbidden{}
}

/*PatchEmailTemplatesByTemplateNameForbidden handles this case with default header values.

Insufficient scope, expected: update:email_templates.
*/
type PatchEmailTemplatesByTemplateNameForbidden struct {
}

func (o *PatchEmailTemplatesByTemplateNameForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/email-templates/{templateName}][%d] patchEmailTemplatesByTemplateNameForbidden ", 403)
}

func (o *PatchEmailTemplatesByTemplateNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailTemplatesByTemplateNameNotFound creates a PatchEmailTemplatesByTemplateNameNotFound with default headers values
func NewPatchEmailTemplatesByTemplateNameNotFound() *PatchEmailTemplatesByTemplateNameNotFound {
	return &PatchEmailTemplatesByTemplateNameNotFound{}
}

/*PatchEmailTemplatesByTemplateNameNotFound handles this case with default header values.

The template does not exist and cannot be updated.
*/
type PatchEmailTemplatesByTemplateNameNotFound struct {
}

func (o *PatchEmailTemplatesByTemplateNameNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/email-templates/{templateName}][%d] patchEmailTemplatesByTemplateNameNotFound ", 404)
}

func (o *PatchEmailTemplatesByTemplateNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEmailTemplatesByTemplateNameTooManyRequests creates a PatchEmailTemplatesByTemplateNameTooManyRequests with default headers values
func NewPatchEmailTemplatesByTemplateNameTooManyRequests() *PatchEmailTemplatesByTemplateNameTooManyRequests {
	return &PatchEmailTemplatesByTemplateNameTooManyRequests{}
}

/*PatchEmailTemplatesByTemplateNameTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type PatchEmailTemplatesByTemplateNameTooManyRequests struct {
}

func (o *PatchEmailTemplatesByTemplateNameTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/email-templates/{templateName}][%d] patchEmailTemplatesByTemplateNameTooManyRequests ", 429)
}

func (o *PatchEmailTemplatesByTemplateNameTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
