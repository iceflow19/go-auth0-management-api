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

// PostTicketReader is a Reader for the PostTicket structure.
type PostTicketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTicketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTicketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewPostTicketNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTicketBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostTicketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostTicketForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostTicketNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTicketOK creates a PostTicketOK with default headers values
func NewPostTicketOK() *PostTicketOK {
	return &PostTicketOK{}
}

/*PostTicketOK handles this case with default header values.

No response was specified
*/
type PostTicketOK struct {
	Payload *models.PostTicketResponse
}

func (o *PostTicketOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/guardian/enrollments/ticket][%d] postTicketOK  %+v", 200, o.Payload)
}

func (o *PostTicketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostTicketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTicketNoContent creates a PostTicketNoContent with default headers values
func NewPostTicketNoContent() *PostTicketNoContent {
	return &PostTicketNoContent{}
}

/*PostTicketNoContent handles this case with default header values.

Ticket created
*/
type PostTicketNoContent struct {
}

func (o *PostTicketNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/guardian/enrollments/ticket][%d] postTicketNoContent ", 204)
}

func (o *PostTicketNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTicketBadRequest creates a PostTicketBadRequest with default headers values
func NewPostTicketBadRequest() *PostTicketBadRequest {
	return &PostTicketBadRequest{}
}

/*PostTicketBadRequest handles this case with default header values.

Invalid input based on schemas.
*/
type PostTicketBadRequest struct {
}

func (o *PostTicketBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/guardian/enrollments/ticket][%d] postTicketBadRequest ", 400)
}

func (o *PostTicketBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTicketUnauthorized creates a PostTicketUnauthorized with default headers values
func NewPostTicketUnauthorized() *PostTicketUnauthorized {
	return &PostTicketUnauthorized{}
}

/*PostTicketUnauthorized handles this case with default header values.

Token has expired or signature is invalid
*/
type PostTicketUnauthorized struct {
}

func (o *PostTicketUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/guardian/enrollments/ticket][%d] postTicketUnauthorized ", 401)
}

func (o *PostTicketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTicketForbidden creates a PostTicketForbidden with default headers values
func NewPostTicketForbidden() *PostTicketForbidden {
	return &PostTicketForbidden{}
}

/*PostTicketForbidden handles this case with default header values.

Insufficient scope
*/
type PostTicketForbidden struct {
}

func (o *PostTicketForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/guardian/enrollments/ticket][%d] postTicketForbidden ", 403)
}

func (o *PostTicketForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTicketNotFound creates a PostTicketNotFound with default headers values
func NewPostTicketNotFound() *PostTicketNotFound {
	return &PostTicketNotFound{}
}

/*PostTicketNotFound handles this case with default header values.

The user does not exist.
*/
type PostTicketNotFound struct {
}

func (o *PostTicketNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/guardian/enrollments/ticket][%d] postTicketNotFound ", 404)
}

func (o *PostTicketNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
