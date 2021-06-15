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

// PartyProductEvent party product event
// swagger:model PartyProductEvent
type PartyProductEvent struct {

	// attributes
	// Required: true
	Attributes *PartyProductEventAttributes `json:"attributes"`

	// created on
	// Format: datetime
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: datetime
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *PartyProductEventRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [party_product_events]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PartyProductEventWithDefaults(defaults client.Defaults) *PartyProductEvent {
	return &PartyProductEvent{

		Attributes: PartyProductEventAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("PartyProductEvent", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("PartyProductEvent", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("PartyProductEvent", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PartyProductEvent", "organisation_id"),

		Relationships: PartyProductEventRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PartyProductEvent", "type"),

		Version: defaults.GetInt64Ptr("PartyProductEvent", "version"),
	}
}

func (m *PartyProductEvent) WithAttributes(attributes PartyProductEventAttributes) *PartyProductEvent {

	m.Attributes = &attributes

	return m
}

func (m *PartyProductEvent) WithoutAttributes() *PartyProductEvent {
	m.Attributes = nil
	return m
}

func (m *PartyProductEvent) WithCreatedOn(createdOn strfmt.DateTime) *PartyProductEvent {

	m.CreatedOn = createdOn

	return m
}

func (m *PartyProductEvent) WithID(id strfmt.UUID) *PartyProductEvent {

	m.ID = &id

	return m
}

func (m *PartyProductEvent) WithoutID() *PartyProductEvent {
	m.ID = nil
	return m
}

func (m *PartyProductEvent) WithModifiedOn(modifiedOn strfmt.DateTime) *PartyProductEvent {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *PartyProductEvent) WithOrganisationID(organisationID strfmt.UUID) *PartyProductEvent {

	m.OrganisationID = &organisationID

	return m
}

func (m *PartyProductEvent) WithoutOrganisationID() *PartyProductEvent {
	m.OrganisationID = nil
	return m
}

func (m *PartyProductEvent) WithRelationships(relationships PartyProductEventRelationships) *PartyProductEvent {

	m.Relationships = &relationships

	return m
}

func (m *PartyProductEvent) WithoutRelationships() *PartyProductEvent {
	m.Relationships = nil
	return m
}

func (m *PartyProductEvent) WithType(typeVar string) *PartyProductEvent {

	m.Type = typeVar

	return m
}

func (m *PartyProductEvent) WithVersion(version int64) *PartyProductEvent {

	m.Version = &version

	return m
}

func (m *PartyProductEvent) WithoutVersion() *PartyProductEvent {
	m.Version = nil
	return m
}

// Validate validates this party product event
func (m *PartyProductEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyProductEvent) validateAttributes(formats strfmt.Registry) error {

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

func (m *PartyProductEvent) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "datetime", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductEvent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductEvent) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "datetime", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductEvent) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductEvent) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

var partyProductEventTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["party_product_events"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyProductEventTypeTypePropEnum = append(partyProductEventTypeTypePropEnum, v)
	}
}

const (

	// PartyProductEventTypePartyProductEvents captures enum value "party_product_events"
	PartyProductEventTypePartyProductEvents string = "party_product_events"
)

// prop value enum
func (m *PartyProductEvent) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyProductEventTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyProductEvent) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductEvent) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyProductEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyProductEvent) UnmarshalBinary(b []byte) error {
	var res PartyProductEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyProductEvent) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
