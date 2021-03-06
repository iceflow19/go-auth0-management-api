// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetProviderResponse get provider response
// swagger:model get_provider_response
type GetProviderResponse struct {

	// credentials
	Credentials *GetProviderResponseCredentials `json:"credentials,omitempty"`

	// The default from address
	DefaultFromAddress string `json:"default_from_address,omitempty"`

	// <code>true</code> if the email provider is enabled, <code>false</code> otherwise
	Enabled *bool `json:"enabled,omitempty"`

	// The name of the email provider
	Name *string `json:"name,omitempty"`

	// settings
	Settings GetProviderResponseSettings `json:"settings,omitempty"`
}

// Validate validates this get provider response
func (m *GetProviderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProviderResponse) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProviderResponse) UnmarshalBinary(b []byte) error {
	var res GetProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
