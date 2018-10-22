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

// PatchProviderReader is a Reader for the PatchProvider structure.
type PatchProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchProviderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchProviderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPatchProviderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchProviderOK creates a PatchProviderOK with default headers values
func NewPatchProviderOK() *PatchProviderOK {
	return &PatchProviderOK{}
}

/*PatchProviderOK handles this case with default header values.

The email provider was updated. See <strong>Response Class</strong> below for schema.
*/
type PatchProviderOK struct {
	Payload *models.PatchProviderResponse
}

func (o *PatchProviderOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/emails/provider][%d] patchProviderOK  %+v", 200, o.Payload)
}

func (o *PatchProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchProviderBadRequest creates a PatchProviderBadRequest with default headers values
func NewPatchProviderBadRequest() *PatchProviderBadRequest {
	return &PatchProviderBadRequest{}
}

/*PatchProviderBadRequest handles this case with default header values.

Invalid request body. The message will vary depending on the cause.
*/
type PatchProviderBadRequest struct {
}

func (o *PatchProviderBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/emails/provider][%d] patchProviderBadRequest ", 400)
}

func (o *PatchProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProviderUnauthorized creates a PatchProviderUnauthorized with default headers values
func NewPatchProviderUnauthorized() *PatchProviderUnauthorized {
	return &PatchProviderUnauthorized{}
}

/*PatchProviderUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type PatchProviderUnauthorized struct {
}

func (o *PatchProviderUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/emails/provider][%d] patchProviderUnauthorized ", 401)
}

func (o *PatchProviderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProviderForbidden creates a PatchProviderForbidden with default headers values
func NewPatchProviderForbidden() *PatchProviderForbidden {
	return &PatchProviderForbidden{}
}

/*PatchProviderForbidden handles this case with default header values.

Insufficient scope, expected any of: update:email_provider
*/
type PatchProviderForbidden struct {
}

func (o *PatchProviderForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/emails/provider][%d] patchProviderForbidden ", 403)
}

func (o *PatchProviderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProviderNotFound creates a PatchProviderNotFound with default headers values
func NewPatchProviderNotFound() *PatchProviderNotFound {
	return &PatchProviderNotFound{}
}

/*PatchProviderNotFound handles this case with default header values.

There is not a configured email provider
*/
type PatchProviderNotFound struct {
}

func (o *PatchProviderNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/emails/provider][%d] patchProviderNotFound ", 404)
}

func (o *PatchProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProviderTooManyRequests creates a PatchProviderTooManyRequests with default headers values
func NewPatchProviderTooManyRequests() *PatchProviderTooManyRequests {
	return &PatchProviderTooManyRequests{}
}

/*PatchProviderTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type PatchProviderTooManyRequests struct {
}

func (o *PatchProviderTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/emails/provider][%d] patchProviderTooManyRequests ", 429)
}

func (o *PatchProviderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}