// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountRequestSubmissionAttributes account request submission attributes
// swagger:model AccountRequestSubmissionAttributes
type AccountRequestSubmissionAttributes struct {

	// status
	Status SubmissionStatus `json:"status,omitempty"`

	// Description of the submission status
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func AccountRequestSubmissionAttributesWithDefaults(defaults client.Defaults) *AccountRequestSubmissionAttributes {
	return &AccountRequestSubmissionAttributes{

		// TODO Status: SubmissionStatus,

		StatusReason: defaults.GetString("AccountRequestSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("AccountRequestSubmissionAttributes", "submission_datetime"),
	}
}

func (m *AccountRequestSubmissionAttributes) WithStatus(status SubmissionStatus) *AccountRequestSubmissionAttributes {

	m.Status = status

	return m
}

func (m *AccountRequestSubmissionAttributes) WithStatusReason(statusReason string) *AccountRequestSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *AccountRequestSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *AccountRequestSubmissionAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *AccountRequestSubmissionAttributes) WithoutSubmissionDatetime() *AccountRequestSubmissionAttributes {
	m.SubmissionDatetime = nil
	return m
}

// Validate validates this account request submission attributes
func (m *AccountRequestSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *AccountRequestSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res AccountRequestSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
