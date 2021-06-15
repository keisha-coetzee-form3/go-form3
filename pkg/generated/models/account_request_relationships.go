// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountRequestRelationships account request relationships
// swagger:model AccountRequestRelationships
type AccountRequestRelationships struct {

	// account
	Account *AccountRequestRelationshipsAccount `json:"account,omitempty"`

	// account request submission
	AccountRequestSubmission *AccountRequestRelationshipsAccountRequestSubmission `json:"account_request_submission,omitempty"`

	// master account
	MasterAccount *AccountRequestRelationshipsMasterAccount `json:"master_account,omitempty"`
}

func AccountRequestRelationshipsWithDefaults(defaults client.Defaults) *AccountRequestRelationships {
	return &AccountRequestRelationships{

		Account: AccountRequestRelationshipsAccountWithDefaults(defaults),

		AccountRequestSubmission: AccountRequestRelationshipsAccountRequestSubmissionWithDefaults(defaults),

		MasterAccount: AccountRequestRelationshipsMasterAccountWithDefaults(defaults),
	}
}

func (m *AccountRequestRelationships) WithAccount(account AccountRequestRelationshipsAccount) *AccountRequestRelationships {

	m.Account = &account

	return m
}

func (m *AccountRequestRelationships) WithoutAccount() *AccountRequestRelationships {
	m.Account = nil
	return m
}

func (m *AccountRequestRelationships) WithAccountRequestSubmission(accountRequestSubmission AccountRequestRelationshipsAccountRequestSubmission) *AccountRequestRelationships {

	m.AccountRequestSubmission = &accountRequestSubmission

	return m
}

func (m *AccountRequestRelationships) WithoutAccountRequestSubmission() *AccountRequestRelationships {
	m.AccountRequestSubmission = nil
	return m
}

func (m *AccountRequestRelationships) WithMasterAccount(masterAccount AccountRequestRelationshipsMasterAccount) *AccountRequestRelationships {

	m.MasterAccount = &masterAccount

	return m
}

func (m *AccountRequestRelationships) WithoutMasterAccount() *AccountRequestRelationships {
	m.MasterAccount = nil
	return m
}

// Validate validates this account request relationships
func (m *AccountRequestRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountRequestSubmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestRelationships) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *AccountRequestRelationships) validateAccountRequestSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountRequestSubmission) { // not required
		return nil
	}

	if m.AccountRequestSubmission != nil {
		if err := m.AccountRequestSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_request_submission")
			}
			return err
		}
	}

	return nil
}

func (m *AccountRequestRelationships) validateMasterAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.MasterAccount) { // not required
		return nil
	}

	if m.MasterAccount != nil {
		if err := m.MasterAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master_account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestRelationships) UnmarshalBinary(b []byte) error {
	var res AccountRequestRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountRequestRelationshipsAccount account request relationships account
// swagger:model AccountRequestRelationshipsAccount
type AccountRequestRelationshipsAccount struct {

	// data
	Data []*AccountReference `json:"data,omitempty"`
}

func AccountRequestRelationshipsAccountWithDefaults(defaults client.Defaults) *AccountRequestRelationshipsAccount {
	return &AccountRequestRelationshipsAccount{

		Data: make([]*AccountReference, 0),
	}
}

func (m *AccountRequestRelationshipsAccount) WithData(data []*AccountReference) *AccountRequestRelationshipsAccount {

	m.Data = data

	return m
}

// Validate validates this account request relationships account
func (m *AccountRequestRelationshipsAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestRelationshipsAccount) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestRelationshipsAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestRelationshipsAccount) UnmarshalBinary(b []byte) error {
	var res AccountRequestRelationshipsAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestRelationshipsAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountRequestRelationshipsAccountRequestSubmission account request relationships account request submission
// swagger:model AccountRequestRelationshipsAccountRequestSubmission
type AccountRequestRelationshipsAccountRequestSubmission struct {

	// Array of Account Request submissions
	// Read Only: true
	Data []*AccountRequestSubmission `json:"data,omitempty"`
}

func AccountRequestRelationshipsAccountRequestSubmissionWithDefaults(defaults client.Defaults) *AccountRequestRelationshipsAccountRequestSubmission {
	return &AccountRequestRelationshipsAccountRequestSubmission{

		Data: make([]*AccountRequestSubmission, 0),
	}
}

func (m *AccountRequestRelationshipsAccountRequestSubmission) WithData(data []*AccountRequestSubmission) *AccountRequestRelationshipsAccountRequestSubmission {

	m.Data = data

	return m
}

// Validate validates this account request relationships account request submission
func (m *AccountRequestRelationshipsAccountRequestSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestRelationshipsAccountRequestSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("account_request_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestRelationshipsAccountRequestSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestRelationshipsAccountRequestSubmission) UnmarshalBinary(b []byte) error {
	var res AccountRequestRelationshipsAccountRequestSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestRelationshipsAccountRequestSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountRequestRelationshipsMasterAccount account request relationships master account
// swagger:model AccountRequestRelationshipsMasterAccount
type AccountRequestRelationshipsMasterAccount struct {

	// data
	Data []*AccountReference `json:"data,omitempty"`
}

func AccountRequestRelationshipsMasterAccountWithDefaults(defaults client.Defaults) *AccountRequestRelationshipsMasterAccount {
	return &AccountRequestRelationshipsMasterAccount{

		Data: make([]*AccountReference, 0),
	}
}

func (m *AccountRequestRelationshipsMasterAccount) WithData(data []*AccountReference) *AccountRequestRelationshipsMasterAccount {

	m.Data = data

	return m
}

// Validate validates this account request relationships master account
func (m *AccountRequestRelationshipsMasterAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountRequestRelationshipsMasterAccount) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("master_account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestRelationshipsMasterAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestRelationshipsMasterAccount) UnmarshalBinary(b []byte) error {
	var res AccountRequestRelationshipsMasterAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestRelationshipsMasterAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
