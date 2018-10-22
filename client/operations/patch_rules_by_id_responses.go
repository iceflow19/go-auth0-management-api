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

// PatchRulesByIDReader is a Reader for the PatchRulesByID structure.
type PatchRulesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRulesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchRulesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchRulesByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchRulesByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchRulesByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchRulesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchRulesByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPatchRulesByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchRulesByIDOK creates a PatchRulesByIDOK with default headers values
func NewPatchRulesByIDOK() *PatchRulesByIDOK {
	return &PatchRulesByIDOK{}
}

/*PatchRulesByIDOK handles this case with default header values.

The rule was updated. See <strong>Response Class</strong> below for schema.
*/
type PatchRulesByIDOK struct {
	Payload *models.PatchRulesByIDResponse
}

func (o *PatchRulesByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdOK  %+v", 200, o.Payload)
}

func (o *PatchRulesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchRulesByIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRulesByIDBadRequest creates a PatchRulesByIDBadRequest with default headers values
func NewPatchRulesByIDBadRequest() *PatchRulesByIDBadRequest {
	return &PatchRulesByIDBadRequest{}
}

/*PatchRulesByIDBadRequest handles this case with default header values.

Invalid request body. The message will vary depending on the cause.
*/
type PatchRulesByIDBadRequest struct {
}

func (o *PatchRulesByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdBadRequest ", 400)
}

func (o *PatchRulesByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRulesByIDUnauthorized creates a PatchRulesByIDUnauthorized with default headers values
func NewPatchRulesByIDUnauthorized() *PatchRulesByIDUnauthorized {
	return &PatchRulesByIDUnauthorized{}
}

/*PatchRulesByIDUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type PatchRulesByIDUnauthorized struct {
}

func (o *PatchRulesByIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdUnauthorized ", 401)
}

func (o *PatchRulesByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRulesByIDForbidden creates a PatchRulesByIDForbidden with default headers values
func NewPatchRulesByIDForbidden() *PatchRulesByIDForbidden {
	return &PatchRulesByIDForbidden{}
}

/*PatchRulesByIDForbidden handles this case with default header values.

Insufficient scope, expected any of: update:rules
*/
type PatchRulesByIDForbidden struct {
}

func (o *PatchRulesByIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdForbidden ", 403)
}

func (o *PatchRulesByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRulesByIDNotFound creates a PatchRulesByIDNotFound with default headers values
func NewPatchRulesByIDNotFound() *PatchRulesByIDNotFound {
	return &PatchRulesByIDNotFound{}
}

/*PatchRulesByIDNotFound handles this case with default header values.

The rule does not exist.
*/
type PatchRulesByIDNotFound struct {
}

func (o *PatchRulesByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdNotFound ", 404)
}

func (o *PatchRulesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRulesByIDConflict creates a PatchRulesByIDConflict with default headers values
func NewPatchRulesByIDConflict() *PatchRulesByIDConflict {
	return &PatchRulesByIDConflict{}
}

/*PatchRulesByIDConflict handles this case with default header values.

A rule with the same name already exists
*/
type PatchRulesByIDConflict struct {
}

func (o *PatchRulesByIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdConflict ", 409)
}

func (o *PatchRulesByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRulesByIDTooManyRequests creates a PatchRulesByIDTooManyRequests with default headers values
func NewPatchRulesByIDTooManyRequests() *PatchRulesByIDTooManyRequests {
	return &PatchRulesByIDTooManyRequests{}
}

/*PatchRulesByIDTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type PatchRulesByIDTooManyRequests struct {
}

func (o *PatchRulesByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/rules/{id}][%d] patchRulesByIdTooManyRequests ", 429)
}

func (o *PatchRulesByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}