// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutTwilioResponse put twilio response
// swagger:model put_twilio_response
type PutTwilioResponse struct {

	// Twilio Authentication token
	AuthToken *string `json:"auth_token,omitempty"`

	// From number
	From *string `json:"from,omitempty"`

	// Copilot SID
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`

	// Twilio SID
	Sid *string `json:"sid,omitempty"`
}

// Validate validates this put twilio response
func (m *PutTwilioResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutTwilioResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutTwilioResponse) UnmarshalBinary(b []byte) error {
	var res PutTwilioResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
