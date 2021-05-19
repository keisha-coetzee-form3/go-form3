// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceType resource type
// swagger:model ResourceType
type ResourceType string

const (

	// ResourceTypePaymentAdviceSubmissionValidations captures enum value "payment_advice_submission_validations"
	ResourceTypePaymentAdviceSubmissionValidations ResourceType = "payment_advice_submission_validations"

	// ResourceTypePaymentBatches captures enum value "payment_batches"
	ResourceTypePaymentBatches ResourceType = "payment_batches"

	// ResourceTypePartyProductUpdatedEvents captures enum value "party_product_updated_events"
	ResourceTypePartyProductUpdatedEvents ResourceType = "party_product_updated_events"

	// ResourceTypePaymentSubmissionValidations captures enum value "payment_submission_validations"
	ResourceTypePaymentSubmissionValidations ResourceType = "payment_submission_validations"

	// ResourceTypeAccountRequests captures enum value "account_requests"
	ResourceTypeAccountRequests ResourceType = "account_requests"

	// ResourceTypePaymentAdvices captures enum value "payment_advices"
	ResourceTypePaymentAdvices ResourceType = "payment_advices"

	// ResourceTypeProductAssociations captures enum value "product_associations"
	ResourceTypeProductAssociations ResourceType = "product_associations"

	// ResourceTypeAccountRequestSubmissionValidations captures enum value "account_request_submission_validations"
	ResourceTypeAccountRequestSubmissionValidations ResourceType = "account_request_submission_validations"

	// ResourceTypeBranches captures enum value "branches"
	ResourceTypeBranches ResourceType = "branches"

	// ResourceTypePartyAccounts captures enum value "party_accounts"
	ResourceTypePartyAccounts ResourceType = "party_accounts"

	// ResourceTypeRecallDecisionSubmissions captures enum value "recall_decision_submissions"
	ResourceTypeRecallDecisionSubmissions ResourceType = "recall_decision_submissions"

	// ResourceTypeContactAccounts captures enum value "contact_accounts"
	ResourceTypeContactAccounts ResourceType = "contact_accounts"

	// ResourceTypeParties captures enum value "parties"
	ResourceTypeParties ResourceType = "parties"

	// ResourceTypeReturns captures enum value "returns"
	ResourceTypeReturns ResourceType = "returns"

	// ResourceTypeRecalls captures enum value "recalls"
	ResourceTypeRecalls ResourceType = "recalls"

	// ResourceTypeAccountRoutings captures enum value "account_routings"
	ResourceTypeAccountRoutings ResourceType = "account_routings"

	// ResourceTypeDirectAccount captures enum value "direct_account"
	ResourceTypeDirectAccount ResourceType = "direct_account"

	// ResourceTypeAccountAmendments captures enum value "account_amendments"
	ResourceTypeAccountAmendments ResourceType = "account_amendments"

	// ResourceTypeReversals captures enum value "reversals"
	ResourceTypeReversals ResourceType = "reversals"

	// ResourceTypeRecallSubmissionValidations captures enum value "recall_submission_validations"
	ResourceTypeRecallSubmissionValidations ResourceType = "recall_submission_validations"

	// ResourceTypeAccountAmendmentSubmissions captures enum value "account_amendment_submissions"
	ResourceTypeAccountAmendmentSubmissions ResourceType = "account_amendment_submissions"

	// ResourceTypePaymentAdviceSubmissions captures enum value "payment_advice_submissions"
	ResourceTypePaymentAdviceSubmissions ResourceType = "payment_advice_submissions"

	// ResourceTypeAccountAmendmentSubmissionValidations captures enum value "account_amendment_submission_validations"
	ResourceTypeAccountAmendmentSubmissionValidations ResourceType = "account_amendment_submission_validations"

	// ResourceTypeRecallReversals captures enum value "recall_reversals"
	ResourceTypeRecallReversals ResourceType = "recall_reversals"

	// ResourceTypeContacts captures enum value "contacts"
	ResourceTypeContacts ResourceType = "contacts"

	// ResourceTypeRecallDecisionSubmissionValidations captures enum value "recall_decision_submission_validations"
	ResourceTypeRecallDecisionSubmissionValidations ResourceType = "recall_decision_submission_validations"

	// ResourceTypeRecallDecisionAdmissions captures enum value "recall_decision_admissions"
	ResourceTypeRecallDecisionAdmissions ResourceType = "recall_decision_admissions"

	// ResourceTypeRecallReversalAdmissions captures enum value "recall_reversal_admissions"
	ResourceTypeRecallReversalAdmissions ResourceType = "recall_reversal_admissions"

	// ResourceTypeBankIds captures enum value "bank_ids"
	ResourceTypeBankIds ResourceType = "bank_ids"

	// ResourceTypeAccountConfigurations captures enum value "account_configurations"
	ResourceTypeAccountConfigurations ResourceType = "account_configurations"

	// ResourceTypePaymentAutomaticReturns captures enum value "payment_automatic_returns"
	ResourceTypePaymentAutomaticReturns ResourceType = "payment_automatic_returns"

	// ResourceTypeAccountIndirects captures enum value "account_indirects"
	ResourceTypeAccountIndirects ResourceType = "account_indirects"

	// ResourceTypeSchemeMessageAdmissions captures enum value "scheme_message_admissions"
	ResourceTypeSchemeMessageAdmissions ResourceType = "scheme_message_admissions"

	// ResourceTypeBics captures enum value "bics"
	ResourceTypeBics ResourceType = "bics"

	// ResourceTypeSchemeMessages captures enum value "scheme_messages"
	ResourceTypeSchemeMessages ResourceType = "scheme_messages"

	// ResourceTypeReversalSubmissionValidations captures enum value "reversal_submission_validations"
	ResourceTypeReversalSubmissionValidations ResourceType = "reversal_submission_validations"

	// ResourceTypePaymentSubmissions captures enum value "payment_submissions"
	ResourceTypePaymentSubmissions ResourceType = "payment_submissions"

	// ResourceTypeReturnSubmissions captures enum value "return_submissions"
	ResourceTypeReturnSubmissions ResourceType = "return_submissions"

	// ResourceTypeReversalAdmissions captures enum value "reversal_admissions"
	ResourceTypeReversalAdmissions ResourceType = "reversal_admissions"

	// ResourceTypeReturnSubmissionValidations captures enum value "return_submission_validations"
	ResourceTypeReturnSubmissionValidations ResourceType = "return_submission_validations"

	// ResourceTypePaymentDefaults captures enum value "payment_defaults"
	ResourceTypePaymentDefaults ResourceType = "payment_defaults"

	// ResourceTypeLimits captures enum value "limits"
	ResourceTypeLimits ResourceType = "limits"

	// ResourceTypePartyProducts captures enum value "party_products"
	ResourceTypePartyProducts ResourceType = "party_products"

	// ResourceTypeReturnReversalAdmissions captures enum value "return_reversal_admissions"
	ResourceTypeReturnReversalAdmissions ResourceType = "return_reversal_admissions"

	// ResourceTypePaymentAdmissionTasks captures enum value "payment_admission_tasks"
	ResourceTypePaymentAdmissionTasks ResourceType = "payment_admission_tasks"

	// ResourceTypePaymentAdmissions captures enum value "payment_admissions"
	ResourceTypePaymentAdmissions ResourceType = "payment_admissions"

	// ResourceTypeReturnAdmissions captures enum value "return_admissions"
	ResourceTypeReturnAdmissions ResourceType = "return_admissions"

	// ResourceTypePayments captures enum value "payments"
	ResourceTypePayments ResourceType = "payments"

	// ResourceTypeAccounts captures enum value "accounts"
	ResourceTypeAccounts ResourceType = "accounts"

	// ResourceTypeRecallDecisions captures enum value "recall_decisions"
	ResourceTypeRecallDecisions ResourceType = "recall_decisions"

	// ResourceTypePartyProductEvents captures enum value "party_product_events"
	ResourceTypePartyProductEvents ResourceType = "party_product_events"

	// ResourceTypeDirectDebit captures enum value "direct_debit"
	ResourceTypeDirectDebit ResourceType = "direct_debit"

	// ResourceTypePositions captures enum value "positions"
	ResourceTypePositions ResourceType = "positions"

	// ResourceTypeReversalSubmissions captures enum value "reversal_submissions"
	ResourceTypeReversalSubmissions ResourceType = "reversal_submissions"

	// ResourceTypeAccountEvents captures enum value "account_events"
	ResourceTypeAccountEvents ResourceType = "account_events"

	// ResourceTypeRecallSubmissions captures enum value "recall_submissions"
	ResourceTypeRecallSubmissions ResourceType = "recall_submissions"

	// ResourceTypeFxDeals captures enum value "fx_deals"
	ResourceTypeFxDeals ResourceType = "fx_deals"

	// ResourceTypeReturnReversals captures enum value "return_reversals"
	ResourceTypeReturnReversals ResourceType = "return_reversals"

	// ResourceTypeRecallAdmissions captures enum value "recall_admissions"
	ResourceTypeRecallAdmissions ResourceType = "recall_admissions"

	// ResourceTypeAccountRequestSubmissions captures enum value "account_request_submissions"
	ResourceTypeAccountRequestSubmissions ResourceType = "account_request_submissions"

	// ResourceTypePartyActors captures enum value "party_actors"
	ResourceTypePartyActors ResourceType = "party_actors"
)

