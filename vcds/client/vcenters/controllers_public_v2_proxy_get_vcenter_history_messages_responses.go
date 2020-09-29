// Code generated by go-swagger; DO NOT EDIT.

package vcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersPublicV2ProxyGetVcenterHistoryMessagesReader is a Reader for the ControllersPublicV2ProxyGetVcenterHistoryMessages structure.
type ControllersPublicV2ProxyGetVcenterHistoryMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllersPublicV2ProxyGetVcenterHistoryMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV2ProxyGetVcenterHistoryMessagesOK creates a ControllersPublicV2ProxyGetVcenterHistoryMessagesOK with default headers values
func NewControllersPublicV2ProxyGetVcenterHistoryMessagesOK() *ControllersPublicV2ProxyGetVcenterHistoryMessagesOK {
	return &ControllersPublicV2ProxyGetVcenterHistoryMessagesOK{}
}

/*ControllersPublicV2ProxyGetVcenterHistoryMessagesOK handles this case with default header values.

Success.
*/
type ControllersPublicV2ProxyGetVcenterHistoryMessagesOK struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.DefHistoryMessages
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesOK) Error() string {
	return fmt.Sprintf("[GET /v2/vcenters/{instance_id}/history][%d] controllersPublicV2ProxyGetVcenterHistoryMessagesOK  %+v", 200, o.Payload)
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.DefHistoryMessages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized creates a ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized with default headers values
func NewControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized() *ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized {
	return &ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized{}
}

/*ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/vcenters/{instance_id}/history][%d] controllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden creates a ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden with default headers values
func NewControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden() *ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden {
	return &ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden{}
}

/*ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden handles this case with default header values.

Forbidden. Access to the specified resource is not authorized. Check the IAM access policy for the `VMware Solutions` service.
*/
type ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/vcenters/{instance_id}/history][%d] controllersPublicV2ProxyGetVcenterHistoryMessagesForbidden  %+v", 403, o.Payload)
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound creates a ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound with default headers values
func NewControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound() *ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound {
	return &ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound{}
}

/*ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound handles this case with default header values.

Not found. The resource cannot be found.
*/
type ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/vcenters/{instance_id}/history][%d] controllersPublicV2ProxyGetVcenterHistoryMessagesNotFound  %+v", 404, o.Payload)
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError creates a ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError with default headers values
func NewControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError() *ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError {
	return &ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError{}
}

/*ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/vcenters/{instance_id}/history][%d] controllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersPublicV2ProxyGetVcenterHistoryMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}