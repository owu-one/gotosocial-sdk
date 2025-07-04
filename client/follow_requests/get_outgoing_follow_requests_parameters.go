// Code generated by go-swagger; DO NOT EDIT.

package follow_requests

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

// NewGetOutgoingFollowRequestsParams creates a new GetOutgoingFollowRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutgoingFollowRequestsParams() *GetOutgoingFollowRequestsParams {
	return &GetOutgoingFollowRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutgoingFollowRequestsParamsWithTimeout creates a new GetOutgoingFollowRequestsParams object
// with the ability to set a timeout on a request.
func NewGetOutgoingFollowRequestsParamsWithTimeout(timeout time.Duration) *GetOutgoingFollowRequestsParams {
	return &GetOutgoingFollowRequestsParams{
		timeout: timeout,
	}
}

// NewGetOutgoingFollowRequestsParamsWithContext creates a new GetOutgoingFollowRequestsParams object
// with the ability to set a context for a request.
func NewGetOutgoingFollowRequestsParamsWithContext(ctx context.Context) *GetOutgoingFollowRequestsParams {
	return &GetOutgoingFollowRequestsParams{
		Context: ctx,
	}
}

// NewGetOutgoingFollowRequestsParamsWithHTTPClient creates a new GetOutgoingFollowRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutgoingFollowRequestsParamsWithHTTPClient(client *http.Client) *GetOutgoingFollowRequestsParams {
	return &GetOutgoingFollowRequestsParams{
		HTTPClient: client,
	}
}

/*
GetOutgoingFollowRequestsParams contains all the parameters to send to the API endpoint

	for the get outgoing follow requests operation.

	Typically these are written to a http.Request.
*/
type GetOutgoingFollowRequestsParams struct {

	/* Limit.

	   Number of follow requested accounts to return.

	   Default: 40
	*/
	Limit *int64

	/* MaxID.

	   Return only follow requested accounts *OLDER* than the given max ID. The follow requestee with the specified ID will not be included in the response. NOTE: the ID is of the internal follow request, NOT any of the returned accounts.
	*/
	MaxID *string

	/* MinID.

	   Return only follow requested accounts *IMMEDIATELY NEWER* than the given min ID. The follow requestee with the specified ID will not be included in the response. NOTE: the ID is of the internal follow request, NOT any of the returned accounts.
	*/
	MinID *string

	/* SinceID.

	   Return only follow requested accounts *NEWER* than the given since ID. The follow requestee with the specified ID will not be included in the response. NOTE: the ID is of the internal follow request, NOT any of the returned accounts.
	*/
	SinceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outgoing follow requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutgoingFollowRequestsParams) WithDefaults() *GetOutgoingFollowRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outgoing follow requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutgoingFollowRequestsParams) SetDefaults() {
	var (
		limitDefault = int64(40)
	)

	val := GetOutgoingFollowRequestsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithTimeout(timeout time.Duration) *GetOutgoingFollowRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithContext(ctx context.Context) *GetOutgoingFollowRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithHTTPClient(client *http.Client) *GetOutgoingFollowRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithLimit(limit *int64) *GetOutgoingFollowRequestsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxID adds the maxID to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithMaxID(maxID *string) *GetOutgoingFollowRequestsParams {
	o.SetMaxID(maxID)
	return o
}

// SetMaxID adds the maxId to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetMaxID(maxID *string) {
	o.MaxID = maxID
}

// WithMinID adds the minID to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithMinID(minID *string) *GetOutgoingFollowRequestsParams {
	o.SetMinID(minID)
	return o
}

// SetMinID adds the minId to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetMinID(minID *string) {
	o.MinID = minID
}

// WithSinceID adds the sinceID to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) WithSinceID(sinceID *string) *GetOutgoingFollowRequestsParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the get outgoing follow requests params
func (o *GetOutgoingFollowRequestsParams) SetSinceID(sinceID *string) {
	o.SinceID = sinceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutgoingFollowRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
