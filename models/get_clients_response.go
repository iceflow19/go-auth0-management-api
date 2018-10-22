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

// GetClientsResponse get clients response
// swagger:model get_clients_response
type GetClientsResponse struct {

	// addons
	Addons GetClientsResponseAddons `json:"addons,omitempty"`

	// allowed clients
	AllowedClients []string `json:"allowed_clients"`

	// allowed logout urls
	AllowedLogoutUrls []string `json:"allowed_logout_urls"`

	// allowed origins
	AllowedOrigins []string `json:"allowed_origins"`

	// The type of application this client represents
	AppType string `json:"app_type,omitempty"`

	// The URLs that Auth0 can use to as a callback for the client
	Callbacks []string `json:"callbacks"`

	// client aliases
	ClientAliases []string `json:"client_aliases"`

	// The id of the client
	ClientID *string `json:"client_id,omitempty"`

	// client metadata
	ClientMetadata GetClientsResponseClientMetadata `json:"client_metadata,omitempty"`

	// The client secret, it must not be public
	ClientSecret *string `json:"client_secret,omitempty"`

	// <code>true</code> if this client can be used to make cross-origin authentication requests, <code>false</code> otherwise (default: <code>false</code>)
	CrossOriginAuth bool `json:"cross_origin_auth,omitempty"`

	// Url fo the location in your site where the cross origin verification takes place for the cross-origin auth flow when performing Auth in your own domain instead of Auth0 hosted login page.
	CrossOriginLoc string `json:"cross_origin_loc,omitempty"`

	// custom login page
	CustomLoginPage string `json:"custom_login_page,omitempty"`

	// <code>true</code> if the custom login page is to be used, <code>false</code> otherwise. Defaults to true
	CustomLoginPageOn *bool `json:"custom_login_page_on,omitempty"`

	// custom login page preview
	CustomLoginPagePreview string `json:"custom_login_page_preview,omitempty"`

	// Free text description of the purpose of the Client. (Max character length: <code>140</code>)
	Description string `json:"description,omitempty"`

	// encryption key
	EncryptionKey *GetClientsResponseEncryptionKey `json:"encryption_key,omitempty"`

	// form template
	FormTemplate string `json:"form_template,omitempty"`

	// Whether this client a first party client or not
	IsFirstParty *bool `json:"is_first_party,omitempty"`

	// jwt configuration
	JwtConfiguration *GetClientsResponseJwtConfiguration `json:"jwt_configuration,omitempty"`

	// The URL of the client logo (recommended size: 150x150)
	LogoURI string `json:"logo_uri,omitempty"`

	// mobile
	Mobile *GetClientsResponseMobile `json:"mobile,omitempty"`

	// The name of the client
	Name *string `json:"name,omitempty"`

	// Whether this client will conform to strict OIDC specifications
	OidcConformant *bool `json:"oidc_conformant,omitempty"`

	// Client signing keys.
	SigningKeys []*GetClientsResponseSigningKeysItems0 `json:"signing_keys"`

	// sso
	Sso *bool `json:"sso,omitempty"`

	// <code>true</code> to disable Single Sign On, <code>false</code> otherwise (default: <code>false</code>)
	SsoDisabled *bool `json:"sso_disabled,omitempty"`

	// Defines the requested authentication method for the token endpoint. Possible values are 'none' (public client without a client secret), 'client_secret_post' (client uses HTTP POST parameters) or 'client_secret_basic' (client uses HTTP Basic)
	// Enum: [none client_secret_post client_secret_basic]
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`

	// A set of URLs that represents valid web origins for use with web message response mode
	WebOrigins []string `json:"web_origins"`
}

// Validate validates this get clients response
func (m *GetClientsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJwtConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSigningKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenEndpointAuthMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClientsResponse) validateEncryptionKey(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionKey) { // not required
		return nil
	}

	if m.EncryptionKey != nil {
		if err := m.EncryptionKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption_key")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientsResponse) validateJwtConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.JwtConfiguration) { // not required
		return nil
	}

	if m.JwtConfiguration != nil {
		if err := m.JwtConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwt_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientsResponse) validateMobile(formats strfmt.Registry) error {

	if swag.IsZero(m.Mobile) { // not required
		return nil
	}

	if m.Mobile != nil {
		if err := m.Mobile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile")
			}
			return err
		}
	}

	return nil
}

func (m *GetClientsResponse) validateSigningKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.SigningKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.SigningKeys); i++ {
		if swag.IsZero(m.SigningKeys[i]) { // not required
			continue
		}

		if m.SigningKeys[i] != nil {
			if err := m.SigningKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signing_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getClientsResponseTypeTokenEndpointAuthMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","client_secret_post","client_secret_basic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getClientsResponseTypeTokenEndpointAuthMethodPropEnum = append(getClientsResponseTypeTokenEndpointAuthMethodPropEnum, v)
	}
}

const (

	// GetClientsResponseTokenEndpointAuthMethodNone captures enum value "none"
	GetClientsResponseTokenEndpointAuthMethodNone string = "none"

	// GetClientsResponseTokenEndpointAuthMethodClientSecretPost captures enum value "client_secret_post"
	GetClientsResponseTokenEndpointAuthMethodClientSecretPost string = "client_secret_post"

	// GetClientsResponseTokenEndpointAuthMethodClientSecretBasic captures enum value "client_secret_basic"
	GetClientsResponseTokenEndpointAuthMethodClientSecretBasic string = "client_secret_basic"
)

// prop value enum
func (m *GetClientsResponse) validateTokenEndpointAuthMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getClientsResponseTypeTokenEndpointAuthMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetClientsResponse) validateTokenEndpointAuthMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenEndpointAuthMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateTokenEndpointAuthMethodEnum("token_endpoint_auth_method", "body", *m.TokenEndpointAuthMethod); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClientsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientsResponse) UnmarshalBinary(b []byte) error {
	var res GetClientsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetClientsResponseSigningKeysItems0 get clients response signing keys items0
// swagger:model GetClientsResponseSigningKeysItems0
type GetClientsResponseSigningKeysItems0 struct {

	// Signing public key.
	Cert string `json:"cert,omitempty"`

	// Signing public key (PKCS#7 format).
	Pkcs7 string `json:"pkcs7,omitempty"`
}

// Validate validates this get clients response signing keys items0
func (m *GetClientsResponseSigningKeysItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClientsResponseSigningKeysItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClientsResponseSigningKeysItems0) UnmarshalBinary(b []byte) error {
	var res GetClientsResponseSigningKeysItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
