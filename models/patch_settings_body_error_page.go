// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchSettingsBodyErrorPage Tenant error page customization
// swagger:model patch_settings_body_error_page
type PatchSettingsBodyErrorPage struct {

	// Replace default error page with a custom HTML (<a href='https://github.com/Shopify/liquid/wiki/Liquid-for-Designers' target='_new'>Liquid syntax</a> is supported)
	HTML string `json:"html,omitempty"`

	// <code>true</code> to show link to log as part of the default error page, <code>false</code> otherwise (default: <code>true</code>)
	ShowLogLink *bool `json:"show_log_link,omitempty"`

	// Redirect to specified url instead of show the default error page
	URL *string `json:"url,omitempty"`
}

// Validate validates this patch settings body error page
func (m *PatchSettingsBodyErrorPage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchSettingsBodyErrorPage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchSettingsBodyErrorPage) UnmarshalBinary(b []byte) error {
	var res PatchSettingsBodyErrorPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}