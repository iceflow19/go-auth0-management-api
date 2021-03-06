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

// GetUserBlocksByIDReader is a Reader for the GetUserBlocksByID structure.
type GetUserBlocksByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserBlocksByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserBlocksByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserBlocksByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUserBlocksByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserBlocksByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserBlocksByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetUserBlocksByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserBlocksByIDOK creates a GetUserBlocksByIDOK with default headers values
func NewGetUserBlocksByIDOK() *GetUserBlocksByIDOK {
	return &GetUserBlocksByIDOK{}
}

/*GetUserBlocksByIDOK handles this case with default header values.

The user was retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetUserBlocksByIDOK struct {
	Payload *models.GetUserBlocksByIDResponse
}

func (o *GetUserBlocksByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/user-blocks/{id}][%d] getUserBlocksByIdOK  %+v", 200, o.Payload)
}

func (o *GetUserBlocksByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUserBlocksByIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserBlocksByIDBadRequest creates a GetUserBlocksByIDBadRequest with default headers values
func NewGetUserBlocksByIDBadRequest() *GetUserBlocksByIDBadRequest {
	return &GetUserBlocksByIDBadRequest{}
}

/*GetUserBlocksByIDBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetUserBlocksByIDBadRequest struct {
}

func (o *GetUserBlocksByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/user-blocks/{id}][%d] getUserBlocksByIdBadRequest ", 400)
}

func (o *GetUserBlocksByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserBlocksByIDUnauthorized creates a GetUserBlocksByIDUnauthorized with default headers values
func NewGetUserBlocksByIDUnauthorized() *GetUserBlocksByIDUnauthorized {
	return &GetUserBlocksByIDUnauthorized{}
}

/*GetUserBlocksByIDUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetUserBlocksByIDUnauthorized struct {
}

func (o *GetUserBlocksByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/user-blocks/{id}][%d] getUserBlocksByIdUnauthorized ", 401)
}

func (o *GetUserBlocksByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserBlocksByIDForbidden creates a GetUserBlocksByIDForbidden with default headers values
func NewGetUserBlocksByIDForbidden() *GetUserBlocksByIDForbidden {
	return &GetUserBlocksByIDForbidden{}
}

/*GetUserBlocksByIDForbidden handles this case with default header values.

Insufficient scope, expected any of: read:users
*/
type GetUserBlocksByIDForbidden struct {
}

func (o *GetUserBlocksByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/user-blocks/{id}][%d] getUserBlocksByIdForbidden ", 403)
}

func (o *GetUserBlocksByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserBlocksByIDNotFound creates a GetUserBlocksByIDNotFound with default headers values
func NewGetUserBlocksByIDNotFound() *GetUserBlocksByIDNotFound {
	return &GetUserBlocksByIDNotFound{}
}

/*GetUserBlocksByIDNotFound handles this case with default header values.

The user does not exist.
*/
type GetUserBlocksByIDNotFound struct {
}

func (o *GetUserBlocksByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/user-blocks/{id}][%d] getUserBlocksByIdNotFound ", 404)
}

func (o *GetUserBlocksByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserBlocksByIDTooManyRequests creates a GetUserBlocksByIDTooManyRequests with default headers values
func NewGetUserBlocksByIDTooManyRequests() *GetUserBlocksByIDTooManyRequests {
	return &GetUserBlocksByIDTooManyRequests{}
}

/*GetUserBlocksByIDTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetUserBlocksByIDTooManyRequests struct {
}

func (o *GetUserBlocksByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/user-blocks/{id}][%d] getUserBlocksByIdTooManyRequests ", 429)
}

func (o *GetUserBlocksByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
