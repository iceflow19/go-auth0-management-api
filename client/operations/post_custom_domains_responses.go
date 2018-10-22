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

// PostCustomDomainsReader is a Reader for the PostCustomDomains structure.
type PostCustomDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCustomDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPostCustomDomainsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCustomDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostCustomDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostCustomDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostCustomDomainsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostCustomDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCustomDomainsOK creates a PostCustomDomainsOK with default headers values
func NewPostCustomDomainsOK() *PostCustomDomainsOK {
	return &PostCustomDomainsOK{}
}

/*PostCustomDomainsOK handles this case with default header values.

No response was specified
*/
type PostCustomDomainsOK struct {
	Payload *models.PostCustomDomainsResponse
}

func (o *PostCustomDomainsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsOK  %+v", 200, o.Payload)
}

func (o *PostCustomDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCustomDomainsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCustomDomainsCreated creates a PostCustomDomainsCreated with default headers values
func NewPostCustomDomainsCreated() *PostCustomDomainsCreated {
	return &PostCustomDomainsCreated{}
}

/*PostCustomDomainsCreated handles this case with default header values.

The custom domain was created (but verification is pending). See <strong>Response Class</strong> below for schema.
*/
type PostCustomDomainsCreated struct {
}

func (o *PostCustomDomainsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsCreated ", 201)
}

func (o *PostCustomDomainsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomDomainsBadRequest creates a PostCustomDomainsBadRequest with default headers values
func NewPostCustomDomainsBadRequest() *PostCustomDomainsBadRequest {
	return &PostCustomDomainsBadRequest{}
}

/*PostCustomDomainsBadRequest handles this case with default header values.

Invalid request body. The message will vary depending on the cause.
*/
type PostCustomDomainsBadRequest struct {
}

func (o *PostCustomDomainsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsBadRequest ", 400)
}

func (o *PostCustomDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomDomainsUnauthorized creates a PostCustomDomainsUnauthorized with default headers values
func NewPostCustomDomainsUnauthorized() *PostCustomDomainsUnauthorized {
	return &PostCustomDomainsUnauthorized{}
}

/*PostCustomDomainsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation.
*/
type PostCustomDomainsUnauthorized struct {
}

func (o *PostCustomDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsUnauthorized ", 401)
}

func (o *PostCustomDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomDomainsForbidden creates a PostCustomDomainsForbidden with default headers values
func NewPostCustomDomainsForbidden() *PostCustomDomainsForbidden {
	return &PostCustomDomainsForbidden{}
}

/*PostCustomDomainsForbidden handles this case with default header values.

The account is not allowed to perform this operation.
*/
type PostCustomDomainsForbidden struct {
}

func (o *PostCustomDomainsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsForbidden ", 403)
}

func (o *PostCustomDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomDomainsConflict creates a PostCustomDomainsConflict with default headers values
func NewPostCustomDomainsConflict() *PostCustomDomainsConflict {
	return &PostCustomDomainsConflict{}
}

/*PostCustomDomainsConflict handles this case with default header values.

You reached the maximum number of custom domains for your account.
*/
type PostCustomDomainsConflict struct {
}

func (o *PostCustomDomainsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsConflict ", 409)
}

func (o *PostCustomDomainsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomDomainsTooManyRequests creates a PostCustomDomainsTooManyRequests with default headers values
func NewPostCustomDomainsTooManyRequests() *PostCustomDomainsTooManyRequests {
	return &PostCustomDomainsTooManyRequests{}
}

/*PostCustomDomainsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type PostCustomDomainsTooManyRequests struct {
}

func (o *PostCustomDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/custom-domains][%d] postCustomDomainsTooManyRequests ", 429)
}

func (o *PostCustomDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}