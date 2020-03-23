// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PartyAccountAttributesAccountWith party account attributes account with
// swagger:model PartyAccountAttributesAccountWith
type PartyAccountAttributesAccountWith struct {

	// bank id
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	BankIDCode string `json:"bank_id_code,omitempty"`

	// bank name
	BankName string `json:"bank_name,omitempty"`
}

func PartyAccountAttributesAccountWithWithDefaults(defaults client.Defaults) *PartyAccountAttributesAccountWith {
	return &PartyAccountAttributesAccountWith{

		BankID: defaults.GetString("PartyAccountAttributesAccountWith", "bank_id"),

		BankIDCode: defaults.GetString("PartyAccountAttributesAccountWith", "bank_id_code"),

		BankName: defaults.GetString("PartyAccountAttributesAccountWith", "bank_name"),
	}
}

func (m *PartyAccountAttributesAccountWith) WithBankID(bankID string) *PartyAccountAttributesAccountWith {

	m.BankID = bankID

	return m
}

func (m *PartyAccountAttributesAccountWith) WithBankIDCode(bankIDCode string) *PartyAccountAttributesAccountWith {

	m.BankIDCode = bankIDCode

	return m
}

func (m *PartyAccountAttributesAccountWith) WithBankName(bankName string) *PartyAccountAttributesAccountWith {

	m.BankName = bankName

	return m
}

// Validate validates this party account attributes account with
func (m *PartyAccountAttributesAccountWith) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartyAccountAttributesAccountWith) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAccountAttributesAccountWith) UnmarshalBinary(b []byte) error {
	var res PartyAccountAttributesAccountWith
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAccountAttributesAccountWith) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
