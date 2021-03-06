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

// PatchClientGrantsByIDBody patch client grants by id body
// swagger:model patch_client-grants_by_id_body
type PatchClientGrantsByIDBody struct {

	// scope
	// Unique: true
	Scope []string `json:"scope"`
}

// Validate validates this patch client grants by id body
func (m *PatchClientGrantsByIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchClientGrantsByIDBody) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := validate.UniqueItems("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchClientGrantsByIDBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchClientGrantsByIDBody) UnmarshalBinary(b []byte) error {
	var res PatchClientGrantsByIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
