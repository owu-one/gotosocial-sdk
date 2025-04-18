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

// NewDomainPermissionDraftGetParams creates a new DomainPermissionDraftGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainPermissionDraftGetParams() *DomainPermissionDraftGetParams {
	return &DomainPermissionDraftGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainPermissionDraftGetParamsWithTimeout creates a new DomainPermissionDraftGetParams object
// with the ability to set a timeout on a request.
func NewDomainPermissionDraftGetParamsWithTimeout(timeout time.Duration) *DomainPermissionDraftGetParams {
	return &DomainPermissionDraftGetParams{
		timeout: timeout,
	}
}

// NewDomainPermissionDraftGetParamsWithContext creates a new DomainPermissionDraftGetParams object
// with the ability to set a context for a request.
func NewDomainPermissionDraftGetParamsWithContext(ctx context.Context) *DomainPermissionDraftGetParams {
	return &DomainPermissionDraftGetParams{
		Context: ctx,
	}
}

// NewDomainPermissionDraftGetParamsWithHTTPClient creates a new DomainPermissionDraftGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainPermissionDraftGetParamsWithHTTPClient(client *http.Client) *DomainPermissionDraftGetParams {
	return &DomainPermissionDraftGetParams{
		HTTPClient: client,
	}
}

/*
DomainPermissionDraftGetParams contains all the parameters to send to the API endpoint

	for the domain permission draft get operation.

	Typically these are written to a http.Request.
*/
type DomainPermissionDraftGetParams struct {

	/* ID.

	   ID of the domain permission draft.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain permission draft get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionDraftGetParams) WithDefaults() *DomainPermissionDraftGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain permission draft get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionDraftGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) WithTimeout(timeout time.Duration) *DomainPermissionDraftGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) WithContext(ctx context.Context) *DomainPermissionDraftGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) WithHTTPClient(client *http.Client) *DomainPermissionDraftGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) WithID(id string) *DomainPermissionDraftGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domain permission draft get params
func (o *DomainPermissionDraftGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DomainPermissionDraftGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
