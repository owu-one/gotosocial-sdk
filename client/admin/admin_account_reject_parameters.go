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
	"github.com/go-openapi/swag"
)

// NewAdminAccountRejectParams creates a new AdminAccountRejectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminAccountRejectParams() *AdminAccountRejectParams {
	return &AdminAccountRejectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminAccountRejectParamsWithTimeout creates a new AdminAccountRejectParams object
// with the ability to set a timeout on a request.
func NewAdminAccountRejectParamsWithTimeout(timeout time.Duration) *AdminAccountRejectParams {
	return &AdminAccountRejectParams{
		timeout: timeout,
	}
}

// NewAdminAccountRejectParamsWithContext creates a new AdminAccountRejectParams object
// with the ability to set a context for a request.
func NewAdminAccountRejectParamsWithContext(ctx context.Context) *AdminAccountRejectParams {
	return &AdminAccountRejectParams{
		Context: ctx,
	}
}

// NewAdminAccountRejectParamsWithHTTPClient creates a new AdminAccountRejectParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminAccountRejectParamsWithHTTPClient(client *http.Client) *AdminAccountRejectParams {
	return &AdminAccountRejectParams{
		HTTPClient: client,
	}
}

/*
AdminAccountRejectParams contains all the parameters to send to the API endpoint

	for the admin account reject operation.

	Typically these are written to a http.Request.
*/
type AdminAccountRejectParams struct {

	/* ID.

	   ID of the account.
	*/
	ID string

	/* Message.

	   Message to include in email to applicant. Will be included only if send_email is true.
	*/
	Message *string

	/* PrivateComment.

	   Comment to leave on why the account was denied. The comment will be visible to admins only.
	*/
	PrivateComment *string

	/* SendEmail.

	   Send an email to the applicant informing them that their sign-up has been rejected.
	*/
	SendEmail *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin account reject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminAccountRejectParams) WithDefaults() *AdminAccountRejectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin account reject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminAccountRejectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin account reject params
func (o *AdminAccountRejectParams) WithTimeout(timeout time.Duration) *AdminAccountRejectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin account reject params
func (o *AdminAccountRejectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin account reject params
func (o *AdminAccountRejectParams) WithContext(ctx context.Context) *AdminAccountRejectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin account reject params
func (o *AdminAccountRejectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin account reject params
func (o *AdminAccountRejectParams) WithHTTPClient(client *http.Client) *AdminAccountRejectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin account reject params
func (o *AdminAccountRejectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the admin account reject params
func (o *AdminAccountRejectParams) WithID(id string) *AdminAccountRejectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the admin account reject params
func (o *AdminAccountRejectParams) SetID(id string) {
	o.ID = id
}

// WithMessage adds the message to the admin account reject params
func (o *AdminAccountRejectParams) WithMessage(message *string) *AdminAccountRejectParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the admin account reject params
func (o *AdminAccountRejectParams) SetMessage(message *string) {
	o.Message = message
}

// WithPrivateComment adds the privateComment to the admin account reject params
func (o *AdminAccountRejectParams) WithPrivateComment(privateComment *string) *AdminAccountRejectParams {
	o.SetPrivateComment(privateComment)
	return o
}

// SetPrivateComment adds the privateComment to the admin account reject params
func (o *AdminAccountRejectParams) SetPrivateComment(privateComment *string) {
	o.PrivateComment = privateComment
}

// WithSendEmail adds the sendEmail to the admin account reject params
func (o *AdminAccountRejectParams) WithSendEmail(sendEmail *bool) *AdminAccountRejectParams {
	o.SetSendEmail(sendEmail)
	return o
}

// SetSendEmail adds the sendEmail to the admin account reject params
func (o *AdminAccountRejectParams) SetSendEmail(sendEmail *bool) {
	o.SendEmail = sendEmail
}

// WriteToRequest writes these params to a swagger request
func (o *AdminAccountRejectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Message != nil {

		// form param message
		var frMessage string
		if o.Message != nil {
			frMessage = *o.Message
		}
		fMessage := frMessage
		if fMessage != "" {
			if err := r.SetFormParam("message", fMessage); err != nil {
				return err
			}
		}
	}

	if o.PrivateComment != nil {

		// form param private_comment
		var frPrivateComment string
		if o.PrivateComment != nil {
			frPrivateComment = *o.PrivateComment
		}
		fPrivateComment := frPrivateComment
		if fPrivateComment != "" {
			if err := r.SetFormParam("private_comment", fPrivateComment); err != nil {
				return err
			}
		}
	}

	if o.SendEmail != nil {

		// form param send_email
		var frSendEmail bool
		if o.SendEmail != nil {
			frSendEmail = *o.SendEmail
		}
		fSendEmail := swag.FormatBool(frSendEmail)
		if fSendEmail != "" {
			if err := r.SetFormParam("send_email", fSendEmail); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
