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

// ReportDetailsListResponse report details list response
// swagger:model ReportDetailsListResponse
type ReportDetailsListResponse struct {

	// data
	// Required: true
	Data []*Report `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func ReportDetailsListResponseWithDefaults(defaults client.Defaults) *ReportDetailsListResponse {
	return &ReportDetailsListResponse{

		Data: make([]*Report, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *ReportDetailsListResponse) WithData(data []*Report) *ReportDetailsListResponse {

	m.Data = data

	return m
}

func (m *ReportDetailsListResponse) WithLinks(links Links) *ReportDetailsListResponse {

	m.Links = &links

	return m
}

func (m *ReportDetailsListResponse) WithoutLinks() *ReportDetailsListResponse {
	m.Links = nil
	return m
}

// Validate validates this report details list response
func (m *ReportDetailsListResponse) Validate(formats strfmt.Registry) error {
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

func (m *ReportDetailsListResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportDetailsListResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *ReportDetailsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportDetailsListResponse) UnmarshalBinary(b []byte) error {
	var res ReportDetailsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportDetailsListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
