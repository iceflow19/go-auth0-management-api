// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchRulesByIDResponse patch rules by id response
// swagger:model patch_rules_by_id_response
type PatchRulesByIDResponse struct {

	// <code>true</code> if the connection is enabled, <code>false</code> otherwise
	Enabled *bool `json:"enabled,omitempty"`

	// The rule's identifier
	ID *string `json:"id,omitempty"`

	// The name of the rule
	Name *string `json:"name,omitempty"`

	// The rule's order in relation to other rules. A rule with a lower order than another rule executes first
	Order *float64 `json:"order,omitempty"`

	// The code to be executed when the rule runs
	Script *string `json:"script,omitempty"`

	// The rule's execution stage
	Stage *string `json:"stage,omitempty"`
}

// Validate validates this patch rules by id response
func (m *PatchRulesByIDResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchRulesByIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchRulesByIDResponse) UnmarshalBinary(b []byte) error {
	var res PatchRulesByIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
