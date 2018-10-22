// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchProviderBody patch provider body
// swagger:model patch_provider_body
type PatchProviderBody struct {

	// credentials
	Credentials *PatchProviderBodyCredentials `json:"credentials,omitempty"`

	// The default from address
	DefaultFromAddress string `json:"default_from_address,omitempty"`

	// <code>true</code> if the email provider is enabled, <code>false</code> otherwise
	Enabled bool `json:"enabled,omitempty"`

	// The name of the email provider
	// Enum: [mandrill sendgrid sparkpost ses smtp]
	Name string `json:"name,omitempty"`

	// settings
	Settings PatchProviderBodySettings `json:"settings,omitempty"`
}

// Validate validates this patch provider body
func (m *PatchProviderBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchProviderBody) validateCredentials(formats strfmt.Registry) error {

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

var patchProviderBodyTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["mandrill","sendgrid","sparkpost","ses","smtp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchProviderBodyTypeNamePropEnum = append(patchProviderBodyTypeNamePropEnum, v)
	}
}

const (

	// PatchProviderBodyNameMandrill captures enum value "mandrill"
	PatchProviderBodyNameMandrill string = "mandrill"

	// PatchProviderBodyNameSendgrid captures enum value "sendgrid"
	PatchProviderBodyNameSendgrid string = "sendgrid"

	// PatchProviderBodyNameSparkpost captures enum value "sparkpost"
	PatchProviderBodyNameSparkpost string = "sparkpost"

	// PatchProviderBodyNameSes captures enum value "ses"
	PatchProviderBodyNameSes string = "ses"

	// PatchProviderBodyNameSMTP captures enum value "smtp"
	PatchProviderBodyNameSMTP string = "smtp"
)

// prop value enum
func (m *PatchProviderBody) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patchProviderBodyTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchProviderBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchProviderBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchProviderBody) UnmarshalBinary(b []byte) error {
	var res PatchProviderBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}