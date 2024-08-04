// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminAccountActionReader is a Reader for the AdminAccountAction structure.
type AdminAccountActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminAccountActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminAccountActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminAccountActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminAccountActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminAccountActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminAccountActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAdminAccountActionNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAdminAccountActionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminAccountActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/accounts/{id}/action] adminAccountAction", response, response.Code())
	}
}

// NewAdminAccountActionOK creates a AdminAccountActionOK with default headers values
func NewAdminAccountActionOK() *AdminAccountActionOK {
	return &AdminAccountActionOK{}
}

/*
AdminAccountActionOK describes a response with status code 200, with default header values.

OK
*/
type AdminAccountActionOK struct {
}

// IsSuccess returns true when this admin account action o k response has a 2xx status code
func (o *AdminAccountActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin account action o k response has a 3xx status code
func (o *AdminAccountActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action o k response has a 4xx status code
func (o *AdminAccountActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin account action o k response has a 5xx status code
func (o *AdminAccountActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action o k response a status code equal to that given
func (o *AdminAccountActionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin account action o k response
func (o *AdminAccountActionOK) Code() int {
	return 200
}

func (o *AdminAccountActionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionOK", 200)
}

func (o *AdminAccountActionOK) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionOK", 200)
}

func (o *AdminAccountActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionBadRequest creates a AdminAccountActionBadRequest with default headers values
func NewAdminAccountActionBadRequest() *AdminAccountActionBadRequest {
	return &AdminAccountActionBadRequest{}
}

/*
AdminAccountActionBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AdminAccountActionBadRequest struct {
}

// IsSuccess returns true when this admin account action bad request response has a 2xx status code
func (o *AdminAccountActionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action bad request response has a 3xx status code
func (o *AdminAccountActionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action bad request response has a 4xx status code
func (o *AdminAccountActionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account action bad request response has a 5xx status code
func (o *AdminAccountActionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action bad request response a status code equal to that given
func (o *AdminAccountActionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin account action bad request response
func (o *AdminAccountActionBadRequest) Code() int {
	return 400
}

func (o *AdminAccountActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionBadRequest", 400)
}

func (o *AdminAccountActionBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionBadRequest", 400)
}

func (o *AdminAccountActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionUnauthorized creates a AdminAccountActionUnauthorized with default headers values
func NewAdminAccountActionUnauthorized() *AdminAccountActionUnauthorized {
	return &AdminAccountActionUnauthorized{}
}

/*
AdminAccountActionUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AdminAccountActionUnauthorized struct {
}

// IsSuccess returns true when this admin account action unauthorized response has a 2xx status code
func (o *AdminAccountActionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action unauthorized response has a 3xx status code
func (o *AdminAccountActionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action unauthorized response has a 4xx status code
func (o *AdminAccountActionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account action unauthorized response has a 5xx status code
func (o *AdminAccountActionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action unauthorized response a status code equal to that given
func (o *AdminAccountActionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin account action unauthorized response
func (o *AdminAccountActionUnauthorized) Code() int {
	return 401
}

func (o *AdminAccountActionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionUnauthorized", 401)
}

func (o *AdminAccountActionUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionUnauthorized", 401)
}

func (o *AdminAccountActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionForbidden creates a AdminAccountActionForbidden with default headers values
func NewAdminAccountActionForbidden() *AdminAccountActionForbidden {
	return &AdminAccountActionForbidden{}
}

/*
AdminAccountActionForbidden describes a response with status code 403, with default header values.

forbidden
*/
type AdminAccountActionForbidden struct {
}

// IsSuccess returns true when this admin account action forbidden response has a 2xx status code
func (o *AdminAccountActionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action forbidden response has a 3xx status code
func (o *AdminAccountActionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action forbidden response has a 4xx status code
func (o *AdminAccountActionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account action forbidden response has a 5xx status code
func (o *AdminAccountActionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action forbidden response a status code equal to that given
func (o *AdminAccountActionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin account action forbidden response
func (o *AdminAccountActionForbidden) Code() int {
	return 403
}

func (o *AdminAccountActionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionForbidden", 403)
}

func (o *AdminAccountActionForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionForbidden", 403)
}

func (o *AdminAccountActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionNotFound creates a AdminAccountActionNotFound with default headers values
func NewAdminAccountActionNotFound() *AdminAccountActionNotFound {
	return &AdminAccountActionNotFound{}
}

/*
AdminAccountActionNotFound describes a response with status code 404, with default header values.

not found
*/
type AdminAccountActionNotFound struct {
}

// IsSuccess returns true when this admin account action not found response has a 2xx status code
func (o *AdminAccountActionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action not found response has a 3xx status code
func (o *AdminAccountActionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action not found response has a 4xx status code
func (o *AdminAccountActionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account action not found response has a 5xx status code
func (o *AdminAccountActionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action not found response a status code equal to that given
func (o *AdminAccountActionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin account action not found response
func (o *AdminAccountActionNotFound) Code() int {
	return 404
}

func (o *AdminAccountActionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionNotFound", 404)
}

func (o *AdminAccountActionNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionNotFound", 404)
}

func (o *AdminAccountActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionNotAcceptable creates a AdminAccountActionNotAcceptable with default headers values
func NewAdminAccountActionNotAcceptable() *AdminAccountActionNotAcceptable {
	return &AdminAccountActionNotAcceptable{}
}

/*
AdminAccountActionNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AdminAccountActionNotAcceptable struct {
}

// IsSuccess returns true when this admin account action not acceptable response has a 2xx status code
func (o *AdminAccountActionNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action not acceptable response has a 3xx status code
func (o *AdminAccountActionNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action not acceptable response has a 4xx status code
func (o *AdminAccountActionNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account action not acceptable response has a 5xx status code
func (o *AdminAccountActionNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action not acceptable response a status code equal to that given
func (o *AdminAccountActionNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the admin account action not acceptable response
func (o *AdminAccountActionNotAcceptable) Code() int {
	return 406
}

func (o *AdminAccountActionNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionNotAcceptable", 406)
}

func (o *AdminAccountActionNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionNotAcceptable", 406)
}

func (o *AdminAccountActionNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionConflict creates a AdminAccountActionConflict with default headers values
func NewAdminAccountActionConflict() *AdminAccountActionConflict {
	return &AdminAccountActionConflict{}
}

/*
AdminAccountActionConflict describes a response with status code 409, with default header values.

Conflict: There is already an admin action running that conflicts with this action. Check the error message in the response body for more information. This is a temporary error; it should be possible to process this action if you try again in a bit.
*/
type AdminAccountActionConflict struct {
}

// IsSuccess returns true when this admin account action conflict response has a 2xx status code
func (o *AdminAccountActionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action conflict response has a 3xx status code
func (o *AdminAccountActionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action conflict response has a 4xx status code
func (o *AdminAccountActionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin account action conflict response has a 5xx status code
func (o *AdminAccountActionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this admin account action conflict response a status code equal to that given
func (o *AdminAccountActionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the admin account action conflict response
func (o *AdminAccountActionConflict) Code() int {
	return 409
}

func (o *AdminAccountActionConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionConflict", 409)
}

func (o *AdminAccountActionConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionConflict", 409)
}

func (o *AdminAccountActionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAccountActionInternalServerError creates a AdminAccountActionInternalServerError with default headers values
func NewAdminAccountActionInternalServerError() *AdminAccountActionInternalServerError {
	return &AdminAccountActionInternalServerError{}
}

/*
AdminAccountActionInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AdminAccountActionInternalServerError struct {
}

// IsSuccess returns true when this admin account action internal server error response has a 2xx status code
func (o *AdminAccountActionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin account action internal server error response has a 3xx status code
func (o *AdminAccountActionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin account action internal server error response has a 4xx status code
func (o *AdminAccountActionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin account action internal server error response has a 5xx status code
func (o *AdminAccountActionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin account action internal server error response a status code equal to that given
func (o *AdminAccountActionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin account action internal server error response
func (o *AdminAccountActionInternalServerError) Code() int {
	return 500
}

func (o *AdminAccountActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionInternalServerError", 500)
}

func (o *AdminAccountActionInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/accounts/{id}/action][%d] adminAccountActionInternalServerError", 500)
}

func (o *AdminAccountActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}