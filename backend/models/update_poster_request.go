// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdatePosterRequest update poster request
//
// swagger:model UpdatePosterRequest
type UpdatePosterRequest struct {

	// active
	Active bool `json:"active,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// photo
	Photo string `json:"photo,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this update poster request
func (m *UpdatePosterRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update poster request based on context it is used
func (m *UpdatePosterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePosterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePosterRequest) UnmarshalBinary(b []byte) error {
	var res UpdatePosterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
