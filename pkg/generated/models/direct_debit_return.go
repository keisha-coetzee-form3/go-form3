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
	"github.com/go-openapi/validate"
)

// DirectDebitReturn direct debit return
// swagger:model DirectDebitReturn
type DirectDebitReturn struct {

	// attributes
	// Required: true
	Attributes *DirectDebitReturnAttributes `json:"attributes"`

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
	Relationships *DirectDebitReturnRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitReturn {
	return &DirectDebitReturn{

		Attributes: DirectDebitReturnAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturn", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitReturn", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturn", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitReturn", "organisation_id"),

		Relationships: DirectDebitReturnRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReturn", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReturn", "version"),
	}
}

func (m *DirectDebitReturn) WithAttributes(attributes DirectDebitReturnAttributes) *DirectDebitReturn {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReturn) WithoutAttributes() *DirectDebitReturn {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReturn) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReturn {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReturn) WithoutCreatedOn() *DirectDebitReturn {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReturn) WithID(id strfmt.UUID) *DirectDebitReturn {

	m.ID = &id

	return m
}

func (m *DirectDebitReturn) WithoutID() *DirectDebitReturn {
	m.ID = nil
	return m
}

func (m *DirectDebitReturn) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReturn {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReturn) WithoutModifiedOn() *DirectDebitReturn {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReturn) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReturn {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitReturn) WithoutOrganisationID() *DirectDebitReturn {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitReturn) WithRelationships(relationships DirectDebitReturnRelationships) *DirectDebitReturn {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReturn) WithoutRelationships() *DirectDebitReturn {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReturn) WithType(typeVar string) *DirectDebitReturn {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReturn) WithVersion(version int64) *DirectDebitReturn {

	m.Version = &version

	return m
}

func (m *DirectDebitReturn) WithoutVersion() *DirectDebitReturn {
	m.Version = nil
	return m
}

