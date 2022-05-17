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

// QueryResponseAdmissionStatus query response admission status
// swagger:model QueryResponseAdmissionStatus
type QueryResponseAdmissionStatus string

const (

	// QueryResponseAdmissionStatusConfirmed captures enum value "confirmed"
	QueryResponseAdmissionStatusConfirmed QueryResponseAdmissionStatus = "confirmed"

	// QueryResponseAdmissionStatusFailed captures enum value "failed"
	QueryResponseAdmissionStatusFailed QueryResponseAdmissionStatus = "failed"
)

// for schema
var queryResponseAdmissionStatusEnum []interface{}

func init() {
	var res []QueryResponseAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryResponseAdmissionStatusEnum = append(queryResponseAdmissionStatusEnum, v)
	}
}

func (m QueryResponseAdmissionStatus) validateQueryResponseAdmissionStatusEnum(path, location string, value QueryResponseAdmissionStatus) error {
	if err := validate.Enum(path, location, value, queryResponseAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this query response admission status
func (m QueryResponseAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQueryResponseAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}