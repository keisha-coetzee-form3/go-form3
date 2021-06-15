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

// Client.CreatePaymentRecall creates a new CreatePaymentRecallRequest object
// with the default values initialized.
func (c *Client) CreatePaymentRecall() *CreatePaymentRecallRequest {
	var ()
	return &CreatePaymentRecallRequest{

		RecallCreation: models.RecallCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentRecall", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentRecallRequest struct {

	/*RecallCreationRequest*/

	*models.RecallCreation

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentRecallRequest) FromJson(j string) (*CreatePaymentRecallRequest, error) {

	var m models.RecallCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.RecallCreation = &m

	return o, nil
}

func (o *CreatePaymentRecallRequest) WithRecallCreationRequest(recallCreationRequest models.RecallCreation) *CreatePaymentRecallRequest {

	o.RecallCreation = &recallCreationRequest

	return o
}

func (o *CreatePaymentRecallRequest) WithoutRecallCreationRequest() *CreatePaymentRecallRequest {

	o.RecallCreation = &models.RecallCreation{}

	return o
}

func (o *CreatePaymentRecallRequest) WithID(id strfmt.UUID) *CreatePaymentRecallRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create payment recall Request
func (o *CreatePaymentRecallRequest) WithContext(ctx context.Context) *CreatePaymentRecallRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment recall Request
func (o *CreatePaymentRecallRequest) WithHTTPClient(client *http.Client) *CreatePaymentRecallRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentRecallRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.RecallCreation != nil {
		if err := r.SetBodyParam(o.RecallCreation); err != nil {
			return err
		}
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
