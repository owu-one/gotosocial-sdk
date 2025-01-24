// Code generated by go-swagger; DO NOT EDIT.

package push

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new push API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new push API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new push API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for push API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationxWwwFormUrlencoded sets the Content-Type header to "application/x-www-form-urlencoded".
func WithContentTypeApplicationxWwwFormUrlencoded(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	PushSubscriptionDelete(params *PushSubscriptionDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionDeleteOK, error)

	PushSubscriptionGet(params *PushSubscriptionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionGetOK, error)

	PushSubscriptionPost(params *PushSubscriptionPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionPostOK, error)

	PushSubscriptionPut(params *PushSubscriptionPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PushSubscriptionDelete deletes the web push subscription associated with the current auth token

If there is no subscription, returns successfully anyway.
*/
func (a *Client) PushSubscriptionDelete(params *PushSubscriptionDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPushSubscriptionDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pushSubscriptionDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/push/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PushSubscriptionDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushSubscriptionDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pushSubscriptionDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PushSubscriptionGet gets the push subscription for the current access token
*/
func (a *Client) PushSubscriptionGet(params *PushSubscriptionGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPushSubscriptionGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pushSubscriptionGet",
		Method:             "GET",
		PathPattern:        "/api/v1/push/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PushSubscriptionGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushSubscriptionGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pushSubscriptionGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PushSubscriptionPost creates a new web push subscription for the current access token or replace the existing one
*/
func (a *Client) PushSubscriptionPost(params *PushSubscriptionPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPushSubscriptionPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pushSubscriptionPost",
		Method:             "POST",
		PathPattern:        "/api/v1/push/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PushSubscriptionPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushSubscriptionPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pushSubscriptionPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PushSubscriptionPut updates the web push subscription for the current access token

Only which notifications you receive can be updated.
*/
func (a *Client) PushSubscriptionPut(params *PushSubscriptionPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PushSubscriptionPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPushSubscriptionPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pushSubscriptionPut",
		Method:             "PUT",
		PathPattern:        "/api/v1/push/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PushSubscriptionPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushSubscriptionPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pushSubscriptionPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}