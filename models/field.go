// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Field Field represents a name/value pair to display on an account's profile.
//
// swagger:model Field
type Field struct {

	// The key/name of this field.
	// Example: pronouns
	Name string `json:"name,omitempty"`

	// The value of this field.
	// Example: they/them
	Value string `json:"value,omitempty"`

	// If this field has been verified, when did this occur? (ISO 8601 Datetime).
	// Example: 2021-07-30T09:20:25+00:00
	VerifiedAt string `json:"verified_at,omitempty"`
}

// Validate validates this field
func (m *Field) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this field based on context it is used
func (m *Field) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Field) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Field) UnmarshalBinary(b []byte) error {
	var res Field
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
