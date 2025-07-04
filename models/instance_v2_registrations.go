// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceV2Registrations Information about registering for this instance.
//
// swagger:model InstanceV2Registrations
type InstanceV2Registrations struct {

	// Whether registrations require moderator approval.
	// Example: true
	ApprovalRequired bool `json:"approval_required,omitempty"`

	// Whether registrations are enabled.
	// Example: false
	Enabled bool `json:"enabled,omitempty"`

	// A custom message (html string) to be shown when registrations are closed.
	// Value will be null if no message is set.
	// Example: \u003cp\u003eRegistrations are currently closed on example.org because of spam bots!\u003c/p\u003e
	Message string `json:"message,omitempty"`

	// A minimum age required to register, if configured.
	// Currently not implemented: will always be null.
	MinAge int64 `json:"min_age,omitempty"`

	// Whether registrations require the user to provide a reason for joining.
	// Only applicable when ApprovalRequired is true.
	// Example: true
	ReasonRequired bool `json:"reason_required,omitempty"`
}

// Validate validates this instance v2 registrations
func (m *InstanceV2Registrations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance v2 registrations based on context it is used
func (m *InstanceV2Registrations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceV2Registrations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceV2Registrations) UnmarshalBinary(b []byte) error {
	var res InstanceV2Registrations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
