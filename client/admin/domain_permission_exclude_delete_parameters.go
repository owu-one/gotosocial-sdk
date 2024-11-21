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

// NewDomainPermissionExcludeDeleteParams creates a new DomainPermissionExcludeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainPermissionExcludeDeleteParams() *DomainPermissionExcludeDeleteParams {
	return &DomainPermissionExcludeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainPermissionExcludeDeleteParamsWithTimeout creates a new DomainPermissionExcludeDeleteParams object
// with the ability to set a timeout on a request.
func NewDomainPermissionExcludeDeleteParamsWithTimeout(timeout time.Duration) *DomainPermissionExcludeDeleteParams {
	return &DomainPermissionExcludeDeleteParams{
		timeout: timeout,
	}
}

// NewDomainPermissionExcludeDeleteParamsWithContext creates a new DomainPermissionExcludeDeleteParams object
// with the ability to set a context for a request.
func NewDomainPermissionExcludeDeleteParamsWithContext(ctx context.Context) *DomainPermissionExcludeDeleteParams {
	return &DomainPermissionExcludeDeleteParams{
		Context: ctx,
	}
}

// NewDomainPermissionExcludeDeleteParamsWithHTTPClient creates a new DomainPermissionExcludeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainPermissionExcludeDeleteParamsWithHTTPClient(client *http.Client) *DomainPermissionExcludeDeleteParams {
	return &DomainPermissionExcludeDeleteParams{
		HTTPClient: client,
	}
}

/*
DomainPermissionExcludeDeleteParams contains all the parameters to send to the API endpoint

	for the domain permission exclude delete operation.

	Typically these are written to a http.Request.
*/
type DomainPermissionExcludeDeleteParams struct {

	/* ID.

	   ID of the domain permission exclude.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain permission exclude delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionExcludeDeleteParams) WithDefaults() *DomainPermissionExcludeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain permission exclude delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionExcludeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) WithTimeout(timeout time.Duration) *DomainPermissionExcludeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) WithContext(ctx context.Context) *DomainPermissionExcludeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) WithHTTPClient(client *http.Client) *DomainPermissionExcludeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) WithID(id string) *DomainPermissionExcludeDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domain permission exclude delete params
func (o *DomainPermissionExcludeDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DomainPermissionExcludeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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