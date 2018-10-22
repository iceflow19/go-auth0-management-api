// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetUserBlocksResponse get user blocks response
// swagger:model get_user-blocks_response
type GetUserBlocksResponse struct {

	// Array of identifier + ip pairs
	BlockedFor []*GetUserBlocksResponseBlockedForItems0 `json:"blocked_for"`
}

// Validate validates this get user blocks response
func (m *GetUserBlocksResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockedFor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUserBlocksResponse) validateBlockedFor(formats strfmt.Registry) error {

	if swag.IsZero(m.BlockedFor) { // not required
		return nil
	}

	for i := 0; i < len(m.BlockedFor); i++ {
		if swag.IsZero(m.BlockedFor[i]) { // not required
			continue
		}

		if m.BlockedFor[i] != nil {
			if err := m.BlockedFor[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blocked_for" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUserBlocksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserBlocksResponse) UnmarshalBinary(b []byte) error {
	var res GetUserBlocksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetUserBlocksResponseBlockedForItems0 get user blocks response blocked for items0
// swagger:model GetUserBlocksResponseBlockedForItems0
type GetUserBlocksResponseBlockedForItems0 struct {

	// Connection identifier
	Connection string `json:"connection,omitempty"`

	// Identifier (can be the user's `email`, `username` or `phone_number`)
	Identifier *string `json:"identifier,omitempty"`

	// IP Address
	IP *string `json:"ip,omitempty"`
}

// Validate validates this get user blocks response blocked for items0
func (m *GetUserBlocksResponseBlockedForItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetUserBlocksResponseBlockedForItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserBlocksResponseBlockedForItems0) UnmarshalBinary(b []byte) error {
	var res GetUserBlocksResponseBlockedForItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}