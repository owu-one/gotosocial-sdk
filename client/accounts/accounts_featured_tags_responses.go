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
)

// AccountsFeaturedTagsReader is a Reader for the AccountsFeaturedTags structure.
type AccountsFeaturedTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountsFeaturedTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountsFeaturedTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccountsFeaturedTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccountsFeaturedTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccountsFeaturedTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAccountsFeaturedTagsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccountsFeaturedTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/accounts/{id}/featured_tags] accountsFeaturedTags", response, response.Code())
	}
}

// NewAccountsFeaturedTagsOK creates a AccountsFeaturedTagsOK with default headers values
func NewAccountsFeaturedTagsOK() *AccountsFeaturedTagsOK {
	return &AccountsFeaturedTagsOK{}
}

/*
AccountsFeaturedTagsOK describes a response with status code 200, with default header values.

AccountsFeaturedTagsOK accounts featured tags o k
*/
type AccountsFeaturedTagsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this accounts featured tags o k response has a 2xx status code
func (o *AccountsFeaturedTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this accounts featured tags o k response has a 3xx status code
func (o *AccountsFeaturedTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts featured tags o k response has a 4xx status code
func (o *AccountsFeaturedTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this accounts featured tags o k response has a 5xx status code
func (o *AccountsFeaturedTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this accounts featured tags o k response a status code equal to that given
func (o *AccountsFeaturedTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the accounts featured tags o k response
func (o *AccountsFeaturedTagsOK) Code() int {
	return 200
}

func (o *AccountsFeaturedTagsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsOK %s", 200, payload)
}

func (o *AccountsFeaturedTagsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsOK %s", 200, payload)
}

func (o *AccountsFeaturedTagsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *AccountsFeaturedTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountsFeaturedTagsBadRequest creates a AccountsFeaturedTagsBadRequest with default headers values
func NewAccountsFeaturedTagsBadRequest() *AccountsFeaturedTagsBadRequest {
	return &AccountsFeaturedTagsBadRequest{}
}

/*
AccountsFeaturedTagsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AccountsFeaturedTagsBadRequest struct {
}

// IsSuccess returns true when this accounts featured tags bad request response has a 2xx status code
func (o *AccountsFeaturedTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this accounts featured tags bad request response has a 3xx status code
func (o *AccountsFeaturedTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts featured tags bad request response has a 4xx status code
func (o *AccountsFeaturedTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this accounts featured tags bad request response has a 5xx status code
func (o *AccountsFeaturedTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this accounts featured tags bad request response a status code equal to that given
func (o *AccountsFeaturedTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the accounts featured tags bad request response
func (o *AccountsFeaturedTagsBadRequest) Code() int {
	return 400
}

func (o *AccountsFeaturedTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsBadRequest", 400)
}

func (o *AccountsFeaturedTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsBadRequest", 400)
}

func (o *AccountsFeaturedTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountsFeaturedTagsUnauthorized creates a AccountsFeaturedTagsUnauthorized with default headers values
func NewAccountsFeaturedTagsUnauthorized() *AccountsFeaturedTagsUnauthorized {
	return &AccountsFeaturedTagsUnauthorized{}
}

/*
AccountsFeaturedTagsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AccountsFeaturedTagsUnauthorized struct {
}

// IsSuccess returns true when this accounts featured tags unauthorized response has a 2xx status code
func (o *AccountsFeaturedTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this accounts featured tags unauthorized response has a 3xx status code
func (o *AccountsFeaturedTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts featured tags unauthorized response has a 4xx status code
func (o *AccountsFeaturedTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this accounts featured tags unauthorized response has a 5xx status code
func (o *AccountsFeaturedTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this accounts featured tags unauthorized response a status code equal to that given
func (o *AccountsFeaturedTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the accounts featured tags unauthorized response
func (o *AccountsFeaturedTagsUnauthorized) Code() int {
	return 401
}

func (o *AccountsFeaturedTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsUnauthorized", 401)
}

func (o *AccountsFeaturedTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsUnauthorized", 401)
}

func (o *AccountsFeaturedTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountsFeaturedTagsNotFound creates a AccountsFeaturedTagsNotFound with default headers values
func NewAccountsFeaturedTagsNotFound() *AccountsFeaturedTagsNotFound {
	return &AccountsFeaturedTagsNotFound{}
}

/*
AccountsFeaturedTagsNotFound describes a response with status code 404, with default header values.

not found
*/
type AccountsFeaturedTagsNotFound struct {
}

// IsSuccess returns true when this accounts featured tags not found response has a 2xx status code
func (o *AccountsFeaturedTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this accounts featured tags not found response has a 3xx status code
func (o *AccountsFeaturedTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts featured tags not found response has a 4xx status code
func (o *AccountsFeaturedTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this accounts featured tags not found response has a 5xx status code
func (o *AccountsFeaturedTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this accounts featured tags not found response a status code equal to that given
func (o *AccountsFeaturedTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the accounts featured tags not found response
func (o *AccountsFeaturedTagsNotFound) Code() int {
	return 404
}

func (o *AccountsFeaturedTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsNotFound", 404)
}

func (o *AccountsFeaturedTagsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsNotFound", 404)
}

func (o *AccountsFeaturedTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountsFeaturedTagsNotAcceptable creates a AccountsFeaturedTagsNotAcceptable with default headers values
func NewAccountsFeaturedTagsNotAcceptable() *AccountsFeaturedTagsNotAcceptable {
	return &AccountsFeaturedTagsNotAcceptable{}
}

/*
AccountsFeaturedTagsNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AccountsFeaturedTagsNotAcceptable struct {
}

// IsSuccess returns true when this accounts featured tags not acceptable response has a 2xx status code
func (o *AccountsFeaturedTagsNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this accounts featured tags not acceptable response has a 3xx status code
func (o *AccountsFeaturedTagsNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts featured tags not acceptable response has a 4xx status code
func (o *AccountsFeaturedTagsNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this accounts featured tags not acceptable response has a 5xx status code
func (o *AccountsFeaturedTagsNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this accounts featured tags not acceptable response a status code equal to that given
func (o *AccountsFeaturedTagsNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the accounts featured tags not acceptable response
func (o *AccountsFeaturedTagsNotAcceptable) Code() int {
	return 406
}

func (o *AccountsFeaturedTagsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsNotAcceptable", 406)
}

func (o *AccountsFeaturedTagsNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsNotAcceptable", 406)
}

func (o *AccountsFeaturedTagsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountsFeaturedTagsInternalServerError creates a AccountsFeaturedTagsInternalServerError with default headers values
func NewAccountsFeaturedTagsInternalServerError() *AccountsFeaturedTagsInternalServerError {
	return &AccountsFeaturedTagsInternalServerError{}
}

/*
AccountsFeaturedTagsInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AccountsFeaturedTagsInternalServerError struct {
}

// IsSuccess returns true when this accounts featured tags internal server error response has a 2xx status code
func (o *AccountsFeaturedTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this accounts featured tags internal server error response has a 3xx status code
func (o *AccountsFeaturedTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accounts featured tags internal server error response has a 4xx status code
func (o *AccountsFeaturedTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this accounts featured tags internal server error response has a 5xx status code
func (o *AccountsFeaturedTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this accounts featured tags internal server error response a status code equal to that given
func (o *AccountsFeaturedTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the accounts featured tags internal server error response
func (o *AccountsFeaturedTagsInternalServerError) Code() int {
	return 500
}

func (o *AccountsFeaturedTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsInternalServerError", 500)
}

func (o *AccountsFeaturedTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/accounts/{id}/featured_tags][%d] accountsFeaturedTagsInternalServerError", 500)
}

func (o *AccountsFeaturedTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
