// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// PostAddressbookPartiesIDAccountsReader is a Reader for the PostAddressbookPartiesIDAccounts structure.
type PostAddressbookPartiesIDAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAddressbookPartiesIDAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAddressbookPartiesIDAccountsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAddressbookPartiesIDAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAddressbookPartiesIDAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostAddressbookPartiesIDAccountsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAddressbookPartiesIDAccountsCreated creates a PostAddressbookPartiesIDAccountsCreated with default headers values
func NewPostAddressbookPartiesIDAccountsCreated() *PostAddressbookPartiesIDAccountsCreated {
	return &PostAddressbookPartiesIDAccountsCreated{}
}

/*PostAddressbookPartiesIDAccountsCreated handles this case with default header values.

creation response
*/
type PostAddressbookPartiesIDAccountsCreated struct {

	//Payload

	// isStream: false
	*models.PartyAccountCreationResponse
}

func (o *PostAddressbookPartiesIDAccountsCreated) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/accounts][%d] postAddressbookPartiesIdAccountsCreated", 201)
}

func (o *PostAddressbookPartiesIDAccountsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PartyAccountCreationResponse = new(models.PartyAccountCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PartyAccountCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDAccountsBadRequest creates a PostAddressbookPartiesIDAccountsBadRequest with default headers values
func NewPostAddressbookPartiesIDAccountsBadRequest() *PostAddressbookPartiesIDAccountsBadRequest {
	return &PostAddressbookPartiesIDAccountsBadRequest{}
}

/*PostAddressbookPartiesIDAccountsBadRequest handles this case with default header values.

Bad Request
*/
type PostAddressbookPartiesIDAccountsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDAccountsBadRequest) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/accounts][%d] postAddressbookPartiesIdAccountsBadRequest", 400)
}

func (o *PostAddressbookPartiesIDAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDAccountsForbidden creates a PostAddressbookPartiesIDAccountsForbidden with default headers values
func NewPostAddressbookPartiesIDAccountsForbidden() *PostAddressbookPartiesIDAccountsForbidden {
	return &PostAddressbookPartiesIDAccountsForbidden{}
}

/*PostAddressbookPartiesIDAccountsForbidden handles this case with default header values.

Forbidden
*/
type PostAddressbookPartiesIDAccountsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDAccountsForbidden) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/accounts][%d] postAddressbookPartiesIdAccountsForbidden", 403)
}

func (o *PostAddressbookPartiesIDAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDAccountsConflict creates a PostAddressbookPartiesIDAccountsConflict with default headers values
func NewPostAddressbookPartiesIDAccountsConflict() *PostAddressbookPartiesIDAccountsConflict {
	return &PostAddressbookPartiesIDAccountsConflict{}
}

/*PostAddressbookPartiesIDAccountsConflict handles this case with default header values.

Conflict
*/
type PostAddressbookPartiesIDAccountsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDAccountsConflict) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/accounts][%d] postAddressbookPartiesIdAccountsConflict", 409)
}

func (o *PostAddressbookPartiesIDAccountsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
