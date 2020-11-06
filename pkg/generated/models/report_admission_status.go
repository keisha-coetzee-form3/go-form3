// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ReportAdmissionStatus report admission status
// swagger:model ReportAdmissionStatus
type ReportAdmissionStatus string

const (

	// ReportAdmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	ReportAdmissionStatusDeliveryConfirmed ReportAdmissionStatus = "delivery_confirmed"

	// ReportAdmissionStatusConfirmed captures enum value "confirmed"
	ReportAdmissionStatusConfirmed ReportAdmissionStatus = "confirmed"

	// ReportAdmissionStatusFailed captures enum value "failed"
	ReportAdmissionStatusFailed ReportAdmissionStatus = "failed"
)

// for schema
var reportAdmissionStatusEnum []interface{}

func init() {
	var res []ReportAdmissionStatus
	if err := json.Unmarshal([]byte(`["delivery_confirmed","confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportAdmissionStatusEnum = append(reportAdmissionStatusEnum, v)
	}
}

func (m ReportAdmissionStatus) validateReportAdmissionStatusEnum(path, location string, value ReportAdmissionStatus) error {
	if err := validate.Enum(path, location, value, reportAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this report admission status
func (m ReportAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReportAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
