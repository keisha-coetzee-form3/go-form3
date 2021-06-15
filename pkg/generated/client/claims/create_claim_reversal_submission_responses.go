// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// CreateClaimReversalSubmissionReader is a Reader for the CreateClaimReversalSubmission structure.
type CreateClaimReversalSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClaimReversalSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateClaimReversalSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClaimReversalSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateClaimReversalSubmissionCreated creates a CreateClaimReversalSubmissionCreated with default headers values
func NewCreateClaimReversalSubmissionCreated() *CreateClaimReversalSubmissionCreated {
	return &CreateClaimReversalSubmissionCreated{}
}

/*CreateClaimReversalSubmissionCreated handles this case with default header values.

Claim Reversal Submission creation response
*/
type CreateClaimReversalSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.ClaimReversalSubmissionDetailsResponse
}

func (o *CreateClaimReversalSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/claims/{id}/reversals/{reversalId}/submissions][%d] createClaimReversalSubmissionCreated", 201)
}

func (o *CreateClaimReversalSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimReversalSubmissionDetailsResponse = new(models.ClaimReversalSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ClaimReversalSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClaimReversalSubmissionBadRequest creates a CreateClaimReversalSubmissionBadRequest with default headers values
func NewCreateClaimReversalSubmissionBadRequest() *CreateClaimReversalSubmissionBadRequest {
	return &CreateClaimReversalSubmissionBadRequest{}
}

/*CreateClaimReversalSubmissionBadRequest handles this case with default header values.

Claim Reversal creation error
*/
type CreateClaimReversalSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateClaimReversalSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/claims/{id}/reversals/{reversalId}/submissions][%d] createClaimReversalSubmissionBadRequest", 400)
}

func (o *CreateClaimReversalSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
