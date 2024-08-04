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
)

// NewThreadContextParams creates a new ThreadContextParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewThreadContextParams() *ThreadContextParams {
	return &ThreadContextParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewThreadContextParamsWithTimeout creates a new ThreadContextParams object
// with the ability to set a timeout on a request.
func NewThreadContextParamsWithTimeout(timeout time.Duration) *ThreadContextParams {
	return &ThreadContextParams{
		timeout: timeout,
	}
}

// NewThreadContextParamsWithContext creates a new ThreadContextParams object
// with the ability to set a context for a request.
func NewThreadContextParamsWithContext(ctx context.Context) *ThreadContextParams {
	return &ThreadContextParams{
		Context: ctx,
	}
}

// NewThreadContextParamsWithHTTPClient creates a new ThreadContextParams object
// with the ability to set a custom HTTPClient for a request.
func NewThreadContextParamsWithHTTPClient(client *http.Client) *ThreadContextParams {
	return &ThreadContextParams{
		HTTPClient: client,
	}
}

/*
ThreadContextParams contains all the parameters to send to the API endpoint

	for the thread context operation.

	Typically these are written to a http.Request.
*/
type ThreadContextParams struct {

	/* ID.

	   Target status ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the thread context params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThreadContextParams) WithDefaults() *ThreadContextParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the thread context params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ThreadContextParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the thread context params
func (o *ThreadContextParams) WithTimeout(timeout time.Duration) *ThreadContextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the thread context params
func (o *ThreadContextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the thread context params
func (o *ThreadContextParams) WithContext(ctx context.Context) *ThreadContextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the thread context params
func (o *ThreadContextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the thread context params
func (o *ThreadContextParams) WithHTTPClient(client *http.Client) *ThreadContextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the thread context params
func (o *ThreadContextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the thread context params
func (o *ThreadContextParams) WithID(id string) *ThreadContextParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the thread context params
func (o *ThreadContextParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ThreadContextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
