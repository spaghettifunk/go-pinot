package models

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MultiPart multi part
type MultiPart struct {

	// body parts
	BodyParts []*BodyPart `json:"bodyParts"`

	// content disposition
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`

	// entity
	Entity interface{} `json:"entity,omitempty"`

	// headers
	Headers map[string][]string `json:"headers,omitempty"`

	// media type
	MediaType *MediaType `json:"mediaType,omitempty"`

	// message body workers
	MessageBodyWorkers MessageBodyWorkers `json:"messageBodyWorkers,omitempty"`

	// parameterized headers
	ParameterizedHeaders map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`

	// parent
	Parent *MultiPart `json:"parent,omitempty"`

	// providers
	Providers Providers `json:"providers,omitempty"`
}

// Validate validates this multi part
func (m *MultiPart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBodyParts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterizedHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiPart) validateBodyParts(formats strfmt.Registry) error {

	if swag.IsZero(m.BodyParts) { // not required
		return nil
	}

	for i := 0; i < len(m.BodyParts); i++ {
		if swag.IsZero(m.BodyParts[i]) { // not required
			continue
		}

		if m.BodyParts[i] != nil {
			if err := m.BodyParts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bodyParts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiPart) validateContentDisposition(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentDisposition) { // not required
		return nil
	}

	if m.ContentDisposition != nil {
		if err := m.ContentDisposition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentDisposition")
			}
			return err
		}
	}

	return nil
}

func (m *MultiPart) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	if m.MediaType != nil {
		if err := m.MediaType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaType")
			}
			return err
		}
	}

	return nil
}

func (m *MultiPart) validateParameterizedHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.ParameterizedHeaders) { // not required
		return nil
	}

	for k := range m.ParameterizedHeaders {

		if err := validate.Required("parameterizedHeaders"+"."+k, "body", m.ParameterizedHeaders[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.ParameterizedHeaders[k]); i++ {

			if err := m.ParameterizedHeaders[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterizedHeaders" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *MultiPart) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiPart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiPart) UnmarshalBinary(b []byte) error {
	var res MultiPart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}