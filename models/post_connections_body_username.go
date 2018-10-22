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

// PostConnectionsBodyUsername post connections body username
// swagger:model post_connections_body_username
type PostConnectionsBodyUsername struct {

	// max
	// Required: true
	// Maximum: 128
	Max *int64 `json:"max"`

	// min
	// Required: true
	// Minimum: 1
	Min *int64 `json:"min"`
}

// Validate validates this post connections body username
func (m *PostConnectionsBodyUsername) Validate(formats strfmt.Registry) error {
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

func (m *PostConnectionsBodyUsername) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	if err := validate.MaximumInt("max", "body", int64(*m.Max), 128, false); err != nil {
		return err
	}

	return nil
}

func (m *PostConnectionsBodyUsername) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	if err := validate.MinimumInt("min", "body", int64(*m.Min), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostConnectionsBodyUsername) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostConnectionsBodyUsername) UnmarshalBinary(b []byte) error {
	var res PostConnectionsBodyUsername
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
