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

// MutedAccount MutedAccount extends Account with a field used only by the muted user list.
//
// swagger:model MutedAccount
type MutedAccount struct {

	// The account URI as discovered via webfinger.
	// Equal to username for local users, or username@domain for remote users.
	// Example: some_user@example.org
	Acct string `json:"acct,omitempty"`

	// Web location of the account's avatar.
	// Example: https://example.org/media/some_user/avatar/original/avatar.jpeg
	Avatar string `json:"avatar,omitempty"`

	// Description of this account's avatar, for alt text.
	// Example: A cute drawing of a smiling sloth.
	AvatarDescription string `json:"avatar_description,omitempty"`

	// Database ID of the media attachment for this account's avatar image.
	// Omitted if no avatar uploaded for this account (ie., default avatar).
	// Example: 01JAJ3XCD66K3T99JZESCR137W
	AvatarMediaID string `json:"avatar_media_id,omitempty"`

	// Web location of a static version of the account's avatar.
	// Only relevant when the account's main avatar is a video or a gif.
	// Example: https://example.org/media/some_user/avatar/static/avatar.png
	AvatarStatic string `json:"avatar_static,omitempty"`

	// Account identifies as a bot.
	Bot bool `json:"bot,omitempty"`

	// When the account was created (ISO 8601 Datetime).
	// Example: 2021-07-30T09:20:25+00:00
	CreatedAt string `json:"created_at,omitempty"`

	// CustomCSS to include when rendering this account's profile or statuses.
	CustomCSS string `json:"custom_css,omitempty"`

	// Account has opted into discovery features.
	Discoverable bool `json:"discoverable,omitempty"`

	// The account's display name.
	// Example: big jeff (he/him)
	DisplayName string `json:"display_name,omitempty"`

	// Array of custom emojis used in this account's note or display name.
	// Empty for blocked accounts.
	Emojis []*Emoji `json:"emojis"`

	// Account has enabled RSS feed.
	// Key/value omitted if false.
	EnableRSS bool `json:"enable_rss,omitempty"`

	// Additional metadata attached to this account's profile.
	// Empty for blocked accounts.
	Fields []*Field `json:"fields"`

	// Number of accounts following this account, according to our instance.
	FollowersCount int64 `json:"followers_count,omitempty"`

	// Number of account's followed by this account, according to our instance.
	FollowingCount int64 `json:"following_count,omitempty"`

	// Web location of the account's header image.
	// Example: https://example.org/media/some_user/header/original/header.jpeg
	Header string `json:"header,omitempty"`

	// Description of this account's header, for alt text.
	// Example: A sunlit field with purple flowers.
	HeaderDescription string `json:"header_description,omitempty"`

	// Database ID of the media attachment for this account's header image.
	// Omitted if no header uploaded for this account (ie., default header).
	// Example: 01JAJ3XCD66K3T99JZESCR137W
	HeaderMediaID string `json:"header_media_id,omitempty"`

	// Web location of a static version of the account's header.
	// Only relevant when the account's main header is a video or a gif.
	// Example: https://example.org/media/some_user/header/static/header.png
	HeaderStatic string `json:"header_static,omitempty"`

	// Account has opted to hide their followers/following collections.
	// Key/value omitted if false.
	HideCollections bool `json:"hide_collections,omitempty"`

	// The account id.
	// Example: 01FBVD42CQ3ZEEVMW180SBX03B
	ID string `json:"id,omitempty"`

	// When the account's most recent status was posted (ISO 8601 Date).
	// Example: 2021-07-30
	LastStatusAt string `json:"last_status_at,omitempty"`

	// Account manually approves follow requests.
	Locked bool `json:"locked,omitempty"`

	// If this account has been muted, when will the mute expire (ISO 8601 Datetime).
	// If the mute is indefinite, this will be null.
	// Example: 2021-07-30T09:20:25+00:00
	MuteExpiresAt string `json:"mute_expires_at,omitempty"`

	// Bio/description of this account.
	Note string `json:"note,omitempty"`

	// Roles lists the public roles of the account on this instance.
	// Unlike Role, this is always available, but never includes permissions details.
	// Key/value omitted for remote accounts.
	Roles []*AccountDisplayRole `json:"roles"`

	// Number of statuses posted by this account, according to our instance.
	StatusesCount int64 `json:"statuses_count,omitempty"`

	// Account has been suspended by our instance.
	Suspended bool `json:"suspended,omitempty"`

	// Filename of user-selected CSS theme to include when rendering this account's profile or statuses. Eg., `blurple-light.css`.
	Theme string `json:"theme,omitempty"`

	// Web location of the account's profile page.
	// Example: https://example.org/@some_user
	URL string `json:"url,omitempty"`

	// The username of the account, not including domain.
	// Example: some_user
	Username string `json:"username,omitempty"`

	// moved
	Moved *Account `json:"moved,omitempty"`

	// role
	Role *AccountRole `json:"role,omitempty"`

	// source
	Source *Source `json:"source,omitempty"`
}

// Validate validates this muted account
func (m *MutedAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmojis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutedAccount) validateEmojis(formats strfmt.Registry) error {
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

func (m *MutedAccount) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MutedAccount) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MutedAccount) validateMoved(formats strfmt.Registry) error {
	if swag.IsZero(m.Moved) { // not required
		return nil
	}

	if m.Moved != nil {
		if err := m.Moved.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moved")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moved")
			}
			return err
		}
	}

	return nil
}

func (m *MutedAccount) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *MutedAccount) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this muted account based on the context it is used
func (m *MutedAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmojis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMoved(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutedAccount) contextValidateEmojis(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MutedAccount) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fields); i++ {

		if m.Fields[i] != nil {

			if swag.IsZero(m.Fields[i]) { // not required
				return nil
			}

			if err := m.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MutedAccount) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {

			if swag.IsZero(m.Roles[i]) { // not required
				return nil
			}

			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MutedAccount) contextValidateMoved(ctx context.Context, formats strfmt.Registry) error {

	if m.Moved != nil {

		if swag.IsZero(m.Moved) { // not required
			return nil
		}

		if err := m.Moved.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moved")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moved")
			}
			return err
		}
	}

	return nil
}

func (m *MutedAccount) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {

		if swag.IsZero(m.Role) { // not required
			return nil
		}

		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *MutedAccount) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MutedAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MutedAccount) UnmarshalBinary(b []byte) error {
	var res MutedAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
