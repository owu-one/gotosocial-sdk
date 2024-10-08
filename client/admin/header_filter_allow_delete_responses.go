// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HeaderFilterAllowDeleteReader is a Reader for the HeaderFilterAllowDelete structure.
type HeaderFilterAllowDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeaderFilterAllowDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewHeaderFilterAllowDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHeaderFilterAllowDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHeaderFilterAllowDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHeaderFilterAllowDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHeaderFilterAllowDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHeaderFilterAllowDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v1/admin/header_allows/{id}] headerFilterAllowDelete", response, response.Code())
	}
}

// NewHeaderFilterAllowDeleteAccepted creates a HeaderFilterAllowDeleteAccepted with default headers values
func NewHeaderFilterAllowDeleteAccepted() *HeaderFilterAllowDeleteAccepted {
	return &HeaderFilterAllowDeleteAccepted{}
}

/*
HeaderFilterAllowDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type HeaderFilterAllowDeleteAccepted struct {
}

// IsSuccess returns true when this header filter allow delete accepted response has a 2xx status code
func (o *HeaderFilterAllowDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this header filter allow delete accepted response has a 3xx status code
func (o *HeaderFilterAllowDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter allow delete accepted response has a 4xx status code
func (o *HeaderFilterAllowDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter allow delete accepted response has a 5xx status code
func (o *HeaderFilterAllowDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter allow delete accepted response a status code equal to that given
func (o *HeaderFilterAllowDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the header filter allow delete accepted response
func (o *HeaderFilterAllowDeleteAccepted) Code() int {
	return 202
}

func (o *HeaderFilterAllowDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteAccepted", 202)
}

func (o *HeaderFilterAllowDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteAccepted", 202)
}

func (o *HeaderFilterAllowDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterAllowDeleteBadRequest creates a HeaderFilterAllowDeleteBadRequest with default headers values
func NewHeaderFilterAllowDeleteBadRequest() *HeaderFilterAllowDeleteBadRequest {
	return &HeaderFilterAllowDeleteBadRequest{}
}

/*
HeaderFilterAllowDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type HeaderFilterAllowDeleteBadRequest struct {
}

// IsSuccess returns true when this header filter allow delete bad request response has a 2xx status code
func (o *HeaderFilterAllowDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter allow delete bad request response has a 3xx status code
func (o *HeaderFilterAllowDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter allow delete bad request response has a 4xx status code
func (o *HeaderFilterAllowDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter allow delete bad request response has a 5xx status code
func (o *HeaderFilterAllowDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter allow delete bad request response a status code equal to that given
func (o *HeaderFilterAllowDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the header filter allow delete bad request response
func (o *HeaderFilterAllowDeleteBadRequest) Code() int {
	return 400
}

func (o *HeaderFilterAllowDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteBadRequest", 400)
}

func (o *HeaderFilterAllowDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteBadRequest", 400)
}

func (o *HeaderFilterAllowDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterAllowDeleteUnauthorized creates a HeaderFilterAllowDeleteUnauthorized with default headers values
func NewHeaderFilterAllowDeleteUnauthorized() *HeaderFilterAllowDeleteUnauthorized {
	return &HeaderFilterAllowDeleteUnauthorized{}
}

/*
HeaderFilterAllowDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type HeaderFilterAllowDeleteUnauthorized struct {
}

// IsSuccess returns true when this header filter allow delete unauthorized response has a 2xx status code
func (o *HeaderFilterAllowDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter allow delete unauthorized response has a 3xx status code
func (o *HeaderFilterAllowDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter allow delete unauthorized response has a 4xx status code
func (o *HeaderFilterAllowDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter allow delete unauthorized response has a 5xx status code
func (o *HeaderFilterAllowDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter allow delete unauthorized response a status code equal to that given
func (o *HeaderFilterAllowDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the header filter allow delete unauthorized response
func (o *HeaderFilterAllowDeleteUnauthorized) Code() int {
	return 401
}

func (o *HeaderFilterAllowDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteUnauthorized", 401)
}

func (o *HeaderFilterAllowDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteUnauthorized", 401)
}

func (o *HeaderFilterAllowDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterAllowDeleteForbidden creates a HeaderFilterAllowDeleteForbidden with default headers values
func NewHeaderFilterAllowDeleteForbidden() *HeaderFilterAllowDeleteForbidden {
	return &HeaderFilterAllowDeleteForbidden{}
}

/*
HeaderFilterAllowDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type HeaderFilterAllowDeleteForbidden struct {
}

// IsSuccess returns true when this header filter allow delete forbidden response has a 2xx status code
func (o *HeaderFilterAllowDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter allow delete forbidden response has a 3xx status code
func (o *HeaderFilterAllowDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter allow delete forbidden response has a 4xx status code
func (o *HeaderFilterAllowDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter allow delete forbidden response has a 5xx status code
func (o *HeaderFilterAllowDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter allow delete forbidden response a status code equal to that given
func (o *HeaderFilterAllowDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the header filter allow delete forbidden response
func (o *HeaderFilterAllowDeleteForbidden) Code() int {
	return 403
}

func (o *HeaderFilterAllowDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteForbidden", 403)
}

func (o *HeaderFilterAllowDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteForbidden", 403)
}

func (o *HeaderFilterAllowDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterAllowDeleteNotFound creates a HeaderFilterAllowDeleteNotFound with default headers values
func NewHeaderFilterAllowDeleteNotFound() *HeaderFilterAllowDeleteNotFound {
	return &HeaderFilterAllowDeleteNotFound{}
}

/*
HeaderFilterAllowDeleteNotFound describes a response with status code 404, with default header values.

not found
*/
type HeaderFilterAllowDeleteNotFound struct {
}

