// Code generated by go-swagger; DO NOT EDIT.

package filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new filters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new filters API client with basic auth credentials.
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

// New creates a new filters API client with a bearer token for authentication.
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
Client for filters API
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

// WithContentTypeApplicationXML sets the Content-Type header to "application/xml".
func WithContentTypeApplicationXML(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/xml"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	FilterKeywordDelete(params *FilterKeywordDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordDeleteOK, error)

	FilterKeywordGet(params *FilterKeywordGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordGetOK, error)

	FilterKeywordPost(params *FilterKeywordPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordPostOK, error)

	FilterKeywordPut(params *FilterKeywordPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordPutOK, error)

	FilterKeywordsGet(params *FilterKeywordsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordsGetOK, error)

	FilterStatusDelete(params *FilterStatusDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusDeleteOK, error)

	FilterStatusGet(params *FilterStatusGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusGetOK, error)

	FilterStatusPost(params *FilterStatusPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusPostOK, error)

	FilterStatusesGet(params *FilterStatusesGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusesGetOK, error)

	FilterV1Delete(params *FilterV1DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1DeleteOK, error)

	FilterV1Get(params *FilterV1GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1GetOK, error)

	FilterV1Post(params *FilterV1PostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1PostOK, error)

	FilterV1Put(params *FilterV1PutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1PutOK, error)

	FilterV2Delete(params *FilterV2DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2DeleteOK, error)

	FilterV2Get(params *FilterV2GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2GetOK, error)

	FilterV2Post(params *FilterV2PostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2PostOK, error)

	FilterV2Put(params *FilterV2PutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2PutOK, error)

	FiltersV1Get(params *FiltersV1GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FiltersV1GetOK, error)

	FiltersV2Get(params *FiltersV2GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FiltersV2GetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
FilterKeywordDelete deletes a single filter keyword with the given ID
*/
func (a *Client) FilterKeywordDelete(params *FilterKeywordDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterKeywordDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterKeywordDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v2/filters/keywords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterKeywordDeleteReader{formats: a.formats},
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
	success, ok := result.(*FilterKeywordDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterKeywordDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterKeywordGet gets a single filter keyword with the given ID
*/
func (a *Client) FilterKeywordGet(params *FilterKeywordGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterKeywordGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterKeywordGet",
		Method:             "GET",
		PathPattern:        "/api/v2/filters/keywords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterKeywordGetReader{formats: a.formats},
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
	success, ok := result.(*FilterKeywordGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterKeywordGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterKeywordPost adds a filter keyword to an existing filter
*/
func (a *Client) FilterKeywordPost(params *FilterKeywordPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterKeywordPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterKeywordPost",
		Method:             "POST",
		PathPattern:        "/api/v2/filters/{id}/keywords",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterKeywordPostReader{formats: a.formats},
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
	success, ok := result.(*FilterKeywordPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterKeywordPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterKeywordPut updates a single filter keyword with the given ID
*/
func (a *Client) FilterKeywordPut(params *FilterKeywordPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterKeywordPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterKeywordPut",
		Method:             "PUT",
		PathPattern:        "/api/v2/filters/keywords{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterKeywordPutReader{formats: a.formats},
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
	success, ok := result.(*FilterKeywordPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterKeywordPut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterKeywordsGet gets all filter keywords for a given filter
*/
func (a *Client) FilterKeywordsGet(params *FilterKeywordsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterKeywordsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterKeywordsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterKeywordsGet",
		Method:             "GET",
		PathPattern:        "/api/v2/filters/{id}/keywords",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterKeywordsGetReader{formats: a.formats},
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
	success, ok := result.(*FilterKeywordsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterKeywordsGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterStatusDelete deletes a single filter status with the given ID
*/
func (a *Client) FilterStatusDelete(params *FilterStatusDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterStatusDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterStatusDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v2/filters/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterStatusDeleteReader{formats: a.formats},
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
	success, ok := result.(*FilterStatusDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterStatusDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterStatusGet gets a single filter status with the given ID
*/
func (a *Client) FilterStatusGet(params *FilterStatusGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterStatusGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterStatusGet",
		Method:             "GET",
		PathPattern:        "/api/v2/filters/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterStatusGetReader{formats: a.formats},
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
	success, ok := result.(*FilterStatusGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterStatusGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterStatusPost adds a filter status to an existing filter
*/
func (a *Client) FilterStatusPost(params *FilterStatusPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterStatusPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterStatusPost",
		Method:             "POST",
		PathPattern:        "/api/v2/filters/{id}/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterStatusPostReader{formats: a.formats},
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
	success, ok := result.(*FilterStatusPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterStatusPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterStatusesGet gets all filter statuses for a given filter
*/
func (a *Client) FilterStatusesGet(params *FilterStatusesGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterStatusesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterStatusesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterStatusesGet",
		Method:             "GET",
		PathPattern:        "/api/v2/filters/{id}/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterStatusesGetReader{formats: a.formats},
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
	success, ok := result.(*FilterStatusesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterStatusesGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV1Delete deletes a single filter with the given ID
*/
func (a *Client) FilterV1Delete(params *FilterV1DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV1DeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV1Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/filters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV1DeleteReader{formats: a.formats},
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
	success, ok := result.(*FilterV1DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV1Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV1Get gets a single filter with the given ID
*/
func (a *Client) FilterV1Get(params *FilterV1GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV1GetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV1Get",
		Method:             "GET",
		PathPattern:        "/api/v1/filters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV1GetReader{formats: a.formats},
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
	success, ok := result.(*FilterV1GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV1Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV1Post creates a single filter
*/
func (a *Client) FilterV1Post(params *FilterV1PostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1PostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV1PostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV1Post",
		Method:             "POST",
		PathPattern:        "/api/v1/filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV1PostReader{formats: a.formats},
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
	success, ok := result.(*FilterV1PostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV1Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV1Put updates a single filter with the given ID
*/
func (a *Client) FilterV1Put(params *FilterV1PutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV1PutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV1PutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV1Put",
		Method:             "PUT",
		PathPattern:        "/api/v1/filters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV1PutReader{formats: a.formats},
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
	success, ok := result.(*FilterV1PutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV1Put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV2Delete deletes a single filter with the given ID
*/
func (a *Client) FilterV2Delete(params *FilterV2DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV2DeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV2Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v2/filters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV2DeleteReader{formats: a.formats},
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
	success, ok := result.(*FilterV2DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV2Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV2Get gets a single filter with the given ID
*/
func (a *Client) FilterV2Get(params *FilterV2GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV2GetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV2Get",
		Method:             "GET",
		PathPattern:        "/api/v2/filters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV2GetReader{formats: a.formats},
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
	success, ok := result.(*FilterV2GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV2Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FilterV2Post creates a single filter
*/
func (a *Client) FilterV2Post(params *FilterV2PostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2PostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV2PostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV2Post",
		Method:             "POST",
		PathPattern:        "/api/v2/filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV2PostReader{formats: a.formats},
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
	success, ok := result.(*FilterV2PostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV2Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	FilterV2Put updates a single filter with the given ID

	Note that this is actually closer to a PATCH operation:

only provided fields will be updated, and omitted fields will remain set to previous values.
*/
func (a *Client) FilterV2Put(params *FilterV2PutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FilterV2PutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFilterV2PutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filterV2Put",
		Method:             "PUT",
		PathPattern:        "/api/v2/filters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FilterV2PutReader{formats: a.formats},
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
	success, ok := result.(*FilterV2PutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filterV2Put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FiltersV1Get gets all filters for the authenticated account
*/
func (a *Client) FiltersV1Get(params *FiltersV1GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FiltersV1GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFiltersV1GetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filtersV1Get",
		Method:             "GET",
		PathPattern:        "/api/v1/filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FiltersV1GetReader{formats: a.formats},
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
	success, ok := result.(*FiltersV1GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filtersV1Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FiltersV2Get gets all filters for the authenticated account
*/
func (a *Client) FiltersV2Get(params *FiltersV2GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FiltersV2GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFiltersV2GetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "filtersV2Get",
		Method:             "GET",
		PathPattern:        "/api/v2/filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FiltersV2GetReader{formats: a.formats},
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
	success, ok := result.(*FiltersV2GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for filtersV2Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
