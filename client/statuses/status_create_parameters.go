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

	/* Boostable.

	   This status can be boosted/reblogged.
	*/
	Boostable *bool

	/* ContentType.

	   Content type to use when parsing this status.
	*/
	ContentType *string

	/* Federated.

	   This status will be federated beyond the local timeline(s).
	*/
	Federated *bool

	/* InReplyToID.

	   ID of the status being replied to, if status is a reply.
	*/
	InReplyToID *string

	/* Language.

	   ISO 639 language code for this status.
	*/
	Language *string

	/* Likeable.

	   This status can be liked/faved.
	*/
	Likeable *bool

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

	/* Replyable.

	   This status can be replied to.
	*/
	Replyable *bool

	/* ScheduledAt.

	     ISO 8601 Datetime at which to schedule a status.
	Providing this parameter will cause ScheduledStatus to be returned instead of Status.
	Must be at least 5 minutes in the future.

	This feature isn't implemented yet.
	*/
	ScheduledAt *string

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
		pollHideTotalsDefault = bool(true)

		pollMultipleDefault = bool(false)
	)

	val := StatusCreateParams{
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

// WithBoostable adds the boostable to the status create params
func (o *StatusCreateParams) WithBoostable(boostable *bool) *StatusCreateParams {
	o.SetBoostable(boostable)
	return o
}

// SetBoostable adds the boostable to the status create params
func (o *StatusCreateParams) SetBoostable(boostable *bool) {
	o.Boostable = boostable
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

// WithLanguage adds the language to the status create params
func (o *StatusCreateParams) WithLanguage(language *string) *StatusCreateParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the status create params
func (o *StatusCreateParams) SetLanguage(language *string) {
	o.Language = language
}

// WithLikeable adds the likeable to the status create params
func (o *StatusCreateParams) WithLikeable(likeable *bool) *StatusCreateParams {
	o.SetLikeable(likeable)
	return o
}

// SetLikeable adds the likeable to the status create params
func (o *StatusCreateParams) SetLikeable(likeable *bool) {
	o.Likeable = likeable
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

// WithReplyable adds the replyable to the status create params
func (o *StatusCreateParams) WithReplyable(replyable *bool) *StatusCreateParams {
	o.SetReplyable(replyable)
	return o
}

// SetReplyable adds the replyable to the status create params
func (o *StatusCreateParams) SetReplyable(replyable *bool) {
	o.Replyable = replyable
}

// WithScheduledAt adds the scheduledAt to the status create params
func (o *StatusCreateParams) WithScheduledAt(scheduledAt *string) *StatusCreateParams {
	o.SetScheduledAt(scheduledAt)
	return o
}

// SetScheduledAt adds the scheduledAt to the status create params
func (o *StatusCreateParams) SetScheduledAt(scheduledAt *string) {
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

	if o.Boostable != nil {

		// form param boostable
		var frBoostable bool
		if o.Boostable != nil {
			frBoostable = *o.Boostable
		}
		fBoostable := swag.FormatBool(frBoostable)
		if fBoostable != "" {
			if err := r.SetFormParam("boostable", fBoostable); err != nil {
				return err
			}
		}
	}

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

	if o.Likeable != nil {

		// form param likeable
		var frLikeable bool
		if o.Likeable != nil {
			frLikeable = *o.Likeable
		}
		fLikeable := swag.FormatBool(frLikeable)
		if fLikeable != "" {
			if err := r.SetFormParam("likeable", fLikeable); err != nil {
				return err
			}
		}
	}

	if o.MediaIDs != nil {

		// binding items for media_ids
		joinedMediaIds := o.bindParamMediaIds(reg)

		// form array param media_ids
		if err := r.SetFormParam("media_ids", joinedMediaIds...); err != nil {
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

	if o.Replyable != nil {

		// form param replyable
		var frReplyable bool
		if o.Replyable != nil {
			frReplyable = *o.Replyable
		}
		fReplyable := swag.FormatBool(frReplyable)
		if fReplyable != "" {
			if err := r.SetFormParam("replyable", fReplyable); err != nil {
				return err
			}
		}
	}

	if o.ScheduledAt != nil {

		// form param scheduled_at
		var frScheduledAt string
		if o.ScheduledAt != nil {
			frScheduledAt = *o.ScheduledAt
		}
		fScheduledAt := frScheduledAt
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

// bindParamStatusCreate binds the parameter media_ids
func (o *StatusCreateParams) bindParamMediaIds(formats strfmt.Registry) []string {
	mediaIdsIR := o.MediaIDs

	var mediaIdsIC []string
	for _, mediaIdsIIR := range mediaIdsIR { // explode []string

		mediaIdsIIV := mediaIdsIIR // string as string
		mediaIdsIC = append(mediaIdsIC, mediaIdsIIV)
	}

	// items.CollectionFormat: ""
	mediaIdsIS := swag.JoinByFormat(mediaIdsIC, "")

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

	// items.CollectionFormat: ""
	pollOptionsIS := swag.JoinByFormat(pollOptionsIC, "")

	return pollOptionsIS
}
