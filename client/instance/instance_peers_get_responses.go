// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// InstancePeersGetReader is a Reader for the InstancePeersGet structure.
type InstancePeersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstancePeersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstancePeersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstancePeersGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInstancePeersGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstancePeersGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstancePeersGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewInstancePeersGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInstancePeersGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/instance/peers] instancePeersGet", response, response.Code())
	}
}

// NewInstancePeersGetOK creates a InstancePeersGetOK with default headers values
func NewInstancePeersGetOK() *InstancePeersGetOK {
	return &InstancePeersGetOK{}
}

/*
	InstancePeersGetOK describes a response with status code 200, with default header values.

	If no filter parameter is provided, or filter is empty, then a legacy, Mastodon-API compatible response will be returned. This will consist of just a 'flat' array of strings like `["example.com", "example.org"]`, which corresponds to domains this instance peers with.

If a filter parameter is provided, then an array of objects with at least a `domain` key set on each object will be returned.

Domains that are silenced or suspended will also have a key `suspended_at` or `silenced_at` that contains an iso8601 date string. If one of these keys is not present on the domain object, it is open. Suspended instances may in some cases be obfuscated, which means they will have some letters replaced by `*` to make it more difficult for bad actors to target instances with harassment.

Whether a flat response or a more detailed response is returned, domains will be sorted alphabetically by hostname.
*/
type InstancePeersGetOK struct {
	Payload []*models.Domain
}

