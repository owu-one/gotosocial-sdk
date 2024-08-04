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

// ThreadContext ThreadContext models the tree or
// "thread" around a given status.
//
// swagger:model ThreadContext
type ThreadContext struct {

	// Parents in the thread.
	Ancestors []*Status `json:"ancestors"`

	// Children in the thread.
	Descendants []*Status `json:"descendants"`
}

// Validate validates this thread context
func (m *ThreadContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAncestors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescendants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThreadContext) validateAncestors(formats strfmt.Registry) error {
	if swag.IsZero(m.Ancestors) { // not required
		return nil
	}

	for i := 0; i < len(m.Ancestors); i++ {
		if swag.IsZero(m.Ancestors[i]) { // not required
			continue
		}

		if m.Ancestors[i] != nil {
			if err := m.Ancestors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ancestors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ancestors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThreadContext) validateDescendants(formats strfmt.Registry) error {
	if swag.IsZero(m.Descendants) { // not required
		return nil
	}

	for i := 0; i < len(m.Descendants); i++ {
		if swag.IsZero(m.Descendants[i]) { // not required
			continue
		}

		if m.Descendants[i] != nil {
			if err := m.Descendants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("descendants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("descendants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this thread context based on the context it is used
func (m *ThreadContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAncestors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescendants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThreadContext) contextValidateAncestors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ancestors); i++ {

		if m.Ancestors[i] != nil {

			if swag.IsZero(m.Ancestors[i]) { // not required
				return nil
			}

			if err := m.Ancestors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ancestors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ancestors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThreadContext) contextValidateDescendants(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Descendants); i++ {

		if m.Descendants[i] != nil {

			if swag.IsZero(m.Descendants[i]) { // not required
				return nil
			}

			if err := m.Descendants[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("descendants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("descendants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThreadContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThreadContext) UnmarshalBinary(b []byte) error {
	var res ThreadContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}