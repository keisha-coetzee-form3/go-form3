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
	"github.com/go-openapi/validate"
)

// NewPartyActor new party actor
// swagger:model NewPartyActor
type NewPartyActor struct {

	// attributes
	// Required: true
	Attributes *PartyActorAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// type
	// Enum: [party_actors]
	Type string `json:"type,omitempty"`
}

func NewPartyActorWithDefaults(defaults client.Defaults) *NewPartyActor {
	return &NewPartyActor{

		Attributes: PartyActorAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewPartyActor", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewPartyActor", "organisation_id"),

		Type: defaults.GetString("NewPartyActor", "type"),
	}
}

func (m *NewPartyActor) WithAttributes(attributes PartyActorAttributes) *NewPartyActor {

	m.Attributes = &attributes

	return m
}

func (m *NewPartyActor) WithoutAttributes() *NewPartyActor {
	m.Attributes = nil
	return m
}

func (m *NewPartyActor) WithID(id strfmt.UUID) *NewPartyActor {

	m.ID = &id

	return m
}

func (m *NewPartyActor) WithoutID() *NewPartyActor {
	m.ID = nil
	return m
}

func (m *NewPartyActor) WithOrganisationID(organisationID strfmt.UUID) *NewPartyActor {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewPartyActor) WithoutOrganisationID() *NewPartyActor {
	m.OrganisationID = nil
	return m
}

func (m *NewPartyActor) WithType(typeVar string) *NewPartyActor {

	m.Type = typeVar

	return m
}

// Validate validates this new party actor
func (m *NewPartyActor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPartyActor) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *NewPartyActor) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewPartyActor) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var newPartyActorTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["party_actors"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newPartyActorTypeTypePropEnum = append(newPartyActorTypeTypePropEnum, v)
	}
}

const (

	// NewPartyActorTypePartyActors captures enum value "party_actors"
	NewPartyActorTypePartyActors string = "party_actors"
)

// prop value enum
func (m *NewPartyActor) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newPartyActorTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewPartyActor) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPartyActor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPartyActor) UnmarshalBinary(b []byte) error {
	var res NewPartyActor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPartyActor) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
