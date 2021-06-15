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

// DirectDebitReturnAdmission direct debit return admission
// swagger:model DirectDebitReturnAdmission
type DirectDebitReturnAdmission struct {

	// attributes
	// Required: true
	Attributes *DirectDebitReturnAdmissionAttributes `json:"attributes"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *DirectDebitReturnAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReturnAdmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnAdmission {
	return &DirectDebitReturnAdmission{

		Attributes: DirectDebitReturnAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitReturnAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitReturnAdmission", "organisation_id"),

		Relationships: DirectDebitReturnAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReturnAdmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReturnAdmission", "version"),
	}
}

func (m *DirectDebitReturnAdmission) WithAttributes(attributes DirectDebitReturnAdmissionAttributes) *DirectDebitReturnAdmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReturnAdmission) WithoutAttributes() *DirectDebitReturnAdmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReturnAdmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReturnAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReturnAdmission) WithoutCreatedOn() *DirectDebitReturnAdmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReturnAdmission) WithID(id strfmt.UUID) *DirectDebitReturnAdmission {

	m.ID = &id

	return m
}

func (m *DirectDebitReturnAdmission) WithoutID() *DirectDebitReturnAdmission {
	m.ID = nil
	return m
}

func (m *DirectDebitReturnAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReturnAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReturnAdmission) WithoutModifiedOn() *DirectDebitReturnAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReturnAdmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReturnAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitReturnAdmission) WithoutOrganisationID() *DirectDebitReturnAdmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitReturnAdmission) WithRelationships(relationships DirectDebitReturnAdmissionRelationships) *DirectDebitReturnAdmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReturnAdmission) WithoutRelationships() *DirectDebitReturnAdmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReturnAdmission) WithType(typeVar string) *DirectDebitReturnAdmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReturnAdmission) WithVersion(version int64) *DirectDebitReturnAdmission {

	m.Version = &version

	return m
}

func (m *DirectDebitReturnAdmission) WithoutVersion() *DirectDebitReturnAdmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit return admission
func (m *DirectDebitReturnAdmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReturnAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnAdmissionAttributes direct debit return admission attributes
// swagger:model DirectDebitReturnAdmissionAttributes
type DirectDebitReturnAdmissionAttributes struct {

	// admission datetime
	// Read Only: true
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// settlement cycle
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// settlement date
	// Format: date
	SettlementDate *strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status DirectDebitReturnAdmissionStatus `json:"status,omitempty"`
}

func DirectDebitReturnAdmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitReturnAdmissionAttributes {
	return &DirectDebitReturnAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("DirectDebitReturnAdmissionAttributes", "admission_datetime"),

		SchemeStatusCode: defaults.GetString("DirectDebitReturnAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("DirectDebitReturnAdmissionAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("DirectDebitReturnAdmissionAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("DirectDebitReturnAdmissionAttributes", "settlement_date"),

		// TODO Status: DirectDebitReturnAdmissionStatus,

	}
}

func (m *DirectDebitReturnAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *DirectDebitReturnAdmissionAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitReturnAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *DirectDebitReturnAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithSettlementCycle(settlementCycle int64) *DirectDebitReturnAdmissionAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithoutSettlementCycle() *DirectDebitReturnAdmissionAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithSettlementDate(settlementDate strfmt.Date) *DirectDebitReturnAdmissionAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithoutSettlementDate() *DirectDebitReturnAdmissionAttributes {
	m.SettlementDate = nil
	return m
}

func (m *DirectDebitReturnAdmissionAttributes) WithStatus(status DirectDebitReturnAdmissionStatus) *DirectDebitReturnAdmissionAttributes {

	m.Status = status

	return m
}

// Validate validates this direct debit return admission attributes
func (m *DirectDebitReturnAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmissionAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmissionAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
