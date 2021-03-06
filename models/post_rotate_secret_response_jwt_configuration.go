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

// PostRotateSecretResponseJwtConfiguration Configuration related to JWTs for the client
// swagger:model post_rotate-secret_response_jwt_configuration
type PostRotateSecretResponseJwtConfiguration struct {

	// Algorithm uses to sign JWTs
	// Enum: [HS256 RS256]
	Alg *string `json:"alg,omitempty"`

	// The amount of seconds the JWT will be valid (affects <code>exp</code> claim)
	LifetimeInSeconds *int64 `json:"lifetime_in_seconds,omitempty"`

	// scopes
	Scopes interface{} `json:"scopes,omitempty"`

	// <code>true</code> if the client secret is base64 encoded, <code>false</code> otherwise. Defaults to <code>true</code>
	SecretEncoded *bool `json:"secret_encoded,omitempty"`
}

// Validate validates this post rotate secret response jwt configuration
func (m *PostRotateSecretResponseJwtConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postRotateSecretResponseJwtConfigurationTypeAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HS256","RS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postRotateSecretResponseJwtConfigurationTypeAlgPropEnum = append(postRotateSecretResponseJwtConfigurationTypeAlgPropEnum, v)
	}
}

const (

	// PostRotateSecretResponseJwtConfigurationAlgHS256 captures enum value "HS256"
	PostRotateSecretResponseJwtConfigurationAlgHS256 string = "HS256"

	// PostRotateSecretResponseJwtConfigurationAlgRS256 captures enum value "RS256"
	PostRotateSecretResponseJwtConfigurationAlgRS256 string = "RS256"
)

// prop value enum
func (m *PostRotateSecretResponseJwtConfiguration) validateAlgEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postRotateSecretResponseJwtConfigurationTypeAlgPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostRotateSecretResponseJwtConfiguration) validateAlg(formats strfmt.Registry) error {

	if swag.IsZero(m.Alg) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlgEnum("alg", "body", *m.Alg); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostRotateSecretResponseJwtConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostRotateSecretResponseJwtConfiguration) UnmarshalBinary(b []byte) error {
	var res PostRotateSecretResponseJwtConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
