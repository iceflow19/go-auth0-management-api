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

// PatchClientsByIDBodyLayer patch clients by id body layer
// swagger:model patch_clients_by_id_body_layer
type PatchClientsByIDBodyLayer struct {

	// An optional expiration in minutes for the generated token. By default it will be 5 minutes
	// Minimum: 0
	Expiration *int64 `json:"expiration,omitempty"`

	// The Authentication Key identifier used to sign the Layer token
	// Required: true
	KeyID *string `json:"keyId"`

	// The name of the property used as the unique user id in Layer. If none is specified user_id will be used
	Principal string `json:"principal,omitempty"`

	// The private key for signing the Layer token
	// Required: true
	PrivateKey *string `json:"privateKey"`

	// The Provider ID of your Layer account
	// Required: true
	ProviderID *string `json:"providerId"`
}

// Validate validates this patch clients by id body layer
func (m *PatchClientsByIDBodyLayer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchClientsByIDBodyLayer) validateExpiration(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.MinimumInt("expiration", "body", int64(*m.Expiration), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PatchClientsByIDBodyLayer) validateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("keyId", "body", m.KeyID); err != nil {
		return err
	}

	return nil
}

func (m *PatchClientsByIDBodyLayer) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *PatchClientsByIDBodyLayer) validateProviderID(formats strfmt.Registry) error {

	if err := validate.Required("providerId", "body", m.ProviderID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchClientsByIDBodyLayer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchClientsByIDBodyLayer) UnmarshalBinary(b []byte) error {
	var res PatchClientsByIDBodyLayer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}