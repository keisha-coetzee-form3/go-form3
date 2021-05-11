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

// PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDReader is a Reader for the PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskID structure.
type PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewPatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated creates a PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated with default headers values
func NewPatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated() *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated {
	return &PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated{}
}

/*PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated handles this case with default header values.

Task creation response
*/
type PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated struct {

	//Payload

	// isStream: false
	*models.TaskDetailsResponse
}

func (o *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated) Error() string {
	return fmt.Sprintf("[PATCH /transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}][%d] patchTransactionPaymentsIdAdmissionsAdmissionIdTasksTaskIdCreated", 201)
}

func (o *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.TaskDetailsResponse = new(models.TaskDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.TaskDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict creates a PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict with default headers values
func NewPatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict() *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict {
	return &PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict{}
}

/*PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict handles this case with default header values.

Conflict
*/
type PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}][%d] patchTransactionPaymentsIdAdmissionsAdmissionIdTasksTaskIdConflict", 409)
}

func (o *PatchTransactionPaymentsIDAdmissionsAdmissionIDTasksTaskIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
