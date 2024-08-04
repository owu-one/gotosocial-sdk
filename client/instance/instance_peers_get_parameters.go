// Code generated by go-swagger; DO NOT EDIT.

package instance

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
)

// NewInstancePeersGetParams creates a new InstancePeersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstancePeersGetParams() *InstancePeersGetParams {
	return &InstancePeersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstancePeersGetParamsWithTimeout creates a new InstancePeersGetParams object
// with the ability to set a timeout on a request.
func NewInstancePeersGetParamsWithTimeout(timeout time.Duration) *InstancePeersGetParams {
	return &InstancePeersGetParams{
		timeout: timeout,
	}
}

// NewInstancePeersGetParamsWithContext creates a new InstancePeersGetParams object
// with the ability to set a context for a request.
func NewInstancePeersGetParamsWithContext(ctx context.Context) *InstancePeersGetParams {
	return &InstancePeersGetParams{
		Context: ctx,
	}
}

// NewInstancePeersGetParamsWithHTTPClient creates a new InstancePeersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstancePeersGetParamsWithHTTPClient(client *http.Client) *InstancePeersGetParams {
	return &InstancePeersGetParams{
		HTTPClient: client,
	}
}

/*
InstancePeersGetParams contains all the parameters to send to the API endpoint

	for the instance peers get operation.

	Typically these are written to a http.Request.
*/
type InstancePeersGetParams struct {

	/* Filter.

	     Comma-separated list of filters to apply to results. Recognized filters are:
	  - `open` -- include peers that are not suspended or silenced
	  - `suspended` -- include peers that have been suspended.

	If filter is `open`, only instances that haven't been suspended or silenced will be returned.

	If filter is `suspended`, only suspended instances will be shown.

	If filter is `open,suspended`, then all known instances will be returned.

	If filter is an empty string or not set, then `open` will be assumed as the default.

	     Default: "open"
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the instance peers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstancePeersGetParams) WithDefaults() *InstancePeersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instance peers get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstancePeersGetParams) SetDefaults() {
	var (
		filterDefault = string("open")
	)

	val := InstancePeersGetParams{
		Filter: &filterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the instance peers get params
func (o *InstancePeersGetParams) WithTimeout(timeout time.Duration) *InstancePeersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance peers get params
func (o *InstancePeersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance peers get params
func (o *InstancePeersGetParams) WithContext(ctx context.Context) *InstancePeersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance peers get params
func (o *InstancePeersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance peers get params
func (o *InstancePeersGetParams) WithHTTPClient(client *http.Client) *InstancePeersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance peers get params
func (o *InstancePeersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the instance peers get params
func (o *InstancePeersGetParams) WithFilter(filter *string) *InstancePeersGetParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the instance peers get params
func (o *InstancePeersGetParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *InstancePeersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}