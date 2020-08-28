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

// ReportAdmissionAttributes report admission attributes
// swagger:model ReportAdmissionAttributes
type ReportAdmissionAttributes struct {

	// admission datetime
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status ReportAdmissionStatus `json:"status,omitempty"`
}

func ReportAdmissionAttributesWithDefaults(defaults client.Defaults) *ReportAdmissionAttributes {
	return &ReportAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("ReportAdmissionAttributes", "admission_datetime"),

		SchemeStatusCode: defaults.GetString("ReportAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReportAdmissionAttributes", "scheme_status_code_description"),

		SourceGateway: defaults.GetString("ReportAdmissionAttributes", "source_gateway"),

		// TODO Status: ReportAdmissionStatus,

	}
}

func (m *ReportAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *ReportAdmissionAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *ReportAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReportAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReportAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReportAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReportAdmissionAttributes) WithSourceGateway(sourceGateway string) *ReportAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *ReportAdmissionAttributes) WithStatus(status ReportAdmissionStatus) *ReportAdmissionAttributes {

	m.Status = status

	return m
}

// Validate validates this report admission attributes
func (m *ReportAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ReportAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ReportAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
