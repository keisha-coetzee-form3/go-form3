// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DirectDebitDecisionCreation direct debit decision creation
// swagger:model DirectDebitDecisionCreation
type DirectDebitDecisionCreation struct {

	// data
	Data *DirectDebitDecision `json:"data,omitempty"`
}

func DirectDebitDecisionCreationWithDefaults(defaults client.Defaults) *DirectDebitDecisionCreation {
	return &DirectDebitDecisionCreation{

		Data: DirectDebitDecisionWithDefaults(defaults),
	}
}

func (m *DirectDebitDecisionCreation) WithData(data DirectDebitDecision) *DirectDebitDecisionCreation {

	m.Data = &data

	return m
}

func (m *DirectDebitDecisionCreation) WithoutData() *DirectDebitDecisionCreation {
	m.Data = nil
	return m
}

// Validate validates this direct debit decision creation
func (m *DirectDebitDecisionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionCreation) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionCreation) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
