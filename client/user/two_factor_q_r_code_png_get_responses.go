// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TwoFactorQRCodePngGetReader is a Reader for the TwoFactorQRCodePngGet structure.
type TwoFactorQRCodePngGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TwoFactorQRCodePngGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTwoFactorQRCodePngGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTwoFactorQRCodePngGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTwoFactorQRCodePngGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewTwoFactorQRCodePngGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTwoFactorQRCodePngGetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTwoFactorQRCodePngGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTwoFactorQRCodePngGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/user/2fa/qr.png] TwoFactorQRCodePngGet", response, response.Code())
	}
}

// NewTwoFactorQRCodePngGetOK creates a TwoFactorQRCodePngGetOK with default headers values
func NewTwoFactorQRCodePngGetOK() *TwoFactorQRCodePngGetOK {
	return &TwoFactorQRCodePngGetOK{}
}

/*
TwoFactorQRCodePngGetOK describes a response with status code 200, with default header values.

QR code png
*/
type TwoFactorQRCodePngGetOK struct {
}

// IsSuccess returns true when this two factor q r code png get o k response has a 2xx status code
func (o *TwoFactorQRCodePngGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this two factor q r code png get o k response has a 3xx status code
func (o *TwoFactorQRCodePngGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get o k response has a 4xx status code
func (o *TwoFactorQRCodePngGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this two factor q r code png get o k response has a 5xx status code
func (o *TwoFactorQRCodePngGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code png get o k response a status code equal to that given
func (o *TwoFactorQRCodePngGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the two factor q r code png get o k response
func (o *TwoFactorQRCodePngGetOK) Code() int {
	return 200
}

func (o *TwoFactorQRCodePngGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetOK", 200)
}

func (o *TwoFactorQRCodePngGetOK) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetOK", 200)
}

func (o *TwoFactorQRCodePngGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodePngGetUnauthorized creates a TwoFactorQRCodePngGetUnauthorized with default headers values
func NewTwoFactorQRCodePngGetUnauthorized() *TwoFactorQRCodePngGetUnauthorized {
	return &TwoFactorQRCodePngGetUnauthorized{}
}

/*
TwoFactorQRCodePngGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type TwoFactorQRCodePngGetUnauthorized struct {
}

// IsSuccess returns true when this two factor q r code png get unauthorized response has a 2xx status code
func (o *TwoFactorQRCodePngGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code png get unauthorized response has a 3xx status code
func (o *TwoFactorQRCodePngGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get unauthorized response has a 4xx status code
func (o *TwoFactorQRCodePngGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code png get unauthorized response has a 5xx status code
func (o *TwoFactorQRCodePngGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code png get unauthorized response a status code equal to that given
func (o *TwoFactorQRCodePngGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the two factor q r code png get unauthorized response
func (o *TwoFactorQRCodePngGetUnauthorized) Code() int {
	return 401
}

func (o *TwoFactorQRCodePngGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetUnauthorized", 401)
}

func (o *TwoFactorQRCodePngGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetUnauthorized", 401)
}

func (o *TwoFactorQRCodePngGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodePngGetForbidden creates a TwoFactorQRCodePngGetForbidden with default headers values
func NewTwoFactorQRCodePngGetForbidden() *TwoFactorQRCodePngGetForbidden {
	return &TwoFactorQRCodePngGetForbidden{}
}

/*
TwoFactorQRCodePngGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type TwoFactorQRCodePngGetForbidden struct {
}

// IsSuccess returns true when this two factor q r code png get forbidden response has a 2xx status code
func (o *TwoFactorQRCodePngGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code png get forbidden response has a 3xx status code
func (o *TwoFactorQRCodePngGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get forbidden response has a 4xx status code
func (o *TwoFactorQRCodePngGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code png get forbidden response has a 5xx status code
func (o *TwoFactorQRCodePngGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code png get forbidden response a status code equal to that given
func (o *TwoFactorQRCodePngGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the two factor q r code png get forbidden response
func (o *TwoFactorQRCodePngGetForbidden) Code() int {
	return 403
}

func (o *TwoFactorQRCodePngGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetForbidden", 403)
}

func (o *TwoFactorQRCodePngGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetForbidden", 403)
}

func (o *TwoFactorQRCodePngGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodePngGetNotAcceptable creates a TwoFactorQRCodePngGetNotAcceptable with default headers values
func NewTwoFactorQRCodePngGetNotAcceptable() *TwoFactorQRCodePngGetNotAcceptable {
	return &TwoFactorQRCodePngGetNotAcceptable{}
}

/*
TwoFactorQRCodePngGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type TwoFactorQRCodePngGetNotAcceptable struct {
}

// IsSuccess returns true when this two factor q r code png get not acceptable response has a 2xx status code
func (o *TwoFactorQRCodePngGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code png get not acceptable response has a 3xx status code
func (o *TwoFactorQRCodePngGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get not acceptable response has a 4xx status code
func (o *TwoFactorQRCodePngGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code png get not acceptable response has a 5xx status code
func (o *TwoFactorQRCodePngGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code png get not acceptable response a status code equal to that given
func (o *TwoFactorQRCodePngGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the two factor q r code png get not acceptable response
func (o *TwoFactorQRCodePngGetNotAcceptable) Code() int {
	return 406
}

func (o *TwoFactorQRCodePngGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetNotAcceptable", 406)
}

func (o *TwoFactorQRCodePngGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetNotAcceptable", 406)
}

func (o *TwoFactorQRCodePngGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodePngGetConflict creates a TwoFactorQRCodePngGetConflict with default headers values
func NewTwoFactorQRCodePngGetConflict() *TwoFactorQRCodePngGetConflict {
	return &TwoFactorQRCodePngGetConflict{}
}

/*
TwoFactorQRCodePngGetConflict describes a response with status code 409, with default header values.

conflict
*/
type TwoFactorQRCodePngGetConflict struct {
}

// IsSuccess returns true when this two factor q r code png get conflict response has a 2xx status code
func (o *TwoFactorQRCodePngGetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code png get conflict response has a 3xx status code
func (o *TwoFactorQRCodePngGetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get conflict response has a 4xx status code
func (o *TwoFactorQRCodePngGetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code png get conflict response has a 5xx status code
func (o *TwoFactorQRCodePngGetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code png get conflict response a status code equal to that given
func (o *TwoFactorQRCodePngGetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the two factor q r code png get conflict response
func (o *TwoFactorQRCodePngGetConflict) Code() int {
	return 409
}

func (o *TwoFactorQRCodePngGetConflict) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetConflict", 409)
}

func (o *TwoFactorQRCodePngGetConflict) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetConflict", 409)
}

func (o *TwoFactorQRCodePngGetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodePngGetUnprocessableEntity creates a TwoFactorQRCodePngGetUnprocessableEntity with default headers values
func NewTwoFactorQRCodePngGetUnprocessableEntity() *TwoFactorQRCodePngGetUnprocessableEntity {
	return &TwoFactorQRCodePngGetUnprocessableEntity{}
}

/*
TwoFactorQRCodePngGetUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type TwoFactorQRCodePngGetUnprocessableEntity struct {
}

// IsSuccess returns true when this two factor q r code png get unprocessable entity response has a 2xx status code
func (o *TwoFactorQRCodePngGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code png get unprocessable entity response has a 3xx status code
func (o *TwoFactorQRCodePngGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get unprocessable entity response has a 4xx status code
func (o *TwoFactorQRCodePngGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this two factor q r code png get unprocessable entity response has a 5xx status code
func (o *TwoFactorQRCodePngGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this two factor q r code png get unprocessable entity response a status code equal to that given
func (o *TwoFactorQRCodePngGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the two factor q r code png get unprocessable entity response
func (o *TwoFactorQRCodePngGetUnprocessableEntity) Code() int {
	return 422
}

func (o *TwoFactorQRCodePngGetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetUnprocessableEntity", 422)
}

func (o *TwoFactorQRCodePngGetUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetUnprocessableEntity", 422)
}

func (o *TwoFactorQRCodePngGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTwoFactorQRCodePngGetInternalServerError creates a TwoFactorQRCodePngGetInternalServerError with default headers values
func NewTwoFactorQRCodePngGetInternalServerError() *TwoFactorQRCodePngGetInternalServerError {
	return &TwoFactorQRCodePngGetInternalServerError{}
}

/*
TwoFactorQRCodePngGetInternalServerError describes a response with status code 500, with default header values.

internal error
*/
type TwoFactorQRCodePngGetInternalServerError struct {
}

// IsSuccess returns true when this two factor q r code png get internal server error response has a 2xx status code
func (o *TwoFactorQRCodePngGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this two factor q r code png get internal server error response has a 3xx status code
func (o *TwoFactorQRCodePngGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this two factor q r code png get internal server error response has a 4xx status code
func (o *TwoFactorQRCodePngGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this two factor q r code png get internal server error response has a 5xx status code
func (o *TwoFactorQRCodePngGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this two factor q r code png get internal server error response a status code equal to that given
func (o *TwoFactorQRCodePngGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the two factor q r code png get internal server error response
func (o *TwoFactorQRCodePngGetInternalServerError) Code() int {
	return 500
}

func (o *TwoFactorQRCodePngGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetInternalServerError", 500)
}

func (o *TwoFactorQRCodePngGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/user/2fa/qr.png][%d] twoFactorQRCodePngGetInternalServerError", 500)
}

func (o *TwoFactorQRCodePngGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
