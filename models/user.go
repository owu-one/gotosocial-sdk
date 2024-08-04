// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// User User models fields relevant to one user.
//
// swagger:model User
type User struct {

	// User is an admin.
	// Example: false
	Admin bool `json:"admin,omitempty"`

	// User was approved by an admin.
	// Example: true
	Approved bool `json:"approved,omitempty"`

	// Time when the last "please confirm your email address" email was sent, if at all. (ISO 8601 Datetime)
	// Example: 2021-07-30T09:20:25+00:00
	ConfirmationSentAt string `json:"confirmation_sent_at,omitempty"`

	// Time at which the email given in the `email` field was confirmed, if at all. (ISO 8601 Datetime)
	// Example: 2021-07-30T09:20:25+00:00
	ConfirmedAt string `json:"confirmed_at,omitempty"`

	// Time this user was created. (ISO 8601 Datetime)
	// Example: 2021-07-30T09:20:25+00:00
	CreatedAt string `json:"created_at,omitempty"`

	// User's account is disabled.
	// Example: false
	Disabled bool `json:"disabled,omitempty"`

	// Confirmed email address of this user, if set.
	// Example: someone@example.org
	Email string `json:"email,omitempty"`

	// Database ID of this user.
	// Example: 01FBVD42CQ3ZEEVMW180SBX03B
	ID string `json:"id,omitempty"`

	// Time at which this user was last emailed, if at all. (ISO 8601 Datetime)
	// Example: 2021-07-30T09:20:25+00:00
	LastEmailedAt string `json:"last_emailed_at,omitempty"`

	// User is a moderator.
	// Example: false
	Moderator bool `json:"moderator,omitempty"`

	// Reason for sign-up, if provided.
	// Example: Please! Pretty please!
	Reason string `json:"reason,omitempty"`

	// Time when the last "please reset your password" email was sent, if at all. (ISO 8601 Datetime)
	// Example: 2021-07-30T09:20:25+00:00
	ResetPasswordSentAt string `json:"reset_password_sent_at,omitempty"`

	// Unconfirmed email address of this user, if set.
	// Example: someone.else@somewhere.else.example.org
	UnconfirmedEmail string `json:"unconfirmed_email,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
