// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserCreateRequest user create request
//
// swagger:model UserCreateRequest
type UserCreateRequest struct {

	// login
	// Example: login
	Login string `json:"login,omitempty"`

	// name
	// Example: пользователь
	Name string `json:"name,omitempty"`

	// pass
	// Example: password
	Pass string `json:"pass,omitempty"`

	// role
	// Example: роль
	Role string `json:"role,omitempty"`
}

// Validate validates this user create request
func (m *UserCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user create request based on context it is used
func (m *UserCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCreateRequest) UnmarshalBinary(b []byte) error {
	var res UserCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
