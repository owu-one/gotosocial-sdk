// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owu-one/gotosocial-sdk/models"
)

// GetFollowedTagsReader is a Reader for the GetFollowedTags structure.
type GetFollowedTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFollowedTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFollowedTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFollowedTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFollowedTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFollowedTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFollowedTagsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFollowedTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/followed_tags] getFollowedTags", response, response.Code())
	}
}

// NewGetFollowedTagsOK creates a GetFollowedTagsOK with default headers values
func NewGetFollowedTagsOK() *GetFollowedTagsOK {
	return &GetFollowedTagsOK{}
}

/*
GetFollowedTagsOK describes a response with status code 200, with default header values.

GetFollowedTagsOK get followed tags o k
*/
type GetFollowedTagsOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Tag
}

// IsSuccess returns true when this get followed tags o k response has a 2xx status code
func (o *GetFollowedTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get followed tags o k response has a 3xx status code
func (o *GetFollowedTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get followed tags o k response has a 4xx status code
func (o *GetFollowedTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get followed tags o k response has a 5xx status code
func (o *GetFollowedTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get followed tags o k response a status code equal to that given
func (o *GetFollowedTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get followed tags o k response
func (o *GetFollowedTagsOK) Code() int {
	return 200
}

func (o *GetFollowedTagsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsOK %s", 200, payload)
}

func (o *GetFollowedTagsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsOK %s", 200, payload)
}

func (o *GetFollowedTagsOK) GetPayload() []*models.Tag {
	return o.Payload
}

func (o *GetFollowedTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFollowedTagsBadRequest creates a GetFollowedTagsBadRequest with default headers values
func NewGetFollowedTagsBadRequest() *GetFollowedTagsBadRequest {
	return &GetFollowedTagsBadRequest{}
}

/*
GetFollowedTagsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetFollowedTagsBadRequest struct {
}

// IsSuccess returns true when this get followed tags bad request response has a 2xx status code
func (o *GetFollowedTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get followed tags bad request response has a 3xx status code
func (o *GetFollowedTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get followed tags bad request response has a 4xx status code
func (o *GetFollowedTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get followed tags bad request response has a 5xx status code
func (o *GetFollowedTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get followed tags bad request response a status code equal to that given
func (o *GetFollowedTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get followed tags bad request response
func (o *GetFollowedTagsBadRequest) Code() int {
	return 400
}

func (o *GetFollowedTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsBadRequest", 400)
}

func (o *GetFollowedTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsBadRequest", 400)
}

func (o *GetFollowedTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFollowedTagsUnauthorized creates a GetFollowedTagsUnauthorized with default headers values
func NewGetFollowedTagsUnauthorized() *GetFollowedTagsUnauthorized {
	return &GetFollowedTagsUnauthorized{}
}

/*
GetFollowedTagsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetFollowedTagsUnauthorized struct {
}

// IsSuccess returns true when this get followed tags unauthorized response has a 2xx status code
func (o *GetFollowedTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get followed tags unauthorized response has a 3xx status code
func (o *GetFollowedTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get followed tags unauthorized response has a 4xx status code
func (o *GetFollowedTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get followed tags unauthorized response has a 5xx status code
func (o *GetFollowedTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get followed tags unauthorized response a status code equal to that given
func (o *GetFollowedTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get followed tags unauthorized response
func (o *GetFollowedTagsUnauthorized) Code() int {
	return 401
}

func (o *GetFollowedTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsUnauthorized", 401)
}

func (o *GetFollowedTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsUnauthorized", 401)
}

func (o *GetFollowedTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFollowedTagsNotFound creates a GetFollowedTagsNotFound with default headers values
func NewGetFollowedTagsNotFound() *GetFollowedTagsNotFound {
	return &GetFollowedTagsNotFound{}
}

/*
GetFollowedTagsNotFound describes a response with status code 404, with default header values.

not found
*/
type GetFollowedTagsNotFound struct {
}

// IsSuccess returns true when this get followed tags not found response has a 2xx status code
func (o *GetFollowedTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get followed tags not found response has a 3xx status code
func (o *GetFollowedTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get followed tags not found response has a 4xx status code
func (o *GetFollowedTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get followed tags not found response has a 5xx status code
func (o *GetFollowedTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get followed tags not found response a status code equal to that given
func (o *GetFollowedTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get followed tags not found response
func (o *GetFollowedTagsNotFound) Code() int {
	return 404
}

func (o *GetFollowedTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsNotFound", 404)
}

func (o *GetFollowedTagsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsNotFound", 404)
}

func (o *GetFollowedTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFollowedTagsNotAcceptable creates a GetFollowedTagsNotAcceptable with default headers values
func NewGetFollowedTagsNotAcceptable() *GetFollowedTagsNotAcceptable {
	return &GetFollowedTagsNotAcceptable{}
}

/*
GetFollowedTagsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type GetFollowedTagsNotAcceptable struct {
}

// IsSuccess returns true when this get followed tags not acceptable response has a 2xx status code
func (o *GetFollowedTagsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get followed tags not acceptable response has a 3xx status code
func (o *GetFollowedTagsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get followed tags not acceptable response has a 4xx status code
func (o *GetFollowedTagsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this get followed tags not acceptable response has a 5xx status code
func (o *GetFollowedTagsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this get followed tags not acceptable response a status code equal to that given
func (o *GetFollowedTagsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the get followed tags not acceptable response
func (o *GetFollowedTagsNotAcceptable) Code() int {
	return 406
}

func (o *GetFollowedTagsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsNotAcceptable", 406)
}

func (o *GetFollowedTagsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsNotAcceptable", 406)
}

func (o *GetFollowedTagsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFollowedTagsInternalServerError creates a GetFollowedTagsInternalServerError with default headers values
func NewGetFollowedTagsInternalServerError() *GetFollowedTagsInternalServerError {
	return &GetFollowedTagsInternalServerError{}
}

/*
GetFollowedTagsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetFollowedTagsInternalServerError struct {
}

// IsSuccess returns true when this get followed tags internal server error response has a 2xx status code
func (o *GetFollowedTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get followed tags internal server error response has a 3xx status code
func (o *GetFollowedTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get followed tags internal server error response has a 4xx status code
func (o *GetFollowedTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get followed tags internal server error response has a 5xx status code
func (o *GetFollowedTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get followed tags internal server error response a status code equal to that given
func (o *GetFollowedTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get followed tags internal server error response
func (o *GetFollowedTagsInternalServerError) Code() int {
	return 500
}

func (o *GetFollowedTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsInternalServerError", 500)
}

func (o *GetFollowedTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/followed_tags][%d] getFollowedTagsInternalServerError", 500)
}

func (o *GetFollowedTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
