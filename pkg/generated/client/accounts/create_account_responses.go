// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// CreateAccountReader is a Reader for the CreateAccount structure.
type CreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountCreated creates a CreateAccountCreated with default headers values
func NewCreateAccountCreated() *CreateAccountCreated {
	return &CreateAccountCreated{}
}

/*CreateAccountCreated handles this case with default header values.

Account creation response
*/
type CreateAccountCreated struct {

	//Payload

	// isStream: false
	*models.AccountCreationResponse
}

func (o *CreateAccountCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/accounts][%d] createAccountCreated", 201)
}

func (o *CreateAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountCreationResponse = new(models.AccountCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
