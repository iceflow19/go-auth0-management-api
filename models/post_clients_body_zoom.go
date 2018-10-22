// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostClientsBodyZoom post clients body zoom
// swagger:model post_clients_body_zoom
type PostClientsBodyZoom struct {

	// The first segment of your Vanity URL, for example: 'https://{account}.zoom.us'
	Account string `json:"account,omitempty"`
}

// Validate validates this post clients body zoom
func (m *PostClientsBodyZoom) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsBodyZoom) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsBodyZoom) UnmarshalBinary(b []byte) error {
	var res PostClientsBodyZoom
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
