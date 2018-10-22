// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetProviderResponseCredentials Provider credentials
// swagger:model get_provider_response_credentials
type GetProviderResponseCredentials struct {

	// AWS Access Key ID
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// API Key
	APIKey *string `json:"api_key,omitempty"`

	// API User
	APIUser string `json:"api_user,omitempty"`

	// AWS default region
	Region string `json:"region,omitempty"`

	// AWS Secret Access Key
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// SMTP host
	SMTPHost string `json:"smtp_host,omitempty"`

	// SMTP password
	SMTPPass string `json:"smtp_pass,omitempty"`

	// SMTP port
	SMTPPort int64 `json:"smtp_port,omitempty"`

	// SMTP user
	SMTPUser string `json:"smtp_user,omitempty"`
}

// Validate validates this get provider response credentials
func (m *GetProviderResponseCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetProviderResponseCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProviderResponseCredentials) UnmarshalBinary(b []byte) error {
	var res GetProviderResponseCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}