// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutTemplatesResponse put templates response
// swagger:model put_templates_response
type PutTemplatesResponse struct {

	// Message sent to the user when they are invited to enroll with a phone number
	// Required: true
	EnrollmentMessage *string `json:"enrollment_message"`

	// Message sent to the user when they are prompted to verify their account
	// Required: true
	VerificationMessage *string `json:"verification_message"`
}

// Validate validates this put templates response
func (m *PutTemplatesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrollmentMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerificationMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutTemplatesResponse) validateEnrollmentMessage(formats strfmt.Registry) error {

	if err := validate.Required("enrollment_message", "body", m.EnrollmentMessage); err != nil {
		return err
	}

	return nil
}

func (m *PutTemplatesResponse) validateVerificationMessage(formats strfmt.Registry) error {

	if err := validate.Required("verification_message", "body", m.VerificationMessage); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutTemplatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutTemplatesResponse) UnmarshalBinary(b []byte) error {
	var res PutTemplatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}