// IsSuccess returns true when this instance peers get o k response has a 2xx status code
func (o *InstancePeersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instance peers get o k response has a 3xx status code
func (o *InstancePeersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get o k response has a 4xx status code
func (o *InstancePeersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance peers get o k response has a 5xx status code
func (o *InstancePeersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this instance peers get o k response a status code equal to that given
func (o *InstancePeersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the instance peers get o k response
func (o *InstancePeersGetOK) Code() int {
	return 200
}

func (o *InstancePeersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetOK %s", 200, payload)
}

func (o *InstancePeersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetOK %s", 200, payload)
}

func (o *InstancePeersGetOK) GetPayload() []*models.Domain {
	return o.Payload
}

func (o *InstancePeersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstancePeersGetBadRequest creates a InstancePeersGetBadRequest with default headers values
func NewInstancePeersGetBadRequest() *InstancePeersGetBadRequest {
	return &InstancePeersGetBadRequest{}
}

/*
InstancePeersGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type InstancePeersGetBadRequest struct {
}

// IsSuccess returns true when this instance peers get bad request response has a 2xx status code
func (o *InstancePeersGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance peers get bad request response has a 3xx status code
func (o *InstancePeersGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get bad request response has a 4xx status code
func (o *InstancePeersGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance peers get bad request response has a 5xx status code
func (o *InstancePeersGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this instance peers get bad request response a status code equal to that given
func (o *InstancePeersGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the instance peers get bad request response
func (o *InstancePeersGetBadRequest) Code() int {
	return 400
}

func (o *InstancePeersGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetBadRequest", 400)
}

func (o *InstancePeersGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetBadRequest", 400)
}

func (o *InstancePeersGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancePeersGetUnauthorized creates a InstancePeersGetUnauthorized with default headers values
func NewInstancePeersGetUnauthorized() *InstancePeersGetUnauthorized {
	return &InstancePeersGetUnauthorized{}
}

/*
InstancePeersGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type InstancePeersGetUnauthorized struct {
}

// IsSuccess returns true when this instance peers get unauthorized response has a 2xx status code
func (o *InstancePeersGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance peers get unauthorized response has a 3xx status code
func (o *InstancePeersGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get unauthorized response has a 4xx status code
func (o *InstancePeersGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance peers get unauthorized response has a 5xx status code
func (o *InstancePeersGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this instance peers get unauthorized response a status code equal to that given
func (o *InstancePeersGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the instance peers get unauthorized response
func (o *InstancePeersGetUnauthorized) Code() int {
	return 401
}

func (o *InstancePeersGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetUnauthorized", 401)
}

func (o *InstancePeersGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetUnauthorized", 401)
}

func (o *InstancePeersGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancePeersGetForbidden creates a InstancePeersGetForbidden with default headers values
func NewInstancePeersGetForbidden() *InstancePeersGetForbidden {
	return &InstancePeersGetForbidden{}
}

/*
InstancePeersGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type InstancePeersGetForbidden struct {
}

// IsSuccess returns true when this instance peers get forbidden response has a 2xx status code
func (o *InstancePeersGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance peers get forbidden response has a 3xx status code
func (o *InstancePeersGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get forbidden response has a 4xx status code
func (o *InstancePeersGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance peers get forbidden response has a 5xx status code
func (o *InstancePeersGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this instance peers get forbidden response a status code equal to that given
func (o *InstancePeersGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the instance peers get forbidden response
func (o *InstancePeersGetForbidden) Code() int {
	return 403
}

func (o *InstancePeersGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetForbidden", 403)
}

func (o *InstancePeersGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetForbidden", 403)
}

func (o *InstancePeersGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancePeersGetNotFound creates a InstancePeersGetNotFound with default headers values
func NewInstancePeersGetNotFound() *InstancePeersGetNotFound {
	return &InstancePeersGetNotFound{}
}

/*
InstancePeersGetNotFound describes a response with status code 404, with default header values.

not found
*/
type InstancePeersGetNotFound struct {
}

// IsSuccess returns true when this instance peers get not found response has a 2xx status code
func (o *InstancePeersGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance peers get not found response has a 3xx status code
func (o *InstancePeersGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get not found response has a 4xx status code
func (o *InstancePeersGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance peers get not found response has a 5xx status code
func (o *InstancePeersGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this instance peers get not found response a status code equal to that given
func (o *InstancePeersGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the instance peers get not found response
func (o *InstancePeersGetNotFound) Code() int {
	return 404
}

func (o *InstancePeersGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetNotFound", 404)
}

func (o *InstancePeersGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetNotFound", 404)
}

func (o *InstancePeersGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancePeersGetNotAcceptable creates a InstancePeersGetNotAcceptable with default headers values
func NewInstancePeersGetNotAcceptable() *InstancePeersGetNotAcceptable {
	return &InstancePeersGetNotAcceptable{}
}

/*
InstancePeersGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type InstancePeersGetNotAcceptable struct {
}

// IsSuccess returns true when this instance peers get not acceptable response has a 2xx status code
func (o *InstancePeersGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance peers get not acceptable response has a 3xx status code
func (o *InstancePeersGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get not acceptable response has a 4xx status code
func (o *InstancePeersGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this instance peers get not acceptable response has a 5xx status code
func (o *InstancePeersGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this instance peers get not acceptable response a status code equal to that given
func (o *InstancePeersGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the instance peers get not acceptable response
func (o *InstancePeersGetNotAcceptable) Code() int {
	return 406
}

func (o *InstancePeersGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetNotAcceptable", 406)
}

func (o *InstancePeersGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetNotAcceptable", 406)
}

func (o *InstancePeersGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstancePeersGetInternalServerError creates a InstancePeersGetInternalServerError with default headers values
func NewInstancePeersGetInternalServerError() *InstancePeersGetInternalServerError {
	return &InstancePeersGetInternalServerError{}
}

/*
InstancePeersGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type InstancePeersGetInternalServerError struct {
}

// IsSuccess returns true when this instance peers get internal server error response has a 2xx status code
func (o *InstancePeersGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instance peers get internal server error response has a 3xx status code
func (o *InstancePeersGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instance peers get internal server error response has a 4xx status code
func (o *InstancePeersGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this instance peers get internal server error response has a 5xx status code
func (o *InstancePeersGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this instance peers get internal server error response a status code equal to that given
func (o *InstancePeersGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the instance peers get internal server error response
func (o *InstancePeersGetInternalServerError) Code() int {
	return 500
}

func (o *InstancePeersGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetInternalServerError", 500)
}

func (o *InstancePeersGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/instance/peers][%d] instancePeersGetInternalServerError", 500)
}

func (o *InstancePeersGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
