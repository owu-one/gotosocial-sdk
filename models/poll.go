// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Poll Poll represents a poll attached to a status.
//
// swagger:model Poll
type Poll struct {

	// Custom emoji to be used for rendering poll options.
	Emojis []*Emoji `json:"emojis"`

	// Is the poll currently expired?
	Expired bool `json:"expired,omitempty"`

	// When the poll ends. (ISO 8601 Datetime).
	ExpiresAt string `json:"expires_at,omitempty"`

	// The ID of the poll in the database.
	// Example: 01FBYKMD1KBMJ0W6JF1YZ3VY5D
	ID string `json:"id,omitempty"`

	// Does the poll allow multiple-choice answers?
	Multiple bool `json:"multiple,omitempty"`

	// Possible answers for the poll.
	Options []*PollOption `json:"options"`

	// When called with a user token, which options has the authorized
	// user chosen? Contains an array of index values for options.
	//
	// Omitted when no user token provided.
	OwnVotes []int64 `json:"own_votes"`

	// When called with a user token, has the authorized user voted?
	//
	// Omitted when no user token provided.
	Voted bool `json:"voted,omitempty"`

	// How many unique accounts have voted on a multiple-choice poll.
	VotersCount int64 `json:"voters_count,omitempty"`

	// How many votes have been received.
	VotesCount int64 `json:"votes_count,omitempty"`
}

// Validate validates this poll
func (m *Poll) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmojis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Poll) validateEmojis(formats strfmt.Registry) error {
	if swag.IsZero(m.Emojis) { // not required
		return nil
	}

	for i := 0; i < len(m.Emojis); i++ {
		if swag.IsZero(m.Emojis[i]) { // not required
			continue
		}

		if m.Emojis[i] != nil {
			if err := m.Emojis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emojis" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emojis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Poll) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {
		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {
			if err := m.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this poll based on the context it is used
func (m *Poll) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmojis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Poll) contextValidateEmojis(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Emojis); i++ {

		if m.Emojis[i] != nil {

			if swag.IsZero(m.Emojis[i]) { // not required
				return nil
			}

			if err := m.Emojis[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emojis" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emojis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Poll) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {

			if swag.IsZero(m.Options[i]) { // not required
				return nil
			}

			if err := m.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Poll) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Poll) UnmarshalBinary(b []byte) error {
	var res Poll
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}