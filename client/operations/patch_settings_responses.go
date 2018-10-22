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

// PatchSettingsReader is a Reader for the PatchSettings structure.
type PatchSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPatchSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSettingsOK creates a PatchSettingsOK with default headers values
func NewPatchSettingsOK() *PatchSettingsOK {
	return &PatchSettingsOK{}
}

/*PatchSettingsOK handles this case with default header values.

The tenant settings were updated. See <strong>Response Class</strong> below for schema.
*/
type PatchSettingsOK struct {
	Payload *models.PatchSettingsResponse
}

func (o *PatchSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/tenants/settings][%d] patchSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSettingsBadRequest creates a PatchSettingsBadRequest with default headers values
func NewPatchSettingsBadRequest() *PatchSettingsBadRequest {
	return &PatchSettingsBadRequest{}
}

/*PatchSettingsBadRequest handles this case with default header values.

Invalid request body. The message will vary depending on the cause.
*/
type PatchSettingsBadRequest struct {
}

func (o *PatchSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/tenants/settings][%d] patchSettingsBadRequest ", 400)
}

func (o *PatchSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSettingsUnauthorized creates a PatchSettingsUnauthorized with default headers values
func NewPatchSettingsUnauthorized() *PatchSettingsUnauthorized {
	return &PatchSettingsUnauthorized{}
}

/*PatchSettingsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type PatchSettingsUnauthorized struct {
}

func (o *PatchSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/tenants/settings][%d] patchSettingsUnauthorized ", 401)
}

func (o *PatchSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSettingsForbidden creates a PatchSettingsForbidden with default headers values
func NewPatchSettingsForbidden() *PatchSettingsForbidden {
	return &PatchSettingsForbidden{}
}

/*PatchSettingsForbidden handles this case with default header values.

Insufficient scope, expected any of: update:tenant_settings
*/
type PatchSettingsForbidden struct {
}

func (o *PatchSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/tenants/settings][%d] patchSettingsForbidden ", 403)
}

func (o *PatchSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSettingsTooManyRequests creates a PatchSettingsTooManyRequests with default headers values
func NewPatchSettingsTooManyRequests() *PatchSettingsTooManyRequests {
	return &PatchSettingsTooManyRequests{}
}

/*PatchSettingsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type PatchSettingsTooManyRequests struct {
}

func (o *PatchSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/tenants/settings][%d] patchSettingsTooManyRequests ", 429)
}

func (o *PatchSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
