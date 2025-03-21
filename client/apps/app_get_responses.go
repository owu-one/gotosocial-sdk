// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// AppGetReader is a Reader for the AppGet structure.
type AppGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAppGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/apps/{id}] appGet", response, response.Code())
	}
}

// NewAppGetOK creates a AppGetOK with default headers values
func NewAppGetOK() *AppGetOK {
	return &AppGetOK{}
}

/*
AppGetOK describes a response with status code 200, with default header values.

The requested application.
*/
type AppGetOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this app get o k response has a 2xx status code
func (o *AppGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app get o k response has a 3xx status code
func (o *AppGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app get o k response has a 4xx status code
func (o *AppGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app get o k response has a 5xx status code
func (o *AppGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app get o k response a status code equal to that given
func (o *AppGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the app get o k response
func (o *AppGetOK) Code() int {
	return 200
}

func (o *AppGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetOK %s", 200, payload)
}

func (o *AppGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetOK %s", 200, payload)
}

func (o *AppGetOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *AppGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppGetBadRequest creates a AppGetBadRequest with default headers values
func NewAppGetBadRequest() *AppGetBadRequest {
	return &AppGetBadRequest{}
}

/*
AppGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AppGetBadRequest struct {
}

// IsSuccess returns true when this app get bad request response has a 2xx status code
func (o *AppGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app get bad request response has a 3xx status code
func (o *AppGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app get bad request response has a 4xx status code
func (o *AppGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this app get bad request response has a 5xx status code
func (o *AppGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this app get bad request response a status code equal to that given
func (o *AppGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the app get bad request response
func (o *AppGetBadRequest) Code() int {
	return 400
}

func (o *AppGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetBadRequest", 400)
}

func (o *AppGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetBadRequest", 400)
}

func (o *AppGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppGetUnauthorized creates a AppGetUnauthorized with default headers values
func NewAppGetUnauthorized() *AppGetUnauthorized {
	return &AppGetUnauthorized{}
}

/*
AppGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AppGetUnauthorized struct {
}

// IsSuccess returns true when this app get unauthorized response has a 2xx status code
func (o *AppGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app get unauthorized response has a 3xx status code
func (o *AppGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app get unauthorized response has a 4xx status code
func (o *AppGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this app get unauthorized response has a 5xx status code
func (o *AppGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this app get unauthorized response a status code equal to that given
func (o *AppGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the app get unauthorized response
func (o *AppGetUnauthorized) Code() int {
	return 401
}

func (o *AppGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetUnauthorized", 401)
}

func (o *AppGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetUnauthorized", 401)
}

func (o *AppGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppGetNotFound creates a AppGetNotFound with default headers values
func NewAppGetNotFound() *AppGetNotFound {
	return &AppGetNotFound{}
}

/*
AppGetNotFound describes a response with status code 404, with default header values.

not found
*/
type AppGetNotFound struct {
}

// IsSuccess returns true when this app get not found response has a 2xx status code
func (o *AppGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app get not found response has a 3xx status code
func (o *AppGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app get not found response has a 4xx status code
func (o *AppGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this app get not found response has a 5xx status code
func (o *AppGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this app get not found response a status code equal to that given
func (o *AppGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the app get not found response
func (o *AppGetNotFound) Code() int {
	return 404
}

func (o *AppGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetNotFound", 404)
}

func (o *AppGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetNotFound", 404)
}

func (o *AppGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppGetNotAcceptable creates a AppGetNotAcceptable with default headers values
func NewAppGetNotAcceptable() *AppGetNotAcceptable {
	return &AppGetNotAcceptable{}
}

/*
AppGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AppGetNotAcceptable struct {
}

// IsSuccess returns true when this app get not acceptable response has a 2xx status code
func (o *AppGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app get not acceptable response has a 3xx status code
func (o *AppGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app get not acceptable response has a 4xx status code
func (o *AppGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this app get not acceptable response has a 5xx status code
func (o *AppGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this app get not acceptable response a status code equal to that given
func (o *AppGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the app get not acceptable response
func (o *AppGetNotAcceptable) Code() int {
	return 406
}

func (o *AppGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetNotAcceptable", 406)
}

func (o *AppGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetNotAcceptable", 406)
}

func (o *AppGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAppGetInternalServerError creates a AppGetInternalServerError with default headers values
func NewAppGetInternalServerError() *AppGetInternalServerError {
	return &AppGetInternalServerError{}
}

/*
AppGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AppGetInternalServerError struct {
}

// IsSuccess returns true when this app get internal server error response has a 2xx status code
func (o *AppGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app get internal server error response has a 3xx status code
func (o *AppGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app get internal server error response has a 4xx status code
func (o *AppGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this app get internal server error response has a 5xx status code
func (o *AppGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this app get internal server error response a status code equal to that given
func (o *AppGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the app get internal server error response
func (o *AppGetInternalServerError) Code() int {
	return 500
}

func (o *AppGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetInternalServerError", 500)
}

func (o *AppGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/apps/{id}][%d] appGetInternalServerError", 500)
}

func (o *AppGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
