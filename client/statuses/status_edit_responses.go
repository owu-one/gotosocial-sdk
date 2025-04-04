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

// StatusEditReader is a Reader for the StatusEdit structure.
type StatusEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStatusEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStatusEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStatusEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStatusEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewStatusEditNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStatusEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v1/statuses/{id}] statusEdit", response, response.Code())
	}
}

// NewStatusEditOK creates a StatusEditOK with default headers values
func NewStatusEditOK() *StatusEditOK {
	return &StatusEditOK{}
}

/*
StatusEditOK describes a response with status code 200, with default header values.

The latest status revision.
*/
type StatusEditOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this status edit o k response has a 2xx status code
func (o *StatusEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status edit o k response has a 3xx status code
func (o *StatusEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit o k response has a 4xx status code
func (o *StatusEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status edit o k response has a 5xx status code
func (o *StatusEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status edit o k response a status code equal to that given
func (o *StatusEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status edit o k response
func (o *StatusEditOK) Code() int {
	return 200
}

func (o *StatusEditOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditOK %s", 200, payload)
}

func (o *StatusEditOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditOK %s", 200, payload)
}

func (o *StatusEditOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StatusEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusEditBadRequest creates a StatusEditBadRequest with default headers values
func NewStatusEditBadRequest() *StatusEditBadRequest {
	return &StatusEditBadRequest{}
}

/*
StatusEditBadRequest describes a response with status code 400, with default header values.

bad request
*/
type StatusEditBadRequest struct {
}

// IsSuccess returns true when this status edit bad request response has a 2xx status code
func (o *StatusEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status edit bad request response has a 3xx status code
func (o *StatusEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit bad request response has a 4xx status code
func (o *StatusEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this status edit bad request response has a 5xx status code
func (o *StatusEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this status edit bad request response a status code equal to that given
func (o *StatusEditBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the status edit bad request response
func (o *StatusEditBadRequest) Code() int {
	return 400
}

func (o *StatusEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditBadRequest", 400)
}

func (o *StatusEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditBadRequest", 400)
}

func (o *StatusEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusEditUnauthorized creates a StatusEditUnauthorized with default headers values
func NewStatusEditUnauthorized() *StatusEditUnauthorized {
	return &StatusEditUnauthorized{}
}

/*
StatusEditUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type StatusEditUnauthorized struct {
}

// IsSuccess returns true when this status edit unauthorized response has a 2xx status code
func (o *StatusEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status edit unauthorized response has a 3xx status code
func (o *StatusEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit unauthorized response has a 4xx status code
func (o *StatusEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this status edit unauthorized response has a 5xx status code
func (o *StatusEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this status edit unauthorized response a status code equal to that given
func (o *StatusEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the status edit unauthorized response
func (o *StatusEditUnauthorized) Code() int {
	return 401
}

func (o *StatusEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditUnauthorized", 401)
}

func (o *StatusEditUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditUnauthorized", 401)
}

func (o *StatusEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusEditForbidden creates a StatusEditForbidden with default headers values
func NewStatusEditForbidden() *StatusEditForbidden {
	return &StatusEditForbidden{}
}

/*
StatusEditForbidden describes a response with status code 403, with default header values.

forbidden
*/
type StatusEditForbidden struct {
}

// IsSuccess returns true when this status edit forbidden response has a 2xx status code
func (o *StatusEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status edit forbidden response has a 3xx status code
func (o *StatusEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit forbidden response has a 4xx status code
func (o *StatusEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this status edit forbidden response has a 5xx status code
func (o *StatusEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this status edit forbidden response a status code equal to that given
func (o *StatusEditForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the status edit forbidden response
func (o *StatusEditForbidden) Code() int {
	return 403
}

func (o *StatusEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditForbidden", 403)
}

func (o *StatusEditForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditForbidden", 403)
}

func (o *StatusEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusEditNotFound creates a StatusEditNotFound with default headers values
func NewStatusEditNotFound() *StatusEditNotFound {
	return &StatusEditNotFound{}
}

/*
StatusEditNotFound describes a response with status code 404, with default header values.

not found
*/
type StatusEditNotFound struct {
}

// IsSuccess returns true when this status edit not found response has a 2xx status code
func (o *StatusEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status edit not found response has a 3xx status code
func (o *StatusEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit not found response has a 4xx status code
func (o *StatusEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this status edit not found response has a 5xx status code
func (o *StatusEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this status edit not found response a status code equal to that given
func (o *StatusEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the status edit not found response
func (o *StatusEditNotFound) Code() int {
	return 404
}

func (o *StatusEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditNotFound", 404)
}

func (o *StatusEditNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditNotFound", 404)
}

func (o *StatusEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusEditNotAcceptable creates a StatusEditNotAcceptable with default headers values
func NewStatusEditNotAcceptable() *StatusEditNotAcceptable {
	return &StatusEditNotAcceptable{}
}

/*
StatusEditNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type StatusEditNotAcceptable struct {
}

// IsSuccess returns true when this status edit not acceptable response has a 2xx status code
func (o *StatusEditNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status edit not acceptable response has a 3xx status code
func (o *StatusEditNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit not acceptable response has a 4xx status code
func (o *StatusEditNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this status edit not acceptable response has a 5xx status code
func (o *StatusEditNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this status edit not acceptable response a status code equal to that given
func (o *StatusEditNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the status edit not acceptable response
func (o *StatusEditNotAcceptable) Code() int {
	return 406
}

func (o *StatusEditNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditNotAcceptable", 406)
}

func (o *StatusEditNotAcceptable) String() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditNotAcceptable", 406)
}

func (o *StatusEditNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStatusEditInternalServerError creates a StatusEditInternalServerError with default headers values
func NewStatusEditInternalServerError() *StatusEditInternalServerError {
	return &StatusEditInternalServerError{}
}

/*
StatusEditInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type StatusEditInternalServerError struct {
}

// IsSuccess returns true when this status edit internal server error response has a 2xx status code
func (o *StatusEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this status edit internal server error response has a 3xx status code
func (o *StatusEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status edit internal server error response has a 4xx status code
func (o *StatusEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this status edit internal server error response has a 5xx status code
func (o *StatusEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this status edit internal server error response a status code equal to that given
func (o *StatusEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the status edit internal server error response
func (o *StatusEditInternalServerError) Code() int {
	return 500
}

func (o *StatusEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditInternalServerError", 500)
}

func (o *StatusEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/statuses/{id}][%d] statusEditInternalServerError", 500)
}

func (o *StatusEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
