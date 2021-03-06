// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProviderReader is a Reader for the DeleteProvider structure.
type DeleteProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProviderOK creates a DeleteProviderOK with default headers values
func NewDeleteProviderOK() *DeleteProviderOK {
	return &DeleteProviderOK{}
}

/*DeleteProviderOK handles this case with default header values.

No response was specified
*/
type DeleteProviderOK struct {
}

func (o *DeleteProviderOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/emails/provider][%d] deleteProviderOK ", 200)
}

func (o *DeleteProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
