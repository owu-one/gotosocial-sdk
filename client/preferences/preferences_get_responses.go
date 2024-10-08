// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PreferencesGetReader is a Reader for the PreferencesGet structure.
type PreferencesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PreferencesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPreferencesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPreferencesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPreferencesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPreferencesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewPreferencesGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPreferencesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/preferences] preferencesGet", response, response.Code())
	}
}

// NewPreferencesGetOK creates a PreferencesGetOK with default headers values
func NewPreferencesGetOK() *PreferencesGetOK {
	return &PreferencesGetOK{}
}

/*
PreferencesGetOK describes a response with status code 200, with default header values.

PreferencesGetOK preferences get o k
*/
type PreferencesGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this preferences get o k response has a 2xx status code
func (o *PreferencesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this preferences get o k response has a 3xx status code
func (o *PreferencesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this preferences get o k response has a 4xx status code
func (o *PreferencesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this preferences get o k response has a 5xx status code
func (o *PreferencesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this preferences get o k response a status code equal to that given
func (o *PreferencesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the preferences get o k response
func (o *PreferencesGetOK) Code() int {
	return 200
}

func (o *PreferencesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetOK %s", 200, payload)
}

func (o *PreferencesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetOK %s", 200, payload)
}

func (o *PreferencesGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PreferencesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreferencesGetBadRequest creates a PreferencesGetBadRequest with default headers values
func NewPreferencesGetBadRequest() *PreferencesGetBadRequest {
	return &PreferencesGetBadRequest{}
}

/*
PreferencesGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PreferencesGetBadRequest struct {
}

// IsSuccess returns true when this preferences get bad request response has a 2xx status code
func (o *PreferencesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this preferences get bad request response has a 3xx status code
func (o *PreferencesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this preferences get bad request response has a 4xx status code
func (o *PreferencesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this preferences get bad request response has a 5xx status code
func (o *PreferencesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this preferences get bad request response a status code equal to that given
func (o *PreferencesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the preferences get bad request response
func (o *PreferencesGetBadRequest) Code() int {
	return 400
}

func (o *PreferencesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetBadRequest", 400)
}

func (o *PreferencesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetBadRequest", 400)
}

func (o *PreferencesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPreferencesGetUnauthorized creates a PreferencesGetUnauthorized with default headers values
func NewPreferencesGetUnauthorized() *PreferencesGetUnauthorized {
	return &PreferencesGetUnauthorized{}
}

/*
PreferencesGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type PreferencesGetUnauthorized struct {
}

// IsSuccess returns true when this preferences get unauthorized response has a 2xx status code
func (o *PreferencesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this preferences get unauthorized response has a 3xx status code
func (o *PreferencesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this preferences get unauthorized response has a 4xx status code
func (o *PreferencesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this preferences get unauthorized response has a 5xx status code
func (o *PreferencesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this preferences get unauthorized response a status code equal to that given
func (o *PreferencesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the preferences get unauthorized response
func (o *PreferencesGetUnauthorized) Code() int {
	return 401
}

func (o *PreferencesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetUnauthorized", 401)
}

func (o *PreferencesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetUnauthorized", 401)
}

func (o *PreferencesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPreferencesGetNotFound creates a PreferencesGetNotFound with default headers values
func NewPreferencesGetNotFound() *PreferencesGetNotFound {
	return &PreferencesGetNotFound{}
}

/*
PreferencesGetNotFound describes a response with status code 404, with default header values.

not found
*/
type PreferencesGetNotFound struct {
}

// IsSuccess returns true when this preferences get not found response has a 2xx status code
func (o *PreferencesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this preferences get not found response has a 3xx status code
func (o *PreferencesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this preferences get not found response has a 4xx status code
func (o *PreferencesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this preferences get not found response has a 5xx status code
func (o *PreferencesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this preferences get not found response a status code equal to that given
func (o *PreferencesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the preferences get not found response
func (o *PreferencesGetNotFound) Code() int {
	return 404
}

func (o *PreferencesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetNotFound", 404)
}

func (o *PreferencesGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetNotFound", 404)
}

func (o *PreferencesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPreferencesGetNotAcceptable creates a PreferencesGetNotAcceptable with default headers values
func NewPreferencesGetNotAcceptable() *PreferencesGetNotAcceptable {
	return &PreferencesGetNotAcceptable{}
}

/*
PreferencesGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type PreferencesGetNotAcceptable struct {
}

// IsSuccess returns true when this preferences get not acceptable response has a 2xx status code
func (o *PreferencesGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this preferences get not acceptable response has a 3xx status code
func (o *PreferencesGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this preferences get not acceptable response has a 4xx status code
func (o *PreferencesGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this preferences get not acceptable response has a 5xx status code
func (o *PreferencesGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this preferences get not acceptable response a status code equal to that given
func (o *PreferencesGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the preferences get not acceptable response
func (o *PreferencesGetNotAcceptable) Code() int {
	return 406
}

func (o *PreferencesGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetNotAcceptable", 406)
}

func (o *PreferencesGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetNotAcceptable", 406)
}

func (o *PreferencesGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPreferencesGetInternalServerError creates a PreferencesGetInternalServerError with default headers values
func NewPreferencesGetInternalServerError() *PreferencesGetInternalServerError {
	return &PreferencesGetInternalServerError{}
}

/*
PreferencesGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type PreferencesGetInternalServerError struct {
}

// IsSuccess returns true when this preferences get internal server error response has a 2xx status code
func (o *PreferencesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this preferences get internal server error response has a 3xx status code
func (o *PreferencesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this preferences get internal server error response has a 4xx status code
func (o *PreferencesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this preferences get internal server error response has a 5xx status code
func (o *PreferencesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this preferences get internal server error response a status code equal to that given
func (o *PreferencesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the preferences get internal server error response
func (o *PreferencesGetInternalServerError) Code() int {
	return 500
}

func (o *PreferencesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetInternalServerError", 500)
}

func (o *PreferencesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/preferences][%d] preferencesGetInternalServerError", 500)
}

func (o *PreferencesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
