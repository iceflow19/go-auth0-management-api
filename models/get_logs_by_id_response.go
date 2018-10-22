// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetLogsByIDResponse get logs by id response
// swagger:model get_logs_by_id_response
type GetLogsByIDResponse struct {

	// The id of the client
	ClientID *string `json:"client_id,omitempty"`

	// The name of the client
	ClientName *string `json:"client_name,omitempty"`

	// The date when the event was created
	Date *string `json:"date,omitempty"`

	// details
	Details GetLogsByIDResponseDetails `json:"details,omitempty"`

	// The IP of the log event source
	IP *string `json:"ip,omitempty"`

	// location info
	LocationInfo GetLogsByIDResponseLocationInfo `json:"location_info,omitempty"`

	// The log event type
	Type *string `json:"type,omitempty"`

	// The user's unique identifier
	UserID *string `json:"user_id,omitempty"`
}

// Validate validates this get logs by id response
func (m *GetLogsByIDResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetLogsByIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogsByIDResponse) UnmarshalBinary(b []byte) error {
	var res GetLogsByIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
