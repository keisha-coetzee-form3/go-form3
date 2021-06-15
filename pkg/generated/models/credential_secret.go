// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CredentialSecret credential secret
// swagger:model CredentialSecret
type CredentialSecret struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`
}

func CredentialSecretWithDefaults(defaults client.Defaults) *CredentialSecret {
	return &CredentialSecret{

		ClientID: defaults.GetString("CredentialSecret", "client_id"),

		ClientSecret: defaults.GetString("CredentialSecret", "client_secret"),
	}
}

func (m *CredentialSecret) WithClientID(clientID string) *CredentialSecret {

	m.ClientID = clientID

	return m
}

func (m *CredentialSecret) WithClientSecret(clientSecret string) *CredentialSecret {

	m.ClientSecret = clientSecret

	return m
}

// Validate validates this credential secret
func (m *CredentialSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CredentialSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialSecret) UnmarshalBinary(b []byte) error {
	var res CredentialSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *CredentialSecret) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
