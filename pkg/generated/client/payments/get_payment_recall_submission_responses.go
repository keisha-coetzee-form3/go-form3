// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetPaymentRecallSubmissionReader is a Reader for the GetPaymentRecallSubmission structure.
type GetPaymentRecallSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallSubmissionOK creates a GetPaymentRecallSubmissionOK with default headers values
func NewGetPaymentRecallSubmissionOK() *GetPaymentRecallSubmissionOK {
	return &GetPaymentRecallSubmissionOK{}
}

/*GetPaymentRecallSubmissionOK handles this case with default header values.

Recall submission details
*/
type GetPaymentRecallSubmissionOK struct {

	//Payload

	// isStream: false
	*models.RecallSubmissionDetailsResponse
}

func (o *GetPaymentRecallSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/submissions/{submissionId}][%d] getPaymentRecallSubmissionOK", 200)
}

func (o *GetPaymentRecallSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallSubmissionDetailsResponse = new(models.RecallSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
