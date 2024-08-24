// Code generated by go-swagger; DO NOT EDIT.

package interaction_requests

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

// AuthorizeInteractionRequestReader is a Reader for the AuthorizeInteractionRequest structure.
type AuthorizeInteractionRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizeInteractionRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizeInteractionRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthorizeInteractionRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthorizeInteractionRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthorizeInteractionRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAuthorizeInteractionRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthorizeInteractionRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/interaction_requests/{id}/authorize] authorizeInteractionRequest", response, response.Code())
	}
}

// NewAuthorizeInteractionRequestOK creates a AuthorizeInteractionRequestOK with default headers values
func NewAuthorizeInteractionRequestOK() *AuthorizeInteractionRequestOK {
	return &AuthorizeInteractionRequestOK{}
}

/*
AuthorizeInteractionRequestOK describes a response with status code 200, with default header values.

The now-approved interaction request.
*/
type AuthorizeInteractionRequestOK struct {
	Payload *models.InteractionRequest
}

// IsSuccess returns true when this authorize interaction request o k response has a 2xx status code
func (o *AuthorizeInteractionRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorize interaction request o k response has a 3xx status code
func (o *AuthorizeInteractionRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize interaction request o k response has a 4xx status code
func (o *AuthorizeInteractionRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorize interaction request o k response has a 5xx status code
func (o *AuthorizeInteractionRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize interaction request o k response a status code equal to that given
func (o *AuthorizeInteractionRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorize interaction request o k response
func (o *AuthorizeInteractionRequestOK) Code() int {
	return 200
}

func (o *AuthorizeInteractionRequestOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestOK %s", 200, payload)
}

func (o *AuthorizeInteractionRequestOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestOK %s", 200, payload)
}

func (o *AuthorizeInteractionRequestOK) GetPayload() *models.InteractionRequest {
	return o.Payload
}

func (o *AuthorizeInteractionRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InteractionRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthorizeInteractionRequestBadRequest creates a AuthorizeInteractionRequestBadRequest with default headers values
func NewAuthorizeInteractionRequestBadRequest() *AuthorizeInteractionRequestBadRequest {
	return &AuthorizeInteractionRequestBadRequest{}
}

/*
AuthorizeInteractionRequestBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AuthorizeInteractionRequestBadRequest struct {
}

// IsSuccess returns true when this authorize interaction request bad request response has a 2xx status code
func (o *AuthorizeInteractionRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize interaction request bad request response has a 3xx status code
func (o *AuthorizeInteractionRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize interaction request bad request response has a 4xx status code
func (o *AuthorizeInteractionRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize interaction request bad request response has a 5xx status code
func (o *AuthorizeInteractionRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize interaction request bad request response a status code equal to that given
func (o *AuthorizeInteractionRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the authorize interaction request bad request response
func (o *AuthorizeInteractionRequestBadRequest) Code() int {
	return 400
}

func (o *AuthorizeInteractionRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestBadRequest", 400)
}

func (o *AuthorizeInteractionRequestBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestBadRequest", 400)
}

func (o *AuthorizeInteractionRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeInteractionRequestUnauthorized creates a AuthorizeInteractionRequestUnauthorized with default headers values
func NewAuthorizeInteractionRequestUnauthorized() *AuthorizeInteractionRequestUnauthorized {
	return &AuthorizeInteractionRequestUnauthorized{}
}

/*
AuthorizeInteractionRequestUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AuthorizeInteractionRequestUnauthorized struct {
}

// IsSuccess returns true when this authorize interaction request unauthorized response has a 2xx status code
func (o *AuthorizeInteractionRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize interaction request unauthorized response has a 3xx status code
func (o *AuthorizeInteractionRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize interaction request unauthorized response has a 4xx status code
func (o *AuthorizeInteractionRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize interaction request unauthorized response has a 5xx status code
func (o *AuthorizeInteractionRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize interaction request unauthorized response a status code equal to that given
func (o *AuthorizeInteractionRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the authorize interaction request unauthorized response
func (o *AuthorizeInteractionRequestUnauthorized) Code() int {
	return 401
}

func (o *AuthorizeInteractionRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestUnauthorized", 401)
}

func (o *AuthorizeInteractionRequestUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestUnauthorized", 401)
}

func (o *AuthorizeInteractionRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeInteractionRequestNotFound creates a AuthorizeInteractionRequestNotFound with default headers values
func NewAuthorizeInteractionRequestNotFound() *AuthorizeInteractionRequestNotFound {
	return &AuthorizeInteractionRequestNotFound{}
}

/*
AuthorizeInteractionRequestNotFound describes a response with status code 404, with default header values.

not found
*/
type AuthorizeInteractionRequestNotFound struct {
}

// IsSuccess returns true when this authorize interaction request not found response has a 2xx status code
func (o *AuthorizeInteractionRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize interaction request not found response has a 3xx status code
func (o *AuthorizeInteractionRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize interaction request not found response has a 4xx status code
func (o *AuthorizeInteractionRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize interaction request not found response has a 5xx status code
func (o *AuthorizeInteractionRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize interaction request not found response a status code equal to that given
func (o *AuthorizeInteractionRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the authorize interaction request not found response
func (o *AuthorizeInteractionRequestNotFound) Code() int {
	return 404
}

func (o *AuthorizeInteractionRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestNotFound", 404)
}

func (o *AuthorizeInteractionRequestNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestNotFound", 404)
}

func (o *AuthorizeInteractionRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeInteractionRequestNotAcceptable creates a AuthorizeInteractionRequestNotAcceptable with default headers values
func NewAuthorizeInteractionRequestNotAcceptable() *AuthorizeInteractionRequestNotAcceptable {
	return &AuthorizeInteractionRequestNotAcceptable{}
}

/*
AuthorizeInteractionRequestNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AuthorizeInteractionRequestNotAcceptable struct {
}

// IsSuccess returns true when this authorize interaction request not acceptable response has a 2xx status code
func (o *AuthorizeInteractionRequestNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize interaction request not acceptable response has a 3xx status code
func (o *AuthorizeInteractionRequestNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize interaction request not acceptable response has a 4xx status code
func (o *AuthorizeInteractionRequestNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this authorize interaction request not acceptable response has a 5xx status code
func (o *AuthorizeInteractionRequestNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this authorize interaction request not acceptable response a status code equal to that given
func (o *AuthorizeInteractionRequestNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the authorize interaction request not acceptable response
func (o *AuthorizeInteractionRequestNotAcceptable) Code() int {
	return 406
}

func (o *AuthorizeInteractionRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestNotAcceptable", 406)
}

func (o *AuthorizeInteractionRequestNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestNotAcceptable", 406)
}

func (o *AuthorizeInteractionRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthorizeInteractionRequestInternalServerError creates a AuthorizeInteractionRequestInternalServerError with default headers values
func NewAuthorizeInteractionRequestInternalServerError() *AuthorizeInteractionRequestInternalServerError {
	return &AuthorizeInteractionRequestInternalServerError{}
}

/*
AuthorizeInteractionRequestInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AuthorizeInteractionRequestInternalServerError struct {
}

// IsSuccess returns true when this authorize interaction request internal server error response has a 2xx status code
func (o *AuthorizeInteractionRequestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this authorize interaction request internal server error response has a 3xx status code
func (o *AuthorizeInteractionRequestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorize interaction request internal server error response has a 4xx status code
func (o *AuthorizeInteractionRequestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorize interaction request internal server error response has a 5xx status code
func (o *AuthorizeInteractionRequestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this authorize interaction request internal server error response a status code equal to that given
func (o *AuthorizeInteractionRequestInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the authorize interaction request internal server error response
func (o *AuthorizeInteractionRequestInternalServerError) Code() int {
	return 500
}

func (o *AuthorizeInteractionRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestInternalServerError", 500)
}

func (o *AuthorizeInteractionRequestInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/interaction_requests/{id}/authorize][%d] authorizeInteractionRequestInternalServerError", 500)
}

func (o *AuthorizeInteractionRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
