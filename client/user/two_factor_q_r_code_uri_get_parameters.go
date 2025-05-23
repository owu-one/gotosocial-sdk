// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewTwoFactorQRCodeURIGetParams creates a new TwoFactorQRCodeURIGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTwoFactorQRCodeURIGetParams() *TwoFactorQRCodeURIGetParams {
	return &TwoFactorQRCodeURIGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTwoFactorQRCodeURIGetParamsWithTimeout creates a new TwoFactorQRCodeURIGetParams object
// with the ability to set a timeout on a request.
func NewTwoFactorQRCodeURIGetParamsWithTimeout(timeout time.Duration) *TwoFactorQRCodeURIGetParams {
	return &TwoFactorQRCodeURIGetParams{
		timeout: timeout,
	}
}

// NewTwoFactorQRCodeURIGetParamsWithContext creates a new TwoFactorQRCodeURIGetParams object
// with the ability to set a context for a request.
func NewTwoFactorQRCodeURIGetParamsWithContext(ctx context.Context) *TwoFactorQRCodeURIGetParams {
	return &TwoFactorQRCodeURIGetParams{
		Context: ctx,
	}
}

// NewTwoFactorQRCodeURIGetParamsWithHTTPClient creates a new TwoFactorQRCodeURIGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTwoFactorQRCodeURIGetParamsWithHTTPClient(client *http.Client) *TwoFactorQRCodeURIGetParams {
	return &TwoFactorQRCodeURIGetParams{
		HTTPClient: client,
	}
}

/*
TwoFactorQRCodeURIGetParams contains all the parameters to send to the API endpoint

	for the two factor q r code URI get operation.

	Typically these are written to a http.Request.
*/
type TwoFactorQRCodeURIGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the two factor q r code URI get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TwoFactorQRCodeURIGetParams) WithDefaults() *TwoFactorQRCodeURIGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the two factor q r code URI get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TwoFactorQRCodeURIGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the two factor q r code URI get params
func (o *TwoFactorQRCodeURIGetParams) WithTimeout(timeout time.Duration) *TwoFactorQRCodeURIGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the two factor q r code URI get params
func (o *TwoFactorQRCodeURIGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the two factor q r code URI get params
func (o *TwoFactorQRCodeURIGetParams) WithContext(ctx context.Context) *TwoFactorQRCodeURIGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the two factor q r code URI get params
func (o *TwoFactorQRCodeURIGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the two factor q r code URI get params
func (o *TwoFactorQRCodeURIGetParams) WithHTTPClient(client *http.Client) *TwoFactorQRCodeURIGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the two factor q r code URI get params
func (o *TwoFactorQRCodeURIGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TwoFactorQRCodeURIGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
