// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// ControllersPublicV1ProxyMutateServiceVariablesReader is a Reader for the ControllersPublicV1ProxyMutateServiceVariables structure.
type ControllersPublicV1ProxyMutateServiceVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControllersPublicV1ProxyMutateServiceVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewControllersPublicV1ProxyMutateServiceVariablesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewControllersPublicV1ProxyMutateServiceVariablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewControllersPublicV1ProxyMutateServiceVariablesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewControllersPublicV1ProxyMutateServiceVariablesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewControllersPublicV1ProxyMutateServiceVariablesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewControllersPublicV1ProxyMutateServiceVariablesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewControllersPublicV1ProxyMutateServiceVariablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewControllersPublicV1ProxyMutateServiceVariablesAccepted creates a ControllersPublicV1ProxyMutateServiceVariablesAccepted with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesAccepted() *ControllersPublicV1ProxyMutateServiceVariablesAccepted {
	return &ControllersPublicV1ProxyMutateServiceVariablesAccepted{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesAccepted handles this case with default header values.

Success. The request for updating or deleting the specified service variables has been accepted.
*/
type ControllersPublicV1ProxyMutateServiceVariablesAccepted struct {
	Payload *ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesAccepted  %+v", 202, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyMutateServiceVariablesBadRequest creates a ControllersPublicV1ProxyMutateServiceVariablesBadRequest with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesBadRequest() *ControllersPublicV1ProxyMutateServiceVariablesBadRequest {
	return &ControllersPublicV1ProxyMutateServiceVariablesBadRequest{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesBadRequest handles this case with default header values.

Bad request. Check your request parameters.
*/
type ControllersPublicV1ProxyMutateServiceVariablesBadRequest struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesBadRequest  %+v", 400, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyMutateServiceVariablesUnauthorized creates a ControllersPublicV1ProxyMutateServiceVariablesUnauthorized with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesUnauthorized() *ControllersPublicV1ProxyMutateServiceVariablesUnauthorized {
	return &ControllersPublicV1ProxyMutateServiceVariablesUnauthorized{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesUnauthorized handles this case with default header values.

Unauthorized. The IAM token is invalid or expired. To retrieve your IAM token, run `ibmcloud login` and then `ibmcloud iam oauth-tokens`.
*/
type ControllersPublicV1ProxyMutateServiceVariablesUnauthorized struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesUnauthorized  %+v", 401, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyMutateServiceVariablesForbidden creates a ControllersPublicV1ProxyMutateServiceVariablesForbidden with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesForbidden() *ControllersPublicV1ProxyMutateServiceVariablesForbidden {
	return &ControllersPublicV1ProxyMutateServiceVariablesForbidden{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesForbidden handles this case with default header values.

Forbidden. Access to the specified resource is not authorized. Check the IAM access policy for the `VMware Solutions` service.
*/
type ControllersPublicV1ProxyMutateServiceVariablesForbidden struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesForbidden  %+v", 403, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyMutateServiceVariablesNotFound creates a ControllersPublicV1ProxyMutateServiceVariablesNotFound with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesNotFound() *ControllersPublicV1ProxyMutateServiceVariablesNotFound {
	return &ControllersPublicV1ProxyMutateServiceVariablesNotFound{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesNotFound handles this case with default header values.

Not found. The resource cannot be found.
*/
type ControllersPublicV1ProxyMutateServiceVariablesNotFound struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesNotFound  %+v", 404, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyMutateServiceVariablesConflict creates a ControllersPublicV1ProxyMutateServiceVariablesConflict with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesConflict() *ControllersPublicV1ProxyMutateServiceVariablesConflict {
	return &ControllersPublicV1ProxyMutateServiceVariablesConflict{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesConflict handles this case with default header values.

Conflict. The request cannot be completed because of a conflict with the current state of the target resource.
*/
type ControllersPublicV1ProxyMutateServiceVariablesConflict struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesConflict) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesConflict  %+v", 409, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControllersPublicV1ProxyMutateServiceVariablesInternalServerError creates a ControllersPublicV1ProxyMutateServiceVariablesInternalServerError with default headers values
func NewControllersPublicV1ProxyMutateServiceVariablesInternalServerError() *ControllersPublicV1ProxyMutateServiceVariablesInternalServerError {
	return &ControllersPublicV1ProxyMutateServiceVariablesInternalServerError{}
}

/*ControllersPublicV1ProxyMutateServiceVariablesInternalServerError handles this case with default header values.

Internal server error. Your request cannot be processed. Please wait a few minutes and try again.
*/
type ControllersPublicV1ProxyMutateServiceVariablesInternalServerError struct {
	/*Global transaction ID for request correlation.
	 */
	XGlobalTransactionID string

	Payload *models.Error
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_instances/{service_instance_id}/service_variables][%d] controllersPublicV1ProxyMutateServiceVariablesInternalServerError  %+v", 500, o.Payload)
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-global-transaction-id
	o.XGlobalTransactionID = response.GetHeader("x-global-transaction-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody controllers public v1 proxy mutate service variables accepted body
swagger:model ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody
*/
type ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody struct {

	// The name list of the variables that were updated or removed.
	Variables []*VariablesItems0 `json:"variables"`
}

// Validate validates this controllers public v1 proxy mutate service variables accepted body
func (o *ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(o.Variables) { // not required
		return nil
	}

	for i := 0; i < len(o.Variables); i++ {
		if swag.IsZero(o.Variables[i]) { // not required
			continue
		}

		if o.Variables[i] != nil {
			if err := o.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllersPublicV1ProxyMutateServiceVariablesAccepted" + "." + "variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody) UnmarshalBinary(b []byte) error {
	var res ControllersPublicV1ProxyMutateServiceVariablesAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ControllersPublicV1ProxyMutateServiceVariablesBody controllers public v1 proxy mutate service variables body
swagger:model ControllersPublicV1ProxyMutateServiceVariablesBody
*/
type ControllersPublicV1ProxyMutateServiceVariablesBody struct {

	// The action to be performed against the specified variables of a service instance. The only supported values are 'update' and 'delete'.
	// Required: true
	// Enum: [update delete]
	Action *string `json:"action"`

	// The list of variables requested to be updated or deleted for a service instance.
	// Required: true
	Variables []*VariablesItems0 `json:"variables"`
}

// Validate validates this controllers public v1 proxy mutate service variables body
func (o *ControllersPublicV1ProxyMutateServiceVariablesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var controllersPublicV1ProxyMutateServiceVariablesBodyTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["update","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		controllersPublicV1ProxyMutateServiceVariablesBodyTypeActionPropEnum = append(controllersPublicV1ProxyMutateServiceVariablesBodyTypeActionPropEnum, v)
	}
}

const (

	// ControllersPublicV1ProxyMutateServiceVariablesBodyActionUpdate captures enum value "update"
	ControllersPublicV1ProxyMutateServiceVariablesBodyActionUpdate string = "update"

	// ControllersPublicV1ProxyMutateServiceVariablesBodyActionDelete captures enum value "delete"
	ControllersPublicV1ProxyMutateServiceVariablesBodyActionDelete string = "delete"
)

// prop value enum
func (o *ControllersPublicV1ProxyMutateServiceVariablesBody) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, controllersPublicV1ProxyMutateServiceVariablesBodyTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesBody) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("service_variable_patch_data"+"."+"action", "body", o.Action); err != nil {
		return err
	}

	// value enum
	if err := o.validateActionEnum("service_variable_patch_data"+"."+"action", "body", *o.Action); err != nil {
		return err
	}

	return nil
}

func (o *ControllersPublicV1ProxyMutateServiceVariablesBody) validateVariables(formats strfmt.Registry) error {

	if err := validate.Required("service_variable_patch_data"+"."+"variables", "body", o.Variables); err != nil {
		return err
	}

	for i := 0; i < len(o.Variables); i++ {
		if swag.IsZero(o.Variables[i]) { // not required
			continue
		}

		if o.Variables[i] != nil {
			if err := o.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_variable_patch_data" + "." + "variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ControllersPublicV1ProxyMutateServiceVariablesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ControllersPublicV1ProxyMutateServiceVariablesBody) UnmarshalBinary(b []byte) error {
	var res ControllersPublicV1ProxyMutateServiceVariablesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VariablesItems0 variables items0
swagger:model VariablesItems0
*/
type VariablesItems0 struct {

	// The name of the variable that was updated or removed.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this variables items0
func (o *VariablesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VariablesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VariablesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VariablesItems0) UnmarshalBinary(b []byte) error {
	var res VariablesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}