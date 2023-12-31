// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// City city
//
// swagger:model City
type City struct {
	ModelTrackTime

	ModelIdentifier

	// city code
	CityCode string `json:"city_code,omitempty" gorm:"type:varchar(5)"`

	// city name
	CityName string `json:"city_name,omitempty" gorm:"type:varchar(50);not null"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *City) UnmarshalJSON(raw []byte) error {
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
		CityCode string `json:"city_code,omitempty"`

		CityName string `json:"city_name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.CityCode = dataAO2.CityCode

	m.CityName = dataAO2.CityName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m City) MarshalJSON() ([]byte, error) {
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
		CityCode string `json:"city_code,omitempty"`

		CityName string `json:"city_name,omitempty"`
	}

	dataAO2.CityCode = m.CityCode

	dataAO2.CityName = m.CityName

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this city
func (m *City) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelTrackTime
	if err := m.ModelTrackTime.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ModelIdentifier
	if err := m.ModelIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this city based on the context it is used
func (m *City) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *City) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *City) UnmarshalBinary(b []byte) error {
	var res City
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
