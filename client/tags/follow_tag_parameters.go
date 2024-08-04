// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewFollowTagParams creates a new FollowTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFollowTagParams() *FollowTagParams {
	return &FollowTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFollowTagParamsWithTimeout creates a new FollowTagParams object
// with the ability to set a timeout on a request.
func NewFollowTagParamsWithTimeout(timeout time.Duration) *FollowTagParams {
	return &FollowTagParams{
		timeout: timeout,
	}
}

// NewFollowTagParamsWithContext creates a new FollowTagParams object
// with the ability to set a context for a request.
func NewFollowTagParamsWithContext(ctx context.Context) *FollowTagParams {
	return &FollowTagParams{
		Context: ctx,
	}
}

// NewFollowTagParamsWithHTTPClient creates a new FollowTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewFollowTagParamsWithHTTPClient(client *http.Client) *FollowTagParams {
	return &FollowTagParams{
		HTTPClient: client,
	}
}

/*
FollowTagParams contains all the parameters to send to the API endpoint

	for the follow tag operation.

	Typically these are written to a http.Request.
*/
type FollowTagParams struct {

	/* TagName.

	   Name of the tag (no leading `#`)
	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the follow tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FollowTagParams) WithDefaults() *FollowTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the follow tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FollowTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the follow tag params
func (o *FollowTagParams) WithTimeout(timeout time.Duration) *FollowTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the follow tag params
func (o *FollowTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the follow tag params
func (o *FollowTagParams) WithContext(ctx context.Context) *FollowTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the follow tag params
func (o *FollowTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the follow tag params
func (o *FollowTagParams) WithHTTPClient(client *http.Client) *FollowTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the follow tag params
func (o *FollowTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagName adds the tagName to the follow tag params
func (o *FollowTagParams) WithTagName(tagName string) *FollowTagParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the follow tag params
func (o *FollowTagParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *FollowTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}