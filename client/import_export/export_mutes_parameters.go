// Code generated by go-swagger; DO NOT EDIT.

package import_export

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

// NewExportMutesParams creates a new ExportMutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportMutesParams() *ExportMutesParams {
	return &ExportMutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportMutesParamsWithTimeout creates a new ExportMutesParams object
// with the ability to set a timeout on a request.
func NewExportMutesParamsWithTimeout(timeout time.Duration) *ExportMutesParams {
	return &ExportMutesParams{
		timeout: timeout,
	}
}

// NewExportMutesParamsWithContext creates a new ExportMutesParams object
// with the ability to set a context for a request.
func NewExportMutesParamsWithContext(ctx context.Context) *ExportMutesParams {
	return &ExportMutesParams{
		Context: ctx,
	}
}

// NewExportMutesParamsWithHTTPClient creates a new ExportMutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportMutesParamsWithHTTPClient(client *http.Client) *ExportMutesParams {
	return &ExportMutesParams{
		HTTPClient: client,
	}
}

/*
ExportMutesParams contains all the parameters to send to the API endpoint

	for the export mutes operation.

	Typically these are written to a http.Request.
*/
type ExportMutesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export mutes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportMutesParams) WithDefaults() *ExportMutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export mutes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportMutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export mutes params
func (o *ExportMutesParams) WithTimeout(timeout time.Duration) *ExportMutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export mutes params
func (o *ExportMutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export mutes params
func (o *ExportMutesParams) WithContext(ctx context.Context) *ExportMutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export mutes params
func (o *ExportMutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export mutes params
func (o *ExportMutesParams) WithHTTPClient(client *http.Client) *ExportMutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export mutes params
func (o *ExportMutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExportMutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
