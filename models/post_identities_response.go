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

// PostIdentitiesResponse post identities response
// swagger:model post_identities_response
type PostIdentitiesResponse struct {

	// IDP access token returned only if scope read:user_idp_token is defined
	AccessToken string `json:"access_token,omitempty"`

	// The name of the connection for the identity.
	// Required: true
	Connection *string `json:"connection"`

	// <code>true</code> if the identity provider is a social provider, <code>false</code>s otherwise
	IsSocial bool `json:"isSocial,omitempty"`

	// profile data
	ProfileData *PostIdentitiesResponseProfileData `json:"profileData,omitempty"`

	// The type of identity provider.
	// Required: true
	Provider *string `json:"provider"`

	// The unique identifier for the user for the identity.
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this post identities response
func (m *PostIdentitiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostIdentitiesResponse) validateConnection(formats strfmt.Registry) error {

	if err := validate.Required("connection", "body", m.Connection); err != nil {
		return err
	}

	return nil
}

func (m *PostIdentitiesResponse) validateProfileData(formats strfmt.Registry) error {

	if swag.IsZero(m.ProfileData) { // not required
		return nil
	}

	if m.ProfileData != nil {
		if err := m.ProfileData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profileData")
			}
			return err
		}
	}

	return nil
}

func (m *PostIdentitiesResponse) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

func (m *PostIdentitiesResponse) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostIdentitiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostIdentitiesResponse) UnmarshalBinary(b []byte) error {
	var res PostIdentitiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
