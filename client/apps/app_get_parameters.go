// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewAppGetParams creates a new AppGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppGetParams() *AppGetParams {
	return &AppGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppGetParamsWithTimeout creates a new AppGetParams object
// with the ability to set a timeout on a request.
func NewAppGetParamsWithTimeout(timeout time.Duration) *AppGetParams {
	return &AppGetParams{
		timeout: timeout,
	}
}

// NewAppGetParamsWithContext creates a new AppGetParams object
// with the ability to set a context for a request.
func NewAppGetParamsWithContext(ctx context.Context) *AppGetParams {
	return &AppGetParams{
		Context: ctx,
	}
}

// NewAppGetParamsWithHTTPClient creates a new AppGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppGetParamsWithHTTPClient(client *http.Client) *AppGetParams {
	return &AppGetParams{
		HTTPClient: client,
	}
}

/*
AppGetParams contains all the parameters to send to the API endpoint

	for the app get operation.

	Typically these are written to a http.Request.
*/
type AppGetParams struct {

	/* ID.

	   The id of the requested application.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the app get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppGetParams) WithDefaults() *AppGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the app get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the app get params
func (o *AppGetParams) WithTimeout(timeout time.Duration) *AppGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app get params
func (o *AppGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app get params
func (o *AppGetParams) WithContext(ctx context.Context) *AppGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app get params
func (o *AppGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app get params
func (o *AppGetParams) WithHTTPClient(client *http.Client) *AppGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app get params
func (o *AppGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the app get params
func (o *AppGetParams) WithID(id string) *AppGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the app get params
func (o *AppGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AppGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
