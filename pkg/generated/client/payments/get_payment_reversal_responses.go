// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// GetPaymentReversalReader is a Reader for the GetPaymentReversal structure.
type GetPaymentReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReversalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReversalOK creates a GetPaymentReversalOK with default headers values
func NewGetPaymentReversalOK() *GetPaymentReversalOK {
	return &GetPaymentReversalOK{}
}

/*GetPaymentReversalOK handles this case with default header values.

Reversal details
*/
type GetPaymentReversalOK struct {

	//Payload

	// isStream: false
	*models.ReversalDetailsResponse
}

func (o *GetPaymentReversalOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/reversals/{reversalId}][%d] getPaymentReversalOK", 200)
}

func (o *GetPaymentReversalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalDetailsResponse = new(models.ReversalDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
