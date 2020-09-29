// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsReader is a Reader for the ControllersV1ServiceControllerRemoveServiceAppliedToVMInsts structure.
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted creates a ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted with default headers values
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted() *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted {
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted{}
}

/*ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted handles this case with default header values.

Delete of service instance initiated successfully
*/
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/service_instances/{service_instance_id}][%d] controllersV1ServiceControllerRemoveServiceAppliedToVmInstsAccepted ", 202)
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	return nil
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized creates a ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized with default headers values
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized() *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized {
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized{}
}

/*ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/service_instances/{service_instance_id}][%d] controllersV1ServiceControllerRemoveServiceAppliedToVmInstsUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden creates a ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden with default headers values
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden() *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden {
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden{}
}

/*ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden handles this case with default header values.

Forbidden. Access to the specified resource is not authorized. Check the IAM access policy for the `VMware Solutions` service.
*/
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/service_instances/{service_instance_id}][%d] controllersV1ServiceControllerRemoveServiceAppliedToVmInstsForbidden  %+v", 403, o.Payload)
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound creates a ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound with default headers values
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound() *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound {
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound{}
}

/*ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound handles this case with default header values.

Not found. The resource cannot be found.
*/
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/service_instances/{service_instance_id}][%d] controllersV1ServiceControllerRemoveServiceAppliedToVmInstsNotFound  %+v", 404, o.Payload)
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError creates a ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError with default headers values
func NewControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError() *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError {
	return &ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError{}
}

/*ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/vcenters/{instance_id}/service_instances/{service_instance_id}][%d] controllersV1ServiceControllerRemoveServiceAppliedToVmInstsInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersV1ServiceControllerRemoveServiceAppliedToVMInstsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
