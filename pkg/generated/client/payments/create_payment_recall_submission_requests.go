// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// Client.CreatePaymentRecallSubmission creates a new CreatePaymentRecallSubmissionRequest object
// with the default values initialized.
func (c *Client) CreatePaymentRecallSubmission() *CreatePaymentRecallSubmissionRequest {
	var ()
	return &CreatePaymentRecallSubmissionRequest{

		RecallSubmissionCreation: models.RecallSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentRecallSubmission", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("CreatePaymentRecallSubmission", "recallId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentRecallSubmissionRequest struct {

	/*RecallSubmissionCreationRequest*/

	*models.RecallSubmissionCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentRecallSubmissionRequest) FromJson(j string) (*CreatePaymentRecallSubmissionRequest, error) {

	var m models.RecallSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.RecallSubmissionCreation = &m

	return o, nil
}

func (o *CreatePaymentRecallSubmissionRequest) WithRecallSubmissionCreationRequest(recallSubmissionCreationRequest models.RecallSubmissionCreation) *CreatePaymentRecallSubmissionRequest {

	o.RecallSubmissionCreation = &recallSubmissionCreationRequest

	return o
}

func (o *CreatePaymentRecallSubmissionRequest) WithoutRecallSubmissionCreationRequest() *CreatePaymentRecallSubmissionRequest {

	o.RecallSubmissionCreation = &models.RecallSubmissionCreation{}

	return o
}

func (o *CreatePaymentRecallSubmissionRequest) WithID(id strfmt.UUID) *CreatePaymentRecallSubmissionRequest {

	o.ID = id

	return o
}

func (o *CreatePaymentRecallSubmissionRequest) WithRecallID(recallID strfmt.UUID) *CreatePaymentRecallSubmissionRequest {

	o.RecallID = recallID

	return o
}

//////////////////
// WithContext adds the context to the create payment recall submission Request
func (o *CreatePaymentRecallSubmissionRequest) WithContext(ctx context.Context) *CreatePaymentRecallSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment recall submission Request
func (o *CreatePaymentRecallSubmissionRequest) WithHTTPClient(client *http.Client) *CreatePaymentRecallSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentRecallSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.RecallSubmissionCreation != nil {
		if err := r.SetBodyParam(o.RecallSubmissionCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
