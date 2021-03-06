// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchClientsByIDBodyEncryptionKey The client's encryption key
// swagger:model patch_clients_by_id_body_encryption_key
type PatchClientsByIDBodyEncryptionKey struct {

	// Encryption certificate
	Cert string `json:"cert,omitempty"`

	// Encryption public key
	Pub string `json:"pub,omitempty"`

	// Certificate subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this patch clients by id body encryption key
func (m *PatchClientsByIDBodyEncryptionKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchClientsByIDBodyEncryptionKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchClientsByIDBodyEncryptionKey) UnmarshalBinary(b []byte) error {
	var res PatchClientsByIDBodyEncryptionKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
