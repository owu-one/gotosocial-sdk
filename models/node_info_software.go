// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeInfoSoftware NodeInfoSoftware represents the name and version number of the software of this node.
//
// swagger:model NodeInfoSoftware
type NodeInfoSoftware struct {

	// name
	// Example: gotosocial
	Name string `json:"name,omitempty"`

	// version
	// Example: 0.1.2 1234567
	Version string `json:"version,omitempty"`
}

// Validate validates this node info software
func (m *NodeInfoSoftware) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node info software based on context it is used
func (m *NodeInfoSoftware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeInfoSoftware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInfoSoftware) UnmarshalBinary(b []byte) error {
	var res NodeInfoSoftware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}