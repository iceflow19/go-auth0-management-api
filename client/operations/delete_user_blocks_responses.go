// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserBlocksReader is a Reader for the DeleteUserBlocks structure.
type DeleteUserBlocksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserBlocksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUserBlocksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteUserBlocksNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserBlocksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteUserBlocksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteUserBlocksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewDeleteUserBlocksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserBlocksOK creates a DeleteUserBlocksOK with default headers values
func NewDeleteUserBlocksOK() *DeleteUserBlocksOK {
	return &DeleteUserBlocksOK{}
}

/*DeleteUserBlocksOK handles this case with default header values.

No response was specified
*/
type DeleteUserBlocksOK struct {
}

func (o *DeleteUserBlocksOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/user-blocks][%d] deleteUserBlocksOK ", 200)
}

func (o *DeleteUserBlocksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBlocksNoContent creates a DeleteUserBlocksNoContent with default headers values
func NewDeleteUserBlocksNoContent() *DeleteUserBlocksNoContent {
	return &DeleteUserBlocksNoContent{}
}

/*DeleteUserBlocksNoContent handles this case with default header values.

The user was unblocked. See <strong>Response Class</strong> below for schema.
*/
type DeleteUserBlocksNoContent struct {
}

func (o *DeleteUserBlocksNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/user-blocks][%d] deleteUserBlocksNoContent ", 204)
}

func (o *DeleteUserBlocksNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBlocksBadRequest creates a DeleteUserBlocksBadRequest with default headers values
func NewDeleteUserBlocksBadRequest() *DeleteUserBlocksBadRequest {
	return &DeleteUserBlocksBadRequest{}
}

/*DeleteUserBlocksBadRequest handles this case with default header values.

Invalid request body. The message will vary depending on the cause.
*/
type DeleteUserBlocksBadRequest struct {
}

func (o *DeleteUserBlocksBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/user-blocks][%d] deleteUserBlocksBadRequest ", 400)
}

func (o *DeleteUserBlocksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBlocksUnauthorized creates a DeleteUserBlocksUnauthorized with default headers values
func NewDeleteUserBlocksUnauthorized() *DeleteUserBlocksUnauthorized {
	return &DeleteUserBlocksUnauthorized{}
}

/*DeleteUserBlocksUnauthorized handles this case with default header values.

Invalid signature received for JSON Web Token validation
*/
type DeleteUserBlocksUnauthorized struct {
}

func (o *DeleteUserBlocksUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/user-blocks][%d] deleteUserBlocksUnauthorized ", 401)
}

func (o *DeleteUserBlocksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBlocksForbidden creates a DeleteUserBlocksForbidden with default headers values
func NewDeleteUserBlocksForbidden() *DeleteUserBlocksForbidden {
	return &DeleteUserBlocksForbidden{}
}

/*DeleteUserBlocksForbidden handles this case with default header values.

Insufficient scope, expected any of: update:users
*/
type DeleteUserBlocksForbidden struct {
}

func (o *DeleteUserBlocksForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/user-blocks][%d] deleteUserBlocksForbidden ", 403)
}

func (o *DeleteUserBlocksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBlocksTooManyRequests creates a DeleteUserBlocksTooManyRequests with default headers values
func NewDeleteUserBlocksTooManyRequests() *DeleteUserBlocksTooManyRequests {
	return &DeleteUserBlocksTooManyRequests{}
}

/*DeleteUserBlocksTooManyRequests handles this case with default header values.

Too many requests. Check the X-RateLimit-Limit, X-RateLimit-Remaining and X-RateLimit-Reset headers.
*/
type DeleteUserBlocksTooManyRequests struct {
}

func (o *DeleteUserBlocksTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/user-blocks][%d] deleteUserBlocksTooManyRequests ", 429)
}

func (o *DeleteUserBlocksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
