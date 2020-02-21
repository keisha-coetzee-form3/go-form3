// Code generated by go-swagger; DO NOT EDIT.

package fx_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.ListFXDeals creates a new ListFXDealsRequest object
// with the default values initialized.
func (c *Client) ListFXDeals() *ListFXDealsRequest {
	var ()
	return &ListFXDealsRequest{

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterPartyID: c.Defaults.GetStrfmtUUIDPtr("ListFXDeals", "filter[party_id]"),

		FilterProcessingDateFrom: c.Defaults.GetStrfmtDatePtr("ListFXDeals", "filter[processing_date_from]"),

		FilterProcessingDateTo: c.Defaults.GetStrfmtDatePtr("ListFXDeals", "filter[processing_date_to]"),

		FilterSourceAmount: c.Defaults.GetStringPtr("ListFXDeals", "filter[source.amount]"),

		FilterSourceCurrency: c.Defaults.GetStringPtr("ListFXDeals", "filter[source.currency]"),

		FilterSubmissionReference: c.Defaults.GetStringPtr("ListFXDeals", "filter[submission.reference]"),

		FilterSubmissionStatus: c.Defaults.GetStringPtr("ListFXDeals", "filter[submission.status]"),

		FilterTargetAmount: c.Defaults.GetStringPtr("ListFXDeals", "filter[target.amount]"),

		FilterTargetCurrency: c.Defaults.GetStringPtr("ListFXDeals", "filter[target.currency]"),

		FilterType: c.Defaults.GetStringPtr("ListFXDeals", "filter[type]"),

		PageNumber: c.Defaults.GetStringPtr("ListFXDeals", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListFXDeals", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListFXDealsRequest struct {

	/*FilterOrganisationID      The organisations to filter on      */

	FilterOrganisationID []strfmt.UUID

	/*FilterPartyID      The Party for which to see FX Deals      */

	FilterPartyID *strfmt.UUID

	/*FilterProcessingDateFrom      The processing date from which, inclusively, to find FX Deals      */

	FilterProcessingDateFrom *strfmt.Date

	/*FilterProcessingDateTo      The processing date until which, inclusively, to find FX Deals      */

	FilterProcessingDateTo *strfmt.Date

	/*FilterSourceAmount      The amount of currency to sell      */

	FilterSourceAmount *string

	/*FilterSourceCurrency      The source currency, for example GBP      */

	FilterSourceCurrency *string

	/*FilterSubmissionReference      Reference of the FX Deal Submission associated to look for      */

	FilterSubmissionReference *string

	/*FilterSubmissionStatus      Status of the FX Deal Submission associated to look for      */

	FilterSubmissionStatus *string

	/*FilterTargetAmount      The amount of currency to buy      */

	FilterTargetAmount *string

	/*FilterTargetCurrency      The target currency, for example GBP      */

	FilterTargetCurrency *string

	/*FilterType      The type of foreign exchange      */

	FilterType *string

	/*PageNumber      Which page to select      */

	PageNumber *string

	/*PageSize      Number of items to select      */

	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListFXDealsRequest) FromJson(j string) *ListFXDealsRequest {

	return o
}

func (o *ListFXDealsRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListFXDealsRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListFXDealsRequest) WithoutFilterOrganisationID() *ListFXDealsRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterPartyID(filterPartyID strfmt.UUID) *ListFXDealsRequest {

	o.FilterPartyID = &filterPartyID

	return o
}

func (o *ListFXDealsRequest) WithoutFilterPartyID() *ListFXDealsRequest {

	o.FilterPartyID = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterProcessingDateFrom(filterProcessingDateFrom strfmt.Date) *ListFXDealsRequest {

	o.FilterProcessingDateFrom = &filterProcessingDateFrom

	return o
}

func (o *ListFXDealsRequest) WithoutFilterProcessingDateFrom() *ListFXDealsRequest {

	o.FilterProcessingDateFrom = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterProcessingDateTo(filterProcessingDateTo strfmt.Date) *ListFXDealsRequest {

	o.FilterProcessingDateTo = &filterProcessingDateTo

	return o
}

func (o *ListFXDealsRequest) WithoutFilterProcessingDateTo() *ListFXDealsRequest {

	o.FilterProcessingDateTo = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterSourceAmount(filterSourceAmount string) *ListFXDealsRequest {

	o.FilterSourceAmount = &filterSourceAmount

	return o
}

func (o *ListFXDealsRequest) WithoutFilterSourceAmount() *ListFXDealsRequest {

	o.FilterSourceAmount = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterSourceCurrency(filterSourceCurrency string) *ListFXDealsRequest {

	o.FilterSourceCurrency = &filterSourceCurrency

	return o
}

func (o *ListFXDealsRequest) WithoutFilterSourceCurrency() *ListFXDealsRequest {

	o.FilterSourceCurrency = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterSubmissionReference(filterSubmissionReference string) *ListFXDealsRequest {

	o.FilterSubmissionReference = &filterSubmissionReference

	return o
}

func (o *ListFXDealsRequest) WithoutFilterSubmissionReference() *ListFXDealsRequest {

	o.FilterSubmissionReference = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterSubmissionStatus(filterSubmissionStatus string) *ListFXDealsRequest {

	o.FilterSubmissionStatus = &filterSubmissionStatus

	return o
}

func (o *ListFXDealsRequest) WithoutFilterSubmissionStatus() *ListFXDealsRequest {

	o.FilterSubmissionStatus = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterTargetAmount(filterTargetAmount string) *ListFXDealsRequest {

	o.FilterTargetAmount = &filterTargetAmount

	return o
}

func (o *ListFXDealsRequest) WithoutFilterTargetAmount() *ListFXDealsRequest {

	o.FilterTargetAmount = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterTargetCurrency(filterTargetCurrency string) *ListFXDealsRequest {

	o.FilterTargetCurrency = &filterTargetCurrency

	return o
}

func (o *ListFXDealsRequest) WithoutFilterTargetCurrency() *ListFXDealsRequest {

	o.FilterTargetCurrency = nil

	return o
}

func (o *ListFXDealsRequest) WithFilterType(filterType string) *ListFXDealsRequest {

	o.FilterType = &filterType

	return o
}

func (o *ListFXDealsRequest) WithoutFilterType() *ListFXDealsRequest {

	o.FilterType = nil

	return o
}

func (o *ListFXDealsRequest) WithPageNumber(pageNumber string) *ListFXDealsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListFXDealsRequest) WithoutPageNumber() *ListFXDealsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListFXDealsRequest) WithPageSize(pageSize int64) *ListFXDealsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListFXDealsRequest) WithoutPageSize() *ListFXDealsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list f x deals Request
func (o *ListFXDealsRequest) WithContext(ctx context.Context) *ListFXDealsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list f x deals Request
func (o *ListFXDealsRequest) WithHTTPClient(client *http.Client) *ListFXDealsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListFXDealsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.FilterPartyID != nil {

		// query param filter[party_id]
		var qrFilterPartyID strfmt.UUID
		if o.FilterPartyID != nil {
			qrFilterPartyID = *o.FilterPartyID
		}
		qFilterPartyID := qrFilterPartyID.String()
		if qFilterPartyID != "" {
			if err := r.SetQueryParam("filter[party_id]", qFilterPartyID); err != nil {
				return err
			}
		}

	}

	if o.FilterProcessingDateFrom != nil {

		// query param filter[processing_date_from]
		var qrFilterProcessingDateFrom strfmt.Date
		if o.FilterProcessingDateFrom != nil {
			qrFilterProcessingDateFrom = *o.FilterProcessingDateFrom
		}
		qFilterProcessingDateFrom := qrFilterProcessingDateFrom.String()
		if qFilterProcessingDateFrom != "" {
			if err := r.SetQueryParam("filter[processing_date_from]", qFilterProcessingDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterProcessingDateTo != nil {

		// query param filter[processing_date_to]
		var qrFilterProcessingDateTo strfmt.Date
		if o.FilterProcessingDateTo != nil {
			qrFilterProcessingDateTo = *o.FilterProcessingDateTo
		}
		qFilterProcessingDateTo := qrFilterProcessingDateTo.String()
		if qFilterProcessingDateTo != "" {
			if err := r.SetQueryParam("filter[processing_date_to]", qFilterProcessingDateTo); err != nil {
				return err
			}
		}

	}

	if o.FilterSourceAmount != nil {

		// query param filter[source.amount]
		var qrFilterSourceAmount string
		if o.FilterSourceAmount != nil {
			qrFilterSourceAmount = *o.FilterSourceAmount
		}
		qFilterSourceAmount := qrFilterSourceAmount
		if qFilterSourceAmount != "" {
			if err := r.SetQueryParam("filter[source.amount]", qFilterSourceAmount); err != nil {
				return err
			}
		}

	}

	if o.FilterSourceCurrency != nil {

		// query param filter[source.currency]
		var qrFilterSourceCurrency string
		if o.FilterSourceCurrency != nil {
			qrFilterSourceCurrency = *o.FilterSourceCurrency
		}
		qFilterSourceCurrency := qrFilterSourceCurrency
		if qFilterSourceCurrency != "" {
			if err := r.SetQueryParam("filter[source.currency]", qFilterSourceCurrency); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionReference != nil {

		// query param filter[submission.reference]
		var qrFilterSubmissionReference string
		if o.FilterSubmissionReference != nil {
			qrFilterSubmissionReference = *o.FilterSubmissionReference
		}
		qFilterSubmissionReference := qrFilterSubmissionReference
		if qFilterSubmissionReference != "" {
			if err := r.SetQueryParam("filter[submission.reference]", qFilterSubmissionReference); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionStatus != nil {

		// query param filter[submission.status]
		var qrFilterSubmissionStatus string
		if o.FilterSubmissionStatus != nil {
			qrFilterSubmissionStatus = *o.FilterSubmissionStatus
		}
		qFilterSubmissionStatus := qrFilterSubmissionStatus
		if qFilterSubmissionStatus != "" {
			if err := r.SetQueryParam("filter[submission.status]", qFilterSubmissionStatus); err != nil {
				return err
			}
		}

	}

	if o.FilterTargetAmount != nil {

		// query param filter[target.amount]
		var qrFilterTargetAmount string
		if o.FilterTargetAmount != nil {
			qrFilterTargetAmount = *o.FilterTargetAmount
		}
		qFilterTargetAmount := qrFilterTargetAmount
		if qFilterTargetAmount != "" {
			if err := r.SetQueryParam("filter[target.amount]", qFilterTargetAmount); err != nil {
				return err
			}
		}

	}

	if o.FilterTargetCurrency != nil {

		// query param filter[target.currency]
		var qrFilterTargetCurrency string
		if o.FilterTargetCurrency != nil {
			qrFilterTargetCurrency = *o.FilterTargetCurrency
		}
		qFilterTargetCurrency := qrFilterTargetCurrency
		if qFilterTargetCurrency != "" {
			if err := r.SetQueryParam("filter[target.currency]", qFilterTargetCurrency); err != nil {
				return err
			}
		}

	}

	if o.FilterType != nil {

		// query param filter[type]
		var qrFilterType string
		if o.FilterType != nil {
			qrFilterType = *o.FilterType
		}
		qFilterType := qrFilterType
		if qFilterType != "" {
			if err := r.SetQueryParam("filter[type]", qFilterType); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}