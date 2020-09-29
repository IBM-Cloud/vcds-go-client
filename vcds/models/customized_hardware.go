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

// CustomizedHardware Customized hardware configuration to be ordered.
// swagger:model customized_hardware
type CustomizedHardware struct {

	// An array of disk groups. If you are not using vSAN, you must specify this setting.
	DiskGroups []*DiskGroup `json:"disk_groups"`

	// An array of disk types for each of the local disks used by vSAN.
	Disks []string `json:"disks"`

	// RAM type of the hardware. To see supported RAM types, use the `GET /v2/ram_types` API.
	// Required: true
	RAM *string `json:"ram"`

	// Server CPU type of the hardware. To see supported server types, use the `GET /v2/server_types` API.
	// Required: true
	Server *string `json:"server"`

	// An array of disk types for each of the local disks used by vSAN cache.
	VsanCacheDisks []string `json:"vsan_cache_disks"`

	// Whether or not to compress the vSAN disks. Applicable only when using vSAN. If true, the compression ratio on vSAN is set to 3.5.
	VsanCompression *bool `json:"vsan_compression,omitempty"`
}

// Validate validates this customized hardware
func (m *CustomizedHardware) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRAM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomizedHardware) validateDiskGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.DiskGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.DiskGroups); i++ {
		if swag.IsZero(m.DiskGroups[i]) { // not required
			continue
		}

		if m.DiskGroups[i] != nil {
			if err := m.DiskGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disk_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomizedHardware) validateRAM(formats strfmt.Registry) error {

	if err := validate.Required("ram", "body", m.RAM); err != nil {
		return err
	}

	return nil
}

func (m *CustomizedHardware) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomizedHardware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomizedHardware) UnmarshalBinary(b []byte) error {
	var res CustomizedHardware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}