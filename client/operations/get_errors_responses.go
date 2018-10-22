// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// GetErrorsReader is a Reader for the GetErrors structure.
type GetErrorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetErrorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetErrorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewGetErrorsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetErrorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetErrorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetErrorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetErrorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetErrorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetErrorsOK creates a GetErrorsOK with default headers values
func NewGetErrorsOK() *GetErrorsOK {
	return &GetErrorsOK{}
}

/*GetErrorsOK handles this case with default header values.

The job was retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetErrorsOK struct {
}

func (o *GetErrorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsOK ", 200)
}

func (o *GetErrorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetErrorsNoContent creates a GetErrorsNoContent with default headers values
func NewGetErrorsNoContent() *GetErrorsNoContent {
	return &GetErrorsNoContent{}
}

/*GetErrorsNoContent handles this case with default header values.

The job was retrieved, but no errors were found.
*/
type GetErrorsNoContent struct {
}

func (o *GetErrorsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsNoContent ", 204)
}

func (o *GetErrorsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetErrorsBadRequest creates a GetErrorsBadRequest with default headers values
func NewGetErrorsBadRequest() *GetErrorsBadRequest {
	return &GetErrorsBadRequest{}
}

/*GetErrorsBadRequest handles this case with default header values.

Invalid request URI. The message will vary depending on the cause.
*/
type GetErrorsBadRequest struct {
}

func (o *GetErrorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsBadRequest ", 400)
}

func (o *GetErrorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetErrorsUnauthorized creates a GetErrorsUnauthorized with default headers values
func NewGetErrorsUnauthorized() *GetErrorsUnauthorized {
	return &GetErrorsUnauthorized{}
}

/*GetErrorsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetErrorsUnauthorized struct {
}

func (o *GetErrorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsUnauthorized ", 401)
}

func (o *GetErrorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetErrorsForbidden creates a GetErrorsForbidden with default headers values
func NewGetErrorsForbidden() *GetErrorsForbidden {
	return &GetErrorsForbidden{}
}

/*GetErrorsForbidden handles this case with default header values.

Insufficient scope, expected any of: create:users, create:passwords_checking_job
*/
type GetErrorsForbidden struct {
}

func (o *GetErrorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsForbidden ", 403)
}

func (o *GetErrorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetErrorsNotFound creates a GetErrorsNotFound with default headers values
func NewGetErrorsNotFound() *GetErrorsNotFound {
	return &GetErrorsNotFound{}
}

/*GetErrorsNotFound handles this case with default header values.

The job does not exist
*/
type GetErrorsNotFound struct {
}

func (o *GetErrorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsNotFound ", 404)
}

func (o *GetErrorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetErrorsTooManyRequests creates a GetErrorsTooManyRequests with default headers values
func NewGetErrorsTooManyRequests() *GetErrorsTooManyRequests {
	return &GetErrorsTooManyRequests{}
}

/*GetErrorsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetErrorsTooManyRequests struct {
}

func (o *GetErrorsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/jobs/{id}/errors][%d] getErrorsTooManyRequests ", 429)
}

func (o *GetErrorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}