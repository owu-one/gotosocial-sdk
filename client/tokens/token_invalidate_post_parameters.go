// Code generated by go-swagger; DO NOT EDIT.

package tokens

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

// NewTokenInvalidatePostParams creates a new TokenInvalidatePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTokenInvalidatePostParams() *TokenInvalidatePostParams {
	return &TokenInvalidatePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTokenInvalidatePostParamsWithTimeout creates a new TokenInvalidatePostParams object
// with the ability to set a timeout on a request.
func NewTokenInvalidatePostParamsWithTimeout(timeout time.Duration) *TokenInvalidatePostParams {
	return &TokenInvalidatePostParams{
		timeout: timeout,
	}
}

// NewTokenInvalidatePostParamsWithContext creates a new TokenInvalidatePostParams object
// with the ability to set a context for a request.
func NewTokenInvalidatePostParamsWithContext(ctx context.Context) *TokenInvalidatePostParams {
	return &TokenInvalidatePostParams{
		Context: ctx,
	}
}

// NewTokenInvalidatePostParamsWithHTTPClient creates a new TokenInvalidatePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewTokenInvalidatePostParamsWithHTTPClient(client *http.Client) *TokenInvalidatePostParams {
	return &TokenInvalidatePostParams{
		HTTPClient: client,
	}
}

/*
TokenInvalidatePostParams contains all the parameters to send to the API endpoint

	for the token invalidate post operation.

	Typically these are written to a http.Request.
*/
type TokenInvalidatePostParams struct {

	/* ID.

	   The id of the target token.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the token invalidate post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenInvalidatePostParams) WithDefaults() *TokenInvalidatePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the token invalidate post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenInvalidatePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the token invalidate post params
func (o *TokenInvalidatePostParams) WithTimeout(timeout time.Duration) *TokenInvalidatePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token invalidate post params
func (o *TokenInvalidatePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token invalidate post params
func (o *TokenInvalidatePostParams) WithContext(ctx context.Context) *TokenInvalidatePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token invalidate post params
func (o *TokenInvalidatePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token invalidate post params
func (o *TokenInvalidatePostParams) WithHTTPClient(client *http.Client) *TokenInvalidatePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token invalidate post params
func (o *TokenInvalidatePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the token invalidate post params
func (o *TokenInvalidatePostParams) WithID(id string) *TokenInvalidatePostParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the token invalidate post params
func (o *TokenInvalidatePostParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TokenInvalidatePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
