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

// PartyAttributes party attributes
// swagger:model PartyAttributes
type PartyAttributes struct {

	// address
	Address []string `json:"address"`

	// city
	City string `json:"city,omitempty"`

	// contact method
	// Enum: [email]
	ContactMethod string `json:"contact_method,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// customer id
	CustomerID string `json:"customer_id,omitempty"`

	// district
	District string `json:"district,omitempty"`

	// email address
	EmailAddress string `json:"email_address,omitempty"`

	// identification
	Identification []string `json:"identification"`

	// identification type
	// Enum: [CompanyRegNumber]
	IdentificationType string `json:"identification_type,omitempty"`

	// name
	Name []string `json:"name"`

	// organisation identification
	OrganisationIdentification *PartyAttributesOrganisationIdentification `json:"organisation_identification,omitempty"`

	// party activity
	PartyActivity *PartyAttributesPartyActivity `json:"party_activity,omitempty"`

	// party type
	// Enum: [organisation]
	PartyType string `json:"party_type,omitempty"`

	// post code
	PostCode string `json:"post_code,omitempty"`

	// province
	Province string `json:"province,omitempty"`

	// risk summary
	RiskSummary []*PartyAttributesRiskSummaryItems0 `json:"risk_summary"`

	// telephone number
	TelephoneNumber string `json:"telephone_number,omitempty"`
}

func PartyAttributesWithDefaults(defaults client.Defaults) *PartyAttributes {
	return &PartyAttributes{

		Address: make([]string, 0),

		City: defaults.GetString("PartyAttributes", "city"),

		ContactMethod: defaults.GetString("PartyAttributes", "contact_method"),

		Country: defaults.GetString("PartyAttributes", "country"),

		CustomerID: defaults.GetString("PartyAttributes", "customer_id"),

		District: defaults.GetString("PartyAttributes", "district"),

		EmailAddress: defaults.GetString("PartyAttributes", "email_address"),

		Identification: make([]string, 0),

		IdentificationType: defaults.GetString("PartyAttributes", "identification_type"),

		Name: make([]string, 0),

		OrganisationIdentification: PartyAttributesOrganisationIdentificationWithDefaults(defaults),

		PartyActivity: PartyAttributesPartyActivityWithDefaults(defaults),

		PartyType: defaults.GetString("PartyAttributes", "party_type"),

		PostCode: defaults.GetString("PartyAttributes", "post_code"),

		Province: defaults.GetString("PartyAttributes", "province"),

		RiskSummary: make([]*PartyAttributesRiskSummaryItems0, 0),

		TelephoneNumber: defaults.GetString("PartyAttributes", "telephone_number"),
	}
}

func (m *PartyAttributes) WithAddress(address []string) *PartyAttributes {

	m.Address = address

	return m
}

func (m *PartyAttributes) WithCity(city string) *PartyAttributes {

	m.City = city

	return m
}

func (m *PartyAttributes) WithContactMethod(contactMethod string) *PartyAttributes {

	m.ContactMethod = contactMethod

	return m
}

func (m *PartyAttributes) WithCountry(country string) *PartyAttributes {

	m.Country = country

	return m
}

func (m *PartyAttributes) WithCustomerID(customerID string) *PartyAttributes {

	m.CustomerID = customerID

	return m
}

func (m *PartyAttributes) WithDistrict(district string) *PartyAttributes {

	m.District = district

	return m
}

func (m *PartyAttributes) WithEmailAddress(emailAddress string) *PartyAttributes {

	m.EmailAddress = emailAddress

	return m
}

func (m *PartyAttributes) WithIdentification(identification []string) *PartyAttributes {

	m.Identification = identification

	return m
}

func (m *PartyAttributes) WithIdentificationType(identificationType string) *PartyAttributes {

	m.IdentificationType = identificationType

	return m
}

func (m *PartyAttributes) WithName(name []string) *PartyAttributes {

	m.Name = name

	return m
}

func (m *PartyAttributes) WithOrganisationIdentification(organisationIdentification PartyAttributesOrganisationIdentification) *PartyAttributes {

	m.OrganisationIdentification = &organisationIdentification

	return m
}

func (m *PartyAttributes) WithoutOrganisationIdentification() *PartyAttributes {
	m.OrganisationIdentification = nil
	return m
}

func (m *PartyAttributes) WithPartyActivity(partyActivity PartyAttributesPartyActivity) *PartyAttributes {

	m.PartyActivity = &partyActivity

	return m
}

func (m *PartyAttributes) WithoutPartyActivity() *PartyAttributes {
	m.PartyActivity = nil
	return m
}

func (m *PartyAttributes) WithPartyType(partyType string) *PartyAttributes {

	m.PartyType = partyType

	return m
}

func (m *PartyAttributes) WithPostCode(postCode string) *PartyAttributes {

	m.PostCode = postCode

	return m
}

func (m *PartyAttributes) WithProvince(province string) *PartyAttributes {

	m.Province = province

	return m
}

func (m *PartyAttributes) WithRiskSummary(riskSummary []*PartyAttributesRiskSummaryItems0) *PartyAttributes {

	m.RiskSummary = riskSummary

	return m
}

func (m *PartyAttributes) WithTelephoneNumber(telephoneNumber string) *PartyAttributes {

	m.TelephoneNumber = telephoneNumber

	return m
}

// Validate validates this party attributes
func (m *PartyAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentificationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRiskSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var partyAttributesTypeContactMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["email"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyAttributesTypeContactMethodPropEnum = append(partyAttributesTypeContactMethodPropEnum, v)
	}
}

const (

	// PartyAttributesContactMethodEmail captures enum value "email"
	PartyAttributesContactMethodEmail string = "email"
)

// prop value enum
func (m *PartyAttributes) validateContactMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyAttributesTypeContactMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyAttributes) validateContactMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateContactMethodEnum("contact_method", "body", m.ContactMethod); err != nil {
		return err
	}

	return nil
}

var partyAttributesTypeIdentificationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CompanyRegNumber"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyAttributesTypeIdentificationTypePropEnum = append(partyAttributesTypeIdentificationTypePropEnum, v)
	}
}

const (

	// PartyAttributesIdentificationTypeCompanyRegNumber captures enum value "CompanyRegNumber"
	PartyAttributesIdentificationTypeCompanyRegNumber string = "CompanyRegNumber"
)

// prop value enum
func (m *PartyAttributes) validateIdentificationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyAttributesTypeIdentificationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyAttributes) validateIdentificationType(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentificationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIdentificationTypeEnum("identification_type", "body", m.IdentificationType); err != nil {
		return err
	}

	return nil
}

func (m *PartyAttributes) validateOrganisationIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationIdentification) { // not required
		return nil
	}

	if m.OrganisationIdentification != nil {
		if err := m.OrganisationIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organisation_identification")
			}
			return err
		}
	}

	return nil
}

func (m *PartyAttributes) validatePartyActivity(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyActivity) { // not required
		return nil
	}

	if m.PartyActivity != nil {
		if err := m.PartyActivity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party_activity")
			}
			return err
		}
	}

	return nil
}

var partyAttributesTypePartyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["organisation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyAttributesTypePartyTypePropEnum = append(partyAttributesTypePartyTypePropEnum, v)
	}
}

const (

	// PartyAttributesPartyTypeOrganisation captures enum value "organisation"
	PartyAttributesPartyTypeOrganisation string = "organisation"
)

// prop value enum
func (m *PartyAttributes) validatePartyTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyAttributesTypePartyTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyAttributes) validatePartyType(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartyTypeEnum("party_type", "body", m.PartyType); err != nil {
		return err
	}

	return nil
}

func (m *PartyAttributes) validateRiskSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.RiskSummary) { // not required
		return nil
	}

	for i := 0; i < len(m.RiskSummary); i++ {
		if swag.IsZero(m.RiskSummary[i]) { // not required
			continue
		}

		if m.RiskSummary[i] != nil {
			if err := m.RiskSummary[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("risk_summary" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAttributes) UnmarshalBinary(b []byte) error {
	var res PartyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PartyAttributesOrganisationIdentification party attributes organisation identification
// swagger:model PartyAttributesOrganisationIdentification
type PartyAttributesOrganisationIdentification struct {

	// identification
	Identification string `json:"identification,omitempty"`

	// industry classifications
	IndustryClassifications []*IndustryClassification `json:"industry_classifications"`

	// organisation description
	OrganisationDescription string `json:"organisation_description,omitempty"`

	// organisation type
	OrganisationType string `json:"organisation_type,omitempty"`

	// registration date
	RegistrationDate string `json:"registration_date,omitempty"`

	// trading address
	TradingAddress []string `json:"trading_address"`

	// trading city
	TradingCity string `json:"trading_city,omitempty"`

	// trading country
	TradingCountry string `json:"trading_country,omitempty"`

	// trading name
	TradingName string `json:"trading_name,omitempty"`

	// trading post code
	TradingPostCode string `json:"trading_post_code,omitempty"`
}

func PartyAttributesOrganisationIdentificationWithDefaults(defaults client.Defaults) *PartyAttributesOrganisationIdentification {
	return &PartyAttributesOrganisationIdentification{

		Identification: defaults.GetString("PartyAttributesOrganisationIdentification", "identification"),

		IndustryClassifications: make([]*IndustryClassification, 0),

		OrganisationDescription: defaults.GetString("PartyAttributesOrganisationIdentification", "organisation_description"),

		OrganisationType: defaults.GetString("PartyAttributesOrganisationIdentification", "organisation_type"),

		RegistrationDate: defaults.GetString("PartyAttributesOrganisationIdentification", "registration_date"),

		TradingAddress: make([]string, 0),

		TradingCity: defaults.GetString("PartyAttributesOrganisationIdentification", "trading_city"),

		TradingCountry: defaults.GetString("PartyAttributesOrganisationIdentification", "trading_country"),

		TradingName: defaults.GetString("PartyAttributesOrganisationIdentification", "trading_name"),

		TradingPostCode: defaults.GetString("PartyAttributesOrganisationIdentification", "trading_post_code"),
	}
}

func (m *PartyAttributesOrganisationIdentification) WithIdentification(identification string) *PartyAttributesOrganisationIdentification {

	m.Identification = identification

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithIndustryClassifications(industryClassifications []*IndustryClassification) *PartyAttributesOrganisationIdentification {

	m.IndustryClassifications = industryClassifications

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithOrganisationDescription(organisationDescription string) *PartyAttributesOrganisationIdentification {

	m.OrganisationDescription = organisationDescription

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithOrganisationType(organisationType string) *PartyAttributesOrganisationIdentification {

	m.OrganisationType = organisationType

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithRegistrationDate(registrationDate string) *PartyAttributesOrganisationIdentification {

	m.RegistrationDate = registrationDate

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithTradingAddress(tradingAddress []string) *PartyAttributesOrganisationIdentification {

	m.TradingAddress = tradingAddress

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithTradingCity(tradingCity string) *PartyAttributesOrganisationIdentification {

	m.TradingCity = tradingCity

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithTradingCountry(tradingCountry string) *PartyAttributesOrganisationIdentification {

	m.TradingCountry = tradingCountry

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithTradingName(tradingName string) *PartyAttributesOrganisationIdentification {

	m.TradingName = tradingName

	return m
}

func (m *PartyAttributesOrganisationIdentification) WithTradingPostCode(tradingPostCode string) *PartyAttributesOrganisationIdentification {

	m.TradingPostCode = tradingPostCode

	return m
}

// Validate validates this party attributes organisation identification
func (m *PartyAttributesOrganisationIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndustryClassifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyAttributesOrganisationIdentification) validateIndustryClassifications(formats strfmt.Registry) error {

	if swag.IsZero(m.IndustryClassifications) { // not required
		return nil
	}

	for i := 0; i < len(m.IndustryClassifications); i++ {
		if swag.IsZero(m.IndustryClassifications[i]) { // not required
			continue
		}

		if m.IndustryClassifications[i] != nil {
			if err := m.IndustryClassifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organisation_identification" + "." + "industry_classifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyAttributesOrganisationIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAttributesOrganisationIdentification) UnmarshalBinary(b []byte) error {
	var res PartyAttributesOrganisationIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAttributesOrganisationIdentification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PartyAttributesPartyActivity party attributes party activity
// swagger:model PartyAttributesPartyActivity
type PartyAttributesPartyActivity struct {

	// annual fx
	AnnualFx *PartyAttributesPartyActivityAnnualFx `json:"annual_fx,omitempty"`

	// annual payment volume
	AnnualPaymentVolume string `json:"annual_payment_volume,omitempty"`

	// currency usage
	CurrencyUsage *PartyAttributesPartyActivityCurrencyUsage `json:"currency_usage,omitempty"`

	// destination of funds countries
	DestinationOfFundsCountries []string `json:"destination_of_funds_countries"`

	// origin of funds
	OriginOfFunds []string `json:"origin_of_funds"`

	// origin of funds countries
	OriginOfFundsCountries []string `json:"origin_of_funds_countries"`

	// reasons for fx
	ReasonsForFx []*ReasonForFX `json:"reasons_for_fx"`
}

func PartyAttributesPartyActivityWithDefaults(defaults client.Defaults) *PartyAttributesPartyActivity {
	return &PartyAttributesPartyActivity{

		AnnualFx: PartyAttributesPartyActivityAnnualFxWithDefaults(defaults),

		AnnualPaymentVolume: defaults.GetString("PartyAttributesPartyActivity", "annual_payment_volume"),

		CurrencyUsage: PartyAttributesPartyActivityCurrencyUsageWithDefaults(defaults),

		DestinationOfFundsCountries: make([]string, 0),

		OriginOfFunds: make([]string, 0),

		OriginOfFundsCountries: make([]string, 0),

		ReasonsForFx: make([]*ReasonForFX, 0),
	}
}

func (m *PartyAttributesPartyActivity) WithAnnualFx(annualFx PartyAttributesPartyActivityAnnualFx) *PartyAttributesPartyActivity {

	m.AnnualFx = &annualFx

	return m
}

func (m *PartyAttributesPartyActivity) WithoutAnnualFx() *PartyAttributesPartyActivity {
	m.AnnualFx = nil
	return m
}

func (m *PartyAttributesPartyActivity) WithAnnualPaymentVolume(annualPaymentVolume string) *PartyAttributesPartyActivity {

	m.AnnualPaymentVolume = annualPaymentVolume

	return m
}

func (m *PartyAttributesPartyActivity) WithCurrencyUsage(currencyUsage PartyAttributesPartyActivityCurrencyUsage) *PartyAttributesPartyActivity {

	m.CurrencyUsage = &currencyUsage

	return m
}

func (m *PartyAttributesPartyActivity) WithoutCurrencyUsage() *PartyAttributesPartyActivity {
	m.CurrencyUsage = nil
	return m
}

func (m *PartyAttributesPartyActivity) WithDestinationOfFundsCountries(destinationOfFundsCountries []string) *PartyAttributesPartyActivity {

	m.DestinationOfFundsCountries = destinationOfFundsCountries

	return m
}

func (m *PartyAttributesPartyActivity) WithOriginOfFunds(originOfFunds []string) *PartyAttributesPartyActivity {

	m.OriginOfFunds = originOfFunds

	return m
}

func (m *PartyAttributesPartyActivity) WithOriginOfFundsCountries(originOfFundsCountries []string) *PartyAttributesPartyActivity {

	m.OriginOfFundsCountries = originOfFundsCountries

	return m
}

func (m *PartyAttributesPartyActivity) WithReasonsForFx(reasonsForFx []*ReasonForFX) *PartyAttributesPartyActivity {

	m.ReasonsForFx = reasonsForFx

	return m
}

// Validate validates this party attributes party activity
func (m *PartyAttributesPartyActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnualFx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrencyUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasonsForFx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyAttributesPartyActivity) validateAnnualFx(formats strfmt.Registry) error {

	if swag.IsZero(m.AnnualFx) { // not required
		return nil
	}

	if m.AnnualFx != nil {
		if err := m.AnnualFx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party_activity" + "." + "annual_fx")
			}
			return err
		}
	}

	return nil
}

func (m *PartyAttributesPartyActivity) validateCurrencyUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrencyUsage) { // not required
		return nil
	}

	if m.CurrencyUsage != nil {
		if err := m.CurrencyUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party_activity" + "." + "currency_usage")
			}
			return err
		}
	}

	return nil
}

func (m *PartyAttributesPartyActivity) validateReasonsForFx(formats strfmt.Registry) error {

	if swag.IsZero(m.ReasonsForFx) { // not required
		return nil
	}

	for i := 0; i < len(m.ReasonsForFx); i++ {
		if swag.IsZero(m.ReasonsForFx[i]) { // not required
			continue
		}

		if m.ReasonsForFx[i] != nil {
			if err := m.ReasonsForFx[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("party_activity" + "." + "reasons_for_fx" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyAttributesPartyActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAttributesPartyActivity) UnmarshalBinary(b []byte) error {
	var res PartyAttributesPartyActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAttributesPartyActivity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PartyAttributesPartyActivityAnnualFx party attributes party activity annual fx
// swagger:model PartyAttributesPartyActivityAnnualFx
type PartyAttributesPartyActivityAnnualFx struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

func PartyAttributesPartyActivityAnnualFxWithDefaults(defaults client.Defaults) *PartyAttributesPartyActivityAnnualFx {
	return &PartyAttributesPartyActivityAnnualFx{

		Amount: defaults.GetString("PartyAttributesPartyActivityAnnualFx", "amount"),

		Currency: defaults.GetString("PartyAttributesPartyActivityAnnualFx", "currency"),
	}
}

func (m *PartyAttributesPartyActivityAnnualFx) WithAmount(amount string) *PartyAttributesPartyActivityAnnualFx {

	m.Amount = amount

	return m
}

func (m *PartyAttributesPartyActivityAnnualFx) WithCurrency(currency string) *PartyAttributesPartyActivityAnnualFx {

	m.Currency = currency

	return m
}

// Validate validates this party attributes party activity annual fx
func (m *PartyAttributesPartyActivityAnnualFx) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartyAttributesPartyActivityAnnualFx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAttributesPartyActivityAnnualFx) UnmarshalBinary(b []byte) error {
	var res PartyAttributesPartyActivityAnnualFx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAttributesPartyActivityAnnualFx) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PartyAttributesPartyActivityCurrencyUsage party attributes party activity currency usage
// swagger:model PartyAttributesPartyActivityCurrencyUsage
type PartyAttributesPartyActivityCurrencyUsage struct {

	// source currencies
	SourceCurrencies []string `json:"source_currencies"`

	// target currencies
	TargetCurrencies []string `json:"target_currencies"`
}

func PartyAttributesPartyActivityCurrencyUsageWithDefaults(defaults client.Defaults) *PartyAttributesPartyActivityCurrencyUsage {
	return &PartyAttributesPartyActivityCurrencyUsage{

		SourceCurrencies: make([]string, 0),

		TargetCurrencies: make([]string, 0),
	}
}

func (m *PartyAttributesPartyActivityCurrencyUsage) WithSourceCurrencies(sourceCurrencies []string) *PartyAttributesPartyActivityCurrencyUsage {

	m.SourceCurrencies = sourceCurrencies

	return m
}

func (m *PartyAttributesPartyActivityCurrencyUsage) WithTargetCurrencies(targetCurrencies []string) *PartyAttributesPartyActivityCurrencyUsage {

	m.TargetCurrencies = targetCurrencies

	return m
}

// Validate validates this party attributes party activity currency usage
func (m *PartyAttributesPartyActivityCurrencyUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartyAttributesPartyActivityCurrencyUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAttributesPartyActivityCurrencyUsage) UnmarshalBinary(b []byte) error {
	var res PartyAttributesPartyActivityCurrencyUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAttributesPartyActivityCurrencyUsage) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PartyAttributesRiskSummaryItems0 party attributes risk summary items0
// swagger:model PartyAttributesRiskSummaryItems0
type PartyAttributesRiskSummaryItems0 struct {

	// risk score
	RiskScore string `json:"risk_score,omitempty"`

	// risk score alignment
	RiskScoreAlignment string `json:"risk_score_alignment,omitempty"`
}

func PartyAttributesRiskSummaryItems0WithDefaults(defaults client.Defaults) *PartyAttributesRiskSummaryItems0 {
	return &PartyAttributesRiskSummaryItems0{

		RiskScore: defaults.GetString("PartyAttributesRiskSummaryItems0", "risk_score"),

		RiskScoreAlignment: defaults.GetString("PartyAttributesRiskSummaryItems0", "risk_score_alignment"),
	}
}

func (m *PartyAttributesRiskSummaryItems0) WithRiskScore(riskScore string) *PartyAttributesRiskSummaryItems0 {

	m.RiskScore = riskScore

	return m
}

func (m *PartyAttributesRiskSummaryItems0) WithRiskScoreAlignment(riskScoreAlignment string) *PartyAttributesRiskSummaryItems0 {

	m.RiskScoreAlignment = riskScoreAlignment

	return m
}

// Validate validates this party attributes risk summary items0
func (m *PartyAttributesRiskSummaryItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartyAttributesRiskSummaryItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyAttributesRiskSummaryItems0) UnmarshalBinary(b []byte) error {
	var res PartyAttributesRiskSummaryItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyAttributesRiskSummaryItems0) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
