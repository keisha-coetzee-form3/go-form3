// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// GetAddressbookPartiesIDActorsReader is a Reader for the GetAddressbookPartiesIDActors structure.
type GetAddressbookPartiesIDActorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesIDActorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesIDActorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesIDActorsOK creates a GetAddressbookPartiesIDActorsOK with default headers values
func NewGetAddressbookPartiesIDActorsOK() *GetAddressbookPartiesIDActorsOK {
	return &GetAddressbookPartiesIDActorsOK{}
}

/*GetAddressbookPartiesIDActorsOK handles this case with default header values.

list party actors reponse
*/
type GetAddressbookPartiesIDActorsOK struct {

	//Payload

	// isStream: false
	*models.ListPartyActorsResponse
}

func (o *GetAddressbookPartiesIDActorsOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/actors][%d] getAddressbookPartiesIdActorsOK", 200)
}

func (o *GetAddressbookPartiesIDActorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ListPartyActorsResponse = new(models.ListPartyActorsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ListPartyActorsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
