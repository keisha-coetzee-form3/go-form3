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

// QueryAdmission query admission
// swagger:model QueryAdmission
type QueryAdmission struct {

	// type
	// Required: true
	Type QueryAdmissionResourceType `json:"type"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// version
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// attributes
	// Required: true
	Attributes *QueryAdmissionAttributes `json:"attributes"`

	// relationships
	// Required: true
	Relationships *QueryAdmissionRelationships `json:"relationships"`
}

func QueryAdmissionWithDefaults(defaults client.Defaults) *QueryAdmission {
	return &QueryAdmission{

		// TODO Type: QueryAdmissionResourceType,

		ID: defaults.GetStrfmtUUIDPtr("QueryAdmission", "id"),

		Version: defaults.GetInt64Ptr("QueryAdmission", "version"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("QueryAdmission", "organisation_id"),

		CreatedOn: defaults.GetStrfmtDateTime("QueryAdmission", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("QueryAdmission", "modified_on"),

		Attributes: QueryAdmissionAttributesWithDefaults(defaults),

		Relationships: QueryAdmissionRelationshipsWithDefaults(defaults),
	}
}

func (m *QueryAdmission) WithType(typeVar QueryAdmissionResourceType) *QueryAdmission {

	m.Type = typeVar

	return m
}

func (m *QueryAdmission) WithID(id strfmt.UUID) *QueryAdmission {

	m.ID = &id

	return m
}

func (m *QueryAdmission) WithoutID() *QueryAdmission {
	m.ID = nil
	return m
}

func (m *QueryAdmission) WithVersion(version int64) *QueryAdmission {

	m.Version = &version

	return m
}

func (m *QueryAdmission) WithoutVersion() *QueryAdmission {
	m.Version = nil
	return m
}

func (m *QueryAdmission) WithOrganisationID(organisationID strfmt.UUID) *QueryAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *QueryAdmission) WithoutOrganisationID() *QueryAdmission {
	m.OrganisationID = nil
	return m
}

func (m *QueryAdmission) WithCreatedOn(createdOn strfmt.DateTime) *QueryAdmission {

	m.CreatedOn = createdOn

	return m
}

func (m *QueryAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *QueryAdmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *QueryAdmission) WithAttributes(attributes QueryAdmissionAttributes) *QueryAdmission {

	m.Attributes = &attributes

	return m
}

func (m *QueryAdmission) WithoutAttributes() *QueryAdmission {
	m.Attributes = nil
	return m
}

func (m *QueryAdmission) WithRelationships(relationships QueryAdmissionRelationships) *QueryAdmission {

	m.Relationships = &relationships

	return m
}

func (m *QueryAdmission) WithoutRelationships() *QueryAdmission {
	m.Relationships = nil
	return m
}

// Validate validates this query admission
func (m *QueryAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryAdmission) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *QueryAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryAdmission) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *QueryAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *QueryAdmission) validateRelationships(formats strfmt.Registry) error {

	if err := validate.Required("relationships", "body", m.Relationships); err != nil {
		return err
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

// MarshalBinary interface implementation
func (m *QueryAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryAdmission) UnmarshalBinary(b []byte) error {
	var res QueryAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
