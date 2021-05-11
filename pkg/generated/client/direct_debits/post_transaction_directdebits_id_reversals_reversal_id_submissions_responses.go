// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsReader is a Reader for the PostTransactionDirectdebitsIDReversalsReversalIDSubmissions structure.
type PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated creates a PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated with default headers values
func NewPostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated() *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated {
	return &PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated{}
}

/*PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated handles this case with default header values.

Reversal submission creation response
*/
type PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitReversalSubmissionCreationResponse
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/reversals/{reversalId}/submissions][%d] postTransactionDirectdebitsIdReversalsReversalIdSubmissionsCreated", 201)
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReversalSubmissionCreationResponse = new(models.DirectDebitReversalSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReversalSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest creates a PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest with default headers values
func NewPostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest() *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest {
	return &PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest{}
}

/*PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest handles this case with default header values.

Reversal submission creation error
*/
type PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/reversals/{reversalId}/submissions][%d] postTransactionDirectdebitsIdReversalsReversalIdSubmissionsBadRequest", 400)
}

func (o *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}