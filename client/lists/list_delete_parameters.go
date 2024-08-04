// Code generated by go-swagger; DO NOT EDIT.

package lists

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

// NewListDeleteParams creates a new ListDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDeleteParams() *ListDeleteParams {
	return &ListDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDeleteParamsWithTimeout creates a new ListDeleteParams object
// with the ability to set a timeout on a request.
func NewListDeleteParamsWithTimeout(timeout time.Duration) *ListDeleteParams {
	return &ListDeleteParams{
		timeout: timeout,
	}
}

// NewListDeleteParamsWithContext creates a new ListDeleteParams object
// with the ability to set a context for a request.
func NewListDeleteParamsWithContext(ctx context.Context) *ListDeleteParams {
	return &ListDeleteParams{
		Context: ctx,
	}
}

// NewListDeleteParamsWithHTTPClient creates a new ListDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDeleteParamsWithHTTPClient(client *http.Client) *ListDeleteParams {
	return &ListDeleteParams{
		HTTPClient: client,
	}
}

/*
ListDeleteParams contains all the parameters to send to the API endpoint

	for the list delete operation.

	Typically these are written to a http.Request.
*/
type ListDeleteParams struct {

	/* ID.

	   ID of the list
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDeleteParams) WithDefaults() *ListDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list delete params
func (o *ListDeleteParams) WithTimeout(timeout time.Duration) *ListDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list delete params
func (o *ListDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list delete params
func (o *ListDeleteParams) WithContext(ctx context.Context) *ListDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list delete params
func (o *ListDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list delete params
func (o *ListDeleteParams) WithHTTPClient(client *http.Client) *ListDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list delete params
func (o *ListDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list delete params
func (o *ListDeleteParams) WithID(id string) *ListDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list delete params
func (o *ListDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
