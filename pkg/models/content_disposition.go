package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentDisposition content disposition
type ContentDisposition struct {
	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`
	// file name
	FileName string `json:"fileName,omitempty"`
	// modification date
	// Format: date-time
	ModificationDate strfmt.DateTime `json:"modificationDate,omitempty"`
	// parameters
	Parameters map[string]string `json:"parameters,omitempty"`
	// read date
	// Format: date-time
	ReadDate strfmt.DateTime `json:"readDate,omitempty"`
	// size
	Size int64 `json:"size,omitempty"`
	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this content disposition
func (m *ContentDisposition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModificationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentDisposition) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentDisposition) validateModificationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModificationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modificationDate", "body", "date-time", m.ModificationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentDisposition) validateReadDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadDate) { // not required
		return nil
	}

	if err := validate.FormatOf("readDate", "body", "date-time", m.ReadDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentDisposition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentDisposition) UnmarshalBinary(b []byte) error {
	var res ContentDisposition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}