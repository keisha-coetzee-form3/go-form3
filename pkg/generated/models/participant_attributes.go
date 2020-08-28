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

// ParticipantAttributes participant attributes
// swagger:model ParticipantAttributes
type ParticipantAttributes struct {

	// Participant's bank code(s)
	BankCode []string `json:"bank_code"`

	// URL where the participant Confirmation of Payee endpoint may be found
	// Required: true
	CopEndpoint *string `json:"cop_endpoint"`

	// Whether the participant is currently accepting Confirmation of Payee requests
	// Required: true
	CopRequestsActive *bool `json:"cop_requests_active"`

	// Name of the participant
	// Required: true
	// Pattern: ^[A-Za-z]*$
	ParticipantName *string `json:"participant_name"`
}

func ParticipantAttributesWithDefaults(defaults client.Defaults) *ParticipantAttributes {
	return &ParticipantAttributes{

		BankCode: make([]string, 0),

		CopEndpoint: defaults.GetStringPtr("ParticipantAttributes", "cop_endpoint"),

		CopRequestsActive: defaults.GetBoolPtr("ParticipantAttributes", "cop_requests_active"),

		ParticipantName: defaults.GetStringPtr("ParticipantAttributes", "participant_name"),
	}
}

func (m *ParticipantAttributes) WithBankCode(bankCode []string) *ParticipantAttributes {

	m.BankCode = bankCode

	return m
}

func (m *ParticipantAttributes) WithCopEndpoint(copEndpoint string) *ParticipantAttributes {

	m.CopEndpoint = &copEndpoint

	return m
}

func (m *ParticipantAttributes) WithoutCopEndpoint() *ParticipantAttributes {
	m.CopEndpoint = nil
	return m
}

func (m *ParticipantAttributes) WithCopRequestsActive(copRequestsActive bool) *ParticipantAttributes {

	m.CopRequestsActive = &copRequestsActive

	return m
}

func (m *ParticipantAttributes) WithoutCopRequestsActive() *ParticipantAttributes {
	m.CopRequestsActive = nil
	return m
}

func (m *ParticipantAttributes) WithParticipantName(participantName string) *ParticipantAttributes {

	m.ParticipantName = &participantName

	return m
}

func (m *ParticipantAttributes) WithoutParticipantName() *ParticipantAttributes {
	m.ParticipantName = nil
	return m
}

// Validate validates this participant attributes
func (m *ParticipantAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCopEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCopRequestsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipantName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParticipantAttributes) validateCopEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("cop_endpoint", "body", m.CopEndpoint); err != nil {
		return err
	}

	return nil
}

func (m *ParticipantAttributes) validateCopRequestsActive(formats strfmt.Registry) error {

	if err := validate.Required("cop_requests_active", "body", m.CopRequestsActive); err != nil {
		return err
	}

	return nil
}

func (m *ParticipantAttributes) validateParticipantName(formats strfmt.Registry) error {

	if err := validate.Required("participant_name", "body", m.ParticipantName); err != nil {
		return err
	}

	if err := validate.Pattern("participant_name", "body", string(*m.ParticipantName), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParticipantAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParticipantAttributes) UnmarshalBinary(b []byte) error {
	var res ParticipantAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ParticipantAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}