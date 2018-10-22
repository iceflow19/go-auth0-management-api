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

// GetLogsByUserReader is a Reader for the GetLogsByUser structure.
type GetLogsByUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsByUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogsByUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLogsByUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetLogsByUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetLogsByUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetLogsByUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsByUserOK creates a GetLogsByUserOK with default headers values
func NewGetLogsByUserOK() *GetLogsByUserOK {
	return &GetLogsByUserOK{}
}

/*GetLogsByUserOK handles this case with default header values.

The log was retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetLogsByUserOK struct {
	Payload []*models.GetLogsByUserResponse
}

func (o *GetLogsByUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user_id}/logs][%d] getLogsByUserOK  %+v", 200, o.Payload)
}

func (o *GetLogsByUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsByUserBadRequest creates a GetLogsByUserBadRequest with default headers values
func NewGetLogsByUserBadRequest() *GetLogsByUserBadRequest {
	return &GetLogsByUserBadRequest{}
}

/*GetLogsByUserBadRequest handles this case with default header values.

Invalid request URI. The message will vary depending on the cause.
*/
type GetLogsByUserBadRequest struct {
}

func (o *GetLogsByUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user_id}/logs][%d] getLogsByUserBadRequest ", 400)
}

func (o *GetLogsByUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByUserUnauthorized creates a GetLogsByUserUnauthorized with default headers values
func NewGetLogsByUserUnauthorized() *GetLogsByUserUnauthorized {
	return &GetLogsByUserUnauthorized{}
}

/*GetLogsByUserUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetLogsByUserUnauthorized struct {
}

func (o *GetLogsByUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user_id}/logs][%d] getLogsByUserUnauthorized ", 401)
}

func (o *GetLogsByUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByUserForbidden creates a GetLogsByUserForbidden with default headers values
func NewGetLogsByUserForbidden() *GetLogsByUserForbidden {
	return &GetLogsByUserForbidden{}
}

/*GetLogsByUserForbidden handles this case with default header values.

Insufficient scope, expected any of: read:logs
*/
type GetLogsByUserForbidden struct {
}

func (o *GetLogsByUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user_id}/logs][%d] getLogsByUserForbidden ", 403)
}

func (o *GetLogsByUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogsByUserTooManyRequests creates a GetLogsByUserTooManyRequests with default headers values
func NewGetLogsByUserTooManyRequests() *GetLogsByUserTooManyRequests {
	return &GetLogsByUserTooManyRequests{}
}

/*GetLogsByUserTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetLogsByUserTooManyRequests struct {
}

func (o *GetLogsByUserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user_id}/logs][%d] getLogsByUserTooManyRequests ", 429)
}

func (o *GetLogsByUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}