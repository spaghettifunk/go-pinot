package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MediaType media type
//
// swagger:model MediaType
type MediaType struct {
	// parameters
	Parameters map[string]string `json:"parameters,omitempty"`
	// subtype
	Subtype string `json:"subtype,omitempty"`
	// type
	Type string `json:"type,omitempty"`
	// wildcard subtype
	WildcardSubtype bool `json:"wildcardSubtype,omitempty"`
	// wildcard type
	WildcardType bool `json:"wildcardType,omitempty"`
}

// Validate validates this media type
func (m *MediaType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MediaType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaType) UnmarshalBinary(b []byte) error {
	var res MediaType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
