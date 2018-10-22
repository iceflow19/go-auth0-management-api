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

// GetResourceServersReader is a Reader for the GetResourceServers structure.
type GetResourceServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResourceServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetResourceServersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetResourceServersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetResourceServersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetResourceServersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResourceServersOK creates a GetResourceServersOK with default headers values
func NewGetResourceServersOK() *GetResourceServersOK {
	return &GetResourceServersOK{}
}

/*GetResourceServersOK handles this case with default header values.

The resource servers were retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetResourceServersOK struct {
	Payload []*models.GetResourceServersResponse
}

func (o *GetResourceServersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/resource-servers][%d] getResourceServersOK  %+v", 200, o.Payload)
}

func (o *GetResourceServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceServersBadRequest creates a GetResourceServersBadRequest with default headers values
func NewGetResourceServersBadRequest() *GetResourceServersBadRequest {
	return &GetResourceServersBadRequest{}
}

/*GetResourceServersBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetResourceServersBadRequest struct {
}

func (o *GetResourceServersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/resource-servers][%d] getResourceServersBadRequest ", 400)
}

func (o *GetResourceServersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceServersUnauthorized creates a GetResourceServersUnauthorized with default headers values
func NewGetResourceServersUnauthorized() *GetResourceServersUnauthorized {
	return &GetResourceServersUnauthorized{}
}

/*GetResourceServersUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetResourceServersUnauthorized struct {
}

func (o *GetResourceServersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/resource-servers][%d] getResourceServersUnauthorized ", 401)
}

func (o *GetResourceServersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceServersForbidden creates a GetResourceServersForbidden with default headers values
func NewGetResourceServersForbidden() *GetResourceServersForbidden {
	return &GetResourceServersForbidden{}
}

/*GetResourceServersForbidden handles this case with default header values.

Insufficient scope, expected any of: read:resource_servers
*/
type GetResourceServersForbidden struct {
}

func (o *GetResourceServersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/resource-servers][%d] getResourceServersForbidden ", 403)
}

func (o *GetResourceServersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceServersTooManyRequests creates a GetResourceServersTooManyRequests with default headers values
func NewGetResourceServersTooManyRequests() *GetResourceServersTooManyRequests {
	return &GetResourceServersTooManyRequests{}
}

/*GetResourceServersTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetResourceServersTooManyRequests struct {
}

func (o *GetResourceServersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/resource-servers][%d] getResourceServersTooManyRequests ", 429)
}

func (o *GetResourceServersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}