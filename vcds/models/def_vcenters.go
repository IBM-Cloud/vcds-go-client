// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DefVcenters The list of vCenters information in JSON format.
// swagger:model def_vcenters
type DefVcenters struct {

	// The detailed information of the vCenter Server instance.
	// Required: true
	Vcenters []*Vcenter `json:"vcenters"`
}

// Validate validates this def vcenters
func (m *DefVcenters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcenters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefVcenters) validateVcenters(formats strfmt.Registry) error {

	if err := validate.Required("vcenters", "body", m.Vcenters); err != nil {
		return err
	}

	for i := 0; i < len(m.Vcenters); i++ {
		if swag.IsZero(m.Vcenters[i]) { // not required
			continue
		}

		if m.Vcenters[i] != nil {
			if err := m.Vcenters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcenters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DefVcenters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefVcenters) UnmarshalBinary(b []byte) error {
	var res DefVcenters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
