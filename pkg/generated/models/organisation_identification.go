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
	"github.com/go-openapi/validate"
)

// OrganisationIdentification organisation identification
// swagger:model OrganisationIdentification
type OrganisationIdentification struct {

	// actors
	Actors []*OrganisationIdentificationActor `json:"actors"`

	// The street name and house number of the postal address of the account holding organisation.
	Address []string `json:"address"`

	// The city where address of the account holding organisation is located.
	// Max Length: 35
	// Min Length: 1
	City string `json:"city,omitempty"`

	// The country where address of the account holding organisation is located. Must be ISO 3166-1 code.
	// Pattern: ^[A-Z]{2}$
	Country string `json:"country,omitempty"`

	// The registration number used to identify the account holding organisation.
	// Max Length: 140
	// Min Length: 1
	Identification string `json:"identification,omitempty"`

	// ISO 3166-1 code used to identify the domicile of the account
	// Pattern: ^[A-Z]{2}$
	TaxResidency string `json:"tax_residency,omitempty"`
}

func OrganisationIdentificationWithDefaults(defaults client.Defaults) *OrganisationIdentification {
	return &OrganisationIdentification{

		Actors: make([]*OrganisationIdentificationActor, 0),

		Address: make([]string, 0),

		City: defaults.GetString("OrganisationIdentification", "city"),

		Country: defaults.GetString("OrganisationIdentification", "country"),

		Identification: defaults.GetString("OrganisationIdentification", "identification"),

		TaxResidency: defaults.GetString("OrganisationIdentification", "tax_residency"),
	}
}

func (m *OrganisationIdentification) WithActors(actors []*OrganisationIdentificationActor) *OrganisationIdentification {

	m.Actors = actors

	return m
}

func (m *OrganisationIdentification) WithAddress(address []string) *OrganisationIdentification {

	m.Address = address

	return m
}

func (m *OrganisationIdentification) WithCity(city string) *OrganisationIdentification {

	m.City = city

	return m
}

func (m *OrganisationIdentification) WithCountry(country string) *OrganisationIdentification {

	m.Country = country

	return m
}

func (m *OrganisationIdentification) WithIdentification(identification string) *OrganisationIdentification {

	m.Identification = identification

	return m
}

func (m *OrganisationIdentification) WithTaxResidency(taxResidency string) *OrganisationIdentification {

	m.TaxResidency = taxResidency

	return m
}

// Validate validates this organisation identification
func (m *OrganisationIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxResidency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganisationIdentification) validateActors(formats strfmt.Registry) error {

	if swag.IsZero(m.Actors) { // not required
		return nil
	}

	for i := 0; i < len(m.Actors); i++ {
		if swag.IsZero(m.Actors[i]) { // not required
			continue
		}

		if m.Actors[i] != nil {
			if err := m.Actors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganisationIdentification) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	for i := 0; i < len(m.Address); i++ {

		if err := validate.MinLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("address"+"."+strconv.Itoa(i), "body", string(m.Address[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *OrganisationIdentification) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", string(m.City), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(m.City), 35); err != nil {
		return err
	}

	return nil
}

func (m *OrganisationIdentification) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.Pattern("country", "body", string(m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *OrganisationIdentification) validateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if err := validate.MinLength("identification", "body", string(m.Identification), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("identification", "body", string(m.Identification), 140); err != nil {
		return err
	}

	return nil
}

func (m *OrganisationIdentification) validateTaxResidency(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxResidency) { // not required
		return nil
	}

	if err := validate.Pattern("tax_residency", "body", string(m.TaxResidency), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganisationIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganisationIdentification) UnmarshalBinary(b []byte) error {
	var res OrganisationIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *OrganisationIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}