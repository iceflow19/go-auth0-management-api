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

// PutFactorsByNameReader is a Reader for the PutFactorsByName structure.
type PutFactorsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFactorsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutFactorsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutFactorsByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutFactorsByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutFactorsByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFactorsByNameOK creates a PutFactorsByNameOK with default headers values
func NewPutFactorsByNameOK() *PutFactorsByNameOK {
	return &PutFactorsByNameOK{}
}

/*PutFactorsByNameOK handles this case with default header values.

Factor was updated
*/
type PutFactorsByNameOK struct {
	Payload *models.PutFactorsByNameResponse
}

func (o *PutFactorsByNameOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/guardian/factors/{name}][%d] putFactorsByNameOK  %+v", 200, o.Payload)
}

func (o *PutFactorsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFactorsByNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFactorsByNameBadRequest creates a PutFactorsByNameBadRequest with default headers values
func NewPutFactorsByNameBadRequest() *PutFactorsByNameBadRequest {
	return &PutFactorsByNameBadRequest{}
}

/*PutFactorsByNameBadRequest handles this case with default header values.

Invalid input based on schemas.
*/
type PutFactorsByNameBadRequest struct {
}

func (o *PutFactorsByNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/guardian/factors/{name}][%d] putFactorsByNameBadRequest ", 400)
}

func (o *PutFactorsByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFactorsByNameUnauthorized creates a PutFactorsByNameUnauthorized with default headers values
func NewPutFactorsByNameUnauthorized() *PutFactorsByNameUnauthorized {
	return &PutFactorsByNameUnauthorized{}
}

/*PutFactorsByNameUnauthorized handles this case with default header values.

Token has expired or signature is invalid
*/
type PutFactorsByNameUnauthorized struct {
}

func (o *PutFactorsByNameUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/guardian/factors/{name}][%d] putFactorsByNameUnauthorized ", 401)
}

func (o *PutFactorsByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFactorsByNameForbidden creates a PutFactorsByNameForbidden with default headers values
func NewPutFactorsByNameForbidden() *PutFactorsByNameForbidden {
	return &PutFactorsByNameForbidden{}
}

/*PutFactorsByNameForbidden handles this case with default header values.

Insufficient scope
*/
type PutFactorsByNameForbidden struct {
}

func (o *PutFactorsByNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/guardian/factors/{name}][%d] putFactorsByNameForbidden ", 403)
}

func (o *PutFactorsByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
