// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DirectDebitAdmissionRelationships direct debit admission relationships
// swagger:model DirectDebitAdmissionRelationships
type DirectDebitAdmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitAdmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`
}

func DirectDebitAdmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitAdmissionRelationships {
	return &DirectDebitAdmissionRelationships{

		DirectDebit: DirectDebitAdmissionRelationshipsDirectDebitWithDefaults(defaults),
	}
}

func (m *DirectDebitAdmissionRelationships) WithDirectDebit(directDebit DirectDebitAdmissionRelationshipsDirectDebit) *DirectDebitAdmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitAdmissionRelationships) WithoutDirectDebit() *DirectDebitAdmissionRelationships {
	m.DirectDebit = nil
	return m
}

// Validate validates this direct debit admission relationships
func (m *DirectDebitAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAdmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitAdmissionRelationshipsDirectDebit direct debit admission relationships direct debit
// swagger:model DirectDebitAdmissionRelationshipsDirectDebit
type DirectDebitAdmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitAdmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitAdmissionRelationshipsDirectDebit {
	return &DirectDebitAdmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitAdmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitAdmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit admission relationships direct debit
func (m *DirectDebitAdmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAdmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitAdmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitAdmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitAdmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitAdmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
