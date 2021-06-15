// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// CreateDirectDebitSubmissionReader is a Reader for the CreateDirectDebitSubmission structure.
type CreateDirectDebitSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDirectDebitSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDirectDebitSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDirectDebitSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDirectDebitSubmissionCreated creates a CreateDirectDebitSubmissionCreated with default headers values
func NewCreateDirectDebitSubmissionCreated() *CreateDirectDebitSubmissionCreated {
	return &CreateDirectDebitSubmissionCreated{}
}

/*CreateDirectDebitSubmissionCreated handles this case with default header values.

Direct debit submission creation response
*/
type CreateDirectDebitSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitSubmissionCreationResponse
}

func (o *CreateDirectDebitSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/submissions][%d] createDirectDebitSubmissionCreated", 201)
}

func (o *CreateDirectDebitSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitSubmissionCreationResponse = new(models.DirectDebitSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDirectDebitSubmissionBadRequest creates a CreateDirectDebitSubmissionBadRequest with default headers values
func NewCreateDirectDebitSubmissionBadRequest() *CreateDirectDebitSubmissionBadRequest {
	return &CreateDirectDebitSubmissionBadRequest{}
}

/*CreateDirectDebitSubmissionBadRequest handles this case with default header values.

Return submission creation error
*/
type CreateDirectDebitSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateDirectDebitSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/submissions][%d] createDirectDebitSubmissionBadRequest", 400)
}

func (o *CreateDirectDebitSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
