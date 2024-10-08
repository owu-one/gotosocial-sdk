// Code generated by go-swagger; DO NOT EDIT.

package filters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// FilterKeywordDeleteReader is a Reader for the FilterKeywordDelete structure.
type FilterKeywordDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterKeywordDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterKeywordDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilterKeywordDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFilterKeywordDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFilterKeywordDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFilterKeywordDeleteNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilterKeywordDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v2/filters/keywords/{id}] filterKeywordDelete", response, response.Code())
	}
}

// NewFilterKeywordDeleteOK creates a FilterKeywordDeleteOK with default headers values
func NewFilterKeywordDeleteOK() *FilterKeywordDeleteOK {
	return &FilterKeywordDeleteOK{}
}

/*
FilterKeywordDeleteOK describes a response with status code 200, with default header values.

filter keyword deleted
*/
type FilterKeywordDeleteOK struct {
}

// IsSuccess returns true when this filter keyword delete o k response has a 2xx status code
func (o *FilterKeywordDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this filter keyword delete o k response has a 3xx status code
func (o *FilterKeywordDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword delete o k response has a 4xx status code
func (o *FilterKeywordDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter keyword delete o k response has a 5xx status code
func (o *FilterKeywordDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword delete o k response a status code equal to that given
func (o *FilterKeywordDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the filter keyword delete o k response
func (o *FilterKeywordDeleteOK) Code() int {
	return 200
}

func (o *FilterKeywordDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteOK", 200)
}

func (o *FilterKeywordDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteOK", 200)
}

func (o *FilterKeywordDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordDeleteBadRequest creates a FilterKeywordDeleteBadRequest with default headers values
func NewFilterKeywordDeleteBadRequest() *FilterKeywordDeleteBadRequest {
	return &FilterKeywordDeleteBadRequest{}
}

/*
FilterKeywordDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FilterKeywordDeleteBadRequest struct {
}

// IsSuccess returns true when this filter keyword delete bad request response has a 2xx status code
func (o *FilterKeywordDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword delete bad request response has a 3xx status code
func (o *FilterKeywordDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword delete bad request response has a 4xx status code
func (o *FilterKeywordDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword delete bad request response has a 5xx status code
func (o *FilterKeywordDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword delete bad request response a status code equal to that given
func (o *FilterKeywordDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the filter keyword delete bad request response
func (o *FilterKeywordDeleteBadRequest) Code() int {
	return 400
}

func (o *FilterKeywordDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteBadRequest", 400)
}

func (o *FilterKeywordDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteBadRequest", 400)
}

func (o *FilterKeywordDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordDeleteUnauthorized creates a FilterKeywordDeleteUnauthorized with default headers values
func NewFilterKeywordDeleteUnauthorized() *FilterKeywordDeleteUnauthorized {
	return &FilterKeywordDeleteUnauthorized{}
}

/*
FilterKeywordDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FilterKeywordDeleteUnauthorized struct {
}

// IsSuccess returns true when this filter keyword delete unauthorized response has a 2xx status code
func (o *FilterKeywordDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword delete unauthorized response has a 3xx status code
func (o *FilterKeywordDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword delete unauthorized response has a 4xx status code
func (o *FilterKeywordDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword delete unauthorized response has a 5xx status code
func (o *FilterKeywordDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword delete unauthorized response a status code equal to that given
func (o *FilterKeywordDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the filter keyword delete unauthorized response
func (o *FilterKeywordDeleteUnauthorized) Code() int {
	return 401
}

func (o *FilterKeywordDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteUnauthorized", 401)
}

func (o *FilterKeywordDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteUnauthorized", 401)
}

func (o *FilterKeywordDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordDeleteNotFound creates a FilterKeywordDeleteNotFound with default headers values
func NewFilterKeywordDeleteNotFound() *FilterKeywordDeleteNotFound {
	return &FilterKeywordDeleteNotFound{}
}

/*
FilterKeywordDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type FilterKeywordDeleteNotFound struct {
}

// IsSuccess returns true when this filter keyword delete not found response has a 2xx status code
func (o *FilterKeywordDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword delete not found response has a 3xx status code
func (o *FilterKeywordDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword delete not found response has a 4xx status code
func (o *FilterKeywordDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword delete not found response has a 5xx status code
func (o *FilterKeywordDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword delete not found response a status code equal to that given
func (o *FilterKeywordDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the filter keyword delete not found response
func (o *FilterKeywordDeleteNotFound) Code() int {
	return 404
}

func (o *FilterKeywordDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteNotFound", 404)
}

func (o *FilterKeywordDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteNotFound", 404)
}

func (o *FilterKeywordDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordDeleteNotAcceptable creates a FilterKeywordDeleteNotAcceptable with default headers values
func NewFilterKeywordDeleteNotAcceptable() *FilterKeywordDeleteNotAcceptable {
	return &FilterKeywordDeleteNotAcceptable{}
}

/*
FilterKeywordDeleteNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FilterKeywordDeleteNotAcceptable struct {
}

// IsSuccess returns true when this filter keyword delete not acceptable response has a 2xx status code
func (o *FilterKeywordDeleteNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword delete not acceptable response has a 3xx status code
func (o *FilterKeywordDeleteNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword delete not acceptable response has a 4xx status code
func (o *FilterKeywordDeleteNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this filter keyword delete not acceptable response has a 5xx status code
func (o *FilterKeywordDeleteNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this filter keyword delete not acceptable response a status code equal to that given
func (o *FilterKeywordDeleteNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the filter keyword delete not acceptable response
func (o *FilterKeywordDeleteNotAcceptable) Code() int {
	return 406
}

func (o *FilterKeywordDeleteNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteNotAcceptable", 406)
}

func (o *FilterKeywordDeleteNotAcceptable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteNotAcceptable", 406)
}

func (o *FilterKeywordDeleteNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFilterKeywordDeleteInternalServerError creates a FilterKeywordDeleteInternalServerError with default headers values
func NewFilterKeywordDeleteInternalServerError() *FilterKeywordDeleteInternalServerError {
	return &FilterKeywordDeleteInternalServerError{}
}

/*
FilterKeywordDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FilterKeywordDeleteInternalServerError struct {
}

// IsSuccess returns true when this filter keyword delete internal server error response has a 2xx status code
func (o *FilterKeywordDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this filter keyword delete internal server error response has a 3xx status code
func (o *FilterKeywordDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this filter keyword delete internal server error response has a 4xx status code
func (o *FilterKeywordDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this filter keyword delete internal server error response has a 5xx status code
func (o *FilterKeywordDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this filter keyword delete internal server error response a status code equal to that given
func (o *FilterKeywordDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the filter keyword delete internal server error response
func (o *FilterKeywordDeleteInternalServerError) Code() int {
	return 500
}

func (o *FilterKeywordDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteInternalServerError", 500)
}

func (o *FilterKeywordDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/filters/keywords/{id}][%d] filterKeywordDeleteInternalServerError", 500)
}

func (o *FilterKeywordDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
