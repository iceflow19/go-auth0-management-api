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

// GetClientsByIDReader is a Reader for the GetClientsByID structure.
type GetClientsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClientsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClientsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetClientsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetClientsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetClientsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetClientsByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClientsByIDOK creates a GetClientsByIDOK with default headers values
func NewGetClientsByIDOK() *GetClientsByIDOK {
	return &GetClientsByIDOK{}
}

/*GetClientsByIDOK handles this case with default header values.

The client was retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetClientsByIDOK struct {
	Payload *models.GetClientsByIDResponse
}

func (o *GetClientsByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients/{id}][%d] getClientsByIdOK  %+v", 200, o.Payload)
}

func (o *GetClientsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetClientsByIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientsByIDBadRequest creates a GetClientsByIDBadRequest with default headers values
func NewGetClientsByIDBadRequest() *GetClientsByIDBadRequest {
	return &GetClientsByIDBadRequest{}
}

/*GetClientsByIDBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetClientsByIDBadRequest struct {
}

func (o *GetClientsByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients/{id}][%d] getClientsByIdBadRequest ", 400)
}

func (o *GetClientsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsByIDUnauthorized creates a GetClientsByIDUnauthorized with default headers values
func NewGetClientsByIDUnauthorized() *GetClientsByIDUnauthorized {
	return &GetClientsByIDUnauthorized{}
}

/*GetClientsByIDUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetClientsByIDUnauthorized struct {
}

func (o *GetClientsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients/{id}][%d] getClientsByIdUnauthorized ", 401)
}

func (o *GetClientsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsByIDForbidden creates a GetClientsByIDForbidden with default headers values
func NewGetClientsByIDForbidden() *GetClientsByIDForbidden {
	return &GetClientsByIDForbidden{}
}

/*GetClientsByIDForbidden handles this case with default header values.

Some fields cannot be read with the permissions granted by the bearer token scopes. The message will vary depending on the fields and the scopes.
*/
type GetClientsByIDForbidden struct {
}

func (o *GetClientsByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients/{id}][%d] getClientsByIdForbidden ", 403)
}

func (o *GetClientsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsByIDNotFound creates a GetClientsByIDNotFound with default headers values
func NewGetClientsByIDNotFound() *GetClientsByIDNotFound {
	return &GetClientsByIDNotFound{}
}

/*GetClientsByIDNotFound handles this case with default header values.

The client does not exist.
*/
type GetClientsByIDNotFound struct {
}

func (o *GetClientsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients/{id}][%d] getClientsByIdNotFound ", 404)
}

func (o *GetClientsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsByIDTooManyRequests creates a GetClientsByIDTooManyRequests with default headers values
func NewGetClientsByIDTooManyRequests() *GetClientsByIDTooManyRequests {
	return &GetClientsByIDTooManyRequests{}
}

/*GetClientsByIDTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetClientsByIDTooManyRequests struct {
}

func (o *GetClientsByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients/{id}][%d] getClientsByIdTooManyRequests ", 429)
}

func (o *GetClientsByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
