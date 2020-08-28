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

// AceDetailsResponse ace details response
// swagger:model AceDetailsResponse
type AceDetailsResponse struct {

	// data
	// Required: true
	Data *Ace `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func AceDetailsResponseWithDefaults(defaults client.Defaults) *AceDetailsResponse {
	return &AceDetailsResponse{

		Data: AceWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *AceDetailsResponse) WithData(data Ace) *AceDetailsResponse {

	m.Data = &data

	return m
}

func (m *AceDetailsResponse) WithoutData() *AceDetailsResponse {
	m.Data = nil
	return m
}

func (m *AceDetailsResponse) WithLinks(links Links) *AceDetailsResponse {

	m.Links = &links

	return m
}

func (m *AceDetailsResponse) WithoutLinks() *AceDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this ace details response
func (m *AceDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AceDetailsResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *AceDetailsResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AceDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AceDetailsResponse) UnmarshalBinary(b []byte) error {
	var res AceDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AceDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
