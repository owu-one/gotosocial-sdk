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

// NewAdminRuleGetParams creates a new AdminRuleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminRuleGetParams() *AdminRuleGetParams {
	return &AdminRuleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminRuleGetParamsWithTimeout creates a new AdminRuleGetParams object
// with the ability to set a timeout on a request.
func NewAdminRuleGetParamsWithTimeout(timeout time.Duration) *AdminRuleGetParams {
	return &AdminRuleGetParams{
		timeout: timeout,
	}
}

// NewAdminRuleGetParamsWithContext creates a new AdminRuleGetParams object
// with the ability to set a context for a request.
func NewAdminRuleGetParamsWithContext(ctx context.Context) *AdminRuleGetParams {
	return &AdminRuleGetParams{
		Context: ctx,
	}
}

// NewAdminRuleGetParamsWithHTTPClient creates a new AdminRuleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminRuleGetParamsWithHTTPClient(client *http.Client) *AdminRuleGetParams {
	return &AdminRuleGetParams{
		HTTPClient: client,
	}
}

/*
AdminRuleGetParams contains all the parameters to send to the API endpoint

	for the admin rule get operation.

	Typically these are written to a http.Request.
*/
type AdminRuleGetParams struct {

	/* ID.

	   The id of the rule.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRuleGetParams) WithDefaults() *AdminRuleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRuleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin rule get params
func (o *AdminRuleGetParams) WithTimeout(timeout time.Duration) *AdminRuleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin rule get params
func (o *AdminRuleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin rule get params
func (o *AdminRuleGetParams) WithContext(ctx context.Context) *AdminRuleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin rule get params
func (o *AdminRuleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin rule get params
func (o *AdminRuleGetParams) WithHTTPClient(client *http.Client) *AdminRuleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin rule get params
func (o *AdminRuleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the admin rule get params
func (o *AdminRuleGetParams) WithID(id string) *AdminRuleGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the admin rule get params
func (o *AdminRuleGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AdminRuleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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