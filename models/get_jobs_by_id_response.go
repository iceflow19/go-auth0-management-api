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

// GetJobsByIDResponse get jobs by id response
// swagger:model get_jobs_by_id_response
type GetJobsByIDResponse struct {

	// The id of the connection.
	ConnectionID *string `json:"connection_id,omitempty"`

	// The date when the job was created.
	CreatedAt string `json:"created_at,omitempty"`

	// The job's identifier. Useful to retrieve its status
	// Required: true
	ID *string `json:"id"`

	// The url to download the result of the job.
	Location string `json:"location,omitempty"`

	// The percentage of the work done so far.
	PercentageDone int64 `json:"percentage_done,omitempty"`

	// The job's status
	// Required: true
	Status *string `json:"status"`

	// Estimated amount of time remaining to finish the job.
	TimeLeftSeconds int64 `json:"time_left_seconds,omitempty"`

	// The type of job
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this get jobs by id response
func (m *GetJobsByIDResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetJobsByIDResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetJobsByIDResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *GetJobsByIDResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetJobsByIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetJobsByIDResponse) UnmarshalBinary(b []byte) error {
	var res GetJobsByIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
