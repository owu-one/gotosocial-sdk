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

// TokenInfo TokenInfo represents metadata about one user-level access token.
//
// The actual access token itself will never be sent via the API.
//
// swagger:model TokenInfo
type TokenInfo struct {

	// When the token was created (ISO 8601 Datetime).
	// Example: 2021-07-30T09:20:25+00:00
	CreatedAt string `json:"created_at,omitempty"`

	// Database ID of this token.
	// Example: 01JMW7QBAZYZ8T8H73PCEX12XG
	ID string `json:"id,omitempty"`

	// Approximate time (accurate to within an hour) when the token was last used (ISO 8601 Datetime).
	// Omitted if token has never been used, or it is not known when it was last used (eg., it was last used before tracking "last_used" became a thing).
	// Example: 2021-07-30T09:20:25+00:00
	LastUsed string `json:"last_used,omitempty"`

	// OAuth scopes granted by the token, space-separated.
	// Example: read write admin
	Scope string `json:"scope,omitempty"`

	// application
	Application *Application `json:"application,omitempty"`
}

// Validate validates this token info
func (m *TokenInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenInfo) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this token info based on the context it is used
func (m *TokenInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenInfo) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {

		if swag.IsZero(m.Application) { // not required
			return nil
		}

		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenInfo) UnmarshalBinary(b []byte) error {
	var res TokenInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
