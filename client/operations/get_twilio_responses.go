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

// GetTwilioReader is a Reader for the GetTwilio structure.
type GetTwilioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTwilioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTwilioOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTwilioBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetTwilioUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTwilioForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTwilioOK creates a GetTwilioOK with default headers values
func NewGetTwilioOK() *GetTwilioOK {
	return &GetTwilioOK{}
}

/*GetTwilioOK handles this case with default header values.

Returns Twilio configuration data
*/
type GetTwilioOK struct {
	Payload *models.GetTwilioResponse
}

func (o *GetTwilioOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/guardian/factors/sms/providers/twilio][%d] getTwilioOK  %+v", 200, o.Payload)
}

func (o *GetTwilioOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTwilioResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTwilioBadRequest creates a GetTwilioBadRequest with default headers values
func NewGetTwilioBadRequest() *GetTwilioBadRequest {
	return &GetTwilioBadRequest{}
}

/*GetTwilioBadRequest handles this case with default header values.

Invalid input based on schemas
*/
type GetTwilioBadRequest struct {
}

func (o *GetTwilioBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/guardian/factors/sms/providers/twilio][%d] getTwilioBadRequest ", 400)
}

func (o *GetTwilioBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTwilioUnauthorized creates a GetTwilioUnauthorized with default headers values
func NewGetTwilioUnauthorized() *GetTwilioUnauthorized {
	return &GetTwilioUnauthorized{}
}

/*GetTwilioUnauthorized handles this case with default header values.

Token has expired or signature is invalid
*/
type GetTwilioUnauthorized struct {
}

func (o *GetTwilioUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/guardian/factors/sms/providers/twilio][%d] getTwilioUnauthorized ", 401)
}

func (o *GetTwilioUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTwilioForbidden creates a GetTwilioForbidden with default headers values
func NewGetTwilioForbidden() *GetTwilioForbidden {
	return &GetTwilioForbidden{}
}

/*GetTwilioForbidden handles this case with default header values.

Insufficient scope
*/
type GetTwilioForbidden struct {
}

func (o *GetTwilioForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/guardian/factors/sms/providers/twilio][%d] getTwilioForbidden ", 403)
}

func (o *GetTwilioForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}