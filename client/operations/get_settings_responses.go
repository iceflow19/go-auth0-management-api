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

// GetSettingsReader is a Reader for the GetSettings structure.
type GetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSettingsOK creates a GetSettingsOK with default headers values
func NewGetSettingsOK() *GetSettingsOK {
	return &GetSettingsOK{}
}

/*GetSettingsOK handles this case with default header values.

The tenant settings were retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetSettingsOK struct {
	Payload *models.GetSettingsResponse
}

func (o *GetSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/tenants/settings][%d] getSettingsOK  %+v", 200, o.Payload)
}

func (o *GetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsBadRequest creates a GetSettingsBadRequest with default headers values
func NewGetSettingsBadRequest() *GetSettingsBadRequest {
	return &GetSettingsBadRequest{}
}

/*GetSettingsBadRequest handles this case with default header values.

Invalid request query string. The message will vary depending on the cause.
*/
type GetSettingsBadRequest struct {
}

func (o *GetSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/tenants/settings][%d] getSettingsBadRequest ", 400)
}

func (o *GetSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSettingsUnauthorized creates a GetSettingsUnauthorized with default headers values
func NewGetSettingsUnauthorized() *GetSettingsUnauthorized {
	return &GetSettingsUnauthorized{}
}

/*GetSettingsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetSettingsUnauthorized struct {
}

func (o *GetSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/tenants/settings][%d] getSettingsUnauthorized ", 401)
}

func (o *GetSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSettingsForbidden creates a GetSettingsForbidden with default headers values
func NewGetSettingsForbidden() *GetSettingsForbidden {
	return &GetSettingsForbidden{}
}

/*GetSettingsForbidden handles this case with default header values.

Insufficient scope, expected any of: read:tenant_settings
*/
type GetSettingsForbidden struct {
}

func (o *GetSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/tenants/settings][%d] getSettingsForbidden ", 403)
}

func (o *GetSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSettingsTooManyRequests creates a GetSettingsTooManyRequests with default headers values
func NewGetSettingsTooManyRequests() *GetSettingsTooManyRequests {
	return &GetSettingsTooManyRequests{}
}

/*GetSettingsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetSettingsTooManyRequests struct {
}

func (o *GetSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/tenants/settings][%d] getSettingsTooManyRequests ", 429)
}

func (o *GetSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}