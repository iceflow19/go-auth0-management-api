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

// PostClientGrantsBody post client grants body
// swagger:model post_client-grants_body
type PostClientGrantsBody struct {

	// The audience.
	// Required: true
	Audience *string `json:"audience"`

	// The identifier of the client.
	// Required: true
	ClientID *string `json:"client_id"`

	// scope
	// Required: true
	// Unique: true
	Scope []string `json:"scope"`
}

// Validate validates this post client grants body
func (m *PostClientGrantsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostClientGrantsBody) validateAudience(formats strfmt.Registry) error {

	if err := validate.Required("audience", "body", m.Audience); err != nil {
		return err
	}

	return nil
}

func (m *PostClientGrantsBody) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *PostClientGrantsBody) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	if err := validate.UniqueItems("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostClientGrantsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientGrantsBody) UnmarshalBinary(b []byte) error {
	var res PostClientGrantsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}