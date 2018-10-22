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

// PostClientsBodyMobile Configuration related to native mobile apps
// swagger:model post_clients_body_mobile
type PostClientsBodyMobile struct {

	// android
	Android *PostClientsBodyMobileAndroid `json:"android,omitempty"`

	// ios
	Ios *PostClientsBodyMobileIos `json:"ios,omitempty"`
}

// Validate validates this post clients body mobile
func (m *PostClientsBodyMobile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAndroid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostClientsBodyMobile) validateAndroid(formats strfmt.Registry) error {

	if swag.IsZero(m.Android) { // not required
		return nil
	}

	if m.Android != nil {
		if err := m.Android.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("android")
			}
			return err
		}
	}

	return nil
}

func (m *PostClientsBodyMobile) validateIos(formats strfmt.Registry) error {

	if swag.IsZero(m.Ios) { // not required
		return nil
	}

	if m.Ios != nil {
		if err := m.Ios.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ios")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsBodyMobile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsBodyMobile) UnmarshalBinary(b []byte) error {
	var res PostClientsBodyMobile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostClientsBodyMobileAndroid Configuration related to Android native apps
// swagger:model PostClientsBodyMobileAndroid
type PostClientsBodyMobileAndroid struct {

	// Application package name found in <code>AndroidManifest.xml</code>
	AppPackageName string `json:"app_package_name,omitempty"`

	// The SHA256 fingerprints of your app's signing certificate. Multiple fingerprints can be used to support different versions of your app, such as debug and production builds
	// Min Items: 1
	Sha256CertFingerprints []string `json:"sha256_cert_fingerprints"`
}

func (m *PostClientsBodyMobileAndroid) UnmarshalJSON(b []byte) error {
	type PostClientsBodyMobileAndroidAlias PostClientsBodyMobileAndroid
	var t PostClientsBodyMobileAndroidAlias
	if err := json.Unmarshal([]byte("{\"app_package_name\":\"com.example\",\"sha256_cert_fingerprints\":[\"D8:A0:83:...\"]}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PostClientsBodyMobileAndroid(t)
	return nil
}

// Validate validates this post clients body mobile android
func (m *PostClientsBodyMobileAndroid) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSha256CertFingerprints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostClientsBodyMobileAndroid) validateSha256CertFingerprints(formats strfmt.Registry) error {

	if swag.IsZero(m.Sha256CertFingerprints) { // not required
		return nil
	}

	iSha256CertFingerprintsSize := int64(len(m.Sha256CertFingerprints))

	if err := validate.MinItems("android"+"."+"sha256_cert_fingerprints", "body", iSha256CertFingerprintsSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsBodyMobileAndroid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsBodyMobileAndroid) UnmarshalBinary(b []byte) error {
	var res PostClientsBodyMobileAndroid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostClientsBodyMobileIos Configuration related to iOS native apps
// swagger:model PostClientsBodyMobileIos
type PostClientsBodyMobileIos struct {

	// Assigned by the developer to the app as its unique identifier inside the store, usually is a reverse domain plus the app name: <code>com.you.MyApp</code>
	AppBundleIdentifier string `json:"app_bundle_identifier,omitempty"`

	// Identifier assigned to the account that signs and upload the app to the store
	TeamID string `json:"team_id,omitempty"`
}

func (m *PostClientsBodyMobileIos) UnmarshalJSON(b []byte) error {
	type PostClientsBodyMobileIosAlias PostClientsBodyMobileIos
	var t PostClientsBodyMobileIosAlias
	if err := json.Unmarshal([]byte("{\"app_bundle_identifier\":\"com.my.bundle.id\",\"team_id\":\"9JA89QQLNQ\"}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PostClientsBodyMobileIos(t)
	return nil
}

// Validate validates this post clients body mobile ios
func (m *PostClientsBodyMobileIos) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsBodyMobileIos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsBodyMobileIos) UnmarshalBinary(b []byte) error {
	var res PostClientsBodyMobileIos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