// Validate validates this direct debit return
func (m *DirectDebitReturn) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReturn) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReturn) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReturn) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturn) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnAttributes direct debit return attributes
// swagger:model DirectDebitReturnAttributes
type DirectDebitReturnAttributes struct {

	// charges amount
	ChargesAmount *CurrencyAndAmount `json:"charges_amount,omitempty"`

	// Unique identifier for organisations collecting payments
	ClearingID string `json:"clearing_id,omitempty"`

	// compensation amount
	CompensationAmount *CurrencyAndAmount `json:"compensation_amount,omitempty"`

	// return amount
	ReturnAmount *CurrencyAndAmount `json:"return_amount,omitempty"`

	// return code
	ReturnCode string `json:"return_code,omitempty"`

	// return initiator
	// Enum: [BANK CUSTOMER]
	ReturnInitiator string `json:"return_initiator,omitempty"`

	// scheme transaction id
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`
}

func DirectDebitReturnAttributesWithDefaults(defaults client.Defaults) *DirectDebitReturnAttributes {
	return &DirectDebitReturnAttributes{

		ChargesAmount: CurrencyAndAmountWithDefaults(defaults),

		ClearingID: defaults.GetString("DirectDebitReturnAttributes", "clearing_id"),

		CompensationAmount: CurrencyAndAmountWithDefaults(defaults),

		ReturnAmount: CurrencyAndAmountWithDefaults(defaults),

		ReturnCode: defaults.GetString("DirectDebitReturnAttributes", "return_code"),

		ReturnInitiator: defaults.GetString("DirectDebitReturnAttributes", "return_initiator"),

		SchemeTransactionID: defaults.GetString("DirectDebitReturnAttributes", "scheme_transaction_id"),
	}
}

func (m *DirectDebitReturnAttributes) WithChargesAmount(chargesAmount CurrencyAndAmount) *DirectDebitReturnAttributes {

	m.ChargesAmount = &chargesAmount

	return m
}

func (m *DirectDebitReturnAttributes) WithoutChargesAmount() *DirectDebitReturnAttributes {
	m.ChargesAmount = nil
	return m
}

func (m *DirectDebitReturnAttributes) WithClearingID(clearingID string) *DirectDebitReturnAttributes {

	m.ClearingID = clearingID

	return m
}

func (m *DirectDebitReturnAttributes) WithCompensationAmount(compensationAmount CurrencyAndAmount) *DirectDebitReturnAttributes {

	m.CompensationAmount = &compensationAmount

	return m
}

func (m *DirectDebitReturnAttributes) WithoutCompensationAmount() *DirectDebitReturnAttributes {
	m.CompensationAmount = nil
	return m
}

func (m *DirectDebitReturnAttributes) WithReturnAmount(returnAmount CurrencyAndAmount) *DirectDebitReturnAttributes {

	m.ReturnAmount = &returnAmount

	return m
}

func (m *DirectDebitReturnAttributes) WithoutReturnAmount() *DirectDebitReturnAttributes {
	m.ReturnAmount = nil
	return m
}

func (m *DirectDebitReturnAttributes) WithReturnCode(returnCode string) *DirectDebitReturnAttributes {

	m.ReturnCode = returnCode

	return m
}

func (m *DirectDebitReturnAttributes) WithReturnInitiator(returnInitiator string) *DirectDebitReturnAttributes {

	m.ReturnInitiator = returnInitiator

	return m
}

func (m *DirectDebitReturnAttributes) WithSchemeTransactionID(schemeTransactionID string) *DirectDebitReturnAttributes {

	m.SchemeTransactionID = schemeTransactionID

	return m
}

// Validate validates this direct debit return attributes
func (m *DirectDebitReturnAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChargesAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompensationAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnInitiator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnAttributes) validateChargesAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargesAmount) { // not required
		return nil
	}

	if m.ChargesAmount != nil {
		if err := m.ChargesAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "charges_amount")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnAttributes) validateCompensationAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.CompensationAmount) { // not required
		return nil
	}

	if m.CompensationAmount != nil {
		if err := m.CompensationAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "compensation_amount")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnAttributes) validateReturnAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnAmount) { // not required
		return nil
	}

	if m.ReturnAmount != nil {
		if err := m.ReturnAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "return_amount")
			}
			return err
		}
	}

	return nil
}

var directDebitReturnAttributesTypeReturnInitiatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BANK","CUSTOMER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitReturnAttributesTypeReturnInitiatorPropEnum = append(directDebitReturnAttributesTypeReturnInitiatorPropEnum, v)
	}
}

const (

	// DirectDebitReturnAttributesReturnInitiatorBANK captures enum value "BANK"
	DirectDebitReturnAttributesReturnInitiatorBANK string = "BANK"

	// DirectDebitReturnAttributesReturnInitiatorCUSTOMER captures enum value "CUSTOMER"
	DirectDebitReturnAttributesReturnInitiatorCUSTOMER string = "CUSTOMER"
)

// prop value enum
func (m *DirectDebitReturnAttributes) validateReturnInitiatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, directDebitReturnAttributesTypeReturnInitiatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DirectDebitReturnAttributes) validateReturnInitiator(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnInitiator) { // not required
		return nil
	}

	// value enum
	if err := m.validateReturnInitiatorEnum("attributes"+"."+"return_initiator", "body", m.ReturnInitiator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationships direct debit return relationships
// swagger:model DirectDebitReturnRelationships
type DirectDebitReturnRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReturnRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit return admission
	DirectDebitReturnAdmission *DirectDebitReturnRelationshipsDirectDebitReturnAdmission `json:"direct_debit_return_admission,omitempty"`

	// direct debit return reversal
	DirectDebitReturnReversal *DirectDebitReturnRelationshipsDirectDebitReturnReversal `json:"direct_debit_return_reversal,omitempty"`

	// direct debit return submission
	DirectDebitReturnSubmission *DirectDebitReturnRelationshipsDirectDebitReturnSubmission `json:"direct_debit_return_submission,omitempty"`
}

func DirectDebitReturnRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationships {
	return &DirectDebitReturnRelationships{

		DirectDebit: DirectDebitReturnRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReturnAdmission: DirectDebitReturnRelationshipsDirectDebitReturnAdmissionWithDefaults(defaults),

		DirectDebitReturnReversal: DirectDebitReturnRelationshipsDirectDebitReturnReversalWithDefaults(defaults),

		DirectDebitReturnSubmission: DirectDebitReturnRelationshipsDirectDebitReturnSubmissionWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnRelationships) WithDirectDebit(directDebit DirectDebitReturnRelationshipsDirectDebit) *DirectDebitReturnRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReturnRelationships) WithoutDirectDebit() *DirectDebitReturnRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReturnRelationships) WithDirectDebitReturnAdmission(directDebitReturnAdmission DirectDebitReturnRelationshipsDirectDebitReturnAdmission) *DirectDebitReturnRelationships {

	m.DirectDebitReturnAdmission = &directDebitReturnAdmission

	return m
}

func (m *DirectDebitReturnRelationships) WithoutDirectDebitReturnAdmission() *DirectDebitReturnRelationships {
	m.DirectDebitReturnAdmission = nil
	return m
}

func (m *DirectDebitReturnRelationships) WithDirectDebitReturnReversal(directDebitReturnReversal DirectDebitReturnRelationshipsDirectDebitReturnReversal) *DirectDebitReturnRelationships {

	m.DirectDebitReturnReversal = &directDebitReturnReversal

	return m
}

func (m *DirectDebitReturnRelationships) WithoutDirectDebitReturnReversal() *DirectDebitReturnRelationships {
	m.DirectDebitReturnReversal = nil
	return m
}

func (m *DirectDebitReturnRelationships) WithDirectDebitReturnSubmission(directDebitReturnSubmission DirectDebitReturnRelationshipsDirectDebitReturnSubmission) *DirectDebitReturnRelationships {

	m.DirectDebitReturnSubmission = &directDebitReturnSubmission

	return m
}

func (m *DirectDebitReturnRelationships) WithoutDirectDebitReturnSubmission() *DirectDebitReturnRelationships {
	m.DirectDebitReturnSubmission = nil
	return m
}

// Validate validates this direct debit return relationships
func (m *DirectDebitReturnRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturnAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturnReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturnSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnRelationships) validateDirectDebitReturnAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturnAdmission) { // not required
		return nil
	}

	if m.DirectDebitReturnAdmission != nil {
		if err := m.DirectDebitReturnAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return_admission")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnRelationships) validateDirectDebitReturnReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturnReversal) { // not required
		return nil
	}

	if m.DirectDebitReturnReversal != nil {
		if err := m.DirectDebitReturnReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnRelationships) validateDirectDebitReturnSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturnSubmission) { // not required
		return nil
	}

	if m.DirectDebitReturnSubmission != nil {
		if err := m.DirectDebitReturnSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationshipsDirectDebit direct debit return relationships direct debit
// swagger:model DirectDebitReturnRelationshipsDirectDebit
type DirectDebitReturnRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReturnRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationshipsDirectDebit {
	return &DirectDebitReturnRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReturnRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReturnRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit return relationships direct debit
func (m *DirectDebitReturnRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationshipsDirectDebitReturnAdmission direct debit return relationships direct debit return admission
// swagger:model DirectDebitReturnRelationshipsDirectDebitReturnAdmission
type DirectDebitReturnRelationshipsDirectDebitReturnAdmission struct {

	// data
	Data []*DirectDebitReturnAdmission `json:"data"`
}

func DirectDebitReturnRelationshipsDirectDebitReturnAdmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationshipsDirectDebitReturnAdmission {
	return &DirectDebitReturnRelationshipsDirectDebitReturnAdmission{

		Data: make([]*DirectDebitReturnAdmission, 0),
	}
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnAdmission) WithData(data []*DirectDebitReturnAdmission) *DirectDebitReturnRelationshipsDirectDebitReturnAdmission {

	m.Data = data

	return m
}

// Validate validates this direct debit return relationships direct debit return admission
func (m *DirectDebitReturnRelationshipsDirectDebitReturnAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationshipsDirectDebitReturnAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationshipsDirectDebitReturnAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationshipsDirectDebitReturnReversal direct debit return relationships direct debit return reversal
// swagger:model DirectDebitReturnRelationshipsDirectDebitReturnReversal
type DirectDebitReturnRelationshipsDirectDebitReturnReversal struct {

	// data
	Data []*DirectDebitReturnReversal `json:"data"`
}

func DirectDebitReturnRelationshipsDirectDebitReturnReversalWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationshipsDirectDebitReturnReversal {
	return &DirectDebitReturnRelationshipsDirectDebitReturnReversal{

		Data: make([]*DirectDebitReturnReversal, 0),
	}
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnReversal) WithData(data []*DirectDebitReturnReversal) *DirectDebitReturnRelationshipsDirectDebitReturnReversal {

	m.Data = data

	return m
}

// Validate validates this direct debit return relationships direct debit return reversal
func (m *DirectDebitReturnRelationshipsDirectDebitReturnReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnReversal) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnReversal) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationshipsDirectDebitReturnReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationshipsDirectDebitReturnReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnRelationshipsDirectDebitReturnSubmission direct debit return relationships direct debit return submission
// swagger:model DirectDebitReturnRelationshipsDirectDebitReturnSubmission
type DirectDebitReturnRelationshipsDirectDebitReturnSubmission struct {

	// data
	Data []*DirectDebitReturnSubmission `json:"data"`
}

func DirectDebitReturnRelationshipsDirectDebitReturnSubmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnRelationshipsDirectDebitReturnSubmission {
	return &DirectDebitReturnRelationshipsDirectDebitReturnSubmission{

		Data: make([]*DirectDebitReturnSubmission, 0),
	}
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) WithData(data []*DirectDebitReturnSubmission) *DirectDebitReturnRelationshipsDirectDebitReturnSubmission {

	m.Data = data

	return m
}

// Validate validates this direct debit return relationships direct debit return submission
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnRelationshipsDirectDebitReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnRelationshipsDirectDebitReturnSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
