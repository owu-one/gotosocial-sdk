// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new debug API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new debug API client with basic auth credentials.
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

// New creates a new debug API client with a bearer token for authentication.
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
Client for debug API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DebugAPURL(params *DebugAPURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DebugAPURLOK, error)

	DebugClearCaches(params *DebugClearCachesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DebugClearCachesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DebugAPURL performs a g e t to the specified activity pub URL and return detailed debugging information

Only enabled / exposed if GoToSocial was built and is running with flag DEBUG=1.
*/
func (a *Client) DebugAPURL(params *DebugAPURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DebugAPURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDebugAPURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "debugAPUrl",
		Method:             "GET",
		PathPattern:        "/api/v1/admin/debug/apurl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DebugAPURLReader{formats: a.formats},
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
	success, ok := result.(*DebugAPURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for debugAPUrl: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DebugClearCaches sweeps clear all in memory caches

Only enabled / exposed if GoToSocial was built and is running with flag DEBUG=1.
*/
func (a *Client) DebugClearCaches(params *DebugClearCachesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DebugClearCachesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDebugClearCachesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "debugClearCaches",
		Method:             "POST",
		PathPattern:        "/api/v1/admin/debug/caches/clear",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DebugClearCachesReader{formats: a.formats},
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
	success, ok := result.(*DebugClearCachesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for debugClearCaches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}