// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Schema schema
//
// swagger:model Schema
type Schema struct {

	// date time field specs
	DateTimeFieldSpecs []*DateTimeFieldSpec `json:"dateTimeFieldSpecs"`

	// dimension field specs
	DimensionFieldSpecs []*DimensionFieldSpec `json:"dimensionFieldSpecs"`

	// metric field specs
	MetricFieldSpecs []*MetricFieldSpec `json:"metricFieldSpecs"`

	// primary key columns
	PrimaryKeyColumns []string `json:"primaryKeyColumns"`

	// schema name
	SchemaName string `json:"schemaName,omitempty"`

	// time field spec
	TimeFieldSpec *TimeFieldSpec `json:"timeFieldSpec,omitempty"`
}

// Validate validates this schema
func (m *Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTimeFieldSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensionFieldSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricFieldSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeFieldSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema) validateDateTimeFieldSpecs(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeFieldSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.DateTimeFieldSpecs); i++ {
		if swag.IsZero(m.DateTimeFieldSpecs[i]) { // not required
			continue
		}

		if m.DateTimeFieldSpecs[i] != nil {
			if err := m.DateTimeFieldSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dateTimeFieldSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Schema) validateDimensionFieldSpecs(formats strfmt.Registry) error {

	if swag.IsZero(m.DimensionFieldSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.DimensionFieldSpecs); i++ {
		if swag.IsZero(m.DimensionFieldSpecs[i]) { // not required
			continue
		}

		if m.DimensionFieldSpecs[i] != nil {
			if err := m.DimensionFieldSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dimensionFieldSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Schema) validateMetricFieldSpecs(formats strfmt.Registry) error {

	if swag.IsZero(m.MetricFieldSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricFieldSpecs); i++ {
		if swag.IsZero(m.MetricFieldSpecs[i]) { // not required
			continue
		}

		if m.MetricFieldSpecs[i] != nil {
			if err := m.MetricFieldSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metricFieldSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Schema) validateTimeFieldSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeFieldSpec) { // not required
		return nil
	}

	if m.TimeFieldSpec != nil {
		if err := m.TimeFieldSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeFieldSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schema) UnmarshalBinary(b []byte) error {
	var res Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}