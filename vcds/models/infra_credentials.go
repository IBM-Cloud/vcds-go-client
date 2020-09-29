// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InfraCredentials infra credentials
// swagger:model infra_credentials
type InfraCredentials struct {

	// infra credentials
	// Required: true
	InfraCredentials *CommonInfraCredentials `json:"infra_credentials"`
}

// Validate validates this infra credentials
func (m *InfraCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfraCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfraCredentials) validateInfraCredentials(formats strfmt.Registry) error {

	if err := validate.Required("infra_credentials", "body", m.InfraCredentials); err != nil {
		return err
	}

	if m.InfraCredentials != nil {
		if err := m.InfraCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infra_credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfraCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraCredentials) UnmarshalBinary(b []byte) error {
	var res InfraCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
