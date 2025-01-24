// Code generated by go-swagger; DO NOT EDIT.

package push

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PushSubscriptionDeleteReader is a Reader for the PushSubscriptionDelete structure.
type PushSubscriptionDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PushSubscriptionDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPushSubscriptionDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPushSubscriptionDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPushSubscriptionDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPushSubscriptionDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/push/subscription] pushSubscriptionDelete", response, response.Code())
	}
}

// NewPushSubscriptionDeleteOK creates a PushSubscriptionDeleteOK with default headers values
func NewPushSubscriptionDeleteOK() *PushSubscriptionDeleteOK {
	return &PushSubscriptionDeleteOK{}
}

/*
PushSubscriptionDeleteOK describes a response with status code 200, with default header values.

Push subscription deleted, or did not exist.
*/
type PushSubscriptionDeleteOK struct {
}

// IsSuccess returns true when this push subscription delete o k response has a 2xx status code
func (o *PushSubscriptionDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this push subscription delete o k response has a 3xx status code
func (o *PushSubscriptionDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this push subscription delete o k response has a 4xx status code
func (o *PushSubscriptionDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this push subscription delete o k response has a 5xx status code
func (o *PushSubscriptionDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this push subscription delete o k response a status code equal to that given
func (o *PushSubscriptionDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the push subscription delete o k response
func (o *PushSubscriptionDeleteOK) Code() int {
	return 200
}

func (o *PushSubscriptionDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteOK", 200)
}

func (o *PushSubscriptionDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteOK", 200)
}

func (o *PushSubscriptionDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPushSubscriptionDeleteBadRequest creates a PushSubscriptionDeleteBadRequest with default headers values
func NewPushSubscriptionDeleteBadRequest() *PushSubscriptionDeleteBadRequest {
	return &PushSubscriptionDeleteBadRequest{}
}

/*
PushSubscriptionDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PushSubscriptionDeleteBadRequest struct {
}

// IsSuccess returns true when this push subscription delete bad request response has a 2xx status code
func (o *PushSubscriptionDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this push subscription delete bad request response has a 3xx status code
func (o *PushSubscriptionDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this push subscription delete bad request response has a 4xx status code
func (o *PushSubscriptionDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this push subscription delete bad request response has a 5xx status code
func (o *PushSubscriptionDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this push subscription delete bad request response a status code equal to that given
func (o *PushSubscriptionDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the push subscription delete bad request response
func (o *PushSubscriptionDeleteBadRequest) Code() int {
	return 400
}

func (o *PushSubscriptionDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteBadRequest", 400)
}

func (o *PushSubscriptionDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteBadRequest", 400)
}

func (o *PushSubscriptionDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPushSubscriptionDeleteUnauthorized creates a PushSubscriptionDeleteUnauthorized with default headers values
func NewPushSubscriptionDeleteUnauthorized() *PushSubscriptionDeleteUnauthorized {
	return &PushSubscriptionDeleteUnauthorized{}
}

/*
PushSubscriptionDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type PushSubscriptionDeleteUnauthorized struct {
}

// IsSuccess returns true when this push subscription delete unauthorized response has a 2xx status code
func (o *PushSubscriptionDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this push subscription delete unauthorized response has a 3xx status code
func (o *PushSubscriptionDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this push subscription delete unauthorized response has a 4xx status code
func (o *PushSubscriptionDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this push subscription delete unauthorized response has a 5xx status code
func (o *PushSubscriptionDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this push subscription delete unauthorized response a status code equal to that given
func (o *PushSubscriptionDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the push subscription delete unauthorized response
func (o *PushSubscriptionDeleteUnauthorized) Code() int {
	return 401
}

func (o *PushSubscriptionDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteUnauthorized", 401)
}

func (o *PushSubscriptionDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteUnauthorized", 401)
}

func (o *PushSubscriptionDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPushSubscriptionDeleteInternalServerError creates a PushSubscriptionDeleteInternalServerError with default headers values
func NewPushSubscriptionDeleteInternalServerError() *PushSubscriptionDeleteInternalServerError {
	return &PushSubscriptionDeleteInternalServerError{}
}

/*
PushSubscriptionDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type PushSubscriptionDeleteInternalServerError struct {
}

// IsSuccess returns true when this push subscription delete internal server error response has a 2xx status code
func (o *PushSubscriptionDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this push subscription delete internal server error response has a 3xx status code
func (o *PushSubscriptionDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this push subscription delete internal server error response has a 4xx status code
func (o *PushSubscriptionDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this push subscription delete internal server error response has a 5xx status code
func (o *PushSubscriptionDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this push subscription delete internal server error response a status code equal to that given
func (o *PushSubscriptionDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the push subscription delete internal server error response
func (o *PushSubscriptionDeleteInternalServerError) Code() int {
	return 500
}

func (o *PushSubscriptionDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteInternalServerError", 500)
}

func (o *PushSubscriptionDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/push/subscription][%d] pushSubscriptionDeleteInternalServerError", 500)
}

func (o *PushSubscriptionDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}