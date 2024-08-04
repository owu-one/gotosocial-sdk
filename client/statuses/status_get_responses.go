// Code generated by go-swagger; DO NOT EDIT.

package statuses

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

// StatusGetReader is a Reader for the StatusGet structure.
type StatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/statuses/{id}] statusGet", response, response.Code())
	}
}

// NewStatusGetOK creates a StatusGetOK with default headers values
func NewStatusGetOK() *StatusGetOK {
	return &StatusGetOK{}
}

/*
StatusGetOK describes a response with status code 200, with default header values.

The requested status.
*/
type StatusGetOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this status get o k response has a 2xx status code
func (o *StatusGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status get o k response has a 3xx status code
func (o *StatusGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get o k response has a 4xx status code
func (o *StatusGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status get o k response has a 5xx status code
func (o *StatusGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status get o k response a status code equal to that given
func (o *StatusGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status get o k response
func (o *StatusGetOK) Code() int {
	return 200
}

func (o *StatusGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetOK %s", 200, payload)
}

func (o *StatusGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetOK %s", 200, payload)
}

func (o *StatusGetOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusGetBadRequest creates a StatusGetBadRequest with default headers values
func NewStatusGetBadRequest() *StatusGetBadRequest {
	return &StatusGetBadRequest{}
}

/*
StatusGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusGetBadRequest struct {
}

// IsSuccess returns true when this status get bad request response has a 2xx status code
func (o *StatusGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status get bad request response has a 3xx status code
func (o *StatusGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get bad request response has a 4xx status code
func (o *StatusGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status get bad request response has a 5xx status code
func (o *StatusGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status get bad request response a status code equal to that given
func (o *StatusGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status get bad request response
func (o *StatusGetBadRequest) Code() int {
	return 400
}

func (o *StatusGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetBadRequest", 400)
}

func (o *StatusGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetBadRequest", 400)
}

func (o *StatusGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusGetUnauthorized creates a StatusGetUnauthorized with default headers values
func NewStatusGetUnauthorized() *StatusGetUnauthorized {
	return &StatusGetUnauthorized{}
}

/*
StatusGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusGetUnauthorized struct {
}

// IsSuccess returns true when this status get unauthorized response has a 2xx status code
func (o *StatusGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status get unauthorized response has a 3xx status code
func (o *StatusGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get unauthorized response has a 4xx status code
func (o *StatusGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status get unauthorized response has a 5xx status code
func (o *StatusGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status get unauthorized response a status code equal to that given
func (o *StatusGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status get unauthorized response
func (o *StatusGetUnauthorized) Code() int {
	return 401
}

func (o *StatusGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetUnauthorized", 401)
}

func (o *StatusGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetUnauthorized", 401)
}

func (o *StatusGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusGetForbidden creates a StatusGetForbidden with default headers values
func NewStatusGetForbidden() *StatusGetForbidden {
	return &StatusGetForbidden{}
}

/*
StatusGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusGetForbidden struct {
}

// IsSuccess returns true when this status get forbidden response has a 2xx status code
func (o *StatusGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status get forbidden response has a 3xx status code
func (o *StatusGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get forbidden response has a 4xx status code
func (o *StatusGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status get forbidden response has a 5xx status code
func (o *StatusGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status get forbidden response a status code equal to that given
func (o *StatusGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status get forbidden response
func (o *StatusGetForbidden) Code() int {
	return 403
}

func (o *StatusGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetForbidden", 403)
}

func (o *StatusGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetForbidden", 403)
}

func (o *StatusGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusGetNotFound creates a StatusGetNotFound with default headers values
func NewStatusGetNotFound() *StatusGetNotFound {
	return &StatusGetNotFound{}
}

/*
StatusGetNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusGetNotFound struct {
}

// IsSuccess returns true when this status get not found response has a 2xx status code
func (o *StatusGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status get not found response has a 3xx status code
func (o *StatusGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get not found response has a 4xx status code
func (o *StatusGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status get not found response has a 5xx status code
func (o *StatusGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status get not found response a status code equal to that given
func (o *StatusGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status get not found response
func (o *StatusGetNotFound) Code() int {
	return 404
}

func (o *StatusGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetNotFound", 404)
}

func (o *StatusGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetNotFound", 404)
}

func (o *StatusGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusGetNotAcceptable creates a StatusGetNotAcceptable with default headers values
func NewStatusGetNotAcceptable() *StatusGetNotAcceptable {
	return &StatusGetNotAcceptable{}
}

/*
StatusGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusGetNotAcceptable struct {
}

// IsSuccess returns true when this status get not acceptable response has a 2xx status code
func (o *StatusGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status get not acceptable response has a 3xx status code
func (o *StatusGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get not acceptable response has a 4xx status code
func (o *StatusGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status get not acceptable response has a 5xx status code
func (o *StatusGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status get not acceptable response a status code equal to that given
func (o *StatusGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status get not acceptable response
func (o *StatusGetNotAcceptable) Code() int {
	return 406
}

func (o *StatusGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetNotAcceptable", 406)
}

func (o *StatusGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetNotAcceptable", 406)
}

func (o *StatusGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusGetInternalServerError creates a StatusGetInternalServerError with default headers values
func NewStatusGetInternalServerError() *StatusGetInternalServerError {
	return &StatusGetInternalServerError{}
}

/*
StatusGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusGetInternalServerError struct {
}

// IsSuccess returns true when this status get internal server error response has a 2xx status code
func (o *StatusGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status get internal server error response has a 3xx status code
func (o *StatusGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status get internal server error response has a 4xx status code
func (o *StatusGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status get internal server error response has a 5xx status code
func (o *StatusGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status get internal server error response a status code equal to that given
func (o *StatusGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status get internal server error response
func (o *StatusGetInternalServerError) Code() int {
	return 500
}

func (o *StatusGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetInternalServerError", 500)
}

func (o *StatusGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/statuses/{id}][%d] statusGetInternalServerError", 500)
}

func (o *StatusGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
