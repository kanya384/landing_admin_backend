// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Title title
//
// swagger:model Title
type Title struct {

	// desktop title
	DesktopTitle string `json:"desktopTitle,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// mobile title
	MobileTitle string `json:"mobileTitle,omitempty"`

	// tag name
	TagName string `json:"tagName,omitempty"`

	// tag value
	TagValue string `json:"tagValue,omitempty"`
}

// Validate validates this title
func (m *Title) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this title based on context it is used
func (m *Title) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Title) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Title) UnmarshalBinary(b []byte) error {
	var res Title
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
