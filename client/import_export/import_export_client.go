// Code generated by go-swagger; DO NOT EDIT.

package import_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new import export API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new import export API client with basic auth credentials.
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

// New creates a new import export API client with a bearer token for authentication.
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
Client for import export API
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

// WithContentTypeMultipartFormData sets the Content-Type header to "multipart/form-data".
func WithContentTypeMultipartFormData(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"multipart/form-data"}
}

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptTextCsv sets the Accept header to "text/csv".
func WithAcceptTextCsv(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/csv"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	ExportBlocks(params *ExportBlocksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportBlocksOK, error)

	ExportFollowers(params *ExportFollowersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportFollowersOK, error)

	ExportFollowing(params *ExportFollowingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportFollowingOK, error)

	ExportLists(params *ExportListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportListsOK, error)

	ExportMutes(params *ExportMutesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportMutesOK, error)

	ExportStats(params *ExportStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportStatsOK, error)

	ImportData(params *ImportDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImportDataAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ExportBlocks exports a c s v file of accounts that you block
*/
func (a *Client) ExportBlocks(params *ExportBlocksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportBlocksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportBlocksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportBlocks",
		Method:             "GET",
		PathPattern:        "/api/v1/exports/blocks.csv",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportBlocksReader{formats: a.formats},
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
	success, ok := result.(*ExportBlocksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportBlocks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportFollowers exports a c s v file of accounts that follow you
*/
func (a *Client) ExportFollowers(params *ExportFollowersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportFollowersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportFollowersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportFollowers",
		Method:             "GET",
		PathPattern:        "/api/v1/exports/followers.csv",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportFollowersReader{formats: a.formats},
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
	success, ok := result.(*ExportFollowersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportFollowers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportFollowing exports a c s v file of accounts that you follow
*/
func (a *Client) ExportFollowing(params *ExportFollowingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportFollowingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportFollowingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportFollowing",
		Method:             "GET",
		PathPattern:        "/api/v1/exports/following.csv",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportFollowingReader{formats: a.formats},
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
	success, ok := result.(*ExportFollowingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportFollowing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportLists exports a c s v file of lists created by you
*/
func (a *Client) ExportLists(params *ExportListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportListsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportLists",
		Method:             "GET",
		PathPattern:        "/api/v1/exports/lists.csv",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportListsReader{formats: a.formats},
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
	success, ok := result.(*ExportListsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportLists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportMutes exports a c s v file of accounts that you mute
*/
func (a *Client) ExportMutes(params *ExportMutesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportMutesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportMutesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportMutes",
		Method:             "GET",
		PathPattern:        "/api/v1/exports/mutes.csv",
		ProducesMediaTypes: []string{"text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportMutesReader{formats: a.formats},
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
	success, ok := result.(*ExportMutesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportMutes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportStats returns informational stats on the number of items that can be exported for requesting account
*/
func (a *Client) ExportStats(params *ExportStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExportStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportStats",
		Method:             "GET",
		PathPattern:        "/api/v1/exports/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportStatsReader{formats: a.formats},
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
	success, ok := result.(*ExportStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ImportData uploads some c s v formatted data to your account

	This can be used to migrate data from a Mastodon-compatible CSV file to a GoToSocial account.

Uploaded data will be processed asynchronously, and not all entries may be processed depending
on domain blocks, user-level blocks, network availability of referenced accounts and statuses, etc.
*/
func (a *Client) ImportData(params *ImportDataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImportDataAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportDataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "importData",
		Method:             "POST",
		PathPattern:        "/api/v1/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImportDataReader{formats: a.formats},
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
	success, ok := result.(*ImportDataAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importData: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
