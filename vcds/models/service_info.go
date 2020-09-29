// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceInfo service info
// swagger:model service_info
type ServiceInfo struct {

	// service info
	ServiceInfo *ServiceOrder `json:"service_info,omitempty"`
}

// Validate validates this service info
func (m *ServiceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceInfo) validateServiceInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceInfo) { // not required
		return nil
	}

	if m.ServiceInfo != nil {
		if err := m.ServiceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInfo) UnmarshalBinary(b []byte) error {
	var res ServiceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
