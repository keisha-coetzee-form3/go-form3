// Code generated by go-swagger; DO NOT EDIT.

package addressbook

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

// Client.GetAddressbookPartiesIDProductsProductID creates a new GetAddressbookPartiesIDProductsProductIDRequest object
// with the default values initialized.
func (c *Client) GetAddressbookPartiesIDProductsProductID() *GetAddressbookPartiesIDProductsProductIDRequest {
	var ()
	return &GetAddressbookPartiesIDProductsProductIDRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAddressbookPartiesIDProductsProductID", "id"),

		ProductID: c.Defaults.GetStrfmtUUID("GetAddressbookPartiesIDProductsProductID", "product_id"),

		Version: c.Defaults.GetInt64Ptr("GetAddressbookPartiesIDProductsProductID", "version"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAddressbookPartiesIDProductsProductIDRequest struct {

	/*ID      Id of party      */

	ID strfmt.UUID

	/*ProductID      Id of party product      */

	ProductID strfmt.UUID

	/*Version*/

	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAddressbookPartiesIDProductsProductIDRequest) FromJson(j string) *GetAddressbookPartiesIDProductsProductIDRequest {

	return o
}

func (o *GetAddressbookPartiesIDProductsProductIDRequest) WithID(id strfmt.UUID) *GetAddressbookPartiesIDProductsProductIDRequest {

	o.ID = id

	return o
}

func (o *GetAddressbookPartiesIDProductsProductIDRequest) WithProductID(productID strfmt.UUID) *GetAddressbookPartiesIDProductsProductIDRequest {

	o.ProductID = productID

	return o
}

func (o *GetAddressbookPartiesIDProductsProductIDRequest) WithVersion(version int64) *GetAddressbookPartiesIDProductsProductIDRequest {

	o.Version = &version

	return o
}

func (o *GetAddressbookPartiesIDProductsProductIDRequest) WithoutVersion() *GetAddressbookPartiesIDProductsProductIDRequest {

	o.Version = nil

	return o
}

//////////////////
// WithContext adds the context to the get addressbook parties ID products product ID Request
func (o *GetAddressbookPartiesIDProductsProductIDRequest) WithContext(ctx context.Context) *GetAddressbookPartiesIDProductsProductIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get addressbook parties ID products product ID Request
func (o *GetAddressbookPartiesIDProductsProductIDRequest) WithHTTPClient(client *http.Client) *GetAddressbookPartiesIDProductsProductIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAddressbookPartiesIDProductsProductIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param product_id
	if err := r.SetPathParam("product_id", o.ProductID.String()); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}