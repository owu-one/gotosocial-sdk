// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TwoFactorQRCodeURIGetReader is a Reader for the TwoFactorQRCodeURIGet structure.
type TwoFactorQRCodeURIGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TwoFactorQRCodeURIGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTwoFactorQRCodeURIGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTwoFactorQRCodeURIGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTwoFactorQRCodeURIGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewTwoFactorQRCodeURIGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTwoFactorQRCodeURIGetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTwoFactorQRCodeURIGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTwoFactorQRCodeURIGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/user/2fa/qruri] TwoFactorQRCodeURIGet", response, response.Code())
	}
}

// NewTwoFactorQRCodeURIGetOK creates a TwoFactorQRCodeURIGetOK with default headers values
func NewTwoFactorQRCodeURIGetOK() *TwoFactorQRCodeURIGetOK {
	return &TwoFactorQRCodeURIGetOK{}
}

/*
TwoFactorQRCodeURIGetOK describes a response with status code 200, with default header values.

QR code uri
*/
type TwoFactorQRCodeURIGetOK struct {
}

// IsSuccess returns true when this two factor q r code Uri get o k response has a 2xx status code
func (o *TwoFactorQRCodeURIGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this two factor q r code Uri get o k response has a 3xx status code
func (o *TwoFactorQRCodeURIGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get o k response has a 4xx status code
func (o *TwoFactorQRCodeURIGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this two factor q r code Uri get o k response has a 5xx status code
func (o *TwoFactorQRCodeURIGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code Uri get o k response a status code equal to that given
func (o *TwoFactorQRCodeURIGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the two factor q r code Uri get o k response
func (o *TwoFactorQRCodeURIGetOK) Code() int {
	return 200
}

func (o *TwoFactorQRCodeURIGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetOK", 200)
}

func (o *TwoFactorQRCodeURIGetOK) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetOK", 200)
}

func (o *TwoFactorQRCodeURIGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodeURIGetUnauthorized creates a TwoFactorQRCodeURIGetUnauthorized with default headers values
func NewTwoFactorQRCodeURIGetUnauthorized() *TwoFactorQRCodeURIGetUnauthorized {
	return &TwoFactorQRCodeURIGetUnauthorized{}
}

/*
TwoFactorQRCodeURIGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type TwoFactorQRCodeURIGetUnauthorized struct {
}

// IsSuccess returns true when this two factor q r code Uri get unauthorized response has a 2xx status code
func (o *TwoFactorQRCodeURIGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code Uri get unauthorized response has a 3xx status code
func (o *TwoFactorQRCodeURIGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get unauthorized response has a 4xx status code
func (o *TwoFactorQRCodeURIGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code Uri get unauthorized response has a 5xx status code
func (o *TwoFactorQRCodeURIGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code Uri get unauthorized response a status code equal to that given
func (o *TwoFactorQRCodeURIGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the two factor q r code Uri get unauthorized response
func (o *TwoFactorQRCodeURIGetUnauthorized) Code() int {
	return 401
}

func (o *TwoFactorQRCodeURIGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetUnauthorized", 401)
}

func (o *TwoFactorQRCodeURIGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetUnauthorized", 401)
}

func (o *TwoFactorQRCodeURIGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodeURIGetForbidden creates a TwoFactorQRCodeURIGetForbidden with default headers values
func NewTwoFactorQRCodeURIGetForbidden() *TwoFactorQRCodeURIGetForbidden {
	return &TwoFactorQRCodeURIGetForbidden{}
}

/*
TwoFactorQRCodeURIGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type TwoFactorQRCodeURIGetForbidden struct {
}

// IsSuccess returns true when this two factor q r code Uri get forbidden response has a 2xx status code
func (o *TwoFactorQRCodeURIGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code Uri get forbidden response has a 3xx status code
func (o *TwoFactorQRCodeURIGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get forbidden response has a 4xx status code
func (o *TwoFactorQRCodeURIGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code Uri get forbidden response has a 5xx status code
func (o *TwoFactorQRCodeURIGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code Uri get forbidden response a status code equal to that given
func (o *TwoFactorQRCodeURIGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the two factor q r code Uri get forbidden response
func (o *TwoFactorQRCodeURIGetForbidden) Code() int {
	return 403
}

func (o *TwoFactorQRCodeURIGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetForbidden", 403)
}

func (o *TwoFactorQRCodeURIGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetForbidden", 403)
}

func (o *TwoFactorQRCodeURIGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodeURIGetNotAcceptable creates a TwoFactorQRCodeURIGetNotAcceptable with default headers values
func NewTwoFactorQRCodeURIGetNotAcceptable() *TwoFactorQRCodeURIGetNotAcceptable {
	return &TwoFactorQRCodeURIGetNotAcceptable{}
}

/*
TwoFactorQRCodeURIGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type TwoFactorQRCodeURIGetNotAcceptable struct {
}

// IsSuccess returns true when this two factor q r code Uri get not acceptable response has a 2xx status code
func (o *TwoFactorQRCodeURIGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code Uri get not acceptable response has a 3xx status code
func (o *TwoFactorQRCodeURIGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get not acceptable response has a 4xx status code
func (o *TwoFactorQRCodeURIGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code Uri get not acceptable response has a 5xx status code
func (o *TwoFactorQRCodeURIGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code Uri get not acceptable response a status code equal to that given
func (o *TwoFactorQRCodeURIGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the two factor q r code Uri get not acceptable response
func (o *TwoFactorQRCodeURIGetNotAcceptable) Code() int {
	return 406
}

func (o *TwoFactorQRCodeURIGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetNotAcceptable", 406)
}

func (o *TwoFactorQRCodeURIGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetNotAcceptable", 406)
}

func (o *TwoFactorQRCodeURIGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodeURIGetConflict creates a TwoFactorQRCodeURIGetConflict with default headers values
func NewTwoFactorQRCodeURIGetConflict() *TwoFactorQRCodeURIGetConflict {
	return &TwoFactorQRCodeURIGetConflict{}
}

/*
TwoFactorQRCodeURIGetConflict describes a response with status code 409, with default header values.

conflict
*/
type TwoFactorQRCodeURIGetConflict struct {
}

// IsSuccess returns true when this two factor q r code Uri get conflict response has a 2xx status code
func (o *TwoFactorQRCodeURIGetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code Uri get conflict response has a 3xx status code
func (o *TwoFactorQRCodeURIGetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get conflict response has a 4xx status code
func (o *TwoFactorQRCodeURIGetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code Uri get conflict response has a 5xx status code
func (o *TwoFactorQRCodeURIGetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code Uri get conflict response a status code equal to that given
func (o *TwoFactorQRCodeURIGetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the two factor q r code Uri get conflict response
func (o *TwoFactorQRCodeURIGetConflict) Code() int {
	return 409
}

func (o *TwoFactorQRCodeURIGetConflict) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetConflict", 409)
}

func (o *TwoFactorQRCodeURIGetConflict) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetConflict", 409)
}

func (o *TwoFactorQRCodeURIGetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodeURIGetUnprocessableEntity creates a TwoFactorQRCodeURIGetUnprocessableEntity with default headers values
func NewTwoFactorQRCodeURIGetUnprocessableEntity() *TwoFactorQRCodeURIGetUnprocessableEntity {
	return &TwoFactorQRCodeURIGetUnprocessableEntity{}
}

/*
TwoFactorQRCodeURIGetUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type TwoFactorQRCodeURIGetUnprocessableEntity struct {
}

// IsSuccess returns true when this two factor q r code Uri get unprocessable entity response has a 2xx status code
func (o *TwoFactorQRCodeURIGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code Uri get unprocessable entity response has a 3xx status code
func (o *TwoFactorQRCodeURIGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get unprocessable entity response has a 4xx status code
func (o *TwoFactorQRCodeURIGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code Uri get unprocessable entity response has a 5xx status code
func (o *TwoFactorQRCodeURIGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code Uri get unprocessable entity response a status code equal to that given
func (o *TwoFactorQRCodeURIGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the two factor q r code Uri get unprocessable entity response
func (o *TwoFactorQRCodeURIGetUnprocessableEntity) Code() int {
	return 422
}

func (o *TwoFactorQRCodeURIGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetUnprocessableEntity", 422)
}

func (o *TwoFactorQRCodeURIGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetUnprocessableEntity", 422)
}

func (o *TwoFactorQRCodeURIGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodeURIGetInternalServerError creates a TwoFactorQRCodeURIGetInternalServerError with default headers values
func NewTwoFactorQRCodeURIGetInternalServerError() *TwoFactorQRCodeURIGetInternalServerError {
	return &TwoFactorQRCodeURIGetInternalServerError{}
}

/*
TwoFactorQRCodeURIGetInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type TwoFactorQRCodeURIGetInternalServerError struct {
}

// IsSuccess returns true when this two factor q r code Uri get internal server error response has a 2xx status code
func (o *TwoFactorQRCodeURIGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code Uri get internal server error response has a 3xx status code
func (o *TwoFactorQRCodeURIGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code Uri get internal server error response has a 4xx status code
func (o *TwoFactorQRCodeURIGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this two factor q r code Uri get internal server error response has a 5xx status code
func (o *TwoFactorQRCodeURIGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this two factor q r code Uri get internal server error response a status code equal to that given
func (o *TwoFactorQRCodeURIGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the two factor q r code Uri get internal server error response
func (o *TwoFactorQRCodeURIGetInternalServerError) Code() int {
	return 500
}

func (o *TwoFactorQRCodeURIGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetInternalServerError", 500)
}

func (o *TwoFactorQRCodeURIGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qruri][%d] twoFactorQRCodeUriGetInternalServerError", 500)
}

func (o *TwoFactorQRCodeURIGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
