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

// NewDomainPermissionSubscriptionRemoveParams creates a new DomainPermissionSubscriptionRemoveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainPermissionSubscriptionRemoveParams() *DomainPermissionSubscriptionRemoveParams {
	return &DomainPermissionSubscriptionRemoveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainPermissionSubscriptionRemoveParamsWithTimeout creates a new DomainPermissionSubscriptionRemoveParams object
// with the ability to set a timeout on a request.
func NewDomainPermissionSubscriptionRemoveParamsWithTimeout(timeout time.Duration) *DomainPermissionSubscriptionRemoveParams {
	return &DomainPermissionSubscriptionRemoveParams{
		timeout: timeout,
	}
}

// NewDomainPermissionSubscriptionRemoveParamsWithContext creates a new DomainPermissionSubscriptionRemoveParams object
// with the ability to set a context for a request.
func NewDomainPermissionSubscriptionRemoveParamsWithContext(ctx context.Context) *DomainPermissionSubscriptionRemoveParams {
	return &DomainPermissionSubscriptionRemoveParams{
		Context: ctx,
	}
}

// NewDomainPermissionSubscriptionRemoveParamsWithHTTPClient creates a new DomainPermissionSubscriptionRemoveParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainPermissionSubscriptionRemoveParamsWithHTTPClient(client *http.Client) *DomainPermissionSubscriptionRemoveParams {
	return &DomainPermissionSubscriptionRemoveParams{
		HTTPClient: client,
	}
}

/*
DomainPermissionSubscriptionRemoveParams contains all the parameters to send to the API endpoint

	for the domain permission subscription remove operation.

	Typically these are written to a http.Request.
*/
type DomainPermissionSubscriptionRemoveParams struct {

	/* ID.

	   ID of the domain permission subscription.
	*/
	ID string

	/* RemoveChildren.

	     When removing the domain permission subscription, also remove children of this subscription, ie., domain permissions that are managed by this subscription. If false, then children will instead be orphaned but not removed.
	Note that removed permissions may end up being created again later by another domain permission subscription of lower priority than the removed subscription. Likewise, orphaned children may be later adopted by another subscription.

	     Default: true
	*/
	RemoveChildren *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain permission subscription remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionSubscriptionRemoveParams) WithDefaults() *DomainPermissionSubscriptionRemoveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain permission subscription remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPermissionSubscriptionRemoveParams) SetDefaults() {
	var (
		removeChildrenDefault = bool(true)
	)

	val := DomainPermissionSubscriptionRemoveParams{
		RemoveChildren: &removeChildrenDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) WithTimeout(timeout time.Duration) *DomainPermissionSubscriptionRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) WithContext(ctx context.Context) *DomainPermissionSubscriptionRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) WithHTTPClient(client *http.Client) *DomainPermissionSubscriptionRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) WithID(id string) *DomainPermissionSubscriptionRemoveParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) SetID(id string) {
	o.ID = id
}

// WithRemoveChildren adds the removeChildren to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) WithRemoveChildren(removeChildren *bool) *DomainPermissionSubscriptionRemoveParams {
	o.SetRemoveChildren(removeChildren)
	return o
}

// SetRemoveChildren adds the removeChildren to the domain permission subscription remove params
func (o *DomainPermissionSubscriptionRemoveParams) SetRemoveChildren(removeChildren *bool) {
	o.RemoveChildren = removeChildren
}

// WriteToRequest writes these params to a swagger request
func (o *DomainPermissionSubscriptionRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.RemoveChildren != nil {

		// form param remove_children
		var frRemoveChildren bool
		if o.RemoveChildren != nil {
			frRemoveChildren = *o.RemoveChildren
		}
		fRemoveChildren := swag.FormatBool(frRemoveChildren)
		if fRemoveChildren != "" {
			if err := r.SetFormParam("remove_children", fRemoveChildren); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
