// Code generated by go-swagger; DO NOT EDIT.

package interaction_requests

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

// NewGetInteractionRequestsParams creates a new GetInteractionRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInteractionRequestsParams() *GetInteractionRequestsParams {
	return &GetInteractionRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInteractionRequestsParamsWithTimeout creates a new GetInteractionRequestsParams object
// with the ability to set a timeout on a request.
func NewGetInteractionRequestsParamsWithTimeout(timeout time.Duration) *GetInteractionRequestsParams {
	return &GetInteractionRequestsParams{
		timeout: timeout,
	}
}

// NewGetInteractionRequestsParamsWithContext creates a new GetInteractionRequestsParams object
// with the ability to set a context for a request.
func NewGetInteractionRequestsParamsWithContext(ctx context.Context) *GetInteractionRequestsParams {
	return &GetInteractionRequestsParams{
		Context: ctx,
	}
}

// NewGetInteractionRequestsParamsWithHTTPClient creates a new GetInteractionRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInteractionRequestsParamsWithHTTPClient(client *http.Client) *GetInteractionRequestsParams {
	return &GetInteractionRequestsParams{
		HTTPClient: client,
	}
}

/*
GetInteractionRequestsParams contains all the parameters to send to the API endpoint

	for the get interaction requests operation.

	Typically these are written to a http.Request.
*/
type GetInteractionRequestsParams struct {

	/* Favourites.

	   If true or not set, pending favourites will be included in the results. At least one of favourites, replies, and reblogs must be true.

	   Default: true
	*/
	Favourites *bool

	/* Limit.

	   Number of interaction requests to return.

	   Default: 40
	*/
	Limit *int64

	/* MaxID.

	   Return only interaction requests *OLDER* than the given max ID. The interaction with the specified ID will not be included in the response.
	*/
	MaxID *string

	/* MinID.

	   Return only interaction requests *IMMEDIATELY NEWER* than the given min ID. The interaction with the specified ID will not be included in the response.
	*/
	MinID *string

	/* Reblogs.

	   If true or not set, pending reblogs will be included in the results. At least one of favourites, replies, and reblogs must be true.

	   Default: true
	*/
	Reblogs *bool

	/* Replies.

	   If true or not set, pending replies will be included in the results. At least one of favourites, replies, and reblogs must be true.

	   Default: true
	*/
	Replies *bool

	/* SinceID.

	   Return only interaction requests *NEWER* than the given since ID. The interaction with the specified ID will not be included in the response.
	*/
	SinceID *string

	/* StatusID.

	   If set, then only interactions targeting the given status_id will be included in the results.
	*/
	StatusID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get interaction requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInteractionRequestsParams) WithDefaults() *GetInteractionRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get interaction requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInteractionRequestsParams) SetDefaults() {
	var (
		favouritesDefault = bool(true)

		limitDefault = int64(40)

		reblogsDefault = bool(true)

		repliesDefault = bool(true)
	)

	val := GetInteractionRequestsParams{
		Favourites: &favouritesDefault,
		Limit:      &limitDefault,
		Reblogs:    &reblogsDefault,
		Replies:    &repliesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get interaction requests params
func (o *GetInteractionRequestsParams) WithTimeout(timeout time.Duration) *GetInteractionRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get interaction requests params
func (o *GetInteractionRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get interaction requests params
func (o *GetInteractionRequestsParams) WithContext(ctx context.Context) *GetInteractionRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get interaction requests params
func (o *GetInteractionRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get interaction requests params
func (o *GetInteractionRequestsParams) WithHTTPClient(client *http.Client) *GetInteractionRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get interaction requests params
func (o *GetInteractionRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFavourites adds the favourites to the get interaction requests params
func (o *GetInteractionRequestsParams) WithFavourites(favourites *bool) *GetInteractionRequestsParams {
	o.SetFavourites(favourites)
	return o
}

// SetFavourites adds the favourites to the get interaction requests params
func (o *GetInteractionRequestsParams) SetFavourites(favourites *bool) {
	o.Favourites = favourites
}

// WithLimit adds the limit to the get interaction requests params
func (o *GetInteractionRequestsParams) WithLimit(limit *int64) *GetInteractionRequestsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get interaction requests params
func (o *GetInteractionRequestsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxID adds the maxID to the get interaction requests params
func (o *GetInteractionRequestsParams) WithMaxID(maxID *string) *GetInteractionRequestsParams {
	o.SetMaxID(maxID)
	return o
}

// SetMaxID adds the maxId to the get interaction requests params
func (o *GetInteractionRequestsParams) SetMaxID(maxID *string) {
	o.MaxID = maxID
}

// WithMinID adds the minID to the get interaction requests params
func (o *GetInteractionRequestsParams) WithMinID(minID *string) *GetInteractionRequestsParams {
	o.SetMinID(minID)
	return o
}

// SetMinID adds the minId to the get interaction requests params
func (o *GetInteractionRequestsParams) SetMinID(minID *string) {
	o.MinID = minID
}

// WithReblogs adds the reblogs to the get interaction requests params
func (o *GetInteractionRequestsParams) WithReblogs(reblogs *bool) *GetInteractionRequestsParams {
	o.SetReblogs(reblogs)
	return o
}

// SetReblogs adds the reblogs to the get interaction requests params
func (o *GetInteractionRequestsParams) SetReblogs(reblogs *bool) {
	o.Reblogs = reblogs
}

// WithReplies adds the replies to the get interaction requests params
func (o *GetInteractionRequestsParams) WithReplies(replies *bool) *GetInteractionRequestsParams {
	o.SetReplies(replies)
	return o
}

// SetReplies adds the replies to the get interaction requests params
func (o *GetInteractionRequestsParams) SetReplies(replies *bool) {
	o.Replies = replies
}

// WithSinceID adds the sinceID to the get interaction requests params
func (o *GetInteractionRequestsParams) WithSinceID(sinceID *string) *GetInteractionRequestsParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the get interaction requests params
func (o *GetInteractionRequestsParams) SetSinceID(sinceID *string) {
	o.SinceID = sinceID
}

// WithStatusID adds the statusID to the get interaction requests params
func (o *GetInteractionRequestsParams) WithStatusID(statusID *string) *GetInteractionRequestsParams {
	o.SetStatusID(statusID)
	return o
}

// SetStatusID adds the statusId to the get interaction requests params
func (o *GetInteractionRequestsParams) SetStatusID(statusID *string) {
	o.StatusID = statusID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInteractionRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Favourites != nil {

		// query param favourites
		var qrFavourites bool

		if o.Favourites != nil {
			qrFavourites = *o.Favourites
		}
		qFavourites := swag.FormatBool(qrFavourites)
		if qFavourites != "" {

			if err := r.SetQueryParam("favourites", qFavourites); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.MaxID != nil {

		// query param max_id
		var qrMaxID string

		if o.MaxID != nil {
			qrMaxID = *o.MaxID
		}
		qMaxID := qrMaxID
		if qMaxID != "" {

			if err := r.SetQueryParam("max_id", qMaxID); err != nil {
				return err
			}
		}
	}

	if o.MinID != nil {

		// query param min_id
		var qrMinID string

		if o.MinID != nil {
			qrMinID = *o.MinID
		}
		qMinID := qrMinID
		if qMinID != "" {

			if err := r.SetQueryParam("min_id", qMinID); err != nil {
				return err
			}
		}
	}

	if o.Reblogs != nil {

		// query param reblogs
		var qrReblogs bool

		if o.Reblogs != nil {
			qrReblogs = *o.Reblogs
		}
		qReblogs := swag.FormatBool(qrReblogs)
		if qReblogs != "" {

			if err := r.SetQueryParam("reblogs", qReblogs); err != nil {
				return err
			}
		}
	}

	if o.Replies != nil {

		// query param replies
		var qrReplies bool

		if o.Replies != nil {
			qrReplies = *o.Replies
		}
		qReplies := swag.FormatBool(qrReplies)
		if qReplies != "" {

			if err := r.SetQueryParam("replies", qReplies); err != nil {
				return err
			}
		}
	}

	if o.SinceID != nil {

		// query param since_id
		var qrSinceID string

		if o.SinceID != nil {
			qrSinceID = *o.SinceID
		}
		qSinceID := qrSinceID
		if qSinceID != "" {

			if err := r.SetQueryParam("since_id", qSinceID); err != nil {
				return err
			}
		}
	}

	if o.StatusID != nil {

		// query param status_id
		var qrStatusID string

		if o.StatusID != nil {
			qrStatusID = *o.StatusID
		}
		qStatusID := qrStatusID
		if qStatusID != "" {

			if err := r.SetQueryParam("status_id", qStatusID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
