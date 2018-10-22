// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchClientsByIDBodyAws patch clients by id body aws
// swagger:model patch_clients_by_id_body_aws
type PatchClientsByIDBodyAws struct {

	// For example: 'arn:aws:iam::010616021751:saml-provider/idpname'
	Principal string `json:"principal,omitempty"`

	// For example: 'arn:aws:iam::010616021751:role/foo'
	Role string `json:"role,omitempty"`
}

// Validate validates this patch clients by id body aws
func (m *PatchClientsByIDBodyAws) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchClientsByIDBodyAws) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchClientsByIDBodyAws) UnmarshalBinary(b []byte) error {
	var res PatchClientsByIDBodyAws
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
