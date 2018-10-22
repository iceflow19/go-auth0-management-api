// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchResourceServersByIDBody patch resource servers by id body
// swagger:model patch_resource-servers_by_id_body
type PatchResourceServersByIDBody struct {

	// Allows issuance of refresh tokens for this entity
	AllowOfflineAccess bool `json:"allow_offline_access,omitempty"`

	// The name of the resource server. Must contain at least one character. Does not allow '<' or '>'
	Name string `json:"name,omitempty"`

	// options
	Options PatchResourceServersByIDBodyOptions `json:"options,omitempty"`

	// scopes
	Scopes []*PatchResourceServersByIDBodyScopesItems0 `json:"scopes"`

	// The algorithm used to sign tokens
	// Enum: [HS256 RS256]
	SigningAlg string `json:"signing_alg,omitempty"`

	// The secret used to sign tokens when using symmetric algorithms
	SigningSecret string `json:"signing_secret,omitempty"`

	// Flag this entity as capable of skipping consent
	SkipConsentForVerifiableFirstPartyClients bool `json:"skip_consent_for_verifiable_first_party_clients,omitempty"`

	// The amount of time (in seconds) that the token will be valid after being issued
	// Maximum: 2.592e+06
	// Minimum: 0
	TokenLifetime *int64 `json:"token_lifetime,omitempty"`

	// A uri from which to retrieve JWKs for this resource server used for verifying the JWT sent to Auth0 for token introspection.
	VerificationLocation string `json:"verificationLocation,omitempty"`
}

// Validate validates this patch resource servers by id body
func (m *PatchResourceServersByIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSigningAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenLifetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchResourceServersByIDBody) validateScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.Scopes) { // not required
		return nil
	}

	for i := 0; i < len(m.Scopes); i++ {
		if swag.IsZero(m.Scopes[i]) { // not required
			continue
		}

		if m.Scopes[i] != nil {
			if err := m.Scopes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var patchResourceServersByIdBodyTypeSigningAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HS256","RS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchResourceServersByIdBodyTypeSigningAlgPropEnum = append(patchResourceServersByIdBodyTypeSigningAlgPropEnum, v)
	}
}

const (

	// PatchResourceServersByIDBodySigningAlgHS256 captures enum value "HS256"
	PatchResourceServersByIDBodySigningAlgHS256 string = "HS256"

	// PatchResourceServersByIDBodySigningAlgRS256 captures enum value "RS256"
	PatchResourceServersByIDBodySigningAlgRS256 string = "RS256"
)

// prop value enum
func (m *PatchResourceServersByIDBody) validateSigningAlgEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patchResourceServersByIdBodyTypeSigningAlgPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchResourceServersByIDBody) validateSigningAlg(formats strfmt.Registry) error {

	if swag.IsZero(m.SigningAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateSigningAlgEnum("signing_alg", "body", m.SigningAlg); err != nil {
		return err
	}

	return nil
}

func (m *PatchResourceServersByIDBody) validateTokenLifetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenLifetime) { // not required
		return nil
	}

	if err := validate.MinimumInt("token_lifetime", "body", int64(*m.TokenLifetime), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("token_lifetime", "body", int64(*m.TokenLifetime), 2.592e+06, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchResourceServersByIDBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchResourceServersByIDBody) UnmarshalBinary(b []byte) error {
	var res PatchResourceServersByIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchResourceServersByIDBodyScopesItems0 patch resource servers by ID body scopes items0
// swagger:model PatchResourceServersByIDBodyScopesItems0
type PatchResourceServersByIDBodyScopesItems0 struct {

	// A user-friendly description of the scope.
	Description string `json:"description,omitempty"`

	// The scope value.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this patch resource servers by ID body scopes items0
func (m *PatchResourceServersByIDBodyScopesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchResourceServersByIDBodyScopesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchResourceServersByIDBodyScopesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchResourceServersByIDBodyScopesItems0) UnmarshalBinary(b []byte) error {
	var res PatchResourceServersByIDBodyScopesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}