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

// UnfollowTagReader is a Reader for the UnfollowTag structure.
type UnfollowTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnfollowTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnfollowTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnfollowTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUnfollowTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUnfollowTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnfollowTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnfollowTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/tags/{tag_name}/unfollow] unfollowTag", response, response.Code())
	}
}

// NewUnfollowTagOK creates a UnfollowTagOK with default headers values
func NewUnfollowTagOK() *UnfollowTagOK {
	return &UnfollowTagOK{}
}

/*
UnfollowTagOK describes a response with status code 200, with default header values.

Info about the tag.
*/
type UnfollowTagOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this unfollow tag o k response has a 2xx status code
func (o *UnfollowTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unfollow tag o k response has a 3xx status code
func (o *UnfollowTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unfollow tag o k response has a 4xx status code
func (o *UnfollowTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unfollow tag o k response has a 5xx status code
func (o *UnfollowTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unfollow tag o k response a status code equal to that given
func (o *UnfollowTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unfollow tag o k response
func (o *UnfollowTagOK) Code() int {
	return 200
}

func (o *UnfollowTagOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagOK %s", 200, payload)
}

func (o *UnfollowTagOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagOK %s", 200, payload)
}

func (o *UnfollowTagOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *UnfollowTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnfollowTagBadRequest creates a UnfollowTagBadRequest with default headers values
func NewUnfollowTagBadRequest() *UnfollowTagBadRequest {
	return &UnfollowTagBadRequest{}
}

/*
UnfollowTagBadRequest describes a response with status code 400, with default header values.

bad request
*/
type UnfollowTagBadRequest struct {
}

// IsSuccess returns true when this unfollow tag bad request response has a 2xx status code
func (o *UnfollowTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unfollow tag bad request response has a 3xx status code
func (o *UnfollowTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unfollow tag bad request response has a 4xx status code
func (o *UnfollowTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this unfollow tag bad request response has a 5xx status code
func (o *UnfollowTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this unfollow tag bad request response a status code equal to that given
func (o *UnfollowTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the unfollow tag bad request response
func (o *UnfollowTagBadRequest) Code() int {
	return 400
}

func (o *UnfollowTagBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagBadRequest", 400)
}

func (o *UnfollowTagBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagBadRequest", 400)
}

func (o *UnfollowTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnfollowTagUnauthorized creates a UnfollowTagUnauthorized with default headers values
func NewUnfollowTagUnauthorized() *UnfollowTagUnauthorized {
	return &UnfollowTagUnauthorized{}
}

/*
UnfollowTagUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type UnfollowTagUnauthorized struct {
}

// IsSuccess returns true when this unfollow tag unauthorized response has a 2xx status code
func (o *UnfollowTagUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unfollow tag unauthorized response has a 3xx status code
func (o *UnfollowTagUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unfollow tag unauthorized response has a 4xx status code
func (o *UnfollowTagUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this unfollow tag unauthorized response has a 5xx status code
func (o *UnfollowTagUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this unfollow tag unauthorized response a status code equal to that given
func (o *UnfollowTagUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the unfollow tag unauthorized response
func (o *UnfollowTagUnauthorized) Code() int {
	return 401
}

func (o *UnfollowTagUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagUnauthorized", 401)
}

func (o *UnfollowTagUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagUnauthorized", 401)
}

func (o *UnfollowTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnfollowTagForbidden creates a UnfollowTagForbidden with default headers values
func NewUnfollowTagForbidden() *UnfollowTagForbidden {
	return &UnfollowTagForbidden{}
}

/*
UnfollowTagForbidden describes a response with status code 403, with default header values.

forbidden
*/
type UnfollowTagForbidden struct {
}

// IsSuccess returns true when this unfollow tag forbidden response has a 2xx status code
func (o *UnfollowTagForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unfollow tag forbidden response has a 3xx status code
func (o *UnfollowTagForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unfollow tag forbidden response has a 4xx status code
func (o *UnfollowTagForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this unfollow tag forbidden response has a 5xx status code
func (o *UnfollowTagForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this unfollow tag forbidden response a status code equal to that given
func (o *UnfollowTagForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the unfollow tag forbidden response
func (o *UnfollowTagForbidden) Code() int {
	return 403
}

func (o *UnfollowTagForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagForbidden", 403)
}

func (o *UnfollowTagForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagForbidden", 403)
}

func (o *UnfollowTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnfollowTagNotFound creates a UnfollowTagNotFound with default headers values
func NewUnfollowTagNotFound() *UnfollowTagNotFound {
	return &UnfollowTagNotFound{}
}

/*
UnfollowTagNotFound describes a response with status code 404, with default header values.

unauthorized
*/
type UnfollowTagNotFound struct {
}

// IsSuccess returns true when this unfollow tag not found response has a 2xx status code
func (o *UnfollowTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unfollow tag not found response has a 3xx status code
func (o *UnfollowTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unfollow tag not found response has a 4xx status code
func (o *UnfollowTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unfollow tag not found response has a 5xx status code
func (o *UnfollowTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unfollow tag not found response a status code equal to that given
func (o *UnfollowTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the unfollow tag not found response
func (o *UnfollowTagNotFound) Code() int {
	return 404
}

func (o *UnfollowTagNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagNotFound", 404)
}

func (o *UnfollowTagNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagNotFound", 404)
}

func (o *UnfollowTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnfollowTagInternalServerError creates a UnfollowTagInternalServerError with default headers values
func NewUnfollowTagInternalServerError() *UnfollowTagInternalServerError {
	return &UnfollowTagInternalServerError{}
}

/*
UnfollowTagInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type UnfollowTagInternalServerError struct {
}

// IsSuccess returns true when this unfollow tag internal server error response has a 2xx status code
func (o *UnfollowTagInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unfollow tag internal server error response has a 3xx status code
func (o *UnfollowTagInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unfollow tag internal server error response has a 4xx status code
func (o *UnfollowTagInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this unfollow tag internal server error response has a 5xx status code
func (o *UnfollowTagInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this unfollow tag internal server error response a status code equal to that given
func (o *UnfollowTagInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the unfollow tag internal server error response
func (o *UnfollowTagInternalServerError) Code() int {
	return 500
}

func (o *UnfollowTagInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagInternalServerError", 500)
}

func (o *UnfollowTagInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/tags/{tag_name}/unfollow][%d] unfollowTagInternalServerError", 500)
}

func (o *UnfollowTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}