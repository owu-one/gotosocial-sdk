// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// DebugAPURLReader is a Reader for the DebugAPURL structure.
type DebugAPURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DebugAPURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDebugAPURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDebugAPURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDebugAPURLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDebugAPURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDebugAPURLNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDebugAPURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/debug/apurl] debugAPUrl", response, response.Code())
	}
}

// NewDebugAPURLOK creates a DebugAPURLOK with default headers values
func NewDebugAPURLOK() *DebugAPURLOK {
	return &DebugAPURLOK{}
}

/*
DebugAPURLOK describes a response with status code 200, with default header values.

DebugAPURLOK debug a p Url o k
*/
type DebugAPURLOK struct {
	Payload *models.DebugAPURLResponse
}

// IsSuccess returns true when this debug a p Url o k response has a 2xx status code
func (o *DebugAPURLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this debug a p Url o k response has a 3xx status code
func (o *DebugAPURLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this debug a p Url o k response has a 4xx status code
func (o *DebugAPURLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this debug a p Url o k response has a 5xx status code
func (o *DebugAPURLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this debug a p Url o k response a status code equal to that given
func (o *DebugAPURLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the debug a p Url o k response
func (o *DebugAPURLOK) Code() int {
	return 200
}

func (o *DebugAPURLOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlOK %s", 200, payload)
}

func (o *DebugAPURLOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlOK %s", 200, payload)
}

func (o *DebugAPURLOK) GetPayload() *models.DebugAPURLResponse {
	return o.Payload
}

func (o *DebugAPURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugAPURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebugAPURLBadRequest creates a DebugAPURLBadRequest with default headers values
func NewDebugAPURLBadRequest() *DebugAPURLBadRequest {
	return &DebugAPURLBadRequest{}
}

/*
DebugAPURLBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DebugAPURLBadRequest struct {
}

// IsSuccess returns true when this debug a p Url bad request response has a 2xx status code
func (o *DebugAPURLBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this debug a p Url bad request response has a 3xx status code
func (o *DebugAPURLBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this debug a p Url bad request response has a 4xx status code
func (o *DebugAPURLBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this debug a p Url bad request response has a 5xx status code
func (o *DebugAPURLBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this debug a p Url bad request response a status code equal to that given
func (o *DebugAPURLBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the debug a p Url bad request response
func (o *DebugAPURLBadRequest) Code() int {
	return 400
}

func (o *DebugAPURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlBadRequest", 400)
}

func (o *DebugAPURLBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlBadRequest", 400)
}

func (o *DebugAPURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebugAPURLUnauthorized creates a DebugAPURLUnauthorized with default headers values
func NewDebugAPURLUnauthorized() *DebugAPURLUnauthorized {
	return &DebugAPURLUnauthorized{}
}

/*
DebugAPURLUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DebugAPURLUnauthorized struct {
}

// IsSuccess returns true when this debug a p Url unauthorized response has a 2xx status code
func (o *DebugAPURLUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this debug a p Url unauthorized response has a 3xx status code
func (o *DebugAPURLUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this debug a p Url unauthorized response has a 4xx status code
func (o *DebugAPURLUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this debug a p Url unauthorized response has a 5xx status code
func (o *DebugAPURLUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this debug a p Url unauthorized response a status code equal to that given
func (o *DebugAPURLUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the debug a p Url unauthorized response
func (o *DebugAPURLUnauthorized) Code() int {
	return 401
}

func (o *DebugAPURLUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlUnauthorized", 401)
}

func (o *DebugAPURLUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlUnauthorized", 401)
}

func (o *DebugAPURLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebugAPURLNotFound creates a DebugAPURLNotFound with default headers values
func NewDebugAPURLNotFound() *DebugAPURLNotFound {
	return &DebugAPURLNotFound{}
}

/*
DebugAPURLNotFound describes a response with status code 404, with default header values.

not found
*/
type DebugAPURLNotFound struct {
}

// IsSuccess returns true when this debug a p Url not found response has a 2xx status code
func (o *DebugAPURLNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this debug a p Url not found response has a 3xx status code
func (o *DebugAPURLNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this debug a p Url not found response has a 4xx status code
func (o *DebugAPURLNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this debug a p Url not found response has a 5xx status code
func (o *DebugAPURLNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this debug a p Url not found response a status code equal to that given
func (o *DebugAPURLNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the debug a p Url not found response
func (o *DebugAPURLNotFound) Code() int {
	return 404
}

func (o *DebugAPURLNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlNotFound", 404)
}

func (o *DebugAPURLNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlNotFound", 404)
}

func (o *DebugAPURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebugAPURLNotAcceptable creates a DebugAPURLNotAcceptable with default headers values
func NewDebugAPURLNotAcceptable() *DebugAPURLNotAcceptable {
	return &DebugAPURLNotAcceptable{}
}

/*
DebugAPURLNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DebugAPURLNotAcceptable struct {
}

// IsSuccess returns true when this debug a p Url not acceptable response has a 2xx status code
func (o *DebugAPURLNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this debug a p Url not acceptable response has a 3xx status code
func (o *DebugAPURLNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this debug a p Url not acceptable response has a 4xx status code
func (o *DebugAPURLNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this debug a p Url not acceptable response has a 5xx status code
func (o *DebugAPURLNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this debug a p Url not acceptable response a status code equal to that given
func (o *DebugAPURLNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the debug a p Url not acceptable response
func (o *DebugAPURLNotAcceptable) Code() int {
	return 406
}

func (o *DebugAPURLNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlNotAcceptable", 406)
}

func (o *DebugAPURLNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlNotAcceptable", 406)
}

func (o *DebugAPURLNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDebugAPURLInternalServerError creates a DebugAPURLInternalServerError with default headers values
func NewDebugAPURLInternalServerError() *DebugAPURLInternalServerError {
	return &DebugAPURLInternalServerError{}
}

/*
DebugAPURLInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DebugAPURLInternalServerError struct {
}

// IsSuccess returns true when this debug a p Url internal server error response has a 2xx status code
func (o *DebugAPURLInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this debug a p Url internal server error response has a 3xx status code
func (o *DebugAPURLInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this debug a p Url internal server error response has a 4xx status code
func (o *DebugAPURLInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this debug a p Url internal server error response has a 5xx status code
func (o *DebugAPURLInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this debug a p Url internal server error response a status code equal to that given
func (o *DebugAPURLInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the debug a p Url internal server error response
func (o *DebugAPURLInternalServerError) Code() int {
	return 500
}

func (o *DebugAPURLInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlInternalServerError", 500)
}

func (o *DebugAPURLInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/debug/apurl][%d] debugAPUrlInternalServerError", 500)
}

func (o *DebugAPURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}