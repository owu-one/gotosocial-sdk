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

// NewInstanceGetV2Params creates a new InstanceGetV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstanceGetV2Params() *InstanceGetV2Params {
	return &InstanceGetV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstanceGetV2ParamsWithTimeout creates a new InstanceGetV2Params object
// with the ability to set a timeout on a request.
func NewInstanceGetV2ParamsWithTimeout(timeout time.Duration) *InstanceGetV2Params {
	return &InstanceGetV2Params{
		timeout: timeout,
	}
}

// NewInstanceGetV2ParamsWithContext creates a new InstanceGetV2Params object
// with the ability to set a context for a request.
func NewInstanceGetV2ParamsWithContext(ctx context.Context) *InstanceGetV2Params {
	return &InstanceGetV2Params{
		Context: ctx,
	}
}

// NewInstanceGetV2ParamsWithHTTPClient creates a new InstanceGetV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewInstanceGetV2ParamsWithHTTPClient(client *http.Client) *InstanceGetV2Params {
	return &InstanceGetV2Params{
		HTTPClient: client,
	}
}

/*
InstanceGetV2Params contains all the parameters to send to the API endpoint

	for the instance get v2 operation.

	Typically these are written to a http.Request.
*/
type InstanceGetV2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the instance get v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceGetV2Params) WithDefaults() *InstanceGetV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the instance get v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstanceGetV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the instance get v2 params
func (o *InstanceGetV2Params) WithTimeout(timeout time.Duration) *InstanceGetV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the instance get v2 params
func (o *InstanceGetV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the instance get v2 params
func (o *InstanceGetV2Params) WithContext(ctx context.Context) *InstanceGetV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the instance get v2 params
func (o *InstanceGetV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the instance get v2 params
func (o *InstanceGetV2Params) WithHTTPClient(client *http.Client) *InstanceGetV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the instance get v2 params
func (o *InstanceGetV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InstanceGetV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
