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

// ControllersPublicV1ProxyGetVcenterClusterReader is a Reader for the ControllersPublicV1ProxyGetVcenterCluster structure.
type ControllersPublicV1ProxyGetVcenterClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV1ProxyGetVcenterClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllersPublicV1ProxyGetVcenterClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewControllersPublicV1ProxyGetVcenterClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllersPublicV1ProxyGetVcenterClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllersPublicV1ProxyGetVcenterClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewControllersPublicV1ProxyGetVcenterClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV1ProxyGetVcenterClusterOK creates a ControllersPublicV1ProxyGetVcenterClusterOK with default headers values
func NewControllersPublicV1ProxyGetVcenterClusterOK() *ControllersPublicV1ProxyGetVcenterClusterOK {
	return &ControllersPublicV1ProxyGetVcenterClusterOK{}
}

/*ControllersPublicV1ProxyGetVcenterClusterOK handles this case with default header values.

Success.
*/
type ControllersPublicV1ProxyGetVcenterClusterOK struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.ClusterDetail
}

func (o *ControllersPublicV1ProxyGetVcenterClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyGetVcenterClusterOK  %+v", 200, o.Payload)
}

func (o *ControllersPublicV1ProxyGetVcenterClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.ClusterDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyGetVcenterClusterUnauthorized creates a ControllersPublicV1ProxyGetVcenterClusterUnauthorized with default headers values
func NewControllersPublicV1ProxyGetVcenterClusterUnauthorized() *ControllersPublicV1ProxyGetVcenterClusterUnauthorized {
	return &ControllersPublicV1ProxyGetVcenterClusterUnauthorized{}
}

/*ControllersPublicV1ProxyGetVcenterClusterUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersPublicV1ProxyGetVcenterClusterUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyGetVcenterClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyGetVcenterClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersPublicV1ProxyGetVcenterClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyGetVcenterClusterForbidden creates a ControllersPublicV1ProxyGetVcenterClusterForbidden with default headers values
func NewControllersPublicV1ProxyGetVcenterClusterForbidden() *ControllersPublicV1ProxyGetVcenterClusterForbidden {
	return &ControllersPublicV1ProxyGetVcenterClusterForbidden{}
}

/*ControllersPublicV1ProxyGetVcenterClusterForbidden handles this case with default header values.

Forbidden. Access to the specified resource is not authorized. Check the IAM access policy for the `VMware Solutions` service.
*/
type ControllersPublicV1ProxyGetVcenterClusterForbidden struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyGetVcenterClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyGetVcenterClusterForbidden  %+v", 403, o.Payload)
}

func (o *ControllersPublicV1ProxyGetVcenterClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyGetVcenterClusterNotFound creates a ControllersPublicV1ProxyGetVcenterClusterNotFound with default headers values
func NewControllersPublicV1ProxyGetVcenterClusterNotFound() *ControllersPublicV1ProxyGetVcenterClusterNotFound {
	return &ControllersPublicV1ProxyGetVcenterClusterNotFound{}
}

/*ControllersPublicV1ProxyGetVcenterClusterNotFound handles this case with default header values.

Not found. The resource cannot be found.
*/
type ControllersPublicV1ProxyGetVcenterClusterNotFound struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyGetVcenterClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyGetVcenterClusterNotFound  %+v", 404, o.Payload)
}

func (o *ControllersPublicV1ProxyGetVcenterClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyGetVcenterClusterInternalServerError creates a ControllersPublicV1ProxyGetVcenterClusterInternalServerError with default headers values
func NewControllersPublicV1ProxyGetVcenterClusterInternalServerError() *ControllersPublicV1ProxyGetVcenterClusterInternalServerError {
	return &ControllersPublicV1ProxyGetVcenterClusterInternalServerError{}
}

/*ControllersPublicV1ProxyGetVcenterClusterInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersPublicV1ProxyGetVcenterClusterInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyGetVcenterClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vcenters/{instance_id}/clusters/{cluster_id}][%d] controllersPublicV1ProxyGetVcenterClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersPublicV1ProxyGetVcenterClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
