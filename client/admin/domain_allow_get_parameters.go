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

// NewDomainAllowGetParams creates a new DomainAllowGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainAllowGetParams() *DomainAllowGetParams {
	return &DomainAllowGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainAllowGetParamsWithTimeout creates a new DomainAllowGetParams object
// with the ability to set a timeout on a request.
func NewDomainAllowGetParamsWithTimeout(timeout time.Duration) *DomainAllowGetParams {
	return &DomainAllowGetParams{
		timeout: timeout,
	}
}

// NewDomainAllowGetParamsWithContext creates a new DomainAllowGetParams object
// with the ability to set a context for a request.
func NewDomainAllowGetParamsWithContext(ctx context.Context) *DomainAllowGetParams {
	return &DomainAllowGetParams{
		Context: ctx,
	}
}

// NewDomainAllowGetParamsWithHTTPClient creates a new DomainAllowGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainAllowGetParamsWithHTTPClient(client *http.Client) *DomainAllowGetParams {
	return &DomainAllowGetParams{
		HTTPClient: client,
	}
}

/*
DomainAllowGetParams contains all the parameters to send to the API endpoint

	for the domain allow get operation.

	Typically these are written to a http.Request.
*/
type DomainAllowGetParams struct {

	/* ID.

	   The id of the domain allow.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain allow get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainAllowGetParams) WithDefaults() *DomainAllowGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain allow get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainAllowGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domain allow get params
func (o *DomainAllowGetParams) WithTimeout(timeout time.Duration) *DomainAllowGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain allow get params
func (o *DomainAllowGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain allow get params
func (o *DomainAllowGetParams) WithContext(ctx context.Context) *DomainAllowGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain allow get params
func (o *DomainAllowGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain allow get params
func (o *DomainAllowGetParams) WithHTTPClient(client *http.Client) *DomainAllowGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain allow get params
func (o *DomainAllowGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the domain allow get params
func (o *DomainAllowGetParams) WithID(id string) *DomainAllowGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domain allow get params
func (o *DomainAllowGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DomainAllowGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
