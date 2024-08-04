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
)

// NewHeaderFilterAllowGetParams creates a new HeaderFilterAllowGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeaderFilterAllowGetParams() *HeaderFilterAllowGetParams {
	return &HeaderFilterAllowGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeaderFilterAllowGetParamsWithTimeout creates a new HeaderFilterAllowGetParams object
// with the ability to set a timeout on a request.
func NewHeaderFilterAllowGetParamsWithTimeout(timeout time.Duration) *HeaderFilterAllowGetParams {
	return &HeaderFilterAllowGetParams{
		timeout: timeout,
	}
}

// NewHeaderFilterAllowGetParamsWithContext creates a new HeaderFilterAllowGetParams object
// with the ability to set a context for a request.
func NewHeaderFilterAllowGetParamsWithContext(ctx context.Context) *HeaderFilterAllowGetParams {
	return &HeaderFilterAllowGetParams{
		Context: ctx,
	}
}

// NewHeaderFilterAllowGetParamsWithHTTPClient creates a new HeaderFilterAllowGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeaderFilterAllowGetParamsWithHTTPClient(client *http.Client) *HeaderFilterAllowGetParams {
	return &HeaderFilterAllowGetParams{
		HTTPClient: client,
	}
}

/*
HeaderFilterAllowGetParams contains all the parameters to send to the API endpoint

	for the header filter allow get operation.

	Typically these are written to a http.Request.
*/
type HeaderFilterAllowGetParams struct {

	/* ID.

	   Target header filter ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the header filter allow get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeaderFilterAllowGetParams) WithDefaults() *HeaderFilterAllowGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the header filter allow get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeaderFilterAllowGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the header filter allow get params
func (o *HeaderFilterAllowGetParams) WithTimeout(timeout time.Duration) *HeaderFilterAllowGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the header filter allow get params
func (o *HeaderFilterAllowGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the header filter allow get params
func (o *HeaderFilterAllowGetParams) WithContext(ctx context.Context) *HeaderFilterAllowGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the header filter allow get params
func (o *HeaderFilterAllowGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the header filter allow get params
func (o *HeaderFilterAllowGetParams) WithHTTPClient(client *http.Client) *HeaderFilterAllowGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the header filter allow get params
func (o *HeaderFilterAllowGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the header filter allow get params
func (o *HeaderFilterAllowGetParams) WithID(id string) *HeaderFilterAllowGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the header filter allow get params
func (o *HeaderFilterAllowGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HeaderFilterAllowGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
