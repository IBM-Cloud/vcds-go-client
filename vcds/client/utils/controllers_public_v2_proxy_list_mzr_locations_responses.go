// Code generated by go-swagger; DO NOT EDIT.

package utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersPublicV2ProxyListMzrLocationsReader is a Reader for the ControllersPublicV2ProxyListMzrLocations structure.
type ControllersPublicV2ProxyListMzrLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV2ProxyListMzrLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewControllersPublicV2ProxyListMzrLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewControllersPublicV2ProxyListMzrLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV2ProxyListMzrLocationsOK creates a ControllersPublicV2ProxyListMzrLocationsOK with default headers values
func NewControllersPublicV2ProxyListMzrLocationsOK() *ControllersPublicV2ProxyListMzrLocationsOK {
	return &ControllersPublicV2ProxyListMzrLocationsOK{}
}

/*ControllersPublicV2ProxyListMzrLocationsOK handles this case with default header values.

Success.
*/
type ControllersPublicV2ProxyListMzrLocationsOK struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.DefMzrLocations
}

func (o *ControllersPublicV2ProxyListMzrLocationsOK) Error() string {
	return fmt.Sprintf("[GET /v2/mzr_locations][%d] controllersPublicV2ProxyListMzrLocationsOK  %+v", 200, o.Payload)
}

func (o *ControllersPublicV2ProxyListMzrLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.DefMzrLocations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV2ProxyListMzrLocationsInternalServerError creates a ControllersPublicV2ProxyListMzrLocationsInternalServerError with default headers values
func NewControllersPublicV2ProxyListMzrLocationsInternalServerError() *ControllersPublicV2ProxyListMzrLocationsInternalServerError {
	return &ControllersPublicV2ProxyListMzrLocationsInternalServerError{}
}

/*ControllersPublicV2ProxyListMzrLocationsInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersPublicV2ProxyListMzrLocationsInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV2ProxyListMzrLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/mzr_locations][%d] controllersPublicV2ProxyListMzrLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersPublicV2ProxyListMzrLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}