// Code generated by go-swagger; DO NOT EDIT.

package media

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

// NewMediaCreateParams creates a new MediaCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMediaCreateParams() *MediaCreateParams {
	return &MediaCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMediaCreateParamsWithTimeout creates a new MediaCreateParams object
// with the ability to set a timeout on a request.
func NewMediaCreateParamsWithTimeout(timeout time.Duration) *MediaCreateParams {
	return &MediaCreateParams{
		timeout: timeout,
	}
}

// NewMediaCreateParamsWithContext creates a new MediaCreateParams object
// with the ability to set a context for a request.
func NewMediaCreateParamsWithContext(ctx context.Context) *MediaCreateParams {
	return &MediaCreateParams{
		Context: ctx,
	}
}

// NewMediaCreateParamsWithHTTPClient creates a new MediaCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewMediaCreateParamsWithHTTPClient(client *http.Client) *MediaCreateParams {
	return &MediaCreateParams{
		HTTPClient: client,
	}
}

/*
MediaCreateParams contains all the parameters to send to the API endpoint

	for the media create operation.

	Typically these are written to a http.Request.
*/
type MediaCreateParams struct {

	/* APIVersion.

	   Version of the API to use. Must be either `v1` or `v2`.
	*/
	APIVersion string

	/* Description.

	   Image or media description to use as alt-text on the attachment. This is very useful for users of screenreaders! May or may not be required, depending on your instance settings.
	*/
	Description *string

	/* File.

	   The media attachment to upload.
	*/
	File runtime.NamedReadCloser

	/* Focus.

	   Focus of the media file. If present, it should be in the form of two comma-separated floats between -1 and 1. For example: `-0.5,0.25`.

	   Default: "0,0"
	*/
	Focus *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the media create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaCreateParams) WithDefaults() *MediaCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the media create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MediaCreateParams) SetDefaults() {
	var (
		focusDefault = string("0,0")
	)

	val := MediaCreateParams{
		Focus: &focusDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the media create params
func (o *MediaCreateParams) WithTimeout(timeout time.Duration) *MediaCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media create params
func (o *MediaCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media create params
func (o *MediaCreateParams) WithContext(ctx context.Context) *MediaCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media create params
func (o *MediaCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media create params
func (o *MediaCreateParams) WithHTTPClient(client *http.Client) *MediaCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media create params
func (o *MediaCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the media create params
func (o *MediaCreateParams) WithAPIVersion(aPIVersion string) *MediaCreateParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the media create params
func (o *MediaCreateParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDescription adds the description to the media create params
func (o *MediaCreateParams) WithDescription(description *string) *MediaCreateParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the media create params
func (o *MediaCreateParams) SetDescription(description *string) {
	o.Description = description
}

// WithFile adds the file to the media create params
func (o *MediaCreateParams) WithFile(file runtime.NamedReadCloser) *MediaCreateParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the media create params
func (o *MediaCreateParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithFocus adds the focus to the media create params
func (o *MediaCreateParams) WithFocus(focus *string) *MediaCreateParams {
	o.SetFocus(focus)
	return o
}

// SetFocus adds the focus to the media create params
func (o *MediaCreateParams) SetFocus(focus *string) {
	o.Focus = focus
}

// WriteToRequest writes these params to a swagger request
func (o *MediaCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api_version
	if err := r.SetPathParam("api_version", o.APIVersion); err != nil {
		return err
	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if o.Focus != nil {

		// form param focus
		var frFocus string
		if o.Focus != nil {
			frFocus = *o.Focus
		}
		fFocus := frFocus
		if fFocus != "" {
			if err := r.SetFormParam("focus", fFocus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
