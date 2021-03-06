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

// GetLogsByIDReader is a Reader for the GetLogsByID structure.
type GetLogsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLogsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetLogsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLogsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLogsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetLogsByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsByIDOK creates a GetLogsByIDOK with default headers values
func NewGetLogsByIDOK() *GetLogsByIDOK {
	return &GetLogsByIDOK{}
}

/*GetLogsByIDOK handles this case with default header values.

The log was retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetLogsByIDOK struct {
	Payload *models.GetLogsByIDResponse
}

func (o *GetLogsByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/logs/{id}][%d] getLogsByIdOK  %+v", 200, o.Payload)
}

func (o *GetLogsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLogsByIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsByIDBadRequest creates a GetLogsByIDBadRequest with default headers values
func NewGetLogsByIDBadRequest() *GetLogsByIDBadRequest {
	return &GetLogsByIDBadRequest{}
}

/*GetLogsByIDBadRequest handles this case with default header values.

Invalid request URI. The message will vary depending on the cause.
*/
type GetLogsByIDBadRequest struct {
}

func (o *GetLogsByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/logs/{id}][%d] getLogsByIdBadRequest ", 400)
}

func (o *GetLogsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByIDUnauthorized creates a GetLogsByIDUnauthorized with default headers values
func NewGetLogsByIDUnauthorized() *GetLogsByIDUnauthorized {
	return &GetLogsByIDUnauthorized{}
}

/*GetLogsByIDUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetLogsByIDUnauthorized struct {
}

func (o *GetLogsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/logs/{id}][%d] getLogsByIdUnauthorized ", 401)
}

func (o *GetLogsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByIDForbidden creates a GetLogsByIDForbidden with default headers values
func NewGetLogsByIDForbidden() *GetLogsByIDForbidden {
	return &GetLogsByIDForbidden{}
}

/*GetLogsByIDForbidden handles this case with default header values.

Insufficient scope, expected any of: read:logs
*/
type GetLogsByIDForbidden struct {
}

func (o *GetLogsByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/logs/{id}][%d] getLogsByIdForbidden ", 403)
}

func (o *GetLogsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByIDNotFound creates a GetLogsByIDNotFound with default headers values
func NewGetLogsByIDNotFound() *GetLogsByIDNotFound {
	return &GetLogsByIDNotFound{}
}

/*GetLogsByIDNotFound handles this case with default header values.

The log does not exist.
*/
type GetLogsByIDNotFound struct {
}

func (o *GetLogsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/logs/{id}][%d] getLogsByIdNotFound ", 404)
}

func (o *GetLogsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByIDTooManyRequests creates a GetLogsByIDTooManyRequests with default headers values
func NewGetLogsByIDTooManyRequests() *GetLogsByIDTooManyRequests {
	return &GetLogsByIDTooManyRequests{}
}

/*GetLogsByIDTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetLogsByIDTooManyRequests struct {
}

func (o *GetLogsByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/logs/{id}][%d] getLogsByIdTooManyRequests ", 429)
}

func (o *GetLogsByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