// for schema
var resourceTypeEnum []interface{}

func init() {
	var res []ResourceType
	if err := json.Unmarshal([]byte(`["payment_advice_submission_validations","payment_batches","party_product_updated_events","payment_submission_validations","account_requests","payment_advices","product_associations","account_request_submission_validations","branches","party_accounts","recall_decision_submissions","contact_accounts","parties","returns","recalls","account_routings","direct_account","account_amendments","reversals","recall_submission_validations","account_amendment_submissions","payment_advice_submissions","account_amendment_submission_validations","recall_reversals","contacts","recall_decision_submission_validations","recall_decision_admissions","recall_reversal_admissions","bank_ids","account_configurations","payment_automatic_returns","account_indirects","scheme_message_admissions","bics","scheme_messages","reversal_submission_validations","payment_submissions","return_submissions","reversal_admissions","return_submission_validations","payment_defaults","limits","party_products","return_reversal_admissions","payment_admission_tasks","payment_admissions","return_admissions","payments","accounts","recall_decisions","party_product_events","direct_debit","positions","reversal_submissions","account_events","recall_submissions","fx_deals","return_reversals","recall_admissions","account_request_submissions","party_actors"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeEnum = append(resourceTypeEnum, v)
	}
}

func (m ResourceType) validateResourceTypeEnum(path, location string, value ResourceType) error {
	if err := validate.Enum(path, location, value, resourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this resource type
func (m ResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
