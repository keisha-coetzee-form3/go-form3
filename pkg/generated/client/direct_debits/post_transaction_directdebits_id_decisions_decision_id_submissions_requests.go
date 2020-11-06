// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissions creates a new PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest object
// with the default values initialized.
func (c *Client) PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissions() *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {
	var ()
	return &PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest{

		DirectDebitDecisionSubmissionCreation: models.DirectDebitDecisionSubmissionCreationWithDefaults(c.Defaults),

		DecisionID: c.Defaults.GetStrfmtUUID("PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissions", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissions", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest struct {

	/*DirectDebitDecisionSubmissionCreationRequest*/

	*models.DirectDebitDecisionSubmissionCreation

	/*DecisionID      Direct Debit decision id      */

	DecisionID strfmt.UUID

	/*ID      Direct Debit id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) FromJson(j string) *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {

	var m models.DirectDebitDecisionSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.DirectDebitDecisionSubmissionCreation = &m

	return o
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WithDirectDebitDecisionSubmissionCreationRequest(directDebitDecisionSubmissionCreationRequest models.DirectDebitDecisionSubmissionCreation) *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {

	o.DirectDebitDecisionSubmissionCreation = &directDebitDecisionSubmissionCreationRequest

	return o
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WithoutDirectDebitDecisionSubmissionCreationRequest() *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {

	o.DirectDebitDecisionSubmissionCreation = &models.DirectDebitDecisionSubmissionCreation{}

	return o
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WithDecisionID(decisionID strfmt.UUID) *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {

	o.DecisionID = decisionID

	return o
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WithID(id strfmt.UUID) *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the post transaction directdebits ID decisions decision ID submissions Request
func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WithContext(ctx context.Context) *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the post transaction directdebits ID decisions decision ID submissions Request
func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WithHTTPClient(client *http.Client) *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitDecisionSubmissionCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitDecisionSubmissionCreation); err != nil {
			return err
		}
	}

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
