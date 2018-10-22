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

// GetCustomDomainsReader is a Reader for the GetCustomDomains structure.
type GetCustomDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCustomDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCustomDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetCustomDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomDomainsOK creates a GetCustomDomainsOK with default headers values
func NewGetCustomDomainsOK() *GetCustomDomainsOK {
	return &GetCustomDomainsOK{}
}

/*GetCustomDomainsOK handles this case with default header values.

The custom domains configurations were retrieved. See <strong>Response Class</strong> below for schema.
*/
type GetCustomDomainsOK struct {
	Payload []*models.GetCustomDomainsResponse
}

func (o *GetCustomDomainsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/custom-domains][%d] getCustomDomainsOK  %+v", 200, o.Payload)
}

func (o *GetCustomDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomDomainsUnauthorized creates a GetCustomDomainsUnauthorized with default headers values
func NewGetCustomDomainsUnauthorized() *GetCustomDomainsUnauthorized {
	return &GetCustomDomainsUnauthorized{}
}

/*GetCustomDomainsUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type GetCustomDomainsUnauthorized struct {
}

func (o *GetCustomDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/custom-domains][%d] getCustomDomainsUnauthorized ", 401)
}

func (o *GetCustomDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomDomainsForbidden creates a GetCustomDomainsForbidden with default headers values
func NewGetCustomDomainsForbidden() *GetCustomDomainsForbidden {
	return &GetCustomDomainsForbidden{}
}

/*GetCustomDomainsForbidden handles this case with default header values.

The account is not allowed to perform this operation.
*/
type GetCustomDomainsForbidden struct {
}

func (o *GetCustomDomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/custom-domains][%d] getCustomDomainsForbidden ", 403)
}

func (o *GetCustomDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomDomainsTooManyRequests creates a GetCustomDomainsTooManyRequests with default headers values
func NewGetCustomDomainsTooManyRequests() *GetCustomDomainsTooManyRequests {
	return &GetCustomDomainsTooManyRequests{}
}

/*GetCustomDomainsTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type GetCustomDomainsTooManyRequests struct {
}

func (o *GetCustomDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/custom-domains][%d] getCustomDomainsTooManyRequests ", 429)
}

func (o *GetCustomDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}