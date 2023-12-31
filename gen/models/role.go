// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Role role
//
// swagger:model Role
type Role struct {
	ModelTrackTime

	ModelIdentifier

	// name
	// Required: true
	Name *string `json:"name"`

	// slug
	Slug string `json:"slug,omitempty" gorm:"type:varchar(255);uniqueIndex"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Role) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModelTrackTime
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModelTrackTime = aO0

	// AO1
	var aO1 ModelIdentifier
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ModelIdentifier = aO1

	// AO2
	var dataAO2 struct {
		Name *string `json:"name"`

		Slug string `json:"slug,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Name = dataAO2.Name

	m.Slug = dataAO2.Slug

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Role) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.ModelTrackTime)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ModelIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		Name *string `json:"name"`

		Slug string `json:"slug,omitempty"`
	}

	dataAO2.Name = m.Name

	dataAO2.Slug = m.Slug

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this role based on the context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
