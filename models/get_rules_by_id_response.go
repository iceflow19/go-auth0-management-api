// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetRulesByIDResponse get rules by id response
// swagger:model get_rules_by_id_response
type GetRulesByIDResponse struct {

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

// Validate validates this get rules by id response
func (m *GetRulesByIDResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetRulesByIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRulesByIDResponse) UnmarshalBinary(b []byte) error {
	var res GetRulesByIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
