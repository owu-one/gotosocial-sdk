// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewDomainPermissionExcludesGetParams creates a new DomainPermissionExcludesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainPermissionExcludesGetParams() *DomainPermissionExcludesGetParams {
	return &DomainPermissionExcludesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainPermissionExcludesGetParamsWithTimeout creates a new DomainPermissionExcludesGetParams object
// with the ability to set a timeout on a request.
func NewDomainPermissionExcludesGetParamsWithTimeout(timeout time.Duration) *DomainPermissionExcludesGetParams {
	return &DomainPermissionExcludesGetParams{
		timeout: timeout,
	}
}

// NewDomainPermissionExcludesGetParamsWithContext creates a new DomainPermissionExcludesGetParams object
// with the ability to set a context for a request.
func NewDomainPermissionExcludesGetParamsWithContext(ctx context.Context) *DomainPermissionExcludesGetParams {
	return &DomainPermissionExcludesGetParams{
		Context: ctx,
	}
}

// NewDomainPermissionExcludesGetParamsWithHTTPClient creates a new DomainPermissionExcludesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainPermissionExcludesGetParamsWithHTTPClient(client *http.Client) *DomainPermissionExcludesGetParams {
	return &DomainPermissionExcludesGetParams{
		HTTPClient: client,
	}
}

/*
DomainPermissionExcludesGetParams contains all the parameters to send to the API endpoint

	for the domain permission excludes get operation.

	Typically these are written to a http.Request.
*/
type DomainPermissionExcludesGetParams struct {

	/* Domain.

	   Return only excludes that target the given domain.
	*/
	Domain *string

	/* Limit.

	   Number of items to return.

	   Default: 20
	*/
	Limit *int64

	/* MaxID.

	   Return only items *OLDER* than the given max ID (for paging downwards). The item with the specified ID will not be included in the response.
	*/
	MaxID *string

	/* MinID.

	   Return only items immediately *NEWER* than the given min ID (for paging upwards). The item with the specified ID will not be included in the response.
	*/
	MinID *string

	/* SinceID.

	   Return only items *NEWER* than the given since ID. The item with the specified ID will not be included in the response.
	*/
	SinceID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain permission excludes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionExcludesGetParams) WithDefaults() *DomainPermissionExcludesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain permission excludes get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionExcludesGetParams) SetDefaults() {
	var (
		limitDefault = int64(20)
	)

	val := DomainPermissionExcludesGetParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithTimeout(timeout time.Duration) *DomainPermissionExcludesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithContext(ctx context.Context) *DomainPermissionExcludesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithHTTPClient(client *http.Client) *DomainPermissionExcludesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithDomain(domain *string) *DomainPermissionExcludesGetParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithLimit adds the limit to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithLimit(limit *int64) *DomainPermissionExcludesGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxID adds the maxID to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithMaxID(maxID *string) *DomainPermissionExcludesGetParams {
	o.SetMaxID(maxID)
	return o
}

// SetMaxID adds the maxId to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetMaxID(maxID *string) {
	o.MaxID = maxID
}

// WithMinID adds the minID to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithMinID(minID *string) *DomainPermissionExcludesGetParams {
	o.SetMinID(minID)
	return o
}

// SetMinID adds the minId to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetMinID(minID *string) {
	o.MinID = minID
}

// WithSinceID adds the sinceID to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) WithSinceID(sinceID *string) *DomainPermissionExcludesGetParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the domain permission excludes get params
func (o *DomainPermissionExcludesGetParams) SetSinceID(sinceID *string) {
	o.SinceID = sinceID
}

// WriteToRequest writes these params to a swagger request
func (o *DomainPermissionExcludesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Domain != nil {

		// query param domain
		var qrDomain string

		if o.Domain != nil {
			qrDomain = *o.Domain
		}
		qDomain := qrDomain
		if qDomain != "" {

			if err := r.SetQueryParam("domain", qDomain); err != nil {
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
