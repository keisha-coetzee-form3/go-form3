// Code generated by go-swagger; DO NOT EDIT.

package account_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberReader is a Reader for the GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumber structure.
type GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK creates a GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK with default headers values
func NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK() *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK {
	return &GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK{}
}

/*GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK handles this case with default header values.

Sort code and account number details
*/
type GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK struct {

	//Payload

	// isStream: false
	*models.AccountNumberDetailsResponse
}

func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK) Error() string {
	return fmt.Sprintf("[GET /validations/gbsdc/sortcodes/{sortcode}/accountnumbers/{accountnumber}][%d] getValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK", 200)
}

func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountNumberDetailsResponse = new(models.AccountNumberDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountNumberDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest creates a GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest with default headers values
func NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest() *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest {
	return &GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest{}
}

/*GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest handles this case with default header values.

Validation error
*/
type GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest) Error() string {
	return fmt.Sprintf("[GET /validations/gbsdc/sortcodes/{sortcode}/accountnumbers/{accountnumber}][%d] getValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest", 400)
}

func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound creates a GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound with default headers values
func NewGetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound() *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound {
	return &GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound{}
}

/*GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound handles this case with default header values.

Validation failed
*/
type GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound) Error() string {
	return fmt.Sprintf("[GET /validations/gbsdc/sortcodes/{sortcode}/accountnumbers/{accountnumber}][%d] getValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound", 404)
}

func (o *GetValidationsGbsdcSortcodesSortcodeAccountnumbersAccountnumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
