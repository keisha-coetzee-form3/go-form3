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

// GetPaymentSubmissionReader is a Reader for the GetPaymentSubmission structure.
type GetPaymentSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentSubmissionOK creates a GetPaymentSubmissionOK with default headers values
func NewGetPaymentSubmissionOK() *GetPaymentSubmissionOK {
	return &GetPaymentSubmissionOK{}
}

/*GetPaymentSubmissionOK handles this case with default header values.

Submission details
*/
type GetPaymentSubmissionOK struct {

	//Payload

	// isStream: false
	*models.PaymentSubmissionDetailsResponse
}

func (o *GetPaymentSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/submissions/{submissionId}][%d] getPaymentSubmissionOK", 200)
}

func (o *GetPaymentSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentSubmissionDetailsResponse = new(models.PaymentSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
