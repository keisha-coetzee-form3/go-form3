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

// AccountAmendmentAttributes account amendment attributes
// swagger:model AccountAmendmentAttributes
type AccountAmendmentAttributes struct {

	// A free-format reference that can be used to link this account to an external system.
	CustomerID string `json:"customer_id,omitempty"`

	// Reason for modification.
	// Min Length: 1
	ModifyReason string `json:"modify_reason,omitempty"`

	// Name of the account holder, up to four lines possible.
	// Max Items: 4
	Name []string `json:"name,omitempty"`

	// organisation identification
	OrganisationIdentification *OrganisationIdentification `json:"organisation_identification,omitempty"`

	// private identification
	PrivateIdentification *PrivateIdentification `json:"private_identification,omitempty"`
}

func AccountAmendmentAttributesWithDefaults(defaults client.Defaults) *AccountAmendmentAttributes {
	return &AccountAmendmentAttributes{

		CustomerID: defaults.GetString("AccountAmendmentAttributes", "customer_id"),

		ModifyReason: defaults.GetString("AccountAmendmentAttributes", "modify_reason"),

		Name: make([]string, 0),

		OrganisationIdentification: OrganisationIdentificationWithDefaults(defaults),

		PrivateIdentification: PrivateIdentificationWithDefaults(defaults),
	}
}

func (m *AccountAmendmentAttributes) WithCustomerID(customerID string) *AccountAmendmentAttributes {

	m.CustomerID = customerID

	return m
}

func (m *AccountAmendmentAttributes) WithModifyReason(modifyReason string) *AccountAmendmentAttributes {

	m.ModifyReason = modifyReason

	return m
}

func (m *AccountAmendmentAttributes) WithName(name []string) *AccountAmendmentAttributes {

	m.Name = name

	return m
}

func (m *AccountAmendmentAttributes) WithOrganisationIdentification(organisationIdentification OrganisationIdentification) *AccountAmendmentAttributes {

	m.OrganisationIdentification = &organisationIdentification

	return m
}

func (m *AccountAmendmentAttributes) WithoutOrganisationIdentification() *AccountAmendmentAttributes {
	m.OrganisationIdentification = nil
	return m
}

func (m *AccountAmendmentAttributes) WithPrivateIdentification(privateIdentification PrivateIdentification) *AccountAmendmentAttributes {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *AccountAmendmentAttributes) WithoutPrivateIdentification() *AccountAmendmentAttributes {
	m.PrivateIdentification = nil
	return m
}

// Validate validates this account amendment attributes
func (m *AccountAmendmentAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifyReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAmendmentAttributes) validateModifyReason(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifyReason) { // not required
		return nil
	}

	if err := validate.MinLength("modify_reason", "body", string(m.ModifyReason), 1); err != nil {
		return err
	}

	return nil
}

func (m *AccountAmendmentAttributes) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	iNameSize := int64(len(m.Name))

	if err := validate.MaxItems("name", "body", iNameSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(m.Name); i++ {

		if err := validate.MinLength("name"+"."+strconv.Itoa(i), "body", string(m.Name[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("name"+"."+strconv.Itoa(i), "body", string(m.Name[i]), 140); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountAmendmentAttributes) validateOrganisationIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentification) { // not required
		return nil
	}

	if m.OrganisationIdentification != nil {
		if err := m.OrganisationIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organisation_identification")
			}
			return err
		}
	}

	return nil
}

func (m *AccountAmendmentAttributes) validatePrivateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIdentification) { // not required
		return nil
	}

	if m.PrivateIdentification != nil {
		if err := m.PrivateIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_identification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendmentAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendmentAttributes) UnmarshalBinary(b []byte) error {
	var res AccountAmendmentAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendmentAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
