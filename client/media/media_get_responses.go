// Code generated by go-swagger; DO NOT EDIT.

package media

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

// MediaGetReader is a Reader for the MediaGet structure.
type MediaGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMediaGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMediaGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMediaGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewMediaGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMediaGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/media/{id}] mediaGet", response, response.Code())
	}
}

// NewMediaGetOK creates a MediaGetOK with default headers values
func NewMediaGetOK() *MediaGetOK {
	return &MediaGetOK{}
}

/*
MediaGetOK describes a response with status code 200, with default header values.

The requested media attachment.
*/
type MediaGetOK struct {
	Payload *models.Attachment
}

// IsSuccess returns true when this media get o k response has a 2xx status code
func (o *MediaGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this media get o k response has a 3xx status code
func (o *MediaGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media get o k response has a 4xx status code
func (o *MediaGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this media get o k response has a 5xx status code
func (o *MediaGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this media get o k response a status code equal to that given
func (o *MediaGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the media get o k response
func (o *MediaGetOK) Code() int {
	return 200
}

func (o *MediaGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetOK %s", 200, payload)
}

func (o *MediaGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetOK %s", 200, payload)
}

func (o *MediaGetOK) GetPayload() *models.Attachment {
	return o.Payload
}

func (o *MediaGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Attachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaGetBadRequest creates a MediaGetBadRequest with default headers values
func NewMediaGetBadRequest() *MediaGetBadRequest {
	return &MediaGetBadRequest{}
}

/*
MediaGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type MediaGetBadRequest struct {
}

// IsSuccess returns true when this media get bad request response has a 2xx status code
func (o *MediaGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this media get bad request response has a 3xx status code
func (o *MediaGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media get bad request response has a 4xx status code
func (o *MediaGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this media get bad request response has a 5xx status code
func (o *MediaGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this media get bad request response a status code equal to that given
func (o *MediaGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the media get bad request response
func (o *MediaGetBadRequest) Code() int {
	return 400
}

func (o *MediaGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetBadRequest", 400)
}

func (o *MediaGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetBadRequest", 400)
}

func (o *MediaGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaGetUnauthorized creates a MediaGetUnauthorized with default headers values
func NewMediaGetUnauthorized() *MediaGetUnauthorized {
	return &MediaGetUnauthorized{}
}

/*
MediaGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type MediaGetUnauthorized struct {
}

// IsSuccess returns true when this media get unauthorized response has a 2xx status code
func (o *MediaGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this media get unauthorized response has a 3xx status code
func (o *MediaGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media get unauthorized response has a 4xx status code
func (o *MediaGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this media get unauthorized response has a 5xx status code
func (o *MediaGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this media get unauthorized response a status code equal to that given
func (o *MediaGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the media get unauthorized response
func (o *MediaGetUnauthorized) Code() int {
	return 401
}

func (o *MediaGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetUnauthorized", 401)
}

func (o *MediaGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetUnauthorized", 401)
}

func (o *MediaGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaGetNotFound creates a MediaGetNotFound with default headers values
func NewMediaGetNotFound() *MediaGetNotFound {
	return &MediaGetNotFound{}
}

/*
MediaGetNotFound describes a response with status code 404, with default header values.

not found
*/
type MediaGetNotFound struct {
}

// IsSuccess returns true when this media get not found response has a 2xx status code
func (o *MediaGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this media get not found response has a 3xx status code
func (o *MediaGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media get not found response has a 4xx status code
func (o *MediaGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this media get not found response has a 5xx status code
func (o *MediaGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this media get not found response a status code equal to that given
func (o *MediaGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the media get not found response
func (o *MediaGetNotFound) Code() int {
	return 404
}

func (o *MediaGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetNotFound", 404)
}

func (o *MediaGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetNotFound", 404)
}

func (o *MediaGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaGetNotAcceptable creates a MediaGetNotAcceptable with default headers values
func NewMediaGetNotAcceptable() *MediaGetNotAcceptable {
	return &MediaGetNotAcceptable{}
}

/*
MediaGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type MediaGetNotAcceptable struct {
}

// IsSuccess returns true when this media get not acceptable response has a 2xx status code
func (o *MediaGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this media get not acceptable response has a 3xx status code
func (o *MediaGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media get not acceptable response has a 4xx status code
func (o *MediaGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this media get not acceptable response has a 5xx status code
func (o *MediaGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this media get not acceptable response a status code equal to that given
func (o *MediaGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the media get not acceptable response
func (o *MediaGetNotAcceptable) Code() int {
	return 406
}

func (o *MediaGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetNotAcceptable", 406)
}

func (o *MediaGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetNotAcceptable", 406)
}

func (o *MediaGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMediaGetInternalServerError creates a MediaGetInternalServerError with default headers values
func NewMediaGetInternalServerError() *MediaGetInternalServerError {
	return &MediaGetInternalServerError{}
}

/*
MediaGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type MediaGetInternalServerError struct {
}

// IsSuccess returns true when this media get internal server error response has a 2xx status code
func (o *MediaGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this media get internal server error response has a 3xx status code
func (o *MediaGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media get internal server error response has a 4xx status code
func (o *MediaGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this media get internal server error response has a 5xx status code
func (o *MediaGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this media get internal server error response a status code equal to that given
func (o *MediaGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the media get internal server error response
func (o *MediaGetInternalServerError) Code() int {
	return 500
}

func (o *MediaGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetInternalServerError", 500)
}

func (o *MediaGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/media/{id}][%d] mediaGetInternalServerError", 500)
}

func (o *MediaGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}