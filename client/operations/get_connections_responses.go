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

// GetConnectionsReader is a Reader for the GetConnections structure.
type GetConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetConnectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetConnectionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetConnectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetConnectionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConnectionsOK creates a GetConnectionsOK with default headers values
func NewGetConnectionsOK() *GetConnectionsOK {
	return &GetConnectionsOK{}
}

/*GetConnectionsOK handles this case with default header values.

The conections were retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetConnectionsOK struct {
	Payload []*models.GetConnectionsResponse
}

func (o *GetConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/connections][%d] getConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionsBadRequest creates a GetConnectionsBadRequest with default headers values
func NewGetConnectionsBadRequest() *GetConnectionsBadRequest {
	return &GetConnectionsBadRequest{}
}

/*GetConnectionsBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetConnectionsBadRequest struct {
}

func (o *GetConnectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/connections][%d] getConnectionsBadRequest ", 400)
}

func (o *GetConnectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsUnauthorized creates a GetConnectionsUnauthorized with default headers values
func NewGetConnectionsUnauthorized() *GetConnectionsUnauthorized {
	return &GetConnectionsUnauthorized{}
}

/*GetConnectionsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetConnectionsUnauthorized struct {
}

func (o *GetConnectionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/connections][%d] getConnectionsUnauthorized ", 401)
}

func (o *GetConnectionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsForbidden creates a GetConnectionsForbidden with default headers values
func NewGetConnectionsForbidden() *GetConnectionsForbidden {
	return &GetConnectionsForbidden{}
}

/*GetConnectionsForbidden handles this case with default header values.

Insufficient scope, expected any of: read:connections
*/
type GetConnectionsForbidden struct {
}

func (o *GetConnectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/connections][%d] getConnectionsForbidden ", 403)
}

func (o *GetConnectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsTooManyRequests creates a GetConnectionsTooManyRequests with default headers values
func NewGetConnectionsTooManyRequests() *GetConnectionsTooManyRequests {
	return &GetConnectionsTooManyRequests{}
}

/*GetConnectionsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetConnectionsTooManyRequests struct {
}

func (o *GetConnectionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/connections][%d] getConnectionsTooManyRequests ", 429)
}

func (o *GetConnectionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
