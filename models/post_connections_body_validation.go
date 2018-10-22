// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostConnectionsBodyValidation Options for validation
// swagger:model post_connections_body_validation
type PostConnectionsBodyValidation struct {

	// username
	Username *PostConnectionsBodyValidationUsername `json:"username,omitempty"`
}

// Validate validates this post connections body validation
func (m *PostConnectionsBodyValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostConnectionsBodyValidation) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if m.Username != nil {
		if err := m.Username.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("username")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostConnectionsBodyValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostConnectionsBodyValidation) UnmarshalBinary(b []byte) error {
	var res PostConnectionsBodyValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostConnectionsBodyValidationUsername post connections body validation username
// swagger:model PostConnectionsBodyValidationUsername
type PostConnectionsBodyValidationUsername struct {

	// max
	// Required: true
	// Maximum: 128
	Max *int64 `json:"max"`

	// min
	// Required: true
	// Minimum: 1
	Min *int64 `json:"min"`
}

// Validate validates this post connections body validation username
func (m *PostConnectionsBodyValidationUsername) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostConnectionsBodyValidationUsername) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("username"+"."+"max", "body", m.Max); err != nil {
		return err
	}

	if err := validate.MaximumInt("username"+"."+"max", "body", int64(*m.Max), 128, false); err != nil {
		return err
	}

	return nil
}

func (m *PostConnectionsBodyValidationUsername) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("username"+"."+"min", "body", m.Min); err != nil {
		return err
	}

	if err := validate.MinimumInt("username"+"."+"min", "body", int64(*m.Min), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostConnectionsBodyValidationUsername) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostConnectionsBodyValidationUsername) UnmarshalBinary(b []byte) error {
	var res PostConnectionsBodyValidationUsername
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}