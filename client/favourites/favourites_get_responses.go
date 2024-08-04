// Code generated by go-swagger; DO NOT EDIT.

package favourites

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

// FavouritesGetReader is a Reader for the FavouritesGet structure.
type FavouritesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FavouritesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFavouritesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFavouritesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFavouritesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFavouritesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewFavouritesGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFavouritesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/favourites] favouritesGet", response, response.Code())
	}
}

// NewFavouritesGetOK creates a FavouritesGetOK with default headers values
func NewFavouritesGetOK() *FavouritesGetOK {
	return &FavouritesGetOK{}
}

/*
FavouritesGetOK describes a response with status code 200, with default header values.

FavouritesGetOK favourites get o k
*/
type FavouritesGetOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Status
}

// IsSuccess returns true when this favourites get o k response has a 2xx status code
func (o *FavouritesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this favourites get o k response has a 3xx status code
func (o *FavouritesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this favourites get o k response has a 4xx status code
func (o *FavouritesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this favourites get o k response has a 5xx status code
func (o *FavouritesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this favourites get o k response a status code equal to that given
func (o *FavouritesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the favourites get o k response
func (o *FavouritesGetOK) Code() int {
	return 200
}

func (o *FavouritesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetOK %s", 200, payload)
}

func (o *FavouritesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetOK %s", 200, payload)
}

func (o *FavouritesGetOK) GetPayload() []*models.Status {
	return o.Payload
}

func (o *FavouritesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFavouritesGetBadRequest creates a FavouritesGetBadRequest with default headers values
func NewFavouritesGetBadRequest() *FavouritesGetBadRequest {
	return &FavouritesGetBadRequest{}
}

/*
FavouritesGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type FavouritesGetBadRequest struct {
}

// IsSuccess returns true when this favourites get bad request response has a 2xx status code
func (o *FavouritesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this favourites get bad request response has a 3xx status code
func (o *FavouritesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this favourites get bad request response has a 4xx status code
func (o *FavouritesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this favourites get bad request response has a 5xx status code
func (o *FavouritesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this favourites get bad request response a status code equal to that given
func (o *FavouritesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the favourites get bad request response
func (o *FavouritesGetBadRequest) Code() int {
	return 400
}

func (o *FavouritesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetBadRequest", 400)
}

func (o *FavouritesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetBadRequest", 400)
}

func (o *FavouritesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFavouritesGetUnauthorized creates a FavouritesGetUnauthorized with default headers values
func NewFavouritesGetUnauthorized() *FavouritesGetUnauthorized {
	return &FavouritesGetUnauthorized{}
}

/*
FavouritesGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FavouritesGetUnauthorized struct {
}

// IsSuccess returns true when this favourites get unauthorized response has a 2xx status code
func (o *FavouritesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this favourites get unauthorized response has a 3xx status code
func (o *FavouritesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this favourites get unauthorized response has a 4xx status code
func (o *FavouritesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this favourites get unauthorized response has a 5xx status code
func (o *FavouritesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this favourites get unauthorized response a status code equal to that given
func (o *FavouritesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the favourites get unauthorized response
func (o *FavouritesGetUnauthorized) Code() int {
	return 401
}

func (o *FavouritesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetUnauthorized", 401)
}

func (o *FavouritesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetUnauthorized", 401)
}

func (o *FavouritesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFavouritesGetNotFound creates a FavouritesGetNotFound with default headers values
func NewFavouritesGetNotFound() *FavouritesGetNotFound {
	return &FavouritesGetNotFound{}
}

/*
FavouritesGetNotFound describes a response with status code 404, with default header values.

not found
*/
type FavouritesGetNotFound struct {
}

// IsSuccess returns true when this favourites get not found response has a 2xx status code
func (o *FavouritesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this favourites get not found response has a 3xx status code
func (o *FavouritesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this favourites get not found response has a 4xx status code
func (o *FavouritesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this favourites get not found response has a 5xx status code
func (o *FavouritesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this favourites get not found response a status code equal to that given
func (o *FavouritesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the favourites get not found response
func (o *FavouritesGetNotFound) Code() int {
	return 404
}

func (o *FavouritesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetNotFound", 404)
}

func (o *FavouritesGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetNotFound", 404)
}

func (o *FavouritesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFavouritesGetNotAcceptable creates a FavouritesGetNotAcceptable with default headers values
func NewFavouritesGetNotAcceptable() *FavouritesGetNotAcceptable {
	return &FavouritesGetNotAcceptable{}
}

/*
FavouritesGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type FavouritesGetNotAcceptable struct {
}

// IsSuccess returns true when this favourites get not acceptable response has a 2xx status code
func (o *FavouritesGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this favourites get not acceptable response has a 3xx status code
func (o *FavouritesGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this favourites get not acceptable response has a 4xx status code
func (o *FavouritesGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this favourites get not acceptable response has a 5xx status code
func (o *FavouritesGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this favourites get not acceptable response a status code equal to that given
func (o *FavouritesGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the favourites get not acceptable response
func (o *FavouritesGetNotAcceptable) Code() int {
	return 406
}

func (o *FavouritesGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetNotAcceptable", 406)
}

func (o *FavouritesGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetNotAcceptable", 406)
}

func (o *FavouritesGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFavouritesGetInternalServerError creates a FavouritesGetInternalServerError with default headers values
func NewFavouritesGetInternalServerError() *FavouritesGetInternalServerError {
	return &FavouritesGetInternalServerError{}
}

/*
FavouritesGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type FavouritesGetInternalServerError struct {
}

// IsSuccess returns true when this favourites get internal server error response has a 2xx status code
func (o *FavouritesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this favourites get internal server error response has a 3xx status code
func (o *FavouritesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this favourites get internal server error response has a 4xx status code
func (o *FavouritesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this favourites get internal server error response has a 5xx status code
func (o *FavouritesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this favourites get internal server error response a status code equal to that given
func (o *FavouritesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the favourites get internal server error response
func (o *FavouritesGetInternalServerError) Code() int {
	return 500
}

func (o *FavouritesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetInternalServerError", 500)
}

func (o *FavouritesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/favourites][%d] favouritesGetInternalServerError", 500)
}

func (o *FavouritesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}