// Code generated by go-swagger; DO NOT EDIT.

package interaction_requests

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

// NewAuthorizeInteractionRequestParams creates a new AuthorizeInteractionRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthorizeInteractionRequestParams() *AuthorizeInteractionRequestParams {
	return &AuthorizeInteractionRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthorizeInteractionRequestParamsWithTimeout creates a new AuthorizeInteractionRequestParams object
// with the ability to set a timeout on a request.
func NewAuthorizeInteractionRequestParamsWithTimeout(timeout time.Duration) *AuthorizeInteractionRequestParams {
	return &AuthorizeInteractionRequestParams{
		timeout: timeout,
	}
}

// NewAuthorizeInteractionRequestParamsWithContext creates a new AuthorizeInteractionRequestParams object
// with the ability to set a context for a request.
func NewAuthorizeInteractionRequestParamsWithContext(ctx context.Context) *AuthorizeInteractionRequestParams {
	return &AuthorizeInteractionRequestParams{
		Context: ctx,
	}
}

// NewAuthorizeInteractionRequestParamsWithHTTPClient creates a new AuthorizeInteractionRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthorizeInteractionRequestParamsWithHTTPClient(client *http.Client) *AuthorizeInteractionRequestParams {
	return &AuthorizeInteractionRequestParams{
		HTTPClient: client,
	}
}

/*
AuthorizeInteractionRequestParams contains all the parameters to send to the API endpoint

	for the authorize interaction request operation.

	Typically these are written to a http.Request.
*/
type AuthorizeInteractionRequestParams struct {

	/* ID.

	   ID of the interaction request targeting you.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authorize interaction request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizeInteractionRequestParams) WithDefaults() *AuthorizeInteractionRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authorize interaction request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthorizeInteractionRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) WithTimeout(timeout time.Duration) *AuthorizeInteractionRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) WithContext(ctx context.Context) *AuthorizeInteractionRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) WithHTTPClient(client *http.Client) *AuthorizeInteractionRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) WithID(id string) *AuthorizeInteractionRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the authorize interaction request params
func (o *AuthorizeInteractionRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AuthorizeInteractionRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
