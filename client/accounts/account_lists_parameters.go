// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewAccountListsParams creates a new AccountListsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountListsParams() *AccountListsParams {
	return &AccountListsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountListsParamsWithTimeout creates a new AccountListsParams object
// with the ability to set a timeout on a request.
func NewAccountListsParamsWithTimeout(timeout time.Duration) *AccountListsParams {
	return &AccountListsParams{
		timeout: timeout,
	}
}

// NewAccountListsParamsWithContext creates a new AccountListsParams object
// with the ability to set a context for a request.
func NewAccountListsParamsWithContext(ctx context.Context) *AccountListsParams {
	return &AccountListsParams{
		Context: ctx,
	}
}

// NewAccountListsParamsWithHTTPClient creates a new AccountListsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountListsParamsWithHTTPClient(client *http.Client) *AccountListsParams {
	return &AccountListsParams{
		HTTPClient: client,
	}
}

/*
AccountListsParams contains all the parameters to send to the API endpoint

	for the account lists operation.

	Typically these are written to a http.Request.
*/
type AccountListsParams struct {

	/* ID.

	   Account ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountListsParams) WithDefaults() *AccountListsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountListsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account lists params
func (o *AccountListsParams) WithTimeout(timeout time.Duration) *AccountListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account lists params
func (o *AccountListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account lists params
func (o *AccountListsParams) WithContext(ctx context.Context) *AccountListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account lists params
func (o *AccountListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account lists params
func (o *AccountListsParams) WithHTTPClient(client *http.Client) *AccountListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account lists params
func (o *AccountListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the account lists params
func (o *AccountListsParams) WithID(id string) *AccountListsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the account lists params
func (o *AccountListsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AccountListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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