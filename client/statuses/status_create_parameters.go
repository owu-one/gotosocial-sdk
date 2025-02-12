// Code generated by go-swagger; DO NOT EDIT.

package statuses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewStatusCreateParams creates a new StatusCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusCreateParams() *StatusCreateParams {
	return &StatusCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusCreateParamsWithTimeout creates a new StatusCreateParams object
// with the ability to set a timeout on a request.
func NewStatusCreateParamsWithTimeout(timeout time.Duration) *StatusCreateParams {
	return &StatusCreateParams{
		timeout: timeout,
	}
}

// NewStatusCreateParamsWithContext creates a new StatusCreateParams object
// with the ability to set a context for a request.
func NewStatusCreateParamsWithContext(ctx context.Context) *StatusCreateParams {
	return &StatusCreateParams{
		Context: ctx,
	}
}

// NewStatusCreateParamsWithHTTPClient creates a new StatusCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusCreateParamsWithHTTPClient(client *http.Client) *StatusCreateParams {
	return &StatusCreateParams{
		HTTPClient: client,
	}
}

/*
StatusCreateParams contains all the parameters to send to the API endpoint

	for the status create operation.

	Typically these are written to a http.Request.
*/
type StatusCreateParams struct {

	/* ContentType.

	   Content type to use when parsing this status.
	*/
	ContentType *string

	/* Federated.

	 ***DEPRECATED***. Included for back compat only. Only used if set and local_only is not yet. If set to true, this status will be federated beyond the local timeline(s). If set to false, this status will NOT be federated beyond the local timeline(s).
	 */
	Federated *bool

	/* InReplyToID.

	   ID of the status being replied to, if status is a reply.
	*/
	InReplyToID *string

	/* InteractionPolicyCanFavouriteAlways0.

	   Nth entry for interaction_policy.can_favourite.always.
	*/
	InteractionPolicyCanFavouriteAlways0 *string

	/* InteractionPolicyCanFavouriteWithApproval0.

	   Nth entry for interaction_policy.can_favourite.with_approval.
	*/
	InteractionPolicyCanFavouriteWithApproval0 *string

	/* InteractionPolicyCanReblogAlways0.

	   Nth entry for interaction_policy.can_reblog.always.
	*/
	InteractionPolicyCanReblogAlways0 *string

	/* InteractionPolicyCanReblogWithApproval0.

	   Nth entry for interaction_policy.can_reblog.with_approval.
	*/
	InteractionPolicyCanReblogWithApproval0 *string

	/* InteractionPolicyCanReplyAlways0.

	   Nth entry for interaction_policy.can_reply.always.
	*/
	InteractionPolicyCanReplyAlways0 *string

	/* InteractionPolicyCanReplyWithApproval0.

	   Nth entry for interaction_policy.can_reply.with_approval.
	*/
	InteractionPolicyCanReplyWithApproval0 *string

	/* Language.

	   ISO 639 language code for this status.
	*/
	Language *string

	/* LocalOnly.

	   If set to true, this status will be "local only" and will NOT be federated beyond the local timeline(s). If set to false (default), this status will be federated to your followers beyond the local timeline(s).
	*/
	LocalOnly *bool

	/* MediaIds.

	     Array of Attachment ids to be attached as media.
	If provided, status becomes optional, and poll cannot be used.

	If the status is being submitted as a form, the key is 'media_ids[]',
	but if it's json or xml, the key is 'media_ids'.
	*/
	MediaIDs []string

	/* PollExpiresIn.

	     Duration the poll should be open, in seconds.
	If provided, media_ids cannot be used, and poll[options] must be provided.

	     Format: int64
	*/
	PollExpiresIn *int64

	/* PollHideTotals.

	   Hide vote counts until the poll ends.

	   Default: true
	*/
	PollHideTotals *bool

	/* PollMultiple.

	   Allow multiple choices on this poll.
	*/
	PollMultiple *bool

	/* PollOptions.

	     Array of possible poll answers.
	If provided, media_ids cannot be used, and poll[expires_in] must be provided.
	*/
	PollOptions []string

	/* ScheduledAt.

	     ISO 8601 Datetime at which to schedule a status.

	Providing this parameter with a *future* time will cause ScheduledStatus to be returned instead of Status.
	Must be at least 5 minutes in the future.
	This feature isn't implemented yet.

	Providing this parameter with a *past* time will cause the status to be backdated,
	and will not push it to the user's followers. This is intended for importing old statuses.

	     Format: date-time
	*/
	ScheduledAt *strfmt.DateTime

	/* Sensitive.

	   Status and attached media should be marked as sensitive.
	*/
	Sensitive *bool

	/* SpoilerText.

	     Text to be shown as a warning or subject before the actual content.
	Statuses are generally collapsed behind this field.
	*/
	SpoilerText *string

	/* Status.

	     Text content of the status.
	If media_ids is provided, this becomes optional.
	Attaching a poll is optional while status is provided.
	*/
	Status *string

	/* Visibility.

	   Visibility of the posted status.
	*/
	Visibility *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusCreateParams) WithDefaults() *StatusCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusCreateParams) SetDefaults() {
	var (
		localOnlyDefault = bool(false)

		pollHideTotalsDefault = bool(true)

		pollMultipleDefault = bool(false)
	)

	val := StatusCreateParams{
		LocalOnly:      &localOnlyDefault,
		PollHideTotals: &pollHideTotalsDefault,
		PollMultiple:   &pollMultipleDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the status create params
func (o *StatusCreateParams) WithTimeout(timeout time.Duration) *StatusCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status create params
func (o *StatusCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status create params
func (o *StatusCreateParams) WithContext(ctx context.Context) *StatusCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status create params
func (o *StatusCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status create params
func (o *StatusCreateParams) WithHTTPClient(client *http.Client) *StatusCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status create params
func (o *StatusCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the status create params
func (o *StatusCreateParams) WithContentType(contentType *string) *StatusCreateParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the status create params
func (o *StatusCreateParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithFederated adds the federated to the status create params
func (o *StatusCreateParams) WithFederated(federated *bool) *StatusCreateParams {
	o.SetFederated(federated)
	return o
}

// SetFederated adds the federated to the status create params
func (o *StatusCreateParams) SetFederated(federated *bool) {
	o.Federated = federated
}

// WithInReplyToID adds the inReplyToID to the status create params
func (o *StatusCreateParams) WithInReplyToID(inReplyToID *string) *StatusCreateParams {
	o.SetInReplyToID(inReplyToID)
	return o
}

// SetInReplyToID adds the inReplyToId to the status create params
func (o *StatusCreateParams) SetInReplyToID(inReplyToID *string) {
	o.InReplyToID = inReplyToID
}

// WithInteractionPolicyCanFavouriteAlways0 adds the interactionPolicyCanFavouriteAlways0 to the status create params
func (o *StatusCreateParams) WithInteractionPolicyCanFavouriteAlways0(interactionPolicyCanFavouriteAlways0 *string) *StatusCreateParams {
	o.SetInteractionPolicyCanFavouriteAlways0(interactionPolicyCanFavouriteAlways0)
	return o
}

// SetInteractionPolicyCanFavouriteAlways0 adds the interactionPolicyCanFavouriteAlways0 to the status create params
func (o *StatusCreateParams) SetInteractionPolicyCanFavouriteAlways0(interactionPolicyCanFavouriteAlways0 *string) {
	o.InteractionPolicyCanFavouriteAlways0 = interactionPolicyCanFavouriteAlways0
}

// WithInteractionPolicyCanFavouriteWithApproval0 adds the interactionPolicyCanFavouriteWithApproval0 to the status create params
func (o *StatusCreateParams) WithInteractionPolicyCanFavouriteWithApproval0(interactionPolicyCanFavouriteWithApproval0 *string) *StatusCreateParams {
	o.SetInteractionPolicyCanFavouriteWithApproval0(interactionPolicyCanFavouriteWithApproval0)
	return o
}

// SetInteractionPolicyCanFavouriteWithApproval0 adds the interactionPolicyCanFavouriteWithApproval0 to the status create params
func (o *StatusCreateParams) SetInteractionPolicyCanFavouriteWithApproval0(interactionPolicyCanFavouriteWithApproval0 *string) {
	o.InteractionPolicyCanFavouriteWithApproval0 = interactionPolicyCanFavouriteWithApproval0
}

// WithInteractionPolicyCanReblogAlways0 adds the interactionPolicyCanReblogAlways0 to the status create params
func (o *StatusCreateParams) WithInteractionPolicyCanReblogAlways0(interactionPolicyCanReblogAlways0 *string) *StatusCreateParams {
	o.SetInteractionPolicyCanReblogAlways0(interactionPolicyCanReblogAlways0)
	return o
}

// SetInteractionPolicyCanReblogAlways0 adds the interactionPolicyCanReblogAlways0 to the status create params
func (o *StatusCreateParams) SetInteractionPolicyCanReblogAlways0(interactionPolicyCanReblogAlways0 *string) {
	o.InteractionPolicyCanReblogAlways0 = interactionPolicyCanReblogAlways0
}

// WithInteractionPolicyCanReblogWithApproval0 adds the interactionPolicyCanReblogWithApproval0 to the status create params
func (o *StatusCreateParams) WithInteractionPolicyCanReblogWithApproval0(interactionPolicyCanReblogWithApproval0 *string) *StatusCreateParams {
	o.SetInteractionPolicyCanReblogWithApproval0(interactionPolicyCanReblogWithApproval0)
	return o
}

// SetInteractionPolicyCanReblogWithApproval0 adds the interactionPolicyCanReblogWithApproval0 to the status create params
func (o *StatusCreateParams) SetInteractionPolicyCanReblogWithApproval0(interactionPolicyCanReblogWithApproval0 *string) {
	o.InteractionPolicyCanReblogWithApproval0 = interactionPolicyCanReblogWithApproval0
}

// WithInteractionPolicyCanReplyAlways0 adds the interactionPolicyCanReplyAlways0 to the status create params
func (o *StatusCreateParams) WithInteractionPolicyCanReplyAlways0(interactionPolicyCanReplyAlways0 *string) *StatusCreateParams {
	o.SetInteractionPolicyCanReplyAlways0(interactionPolicyCanReplyAlways0)
	return o
}

// SetInteractionPolicyCanReplyAlways0 adds the interactionPolicyCanReplyAlways0 to the status create params
func (o *StatusCreateParams) SetInteractionPolicyCanReplyAlways0(interactionPolicyCanReplyAlways0 *string) {
	o.InteractionPolicyCanReplyAlways0 = interactionPolicyCanReplyAlways0
}

// WithInteractionPolicyCanReplyWithApproval0 adds the interactionPolicyCanReplyWithApproval0 to the status create params
func (o *StatusCreateParams) WithInteractionPolicyCanReplyWithApproval0(interactionPolicyCanReplyWithApproval0 *string) *StatusCreateParams {
	o.SetInteractionPolicyCanReplyWithApproval0(interactionPolicyCanReplyWithApproval0)
	return o
}

// SetInteractionPolicyCanReplyWithApproval0 adds the interactionPolicyCanReplyWithApproval0 to the status create params
func (o *StatusCreateParams) SetInteractionPolicyCanReplyWithApproval0(interactionPolicyCanReplyWithApproval0 *string) {
	o.InteractionPolicyCanReplyWithApproval0 = interactionPolicyCanReplyWithApproval0
}

// WithLanguage adds the language to the status create params
func (o *StatusCreateParams) WithLanguage(language *string) *StatusCreateParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the status create params
func (o *StatusCreateParams) SetLanguage(language *string) {
	o.Language = language
}

// WithLocalOnly adds the localOnly to the status create params
func (o *StatusCreateParams) WithLocalOnly(localOnly *bool) *StatusCreateParams {
	o.SetLocalOnly(localOnly)
	return o
}

// SetLocalOnly adds the localOnly to the status create params
func (o *StatusCreateParams) SetLocalOnly(localOnly *bool) {
	o.LocalOnly = localOnly
}

// WithMediaIDs adds the mediaIds to the status create params
func (o *StatusCreateParams) WithMediaIDs(mediaIds []string) *StatusCreateParams {
	o.SetMediaIDs(mediaIds)
	return o
}

// SetMediaIDs adds the mediaIds to the status create params
func (o *StatusCreateParams) SetMediaIDs(mediaIds []string) {
	o.MediaIDs = mediaIds
}

// WithPollExpiresIn adds the pollExpiresIn to the status create params
func (o *StatusCreateParams) WithPollExpiresIn(pollExpiresIn *int64) *StatusCreateParams {
	o.SetPollExpiresIn(pollExpiresIn)
	return o
}

// SetPollExpiresIn adds the pollExpiresIn to the status create params
func (o *StatusCreateParams) SetPollExpiresIn(pollExpiresIn *int64) {
	o.PollExpiresIn = pollExpiresIn
}

// WithPollHideTotals adds the pollHideTotals to the status create params
func (o *StatusCreateParams) WithPollHideTotals(pollHideTotals *bool) *StatusCreateParams {
	o.SetPollHideTotals(pollHideTotals)
	return o
}

// SetPollHideTotals adds the pollHideTotals to the status create params
func (o *StatusCreateParams) SetPollHideTotals(pollHideTotals *bool) {
	o.PollHideTotals = pollHideTotals
}

// WithPollMultiple adds the pollMultiple to the status create params
func (o *StatusCreateParams) WithPollMultiple(pollMultiple *bool) *StatusCreateParams {
	o.SetPollMultiple(pollMultiple)
	return o
}

// SetPollMultiple adds the pollMultiple to the status create params
func (o *StatusCreateParams) SetPollMultiple(pollMultiple *bool) {
	o.PollMultiple = pollMultiple
}

// WithPollOptions adds the pollOptions to the status create params
func (o *StatusCreateParams) WithPollOptions(pollOptions []string) *StatusCreateParams {
	o.SetPollOptions(pollOptions)
	return o
}

// SetPollOptions adds the pollOptions to the status create params
func (o *StatusCreateParams) SetPollOptions(pollOptions []string) {
	o.PollOptions = pollOptions
}

// WithScheduledAt adds the scheduledAt to the status create params
func (o *StatusCreateParams) WithScheduledAt(scheduledAt *strfmt.DateTime) *StatusCreateParams {
	o.SetScheduledAt(scheduledAt)
	return o
}

// SetScheduledAt adds the scheduledAt to the status create params
func (o *StatusCreateParams) SetScheduledAt(scheduledAt *strfmt.DateTime) {
	o.ScheduledAt = scheduledAt
}

// WithSensitive adds the sensitive to the status create params
func (o *StatusCreateParams) WithSensitive(sensitive *bool) *StatusCreateParams {
	o.SetSensitive(sensitive)
	return o
}

// SetSensitive adds the sensitive to the status create params
func (o *StatusCreateParams) SetSensitive(sensitive *bool) {
	o.Sensitive = sensitive
}

// WithSpoilerText adds the spoilerText to the status create params
func (o *StatusCreateParams) WithSpoilerText(spoilerText *string) *StatusCreateParams {
	o.SetSpoilerText(spoilerText)
	return o
}

// SetSpoilerText adds the spoilerText to the status create params
func (o *StatusCreateParams) SetSpoilerText(spoilerText *string) {
	o.SpoilerText = spoilerText
}

// WithStatus adds the status to the status create params
func (o *StatusCreateParams) WithStatus(status *string) *StatusCreateParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the status create params
func (o *StatusCreateParams) SetStatus(status *string) {
	o.Status = status
}

// WithVisibility adds the visibility to the status create params
func (o *StatusCreateParams) WithVisibility(visibility *string) *StatusCreateParams {
	o.SetVisibility(visibility)
	return o
}

// SetVisibility adds the visibility to the status create params
func (o *StatusCreateParams) SetVisibility(visibility *string) {
	o.Visibility = visibility
}

// WriteToRequest writes these params to a swagger request
func (o *StatusCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// form param content_type
		var frContentType string
		if o.ContentType != nil {
			frContentType = *o.ContentType
		}
		fContentType := frContentType
		if fContentType != "" {
			if err := r.SetFormParam("content_type", fContentType); err != nil {
				return err
			}
		}
	}

	if o.Federated != nil {

		// form param federated
		var frFederated bool
		if o.Federated != nil {
			frFederated = *o.Federated
		}
		fFederated := swag.FormatBool(frFederated)
		if fFederated != "" {
			if err := r.SetFormParam("federated", fFederated); err != nil {
				return err
			}
		}
	}

	if o.InReplyToID != nil {

		// form param in_reply_to_id
		var frInReplyToID string
		if o.InReplyToID != nil {
			frInReplyToID = *o.InReplyToID
		}
		fInReplyToID := frInReplyToID
		if fInReplyToID != "" {
			if err := r.SetFormParam("in_reply_to_id", fInReplyToID); err != nil {
				return err
			}
		}
	}

	if o.InteractionPolicyCanFavouriteAlways0 != nil {

		// form param interaction_policy[can_favourite][always][0]
		var frInteractionPolicyCanFavouriteAlways0 string
		if o.InteractionPolicyCanFavouriteAlways0 != nil {
			frInteractionPolicyCanFavouriteAlways0 = *o.InteractionPolicyCanFavouriteAlways0
		}
		fInteractionPolicyCanFavouriteAlways0 := frInteractionPolicyCanFavouriteAlways0
		if fInteractionPolicyCanFavouriteAlways0 != "" {
			if err := r.SetFormParam("interaction_policy[can_favourite][always][0]", fInteractionPolicyCanFavouriteAlways0); err != nil {
				return err
			}
		}
	}

	if o.InteractionPolicyCanFavouriteWithApproval0 != nil {

		// form param interaction_policy[can_favourite][with_approval][0]
		var frInteractionPolicyCanFavouriteWithApproval0 string
		if o.InteractionPolicyCanFavouriteWithApproval0 != nil {
			frInteractionPolicyCanFavouriteWithApproval0 = *o.InteractionPolicyCanFavouriteWithApproval0
		}
		fInteractionPolicyCanFavouriteWithApproval0 := frInteractionPolicyCanFavouriteWithApproval0
		if fInteractionPolicyCanFavouriteWithApproval0 != "" {
			if err := r.SetFormParam("interaction_policy[can_favourite][with_approval][0]", fInteractionPolicyCanFavouriteWithApproval0); err != nil {
				return err
			}
		}
	}

	if o.InteractionPolicyCanReblogAlways0 != nil {

		// form param interaction_policy[can_reblog][always][0]
		var frInteractionPolicyCanReblogAlways0 string
		if o.InteractionPolicyCanReblogAlways0 != nil {
			frInteractionPolicyCanReblogAlways0 = *o.InteractionPolicyCanReblogAlways0
		}
		fInteractionPolicyCanReblogAlways0 := frInteractionPolicyCanReblogAlways0
		if fInteractionPolicyCanReblogAlways0 != "" {
			if err := r.SetFormParam("interaction_policy[can_reblog][always][0]", fInteractionPolicyCanReblogAlways0); err != nil {
				return err
			}
		}
	}

	if o.InteractionPolicyCanReblogWithApproval0 != nil {

		// form param interaction_policy[can_reblog][with_approval][0]
		var frInteractionPolicyCanReblogWithApproval0 string
		if o.InteractionPolicyCanReblogWithApproval0 != nil {
			frInteractionPolicyCanReblogWithApproval0 = *o.InteractionPolicyCanReblogWithApproval0
		}
		fInteractionPolicyCanReblogWithApproval0 := frInteractionPolicyCanReblogWithApproval0
		if fInteractionPolicyCanReblogWithApproval0 != "" {
			if err := r.SetFormParam("interaction_policy[can_reblog][with_approval][0]", fInteractionPolicyCanReblogWithApproval0); err != nil {
				return err
			}
		}
	}

	if o.InteractionPolicyCanReplyAlways0 != nil {

		// form param interaction_policy[can_reply][always][0]
		var frInteractionPolicyCanReplyAlways0 string
		if o.InteractionPolicyCanReplyAlways0 != nil {
			frInteractionPolicyCanReplyAlways0 = *o.InteractionPolicyCanReplyAlways0
		}
		fInteractionPolicyCanReplyAlways0 := frInteractionPolicyCanReplyAlways0
		if fInteractionPolicyCanReplyAlways0 != "" {
			if err := r.SetFormParam("interaction_policy[can_reply][always][0]", fInteractionPolicyCanReplyAlways0); err != nil {
				return err
			}
		}
	}

	if o.InteractionPolicyCanReplyWithApproval0 != nil {

		// form param interaction_policy[can_reply][with_approval][0]
		var frInteractionPolicyCanReplyWithApproval0 string
		if o.InteractionPolicyCanReplyWithApproval0 != nil {
			frInteractionPolicyCanReplyWithApproval0 = *o.InteractionPolicyCanReplyWithApproval0
		}
		fInteractionPolicyCanReplyWithApproval0 := frInteractionPolicyCanReplyWithApproval0
		if fInteractionPolicyCanReplyWithApproval0 != "" {
			if err := r.SetFormParam("interaction_policy[can_reply][with_approval][0]", fInteractionPolicyCanReplyWithApproval0); err != nil {
				return err
			}
		}
	}

	if o.Language != nil {

		// form param language
		var frLanguage string
		if o.Language != nil {
			frLanguage = *o.Language
		}
		fLanguage := frLanguage
		if fLanguage != "" {
			if err := r.SetFormParam("language", fLanguage); err != nil {
				return err
			}
		}
	}

	if o.LocalOnly != nil {

		// form param local_only
		var frLocalOnly bool
		if o.LocalOnly != nil {
			frLocalOnly = *o.LocalOnly
		}
		fLocalOnly := swag.FormatBool(frLocalOnly)
		if fLocalOnly != "" {
			if err := r.SetFormParam("local_only", fLocalOnly); err != nil {
				return err
			}
		}
	}

	if o.MediaIDs != nil {

		// binding items for media_ids[]
		joinedMediaIds := o.bindParamMediaIds(reg)

		// form array param media_ids[]
		if err := r.SetFormParam("media_ids[]", joinedMediaIds...); err != nil {
			return err
		}
	}

	if o.PollExpiresIn != nil {

		// form param poll[expires_in]
		var frPollExpiresIn int64
		if o.PollExpiresIn != nil {
			frPollExpiresIn = *o.PollExpiresIn
		}
		fPollExpiresIn := swag.FormatInt64(frPollExpiresIn)
		if fPollExpiresIn != "" {
			if err := r.SetFormParam("poll[expires_in]", fPollExpiresIn); err != nil {
				return err
			}
		}
	}

	if o.PollHideTotals != nil {

		// form param poll[hide_totals]
		var frPollHideTotals bool
		if o.PollHideTotals != nil {
			frPollHideTotals = *o.PollHideTotals
		}
		fPollHideTotals := swag.FormatBool(frPollHideTotals)
		if fPollHideTotals != "" {
			if err := r.SetFormParam("poll[hide_totals]", fPollHideTotals); err != nil {
				return err
			}
		}
	}

	if o.PollMultiple != nil {

		// form param poll[multiple]
		var frPollMultiple bool
		if o.PollMultiple != nil {
			frPollMultiple = *o.PollMultiple
		}
		fPollMultiple := swag.FormatBool(frPollMultiple)
		if fPollMultiple != "" {
			if err := r.SetFormParam("poll[multiple]", fPollMultiple); err != nil {
				return err
			}
		}
	}

	if o.PollOptions != nil {

		// binding items for poll[options][]
		joinedPollOptions := o.bindParamPollOptions(reg)

		// form array param poll[options][]
		if err := r.SetFormParam("poll[options][]", joinedPollOptions...); err != nil {
			return err
		}
	}

	if o.ScheduledAt != nil {

		// form param scheduled_at
		var frScheduledAt strfmt.DateTime
		if o.ScheduledAt != nil {
			frScheduledAt = *o.ScheduledAt
		}
		fScheduledAt := frScheduledAt.String()
		if fScheduledAt != "" {
			if err := r.SetFormParam("scheduled_at", fScheduledAt); err != nil {
				return err
			}
		}
	}

	if o.Sensitive != nil {

		// form param sensitive
		var frSensitive bool
		if o.Sensitive != nil {
			frSensitive = *o.Sensitive
		}
		fSensitive := swag.FormatBool(frSensitive)
		if fSensitive != "" {
			if err := r.SetFormParam("sensitive", fSensitive); err != nil {
				return err
			}
		}
	}

	if o.SpoilerText != nil {

		// form param spoiler_text
		var frSpoilerText string
		if o.SpoilerText != nil {
			frSpoilerText = *o.SpoilerText
		}
		fSpoilerText := frSpoilerText
		if fSpoilerText != "" {
			if err := r.SetFormParam("spoiler_text", fSpoilerText); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// form param status
		var frStatus string
		if o.Status != nil {
			frStatus = *o.Status
		}
		fStatus := frStatus
		if fStatus != "" {
			if err := r.SetFormParam("status", fStatus); err != nil {
				return err
			}
		}
	}

	if o.Visibility != nil {

		// form param visibility
		var frVisibility string
		if o.Visibility != nil {
			frVisibility = *o.Visibility
		}
		fVisibility := frVisibility
		if fVisibility != "" {
			if err := r.SetFormParam("visibility", fVisibility); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamStatusCreate binds the parameter media_ids[]
func (o *StatusCreateParams) bindParamMediaIds(formats strfmt.Registry) []string {
	mediaIdsIR := o.MediaIDs

	var mediaIdsIC []string
	for _, mediaIdsIIR := range mediaIdsIR { // explode []string

		mediaIdsIIV := mediaIdsIIR // string as string
		mediaIdsIC = append(mediaIdsIC, mediaIdsIIV)
	}

	// items.CollectionFormat: "multi"
	mediaIdsIS := swag.JoinByFormat(mediaIdsIC, "multi")

	return mediaIdsIS
}

// bindParamStatusCreate binds the parameter poll[options][]
func (o *StatusCreateParams) bindParamPollOptions(formats strfmt.Registry) []string {
	pollOptionsIR := o.PollOptions

	var pollOptionsIC []string
	for _, pollOptionsIIR := range pollOptionsIR { // explode []string

		pollOptionsIIV := pollOptionsIIR // string as string
		pollOptionsIC = append(pollOptionsIC, pollOptionsIIV)
	}

	// items.CollectionFormat: "multi"
	pollOptionsIS := swag.JoinByFormat(pollOptionsIC, "multi")

	return pollOptionsIS
}
