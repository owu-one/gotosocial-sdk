// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountExportStats AccountExportStats models an account's stats
// specifically for the purpose of informing about
// export sizes at the /api/v1/exports/stats endpoint.
//
// swagger:model AccountExportStats
type AccountExportStats struct {

	// Number of accounts blocked by this account.
	// Example: 15
	BlocksCount int64 `json:"blocks_count,omitempty"`

	// Number of accounts following this account.
	// Example: 50
	FollowersCount int64 `json:"followers_count,omitempty"`

	// Number of accounts followed by this account.
	// Example: 50
	FollowingCount int64 `json:"following_count,omitempty"`

	// Number of lists created by this account.
	// Example: 10
	ListsCount int64 `json:"lists_count,omitempty"`

	// TODO: String representation of media storage size attributed to this account.
	// Example: 500MB
	MediaStorage string `json:"media_storage,omitempty"`

	// Number of accounts muted by this account.
	// Example: 11
	MutesCount int64 `json:"mutes_count,omitempty"`

	// Number of statuses created by this account.
	// Example: 81986
	StatusesCount int64 `json:"statuses_count,omitempty"`
}

// Validate validates this account export stats
func (m *AccountExportStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account export stats based on context it is used
func (m *AccountExportStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountExportStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountExportStats) UnmarshalBinary(b []byte) error {
	var res AccountExportStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
