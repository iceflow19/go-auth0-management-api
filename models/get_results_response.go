// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetResultsResponse get results response
// swagger:model get_results_response
type GetResultsResponse struct {

	// Provided user email
	Email string `json:"email,omitempty"`

	// Whether user was found
	Exist bool `json:"exist,omitempty"`

	// Whether password matched the one provided
	Matched bool `json:"matched,omitempty"`

	// Provided username
	Username string `json:"username,omitempty"`
}

// Validate validates this get results response
func (m *GetResultsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetResultsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetResultsResponse) UnmarshalBinary(b []byte) error {
	var res GetResultsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}