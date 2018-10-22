// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostClientsBodyAzureSb post clients body azure sb
// swagger:model post_clients_body_azure_sb
type PostClientsBodyAzureSb struct {

	// The entity you want to request a token for, for example: 'my-queue'
	EntityPath string `json:"entityPath,omitempty"`

	// Expiration in minutes for the generated token. Default: 5 minutes
	Expiration int64 `json:"expiration,omitempty"`

	// Your Azure Service Bus namespace. The first segment of the Service Bus URL: http://YOUR_NAMESPACE.servicebus.windows.net
	Namespace string `json:"namespace,omitempty"`

	// The Primary Key associated with the shared access policy
	SasKey string `json:"sasKey,omitempty"`

	// The shared access policy name defined in your Service Bus entity
	SasKeyName string `json:"sasKeyName,omitempty"`
}

// Validate validates this post clients body azure sb
func (m *PostClientsBodyAzureSb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostClientsBodyAzureSb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostClientsBodyAzureSb) UnmarshalBinary(b []byte) error {
	var res PostClientsBodyAzureSb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}