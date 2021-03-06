// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LeadControllerEntry lead controller entry
//
// swagger:model LeadControllerEntry
type LeadControllerEntry struct {

	// lead controller Id
	// Read Only: true
	LeadControllerID string `json:"leadControllerId,omitempty"`

	// table names
	// Read Only: true
	TableNames []string `json:"tableNames"`
}

// Validate validates this lead controller entry
func (m *LeadControllerEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeadControllerEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeadControllerEntry) UnmarshalBinary(b []byte) error {
	var res LeadControllerEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
