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

// PostConnectionsBodyPasswordDictionary Options for password dictionary policy
// swagger:model post_connections_body_password_dictionary
type PostConnectionsBodyPasswordDictionary struct {

	// Custom Password Dictionary. An array of up to 200 entries.
	// Max Items: 200
	Dictionary []string `json:"dictionary"`

	// enable
	// Required: true
	Enable *bool `json:"enable"`
}

// Validate validates this post connections body password dictionary
func (m *PostConnectionsBodyPasswordDictionary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDictionary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostConnectionsBodyPasswordDictionary) validateDictionary(formats strfmt.Registry) error {

	if swag.IsZero(m.Dictionary) { // not required
		return nil
	}

	iDictionarySize := int64(len(m.Dictionary))

	if err := validate.MaxItems("dictionary", "body", iDictionarySize, 200); err != nil {
		return err
	}

	return nil
}

func (m *PostConnectionsBodyPasswordDictionary) validateEnable(formats strfmt.Registry) error {

	if err := validate.Required("enable", "body", m.Enable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostConnectionsBodyPasswordDictionary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostConnectionsBodyPasswordDictionary) UnmarshalBinary(b []byte) error {
	var res PostConnectionsBodyPasswordDictionary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
