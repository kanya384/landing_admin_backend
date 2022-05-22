// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DayLeadsInfo day leads info
//
// swagger:model DayLeadsInfo
type DayLeadsInfo struct {

	// count
	Count float64 `json:"count,omitempty"`

	// date
	Date string `json:"date,omitempty"`
}

// Validate validates this day leads info
func (m *DayLeadsInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this day leads info based on context it is used
func (m *DayLeadsInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DayLeadsInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DayLeadsInfo) UnmarshalBinary(b []byte) error {
	var res DayLeadsInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}