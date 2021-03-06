// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDailyResponse get daily response
// swagger:model get_daily_response
type GetDailyResponse struct {

	// The date to which the stats belong
	Date *string `json:"date,omitempty"`

	// The amount of logins on the date
	Logins *float64 `json:"logins,omitempty"`
}

// Validate validates this get daily response
func (m *GetDailyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetDailyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDailyResponse) UnmarshalBinary(b []byte) error {
	var res GetDailyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
