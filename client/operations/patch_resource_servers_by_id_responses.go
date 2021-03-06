// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// PatchResourceServersByIDReader is a Reader for the PatchResourceServersByID structure.
type PatchResourceServersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchResourceServersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchResourceServersByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPatchResourceServersByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchResourceServersByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchResourceServersByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchResourceServersByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchResourceServersByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPatchResourceServersByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchResourceServersByIDOK creates a PatchResourceServersByIDOK with default headers values
func NewPatchResourceServersByIDOK() *PatchResourceServersByIDOK {
	return &PatchResourceServersByIDOK{}
}

/*PatchResourceServersByIDOK handles this case with default header values.

No response was specified
*/
type PatchResourceServersByIDOK struct {
}

func (o *PatchResourceServersByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdOK ", 200)
}

func (o *PatchResourceServersByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchResourceServersByIDCreated creates a PatchResourceServersByIDCreated with default headers values
func NewPatchResourceServersByIDCreated() *PatchResourceServersByIDCreated {
	return &PatchResourceServersByIDCreated{}
}

/*PatchResourceServersByIDCreated handles this case with default header values.

The resource server was updated.
*/
type PatchResourceServersByIDCreated struct {
}

func (o *PatchResourceServersByIDCreated) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdCreated ", 201)
}

func (o *PatchResourceServersByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchResourceServersByIDBadRequest creates a PatchResourceServersByIDBadRequest with default headers values
func NewPatchResourceServersByIDBadRequest() *PatchResourceServersByIDBadRequest {
	return &PatchResourceServersByIDBadRequest{}
}

/*PatchResourceServersByIDBadRequest handles this case with default header values.

System resource servers cannot be patched.
*/
type PatchResourceServersByIDBadRequest struct {
}

func (o *PatchResourceServersByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdBadRequest ", 400)
}

func (o *PatchResourceServersByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchResourceServersByIDUnauthorized creates a PatchResourceServersByIDUnauthorized with default headers values
func NewPatchResourceServersByIDUnauthorized() *PatchResourceServersByIDUnauthorized {
	return &PatchResourceServersByIDUnauthorized{}
}

/*PatchResourceServersByIDUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type PatchResourceServersByIDUnauthorized struct {
}

func (o *PatchResourceServersByIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdUnauthorized ", 401)
}

func (o *PatchResourceServersByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchResourceServersByIDForbidden creates a PatchResourceServersByIDForbidden with default headers values
func NewPatchResourceServersByIDForbidden() *PatchResourceServersByIDForbidden {
	return &PatchResourceServersByIDForbidden{}
}

/*PatchResourceServersByIDForbidden handles this case with default header values.

Insufficient scope, expected any of: update:resource_servers
*/
type PatchResourceServersByIDForbidden struct {
}

func (o *PatchResourceServersByIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdForbidden ", 403)
}

func (o *PatchResourceServersByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchResourceServersByIDConflict creates a PatchResourceServersByIDConflict with default headers values
func NewPatchResourceServersByIDConflict() *PatchResourceServersByIDConflict {
	return &PatchResourceServersByIDConflict{}
}

/*PatchResourceServersByIDConflict handles this case with default header values.

A resource server with the same identifier already exists.
*/
type PatchResourceServersByIDConflict struct {
}

func (o *PatchResourceServersByIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdConflict ", 409)
}

func (o *PatchResourceServersByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchResourceServersByIDTooManyRequests creates a PatchResourceServersByIDTooManyRequests with default headers values
func NewPatchResourceServersByIDTooManyRequests() *PatchResourceServersByIDTooManyRequests {
	return &PatchResourceServersByIDTooManyRequests{}
}

/*PatchResourceServersByIDTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type PatchResourceServersByIDTooManyRequests struct {
}

func (o *PatchResourceServersByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/resource-servers/{id}][%d] patchResourceServersByIdTooManyRequests ", 429)
}

func (o *PatchResourceServersByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
