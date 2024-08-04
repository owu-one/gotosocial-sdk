// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// AccountStatusesReader is a Reader for the AccountStatuses structure.
type AccountStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountStatusesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountStatusesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountStatusesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountStatusesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountStatusesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/accounts/{id}/statuses] accountStatuses", response, response.Code())
	}
}

// NewAccountStatusesOK creates a AccountStatusesOK with default headers values
func NewAccountStatusesOK() *AccountStatusesOK {
	return &AccountStatusesOK{}
}

/*
AccountStatusesOK describes a response with status code 200, with default header values.

Array of statuses.
*/
type AccountStatusesOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Status
}

// IsSuccess returns true when this account statuses o k response has a 2xx status code
func (o *AccountStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account statuses o k response has a 3xx status code
func (o *AccountStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account statuses o k response has a 4xx status code
func (o *AccountStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account statuses o k response has a 5xx status code
func (o *AccountStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account statuses o k response a status code equal to that given
func (o *AccountStatusesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account statuses o k response
func (o *AccountStatusesOK) Code() int {
	return 200
}

func (o *AccountStatusesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesOK %s", 200, payload)
}

func (o *AccountStatusesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesOK %s", 200, payload)
}

func (o *AccountStatusesOK) GetPayload() []*models.Status {
	return o.Payload
}

func (o *AccountStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAccountStatusesBadRequest creates a AccountStatusesBadRequest with default headers values
func NewAccountStatusesBadRequest() *AccountStatusesBadRequest {
	return &AccountStatusesBadRequest{}
}

/*
AccountStatusesBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountStatusesBadRequest struct {
}

// IsSuccess returns true when this account statuses bad request response has a 2xx status code
func (o *AccountStatusesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account statuses bad request response has a 3xx status code
func (o *AccountStatusesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account statuses bad request response has a 4xx status code
func (o *AccountStatusesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this account statuses bad request response has a 5xx status code
func (o *AccountStatusesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this account statuses bad request response a status code equal to that given
func (o *AccountStatusesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the account statuses bad request response
func (o *AccountStatusesBadRequest) Code() int {
	return 400
}

func (o *AccountStatusesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesBadRequest", 400)
}

func (o *AccountStatusesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesBadRequest", 400)
}

func (o *AccountStatusesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountStatusesUnauthorized creates a AccountStatusesUnauthorized with default headers values
func NewAccountStatusesUnauthorized() *AccountStatusesUnauthorized {
	return &AccountStatusesUnauthorized{}
}

/*
AccountStatusesUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountStatusesUnauthorized struct {
}

// IsSuccess returns true when this account statuses unauthorized response has a 2xx status code
func (o *AccountStatusesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account statuses unauthorized response has a 3xx status code
func (o *AccountStatusesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account statuses unauthorized response has a 4xx status code
func (o *AccountStatusesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this account statuses unauthorized response has a 5xx status code
func (o *AccountStatusesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this account statuses unauthorized response a status code equal to that given
func (o *AccountStatusesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the account statuses unauthorized response
func (o *AccountStatusesUnauthorized) Code() int {
	return 401
}

func (o *AccountStatusesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesUnauthorized", 401)
}

func (o *AccountStatusesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesUnauthorized", 401)
}

func (o *AccountStatusesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountStatusesNotFound creates a AccountStatusesNotFound with default headers values
func NewAccountStatusesNotFound() *AccountStatusesNotFound {
	return &AccountStatusesNotFound{}
}

/*
AccountStatusesNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountStatusesNotFound struct {
}

// IsSuccess returns true when this account statuses not found response has a 2xx status code
func (o *AccountStatusesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account statuses not found response has a 3xx status code
func (o *AccountStatusesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account statuses not found response has a 4xx status code
func (o *AccountStatusesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this account statuses not found response has a 5xx status code
func (o *AccountStatusesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this account statuses not found response a status code equal to that given
func (o *AccountStatusesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the account statuses not found response
func (o *AccountStatusesNotFound) Code() int {
	return 404
}

func (o *AccountStatusesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesNotFound", 404)
}

func (o *AccountStatusesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesNotFound", 404)
}

func (o *AccountStatusesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountStatusesNotAcceptable creates a AccountStatusesNotAcceptable with default headers values
func NewAccountStatusesNotAcceptable() *AccountStatusesNotAcceptable {
	return &AccountStatusesNotAcceptable{}
}

/*
AccountStatusesNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountStatusesNotAcceptable struct {
}

// IsSuccess returns true when this account statuses not acceptable response has a 2xx status code
func (o *AccountStatusesNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account statuses not acceptable response has a 3xx status code
func (o *AccountStatusesNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account statuses not acceptable response has a 4xx status code
func (o *AccountStatusesNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this account statuses not acceptable response has a 5xx status code
func (o *AccountStatusesNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this account statuses not acceptable response a status code equal to that given
func (o *AccountStatusesNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the account statuses not acceptable response
func (o *AccountStatusesNotAcceptable) Code() int {
	return 406
}

func (o *AccountStatusesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesNotAcceptable", 406)
}

func (o *AccountStatusesNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesNotAcceptable", 406)
}

func (o *AccountStatusesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountStatusesInternalServerError creates a AccountStatusesInternalServerError with default headers values
func NewAccountStatusesInternalServerError() *AccountStatusesInternalServerError {
	return &AccountStatusesInternalServerError{}
}

/*
AccountStatusesInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountStatusesInternalServerError struct {
}

// IsSuccess returns true when this account statuses internal server error response has a 2xx status code
func (o *AccountStatusesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this account statuses internal server error response has a 3xx status code
func (o *AccountStatusesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account statuses internal server error response has a 4xx status code
func (o *AccountStatusesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this account statuses internal server error response has a 5xx status code
func (o *AccountStatusesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this account statuses internal server error response a status code equal to that given
func (o *AccountStatusesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the account statuses internal server error response
func (o *AccountStatusesInternalServerError) Code() int {
	return 500
}

func (o *AccountStatusesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesInternalServerError", 500)
}

func (o *AccountStatusesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/statuses][%d] accountStatusesInternalServerError", 500)
}

func (o *AccountStatusesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}