// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkConfig Network configuration.
// swagger:model network_config
type NetworkConfig struct {

	// Network Uplink Speed. The default value is 10 GB. 25 GB is available only for certain server types and data centers.
	// Enum: [10GB 25GB]
	UplinkSpeed *string `json:"uplink_speed,omitempty"`
}

// Validate validates this network config
func (m *NetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUplinkSpeed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkConfigTypeUplinkSpeedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["10GB","25GB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkConfigTypeUplinkSpeedPropEnum = append(networkConfigTypeUplinkSpeedPropEnum, v)
	}
}

const (

	// NetworkConfigUplinkSpeedNr10GB captures enum value "10GB"
	NetworkConfigUplinkSpeedNr10GB string = "10GB"

	// NetworkConfigUplinkSpeedNr25GB captures enum value "25GB"
	NetworkConfigUplinkSpeedNr25GB string = "25GB"
)

// prop value enum
func (m *NetworkConfig) validateUplinkSpeedEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkConfigTypeUplinkSpeedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkConfig) validateUplinkSpeed(formats strfmt.Registry) error {

	if swag.IsZero(m.UplinkSpeed) { // not required
		return nil
	}

	// value enum
	if err := m.validateUplinkSpeedEnum("uplink_speed", "body", *m.UplinkSpeed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkConfig) UnmarshalBinary(b []byte) error {
	var res NetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
