// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceV2ThumbnailVersions Links to scaled resolution images, for high DPI screens.
//
// swagger:model InstanceV2ThumbnailVersions
type InstanceV2ThumbnailVersions struct {

	// The URL for the thumbnail image at 1x resolution.
	// Key/value not set if scaled versions not available.
	Size1URL string `json:"@1x,omitempty"`

	// The URL for the thumbnail image at 2x resolution.
	// Key/value not set if scaled versions not available.
	Size2URL string `json:"@2x,omitempty"`
}

// Validate validates this instance v2 thumbnail versions
func (m *InstanceV2ThumbnailVersions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance v2 thumbnail versions based on context it is used
func (m *InstanceV2ThumbnailVersions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceV2ThumbnailVersions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceV2ThumbnailVersions) UnmarshalBinary(b []byte) error {
	var res InstanceV2ThumbnailVersions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
