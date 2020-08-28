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

// MandateAttributesDebtorParty mandate attributes debtor party
// swagger:model MandateAttributesDebtorParty
type MandateAttributesDebtorParty struct {

	// account name
	AccountName string `json:"account_name"`

	// account number
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// address
	Address []string `json:"address"`

	// country
	Country string `json:"country,omitempty"`
}

func MandateAttributesDebtorPartyWithDefaults(defaults client.Defaults) *MandateAttributesDebtorParty {
	return &MandateAttributesDebtorParty{

		AccountName: defaults.GetString("MandateAttributesDebtorParty", "account_name"),

		AccountNumber: defaults.GetString("MandateAttributesDebtorParty", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		AccountWith: AccountHoldingEntityWithDefaults(defaults),

		Address: make([]string, 0),

		Country: defaults.GetString("MandateAttributesDebtorParty", "country"),
	}
}

func (m *MandateAttributesDebtorParty) WithAccountName(accountName string) *MandateAttributesDebtorParty {

	m.AccountName = accountName

	return m
}

func (m *MandateAttributesDebtorParty) WithAccountNumber(accountNumber string) *MandateAttributesDebtorParty {

	m.AccountNumber = accountNumber

	return m
}

func (m *MandateAttributesDebtorParty) WithAccountNumberCode(accountNumberCode AccountNumberCode) *MandateAttributesDebtorParty {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *MandateAttributesDebtorParty) WithAccountWith(accountWith AccountHoldingEntity) *MandateAttributesDebtorParty {

	m.AccountWith = &accountWith

	return m
}

func (m *MandateAttributesDebtorParty) WithoutAccountWith() *MandateAttributesDebtorParty {
	m.AccountWith = nil
	return m
}

func (m *MandateAttributesDebtorParty) WithAddress(address []string) *MandateAttributesDebtorParty {

	m.Address = address

	return m
}

func (m *MandateAttributesDebtorParty) WithCountry(country string) *MandateAttributesDebtorParty {

	m.Country = country

	return m
}

// Validate validates this mandate attributes debtor party
func (m *MandateAttributesDebtorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAttributesDebtorParty) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateAttributesDebtorParty) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_number_code")
		}
		return err
	}

	return nil
}

func (m *MandateAttributesDebtorParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_with")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAttributesDebtorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAttributesDebtorParty) UnmarshalBinary(b []byte) error {
	var res MandateAttributesDebtorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAttributesDebtorParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
