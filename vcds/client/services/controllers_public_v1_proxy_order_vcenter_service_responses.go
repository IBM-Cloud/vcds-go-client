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

// ControllersPublicV1ProxyOrderVcenterServiceReader is a Reader for the ControllersPublicV1ProxyOrderVcenterService structure.
type ControllersPublicV1ProxyOrderVcenterServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV1ProxyOrderVcenterServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllersPublicV1ProxyOrderVcenterServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewControllersPublicV1ProxyOrderVcenterServiceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewControllersPublicV1ProxyOrderVcenterServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewControllersPublicV1ProxyOrderVcenterServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllersPublicV1ProxyOrderVcenterServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllersPublicV1ProxyOrderVcenterServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewControllersPublicV1ProxyOrderVcenterServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV1ProxyOrderVcenterServiceOK creates a ControllersPublicV1ProxyOrderVcenterServiceOK with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceOK() *ControllersPublicV1ProxyOrderVcenterServiceOK {
	return &ControllersPublicV1ProxyOrderVcenterServiceOK{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceOK handles this case with default header values.

Success. The request has been successfully verified.
*/
type ControllersPublicV1ProxyOrderVcenterServiceOK struct {
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceOK ", 200)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControllersPublicV1ProxyOrderVcenterServiceAccepted creates a ControllersPublicV1ProxyOrderVcenterServiceAccepted with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceAccepted() *ControllersPublicV1ProxyOrderVcenterServiceAccepted {
	return &ControllersPublicV1ProxyOrderVcenterServiceAccepted{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceAccepted handles this case with default header values.

Success. The request for adding the new service has been accepted.
*/
type ControllersPublicV1ProxyOrderVcenterServiceAccepted struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.ServiceOrderResponse
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceAccepted  %+v", 202, o.Payload)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.ServiceOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyOrderVcenterServiceBadRequest creates a ControllersPublicV1ProxyOrderVcenterServiceBadRequest with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceBadRequest() *ControllersPublicV1ProxyOrderVcenterServiceBadRequest {
	return &ControllersPublicV1ProxyOrderVcenterServiceBadRequest{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceBadRequest handles this case with default header values.

Bad request. Check your request parameters.
*/
type ControllersPublicV1ProxyOrderVcenterServiceBadRequest struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceBadRequest  %+v", 400, o.Payload)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyOrderVcenterServiceUnauthorized creates a ControllersPublicV1ProxyOrderVcenterServiceUnauthorized with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceUnauthorized() *ControllersPublicV1ProxyOrderVcenterServiceUnauthorized {
	return &ControllersPublicV1ProxyOrderVcenterServiceUnauthorized{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersPublicV1ProxyOrderVcenterServiceUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyOrderVcenterServiceForbidden creates a ControllersPublicV1ProxyOrderVcenterServiceForbidden with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceForbidden() *ControllersPublicV1ProxyOrderVcenterServiceForbidden {
	return &ControllersPublicV1ProxyOrderVcenterServiceForbidden{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceForbidden handles this case with default header values.

Forbidden. Access to the specified resource is not authorized. Check the IAM access policy for the `VMware Solutions` service.
*/
type ControllersPublicV1ProxyOrderVcenterServiceForbidden struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceForbidden  %+v", 403, o.Payload)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyOrderVcenterServiceNotFound creates a ControllersPublicV1ProxyOrderVcenterServiceNotFound with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceNotFound() *ControllersPublicV1ProxyOrderVcenterServiceNotFound {
	return &ControllersPublicV1ProxyOrderVcenterServiceNotFound{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceNotFound handles this case with default header values.

Not found. The resource cannot be found.
*/
type ControllersPublicV1ProxyOrderVcenterServiceNotFound struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceNotFound  %+v", 404, o.Payload)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyOrderVcenterServiceInternalServerError creates a ControllersPublicV1ProxyOrderVcenterServiceInternalServerError with default headers values
func NewControllersPublicV1ProxyOrderVcenterServiceInternalServerError() *ControllersPublicV1ProxyOrderVcenterServiceInternalServerError {
	return &ControllersPublicV1ProxyOrderVcenterServiceInternalServerError{}
}

/*ControllersPublicV1ProxyOrderVcenterServiceInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersPublicV1ProxyOrderVcenterServiceInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vcenters/{instance_id}/service_instances][%d] controllersPublicV1ProxyOrderVcenterServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersPublicV1ProxyOrderVcenterServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
