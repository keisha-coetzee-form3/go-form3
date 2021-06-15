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

// PartyRelationships party relationships
// swagger:model PartyRelationships
type PartyRelationships struct {

	// contact
	Contact *RelationshipsContactRef `json:"contact,omitempty"`
}

func PartyRelationshipsWithDefaults(defaults client.Defaults) *PartyRelationships {
	return &PartyRelationships{

		Contact: RelationshipsContactRefWithDefaults(defaults),
	}
}

func (m *PartyRelationships) WithContact(contact RelationshipsContactRef) *PartyRelationships {

	m.Contact = &contact

	return m
}

func (m *PartyRelationships) WithoutContact() *PartyRelationships {
	m.Contact = nil
	return m
}

// Validate validates this party relationships
func (m *PartyRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyRelationships) validateContact(formats strfmt.Registry) error {

	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyRelationships) UnmarshalBinary(b []byte) error {
	var res PartyRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
