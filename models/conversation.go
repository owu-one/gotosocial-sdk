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

// Conversation Conversation represents a conversation
// with "direct message" visibility.
//
// swagger:model Conversation
type Conversation struct {

	// Participants in the conversation.
	Accounts []*Account `json:"accounts"`

	// Local database ID of the conversation.
	ID string `json:"id,omitempty"`

	// Is the conversation currently marked as unread?
	Unread bool `json:"unread,omitempty"`

	// last status
	LastStatus *Status `json:"last_status,omitempty"`
}

// Validate validates this conversation
func (m *Conversation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Conversation) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Accounts); i++ {
		if swag.IsZero(m.Accounts[i]) { // not required
			continue
		}

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) validateLastStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LastStatus) { // not required
		return nil
	}

	if m.LastStatus != nil {
		if err := m.LastStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conversation based on the context it is used
func (m *Conversation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Conversation) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Accounts); i++ {

		if m.Accounts[i] != nil {

			if swag.IsZero(m.Accounts[i]) { // not required
				return nil
			}

			if err := m.Accounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) contextValidateLastStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.LastStatus != nil {

		if swag.IsZero(m.LastStatus) { // not required
			return nil
		}

		if err := m.LastStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Conversation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Conversation) UnmarshalBinary(b []byte) error {
	var res Conversation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}