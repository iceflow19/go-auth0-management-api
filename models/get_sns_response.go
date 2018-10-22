// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSnsResponse get sns response
// swagger:model get_sns_response
type GetSnsResponse struct {

	// aws access key id
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// aws region
	AwsRegion *string `json:"aws_region,omitempty"`

	// aws secret access key
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`

	// sns apns platform application arn
	SnsApnsPlatformApplicationArn string `json:"sns_apns_platform_application_arn,omitempty"`

	// sns gcm platform application arn
	SnsGcmPlatformApplicationArn *string `json:"sns_gcm_platform_application_arn,omitempty"`
}

// Validate validates this get sns response
func (m *GetSnsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSnsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSnsResponse) UnmarshalBinary(b []byte) error {
	var res GetSnsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
