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

// PatchSettingsBodyFlags Tenant flags
// swagger:model patch_settings_body_flags
type PatchSettingsBodyFlags struct {

	// Enables the first version of the Change Password flow. We've deprecated this option and recommending a safer flow. This flag is only for backwards compatibility
	// Enum: [false]
	ChangePwdFlowV1 bool `json:"change_pwd_flow_v1,omitempty"`

	// This flag enables the API section in the Auth0 Management Dashboard.
	EnableApisSection bool `json:"enable_apis_section,omitempty"`

	// This flag determines whether all current connections shall be enabled when a new <code>client</code> is created. Default value is <code>true</code>.
	EnableClientConnections bool `json:"enable_client_connections,omitempty"`

	// If enabled, All your email links and urls will use your configured custom domain. If no custom domain is found the email operation will fail.
	EnableCustomDomainInEmails bool `json:"enable_custom_domain_in_emails,omitempty"`

	// This flag enables dynamic client registration.
	EnableDynamicClientRegistration bool `json:"enable_dynamic_client_registration,omitempty"`

	// This flag enables advanced API Authorization scenarios.
	EnablePipeline2 bool `json:"enable_pipeline2,omitempty"`
}

// Validate validates this patch settings body flags
func (m *PatchSettingsBodyFlags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangePwdFlowV1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchSettingsBodyFlagsTypeChangePwdFlowV1PropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchSettingsBodyFlagsTypeChangePwdFlowV1PropEnum = append(patchSettingsBodyFlagsTypeChangePwdFlowV1PropEnum, v)
	}
}

// prop value enum
func (m *PatchSettingsBodyFlags) validateChangePwdFlowV1Enum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, patchSettingsBodyFlagsTypeChangePwdFlowV1PropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchSettingsBodyFlags) validateChangePwdFlowV1(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangePwdFlowV1) { // not required
		return nil
	}

	// value enum
	if err := m.validateChangePwdFlowV1Enum("change_pwd_flow_v1", "body", m.ChangePwdFlowV1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchSettingsBodyFlags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchSettingsBodyFlags) UnmarshalBinary(b []byte) error {
	var res PatchSettingsBodyFlags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