// IsSuccess returns true when this header filter allow delete not found response has a 2xx status code
func (o *HeaderFilterAllowDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter allow delete not found response has a 3xx status code
func (o *HeaderFilterAllowDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter allow delete not found response has a 4xx status code
func (o *HeaderFilterAllowDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this header filter allow delete not found response has a 5xx status code
func (o *HeaderFilterAllowDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this header filter allow delete not found response a status code equal to that given
func (o *HeaderFilterAllowDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the header filter allow delete not found response
func (o *HeaderFilterAllowDeleteNotFound) Code() int {
	return 404
}

func (o *HeaderFilterAllowDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteNotFound", 404)
}

func (o *HeaderFilterAllowDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteNotFound", 404)
}

func (o *HeaderFilterAllowDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHeaderFilterAllowDeleteInternalServerError creates a HeaderFilterAllowDeleteInternalServerError with default headers values
func NewHeaderFilterAllowDeleteInternalServerError() *HeaderFilterAllowDeleteInternalServerError {
	return &HeaderFilterAllowDeleteInternalServerError{}
}

/*
HeaderFilterAllowDeleteInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type HeaderFilterAllowDeleteInternalServerError struct {
}

// IsSuccess returns true when this header filter allow delete internal server error response has a 2xx status code
func (o *HeaderFilterAllowDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this header filter allow delete internal server error response has a 3xx status code
func (o *HeaderFilterAllowDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this header filter allow delete internal server error response has a 4xx status code
func (o *HeaderFilterAllowDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this header filter allow delete internal server error response has a 5xx status code
func (o *HeaderFilterAllowDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this header filter allow delete internal server error response a status code equal to that given
func (o *HeaderFilterAllowDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the header filter allow delete internal server error response
func (o *HeaderFilterAllowDeleteInternalServerError) Code() int {
	return 500
}

func (o *HeaderFilterAllowDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteInternalServerError", 500)
}

func (o *HeaderFilterAllowDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/admin/header_allows/{id}][%d] headerFilterAllowDeleteInternalServerError", 500)
}

func (o *HeaderFilterAllowDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
