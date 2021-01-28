package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParameterizedHeader parameterized header
type ParameterizedHeader struct {
	// parameters
	Parameters map[string]string `json:"parameters,omitempty"`
	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this parameterized header
func (m *ParameterizedHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterizedHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterizedHeader) UnmarshalBinary(b []byte) error {
	var res ParameterizedHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
