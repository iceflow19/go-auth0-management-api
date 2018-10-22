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

// PostClientsBodyAndroid Configuration related to Android native apps
// swagger:model post_clients_body_android
type PostClientsBodyAndroid struct {

	// Application package name found in <code>AndroidManifest.xml</code>
	AppPackageName string `json:"app_package_name,omitempty"`

	// The SHA256 fingerprints of your app's signing certificate. Multiple fingerprints can be used to support different versions of your app, such as debug and production builds
	// Min Items: 1
	Sha256CertFingerprints []string `json:"sha256_cert_fingerprints"`
}

func (m *PostClientsBodyAndroid) UnmarshalJSON(b []byte) error {
	type PostClientsBodyAndroidAlias PostClientsBodyAndroid
	var t PostClientsBodyAndroidAlias
	if err := json.Unmarshal([]byte("{\"app_package_name\":\"com.example\",\"sha256_cert_fingerprints\":[\"D8:A0:83:...\"]}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PostClientsBodyAndroid(t)
	return nil
}

// Validate validates this post clients body android
func (m *PostClientsBodyAndroid) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSha256CertFingerprints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostClientsBodyAndroid) validateSha256CertFingerprints(formats strfmt.Registry) error {

	if swag.IsZero(m.Sha256CertFingerprints) { // not required
		return nil
	}

	iSha256CertFingerprintsSize := int64(len(m.Sha256CertFingerprints))

	if err := validate.MinItems("sha256_cert_fingerprints", "body", iSha256CertFingerprintsSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsBodyAndroid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsBodyAndroid) UnmarshalBinary(b []byte) error {
	var res PostClientsBodyAndroid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
