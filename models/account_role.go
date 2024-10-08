// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountRole AccountRole models the role of an account.
//
// swagger:model AccountRole
type AccountRole struct {

	// Color is a 6-digit CSS-style hex color code with leading `#`, or an empty string if this role has no color.
	// Since GotoSocial doesn't use role colors, we leave this empty.
	Color string `json:"color,omitempty"`

	// Highlighted indicates whether the role is publicly visible on the user profile.
	// This is always true for GotoSocial's built-in admin and moderator roles, and false otherwise.
	Highlighted bool `json:"highlighted,omitempty"`

	// ID of the role.
	// Not used by GotoSocial, but we set it to the role name, just in case a client expects a unique ID.
	ID string `json:"id,omitempty"`

	// Name of the role.
	Name string `json:"name,omitempty"`

	// Permissions is a bitmap serialized as a numeric string, indicating which admin/moderation actions a user can perform.
	Permissions string `json:"permissions,omitempty"`
}

// Validate validates this account role
func (m *AccountRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account role based on context it is used
func (m *AccountRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRole) UnmarshalBinary(b []byte) error {
	var res AccountRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
