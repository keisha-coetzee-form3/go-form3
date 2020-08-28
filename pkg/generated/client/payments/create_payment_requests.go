// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// Client.CreatePayment creates a new CreatePaymentRequest object
// with the default values initialized.
func (c *Client) CreatePayment() *CreatePaymentRequest {
	var ()
	return &CreatePaymentRequest{

		PaymentCreation: models.PaymentCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentRequest struct {

	/*PaymentCreationRequest*/

	*models.PaymentCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentRequest) FromJson(j string) *CreatePaymentRequest {

	var m models.PaymentCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.PaymentCreation = &m

	return o
}

func (o *CreatePaymentRequest) WithPaymentCreationRequest(paymentCreationRequest models.PaymentCreation) *CreatePaymentRequest {

	o.PaymentCreation = &paymentCreationRequest

	return o
}

func (o *CreatePaymentRequest) WithoutPaymentCreationRequest() *CreatePaymentRequest {

	o.PaymentCreation = &models.PaymentCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create payment Request
func (o *CreatePaymentRequest) WithContext(ctx context.Context) *CreatePaymentRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment Request
func (o *CreatePaymentRequest) WithHTTPClient(client *http.Client) *CreatePaymentRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.PaymentCreation != nil {
		if err := r.SetBodyParam(o.PaymentCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
