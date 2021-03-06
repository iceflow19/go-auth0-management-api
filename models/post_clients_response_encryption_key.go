// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostClientsResponseEncryptionKey post clients response encryption key
// swagger:model post_clients_response_encryption_key
type PostClientsResponseEncryptionKey struct {

	// cert
	Cert string `json:"cert,omitempty"`

	// pub
	Pub string `json:"pub,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this post clients response encryption key
func (m *PostClientsResponseEncryptionKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsResponseEncryptionKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsResponseEncryptionKey) UnmarshalBinary(b []byte) error {
	var res PostClientsResponseEncryptionKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
