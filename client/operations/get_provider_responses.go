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

// GetProviderReader is a Reader for the GetProvider structure.
type GetProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetProviderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetProviderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetProviderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProviderOK creates a GetProviderOK with default headers values
func NewGetProviderOK() *GetProviderOK {
	return &GetProviderOK{}
}

/*GetProviderOK handles this case with default header values.

The email provider was retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetProviderOK struct {
	Payload *models.GetProviderResponse
}

func (o *GetProviderOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/emails/provider][%d] getProviderOK  %+v", 200, o.Payload)
}

func (o *GetProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProviderBadRequest creates a GetProviderBadRequest with default headers values
func NewGetProviderBadRequest() *GetProviderBadRequest {
	return &GetProviderBadRequest{}
}

/*GetProviderBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetProviderBadRequest struct {
}

func (o *GetProviderBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/emails/provider][%d] getProviderBadRequest ", 400)
}

func (o *GetProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProviderUnauthorized creates a GetProviderUnauthorized with default headers values
func NewGetProviderUnauthorized() *GetProviderUnauthorized {
	return &GetProviderUnauthorized{}
}

/*GetProviderUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetProviderUnauthorized struct {
}

func (o *GetProviderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/emails/provider][%d] getProviderUnauthorized ", 401)
}

func (o *GetProviderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProviderForbidden creates a GetProviderForbidden with default headers values
func NewGetProviderForbidden() *GetProviderForbidden {
	return &GetProviderForbidden{}
}

/*GetProviderForbidden handles this case with default header values.

Insufficient scope, expected any of: read:email_provider
*/
type GetProviderForbidden struct {
}

func (o *GetProviderForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/emails/provider][%d] getProviderForbidden ", 403)
}

func (o *GetProviderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProviderNotFound creates a GetProviderNotFound with default headers values
func NewGetProviderNotFound() *GetProviderNotFound {
	return &GetProviderNotFound{}
}

/*GetProviderNotFound handles this case with default header values.

There is not a configured email provider
*/
type GetProviderNotFound struct {
}

func (o *GetProviderNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/emails/provider][%d] getProviderNotFound ", 404)
}

func (o *GetProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProviderTooManyRequests creates a GetProviderTooManyRequests with default headers values
func NewGetProviderTooManyRequests() *GetProviderTooManyRequests {
	return &GetProviderTooManyRequests{}
}

/*GetProviderTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetProviderTooManyRequests struct {
}

func (o *GetProviderTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/emails/provider][%d] getProviderTooManyRequests ", 429)
}

func (o *GetProviderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
