// Code generated by go-swagger; DO NOT EDIT.

package statuses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new statuses API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new statuses API client with basic auth credentials.
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

// New creates a new statuses API client with a bearer token for authentication.
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
Client for statuses API
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
	StatusBookmark(params *StatusBookmarkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusBookmarkOK, error)

	StatusBoostedBy(params *StatusBoostedByParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusBoostedByOK, error)

	StatusCreate(params *StatusCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusCreateOK, error)

	StatusDelete(params *StatusDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusDeleteOK, error)

	StatusEdit(params *StatusEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusEditOK, error)

	StatusFave(params *StatusFaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusFaveOK, error)

	StatusFavedBy(params *StatusFavedByParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusFavedByOK, error)

	StatusGet(params *StatusGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusGetOK, error)

	StatusHistoryGet(params *StatusHistoryGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusHistoryGetOK, error)

	StatusMute(params *StatusMuteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusMuteOK, error)

	StatusPin(params *StatusPinParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusPinOK, error)

	StatusReblog(params *StatusReblogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusReblogOK, error)

	StatusSourceGet(params *StatusSourceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusSourceGetOK, error)

	StatusUnbookmark(params *StatusUnbookmarkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnbookmarkOK, error)

	StatusUnfave(params *StatusUnfaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnfaveOK, error)

	StatusUnmute(params *StatusUnmuteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnmuteOK, error)

	StatusUnpin(params *StatusUnpinParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnpinOK, error)

	StatusUnreblog(params *StatusUnreblogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnreblogOK, error)

	ThreadContext(params *ThreadContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ThreadContextOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
StatusBookmark bookmarks status with the given ID
*/
func (a *Client) StatusBookmark(params *StatusBookmarkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusBookmarkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusBookmarkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusBookmark",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/bookmark",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusBookmarkReader{formats: a.formats},
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
	success, ok := result.(*StatusBookmarkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusBookmark: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusBoostedBy views accounts that have reblogged boosted the target status
*/
func (a *Client) StatusBoostedBy(params *StatusBoostedByParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusBoostedByOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusBoostedByParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusBoostedBy",
		Method:             "GET",
		PathPattern:        "/api/v1/statuses/{id}/reblogged_by",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusBoostedByReader{formats: a.formats},
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
	success, ok := result.(*StatusBoostedByOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusBoostedBy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StatusCreate creates a new status using the given form field parameters

	The parameters can also be given in the body of the request, as JSON, if the content-type is set to 'application/json'.

The 'interaction_policy' field can be used to set an interaction policy for this status.

If submitting using form data, use the following pattern to set an interaction policy:

`interaction_policy[INTERACTION_TYPE][CONDITION][INDEX]=Value`

For example: `interaction_policy[can_reply][always][0]=author`

Using `curl` this might look something like:

`curl -F 'interaction_policy[can_reply][always][0]=author' -F 'interaction_policy[can_reply][always][1]=followers' [... other form fields ...]`

The JSON equivalent would be:

`curl -H 'Content-Type: application/json' -d '{"interaction_policy":{"can_reply":{"always":["author","followers"]}} [... other json fields ...]}'`

The server will perform some normalization on the submitted policy so that you can't submit something totally invalid.
*/
func (a *Client) StatusCreate(params *StatusCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusCreate",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusCreateReader{formats: a.formats},
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
	success, ok := result.(*StatusCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StatusDelete deletes status with the given ID the status must belong to you

	The deleted status will be returned in the response. The `text` field will contain the original text of the status as it was submitted.

This is useful when doing a 'delete and redraft' type operation.
*/
func (a *Client) StatusDelete(params *StatusDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusDeleteReader{formats: a.formats},
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
	success, ok := result.(*StatusDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusEdit edits an existing status using the given form field parameters

The parameters can also be given in the body of the request, as JSON, if the content-type is set to 'application/json'.
*/
func (a *Client) StatusEdit(params *StatusEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusEditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusEditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusEdit",
		Method:             "PUT",
		PathPattern:        "/api/v1/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusEditReader{formats: a.formats},
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
	success, ok := result.(*StatusEditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusEdit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusFave stars like favourite the given status if permitted
*/
func (a *Client) StatusFave(params *StatusFaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusFaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusFaveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusFave",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/favourite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusFaveReader{formats: a.formats},
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
	success, ok := result.(*StatusFaveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusFave: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusFavedBy views accounts that have faved starred liked the target status
*/
func (a *Client) StatusFavedBy(params *StatusFavedByParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusFavedByOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusFavedByParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusFavedBy",
		Method:             "GET",
		PathPattern:        "/api/v1/statuses/{id}/favourited_by",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusFavedByReader{formats: a.formats},
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
	success, ok := result.(*StatusFavedByOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusFavedBy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusGet views status with the given ID
*/
func (a *Client) StatusGet(params *StatusGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusGet",
		Method:             "GET",
		PathPattern:        "/api/v1/statuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusGetReader{formats: a.formats},
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
	success, ok := result.(*StatusGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusHistoryGet views edit history of status with the given ID
*/
func (a *Client) StatusHistoryGet(params *StatusHistoryGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusHistoryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusHistoryGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusHistoryGet",
		Method:             "GET",
		PathPattern:        "/api/v1/statuses/{id}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusHistoryGetReader{formats: a.formats},
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
	success, ok := result.(*StatusHistoryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusHistoryGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StatusMute mutes a status s thread this prevents notifications from being created for future replies likes boosts etc in the thread of which the target status is a part

	Target status must belong to you or mention you.

Status thread mutes and unmutes are idempotent. If you already mute a thread, muting it again just means it stays muted and you'll get 200 OK back.
*/
func (a *Client) StatusMute(params *StatusMuteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusMuteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusMuteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusMute",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/mute",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusMuteReader{formats: a.formats},
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
	success, ok := result.(*StatusMuteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusMute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StatusPin pins a status to the top of your profile and add it to your featured activity pub collection

	You can only pin original posts (not reblogs) that you authored yourself.

Supported privacy levels for pinned posts are public, unlisted, and private/followers-only,
but only public posts will appear on the web version of your profile.
*/
func (a *Client) StatusPin(params *StatusPinParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusPinOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusPinParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusPin",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/pin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusPinReader{formats: a.formats},
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
	success, ok := result.(*StatusPinOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusPin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StatusReblog reblogs boost status with the given ID

	If the target status is rebloggable/boostable, it will be shared with your followers.

This is equivalent to an ActivityPub 'Announce' activity.
*/
func (a *Client) StatusReblog(params *StatusReblogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusReblogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusReblogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusReblog",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/reblog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusReblogReader{formats: a.formats},
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
	success, ok := result.(*StatusReblogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusReblog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusSourceGet views source text of status with the given ID requester must own the status
*/
func (a *Client) StatusSourceGet(params *StatusSourceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusSourceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusSourceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusSourceGet",
		Method:             "GET",
		PathPattern:        "/api/v1/statuses/{id}/source",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusSourceGetReader{formats: a.formats},
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
	success, ok := result.(*StatusSourceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusSourceGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusUnbookmark unbookmarks status with the given ID
*/
func (a *Client) StatusUnbookmark(params *StatusUnbookmarkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnbookmarkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusUnbookmarkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusUnbookmark",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/unbookmark",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusUnbookmarkReader{formats: a.formats},
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
	success, ok := result.(*StatusUnbookmarkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusUnbookmark: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusUnfave unstars unlike unfavourite the given status
*/
func (a *Client) StatusUnfave(params *StatusUnfaveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnfaveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusUnfaveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusUnfave",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/unfavourite",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusUnfaveReader{formats: a.formats},
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
	success, ok := result.(*StatusUnfaveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusUnfave: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	StatusUnmute unmutes a status s thread this reenables notifications for future replies likes boosts etc in the thread of which the target status is a part

	Target status must belong to you or mention you.

Status thread mutes and unmutes are idempotent. If you already unmuted a thread, unmuting it again just means it stays unmuted and you'll get 200 OK back.
*/
func (a *Client) StatusUnmute(params *StatusUnmuteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnmuteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusUnmuteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusUnmute",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/unmute",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusUnmuteReader{formats: a.formats},
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
	success, ok := result.(*StatusUnmuteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusUnmute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusUnpin unpins one of your pinned statuses
*/
func (a *Client) StatusUnpin(params *StatusUnpinParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnpinOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusUnpinParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusUnpin",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/unpin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusUnpinReader{formats: a.formats},
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
	success, ok := result.(*StatusUnpinOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusUnpin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StatusUnreblog unreblogs unboost status with the given ID
*/
func (a *Client) StatusUnreblog(params *StatusUnreblogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StatusUnreblogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusUnreblogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "statusUnreblog",
		Method:             "POST",
		PathPattern:        "/api/v1/statuses/{id}/unreblog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StatusUnreblogReader{formats: a.formats},
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
	success, ok := result.(*StatusUnreblogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for statusUnreblog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ThreadContext returns ancestors and descendants of the given status

The returned statuses will be ordered in a thread structure, so they are suitable to be displayed in the order in which they were returned.
*/
func (a *Client) ThreadContext(params *ThreadContextParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ThreadContextOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewThreadContextParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "threadContext",
		Method:             "GET",
		PathPattern:        "/api/v1/statuses/{id}/context",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ThreadContextReader{formats: a.formats},
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
	success, ok := result.(*ThreadContextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for threadContext: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
