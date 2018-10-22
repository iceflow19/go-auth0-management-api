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

// GetClientsReader is a Reader for the GetClients structure.
type GetClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetClientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetClientsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClientsOK creates a GetClientsOK with default headers values
func NewGetClientsOK() *GetClientsOK {
	return &GetClientsOK{}
}

/*GetClientsOK handles this case with default header values.

The clients were retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetClientsOK struct {
	Payload []*models.GetClientsResponse
}

func (o *GetClientsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients][%d] getClientsOK  %+v", 200, o.Payload)
}

func (o *GetClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientsBadRequest creates a GetClientsBadRequest with default headers values
func NewGetClientsBadRequest() *GetClientsBadRequest {
	return &GetClientsBadRequest{}
}

/*GetClientsBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetClientsBadRequest struct {
}

func (o *GetClientsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients][%d] getClientsBadRequest ", 400)
}

func (o *GetClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsUnauthorized creates a GetClientsUnauthorized with default headers values
func NewGetClientsUnauthorized() *GetClientsUnauthorized {
	return &GetClientsUnauthorized{}
}

/*GetClientsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetClientsUnauthorized struct {
}

func (o *GetClientsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients][%d] getClientsUnauthorized ", 401)
}

func (o *GetClientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsForbidden creates a GetClientsForbidden with default headers values
func NewGetClientsForbidden() *GetClientsForbidden {
	return &GetClientsForbidden{}
}

/*GetClientsForbidden handles this case with default header values.

Some fields cannot be read with the permissions granted by the bearer token scopes. The message will vary depending on the fields and the scopes.
*/
type GetClientsForbidden struct {
}

func (o *GetClientsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients][%d] getClientsForbidden ", 403)
}

func (o *GetClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClientsTooManyRequests creates a GetClientsTooManyRequests with default headers values
func NewGetClientsTooManyRequests() *GetClientsTooManyRequests {
	return &GetClientsTooManyRequests{}
}

/*GetClientsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetClientsTooManyRequests struct {
}

func (o *GetClientsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/clients][%d] getClientsTooManyRequests ", 429)
}

func (o *GetClientsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}