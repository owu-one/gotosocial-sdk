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

// NewHeaderFilterBlockGetParams creates a new HeaderFilterBlockGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeaderFilterBlockGetParams() *HeaderFilterBlockGetParams {
	return &HeaderFilterBlockGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeaderFilterBlockGetParamsWithTimeout creates a new HeaderFilterBlockGetParams object
// with the ability to set a timeout on a request.
func NewHeaderFilterBlockGetParamsWithTimeout(timeout time.Duration) *HeaderFilterBlockGetParams {
	return &HeaderFilterBlockGetParams{
		timeout: timeout,
	}
}

// NewHeaderFilterBlockGetParamsWithContext creates a new HeaderFilterBlockGetParams object
// with the ability to set a context for a request.
func NewHeaderFilterBlockGetParamsWithContext(ctx context.Context) *HeaderFilterBlockGetParams {
	return &HeaderFilterBlockGetParams{
		Context: ctx,
	}
}

// NewHeaderFilterBlockGetParamsWithHTTPClient creates a new HeaderFilterBlockGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeaderFilterBlockGetParamsWithHTTPClient(client *http.Client) *HeaderFilterBlockGetParams {
	return &HeaderFilterBlockGetParams{
		HTTPClient: client,
	}
}

/*
HeaderFilterBlockGetParams contains all the parameters to send to the API endpoint

	for the header filter block get operation.

	Typically these are written to a http.Request.
*/
type HeaderFilterBlockGetParams struct {

	/* ID.

	   Target header filter ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the header filter block get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeaderFilterBlockGetParams) WithDefaults() *HeaderFilterBlockGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the header filter block get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeaderFilterBlockGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the header filter block get params
func (o *HeaderFilterBlockGetParams) WithTimeout(timeout time.Duration) *HeaderFilterBlockGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the header filter block get params
func (o *HeaderFilterBlockGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the header filter block get params
func (o *HeaderFilterBlockGetParams) WithContext(ctx context.Context) *HeaderFilterBlockGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the header filter block get params
func (o *HeaderFilterBlockGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the header filter block get params
func (o *HeaderFilterBlockGetParams) WithHTTPClient(client *http.Client) *HeaderFilterBlockGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the header filter block get params
func (o *HeaderFilterBlockGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the header filter block get params
func (o *HeaderFilterBlockGetParams) WithID(id string) *HeaderFilterBlockGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the header filter block get params
func (o *HeaderFilterBlockGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HeaderFilterBlockGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
