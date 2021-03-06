// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostTicketResponse post ticket response
// swagger:model post_ticket_response
type PostTicketResponse struct {

	// The ticket_id used to identify the enrollment
	TicketID *string `json:"ticket_id,omitempty"`

	// The url you can use to start enrollment
	TicketURL *string `json:"ticket_url,omitempty"`
}

// Validate validates this post ticket response
func (m *PostTicketResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostTicketResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTicketResponse) UnmarshalBinary(b []byte) error {
	var res PostTicketResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
