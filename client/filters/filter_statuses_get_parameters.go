// Code generated by go-swagger; DO NOT EDIT.

package filters

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

// NewFilterStatusesGetParams creates a new FilterStatusesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFilterStatusesGetParams() *FilterStatusesGetParams {
	return &FilterStatusesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFilterStatusesGetParamsWithTimeout creates a new FilterStatusesGetParams object
// with the ability to set a timeout on a request.
func NewFilterStatusesGetParamsWithTimeout(timeout time.Duration) *FilterStatusesGetParams {
	return &FilterStatusesGetParams{
		timeout: timeout,
	}
}

// NewFilterStatusesGetParamsWithContext creates a new FilterStatusesGetParams object
// with the ability to set a context for a request.
func NewFilterStatusesGetParamsWithContext(ctx context.Context) *FilterStatusesGetParams {
	return &FilterStatusesGetParams{
		Context: ctx,
	}
}

// NewFilterStatusesGetParamsWithHTTPClient creates a new FilterStatusesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFilterStatusesGetParamsWithHTTPClient(client *http.Client) *FilterStatusesGetParams {
	return &FilterStatusesGetParams{
		HTTPClient: client,
	}
}

/*
FilterStatusesGetParams contains all the parameters to send to the API endpoint

	for the filter statuses get operation.

	Typically these are written to a http.Request.
*/
type FilterStatusesGetParams struct {

	/* ID.

	   ID of the filter
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the filter statuses get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FilterStatusesGetParams) WithDefaults() *FilterStatusesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the filter statuses get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FilterStatusesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the filter statuses get params
func (o *FilterStatusesGetParams) WithTimeout(timeout time.Duration) *FilterStatusesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the filter statuses get params
func (o *FilterStatusesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the filter statuses get params
func (o *FilterStatusesGetParams) WithContext(ctx context.Context) *FilterStatusesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the filter statuses get params
func (o *FilterStatusesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the filter statuses get params
func (o *FilterStatusesGetParams) WithHTTPClient(client *http.Client) *FilterStatusesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the filter statuses get params
func (o *FilterStatusesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the filter statuses get params
func (o *FilterStatusesGetParams) WithID(id string) *FilterStatusesGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the filter statuses get params
func (o *FilterStatusesGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FilterStatusesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}