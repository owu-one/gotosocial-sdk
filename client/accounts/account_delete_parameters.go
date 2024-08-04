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

// NewAccountDeleteParams creates a new AccountDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountDeleteParams() *AccountDeleteParams {
	return &AccountDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountDeleteParamsWithTimeout creates a new AccountDeleteParams object
// with the ability to set a timeout on a request.
func NewAccountDeleteParamsWithTimeout(timeout time.Duration) *AccountDeleteParams {
	return &AccountDeleteParams{
		timeout: timeout,
	}
}

// NewAccountDeleteParamsWithContext creates a new AccountDeleteParams object
// with the ability to set a context for a request.
func NewAccountDeleteParamsWithContext(ctx context.Context) *AccountDeleteParams {
	return &AccountDeleteParams{
		Context: ctx,
	}
}

// NewAccountDeleteParamsWithHTTPClient creates a new AccountDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountDeleteParamsWithHTTPClient(client *http.Client) *AccountDeleteParams {
	return &AccountDeleteParams{
		HTTPClient: client,
	}
}

/*
AccountDeleteParams contains all the parameters to send to the API endpoint

	for the account delete operation.

	Typically these are written to a http.Request.
*/
type AccountDeleteParams struct {

	/* Password.

	   Password of the account user, for confirmation.
	*/
	Password string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountDeleteParams) WithDefaults() *AccountDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account delete params
func (o *AccountDeleteParams) WithTimeout(timeout time.Duration) *AccountDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account delete params
func (o *AccountDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account delete params
func (o *AccountDeleteParams) WithContext(ctx context.Context) *AccountDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account delete params
func (o *AccountDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account delete params
func (o *AccountDeleteParams) WithHTTPClient(client *http.Client) *AccountDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account delete params
func (o *AccountDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPassword adds the password to the account delete params
func (o *AccountDeleteParams) WithPassword(password string) *AccountDeleteParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the account delete params
func (o *AccountDeleteParams) SetPassword(password string) {
	o.Password = password
}

// WriteToRequest writes these params to a swagger request
func (o *AccountDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param password
	frPassword := o.Password
	fPassword := frPassword
	if fPassword != "" {
		if err := r.SetFormParam("password", fPassword); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}