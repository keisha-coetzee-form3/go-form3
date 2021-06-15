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

// QueryResponseSubmission query response submission
// swagger:model QueryResponseSubmission
type QueryResponseSubmission struct {

	// type
	// Required: true
	Type *QueryResponseSubmissionResourceType `json:"type"`

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
	Attributes *QueryResponseSubmissionAttributes `json:"attributes"`

	// relationships
	// Required: true
	Relationships *QueryResponseSubmissionRelationships `json:"relationships"`
}

func QueryResponseSubmissionWithDefaults(defaults client.Defaults) *QueryResponseSubmission {
	return &QueryResponseSubmission{

		// TODO Type: QueryResponseSubmissionResourceType,

		ID: defaults.GetStrfmtUUIDPtr("QueryResponseSubmission", "id"),

		Version: defaults.GetInt64Ptr("QueryResponseSubmission", "version"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("QueryResponseSubmission", "organisation_id"),

		CreatedOn: defaults.GetStrfmtDateTime("QueryResponseSubmission", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("QueryResponseSubmission", "modified_on"),

		Attributes: QueryResponseSubmissionAttributesWithDefaults(defaults),

		Relationships: QueryResponseSubmissionRelationshipsWithDefaults(defaults),
	}
}

func (m *QueryResponseSubmission) WithType(typeVar QueryResponseSubmissionResourceType) *QueryResponseSubmission {

	m.Type = &typeVar

	return m
}

func (m *QueryResponseSubmission) WithoutType() *QueryResponseSubmission {
	m.Type = nil
	return m
}

func (m *QueryResponseSubmission) WithID(id strfmt.UUID) *QueryResponseSubmission {

	m.ID = &id

	return m
}

func (m *QueryResponseSubmission) WithoutID() *QueryResponseSubmission {
	m.ID = nil
	return m
}

func (m *QueryResponseSubmission) WithVersion(version int64) *QueryResponseSubmission {

	m.Version = &version

	return m
}

func (m *QueryResponseSubmission) WithoutVersion() *QueryResponseSubmission {
	m.Version = nil
	return m
}

func (m *QueryResponseSubmission) WithOrganisationID(organisationID strfmt.UUID) *QueryResponseSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *QueryResponseSubmission) WithoutOrganisationID() *QueryResponseSubmission {
	m.OrganisationID = nil
	return m
}

func (m *QueryResponseSubmission) WithCreatedOn(createdOn strfmt.DateTime) *QueryResponseSubmission {

	m.CreatedOn = createdOn

	return m
}

func (m *QueryResponseSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *QueryResponseSubmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *QueryResponseSubmission) WithAttributes(attributes QueryResponseSubmissionAttributes) *QueryResponseSubmission {

	m.Attributes = &attributes

	return m
}

func (m *QueryResponseSubmission) WithoutAttributes() *QueryResponseSubmission {
	m.Attributes = nil
	return m
}

func (m *QueryResponseSubmission) WithRelationships(relationships QueryResponseSubmissionRelationships) *QueryResponseSubmission {

	m.Relationships = &relationships

	return m
}

func (m *QueryResponseSubmission) WithoutRelationships() *QueryResponseSubmission {
	m.Relationships = nil
	return m
}

// Validate validates this query response submission
func (m *QueryResponseSubmission) Validate(formats strfmt.Registry) error {
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

func (m *QueryResponseSubmission) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResponseSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseSubmission) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueryResponseSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *QueryResponseSubmission) validateRelationships(formats strfmt.Registry) error {

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
func (m *QueryResponseSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseSubmission) UnmarshalBinary(b []byte) error {
	var res QueryResponseSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
