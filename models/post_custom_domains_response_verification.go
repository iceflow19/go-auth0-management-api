// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCustomDomainsResponseVerification The custom domain verification settings.
// swagger:model post_custom-domains_response_verification
type PostCustomDomainsResponseVerification struct {

	// The custom domain verification methods.
	// Min Items: 1
	Methods []*PostCustomDomainsResponseVerificationMethodsItems0 `json:"methods"`
}

// Validate validates this post custom domains response verification
func (m *PostCustomDomainsResponseVerification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCustomDomainsResponseVerification) validateMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.Methods) { // not required
		return nil
	}

	iMethodsSize := int64(len(m.Methods))

	if err := validate.MinItems("methods", "body", iMethodsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Methods); i++ {
		if swag.IsZero(m.Methods[i]) { // not required
			continue
		}

		if m.Methods[i] != nil {
			if err := m.Methods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCustomDomainsResponseVerification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCustomDomainsResponseVerification) UnmarshalBinary(b []byte) error {
	var res PostCustomDomainsResponseVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostCustomDomainsResponseVerificationMethodsItems0 post custom domains response verification methods items0
// swagger:model PostCustomDomainsResponseVerificationMethodsItems0
type PostCustomDomainsResponseVerificationMethodsItems0 struct {

	// The custom domain verification method.
	// Required: true
	// Enum: [cname txt]
	Name *string `json:"name"`

	// The value used to verify the custom domain.
	// Required: true
	Record *string `json:"record"`
}

// Validate validates this post custom domains response verification methods items0
func (m *PostCustomDomainsResponseVerificationMethodsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postCustomDomainsResponseVerificationMethodsItems0TypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cname","txt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCustomDomainsResponseVerificationMethodsItems0TypeNamePropEnum = append(postCustomDomainsResponseVerificationMethodsItems0TypeNamePropEnum, v)
	}
}

const (

	// PostCustomDomainsResponseVerificationMethodsItems0NameCname captures enum value "cname"
	PostCustomDomainsResponseVerificationMethodsItems0NameCname string = "cname"

	// PostCustomDomainsResponseVerificationMethodsItems0NameTxt captures enum value "txt"
	PostCustomDomainsResponseVerificationMethodsItems0NameTxt string = "txt"
)

// prop value enum
func (m *PostCustomDomainsResponseVerificationMethodsItems0) validateNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCustomDomainsResponseVerificationMethodsItems0TypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostCustomDomainsResponseVerificationMethodsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostCustomDomainsResponseVerificationMethodsItems0) validateRecord(formats strfmt.Registry) error {

	if err := validate.Required("record", "body", m.Record); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCustomDomainsResponseVerificationMethodsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCustomDomainsResponseVerificationMethodsItems0) UnmarshalBinary(b []byte) error {
	var res PostCustomDomainsResponseVerificationMethodsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
