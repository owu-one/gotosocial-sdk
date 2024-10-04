// Code generated by go-swagger; DO NOT EDIT.

package import_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExportFollowersReader is a Reader for the ExportFollowers structure.
type ExportFollowersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportFollowersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportFollowersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportFollowersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewExportFollowersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExportFollowersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/exports/followers.csv] exportFollowers", response, response.Code())
	}
}

// NewExportFollowersOK creates a ExportFollowersOK with default headers values
func NewExportFollowersOK() *ExportFollowersOK {
	return &ExportFollowersOK{}
}

/*
ExportFollowersOK describes a response with status code 200, with default header values.

CSV file of accounts that follow you.
*/
type ExportFollowersOK struct {
}

// IsSuccess returns true when this export followers o k response has a 2xx status code
func (o *ExportFollowersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export followers o k response has a 3xx status code
func (o *ExportFollowersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export followers o k response has a 4xx status code
func (o *ExportFollowersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export followers o k response has a 5xx status code
func (o *ExportFollowersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export followers o k response a status code equal to that given
func (o *ExportFollowersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export followers o k response
func (o *ExportFollowersOK) Code() int {
	return 200
}

func (o *ExportFollowersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersOK", 200)
}

func (o *ExportFollowersOK) String() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersOK", 200)
}

func (o *ExportFollowersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFollowersUnauthorized creates a ExportFollowersUnauthorized with default headers values
func NewExportFollowersUnauthorized() *ExportFollowersUnauthorized {
	return &ExportFollowersUnauthorized{}
}

/*
ExportFollowersUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ExportFollowersUnauthorized struct {
}

// IsSuccess returns true when this export followers unauthorized response has a 2xx status code
func (o *ExportFollowersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export followers unauthorized response has a 3xx status code
func (o *ExportFollowersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export followers unauthorized response has a 4xx status code
func (o *ExportFollowersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export followers unauthorized response has a 5xx status code
func (o *ExportFollowersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export followers unauthorized response a status code equal to that given
func (o *ExportFollowersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the export followers unauthorized response
func (o *ExportFollowersUnauthorized) Code() int {
	return 401
}

func (o *ExportFollowersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersUnauthorized", 401)
}

func (o *ExportFollowersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersUnauthorized", 401)
}

func (o *ExportFollowersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFollowersNotAcceptable creates a ExportFollowersNotAcceptable with default headers values
func NewExportFollowersNotAcceptable() *ExportFollowersNotAcceptable {
	return &ExportFollowersNotAcceptable{}
}

/*
ExportFollowersNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ExportFollowersNotAcceptable struct {
}

// IsSuccess returns true when this export followers not acceptable response has a 2xx status code
func (o *ExportFollowersNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export followers not acceptable response has a 3xx status code
func (o *ExportFollowersNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export followers not acceptable response has a 4xx status code
func (o *ExportFollowersNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this export followers not acceptable response has a 5xx status code
func (o *ExportFollowersNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this export followers not acceptable response a status code equal to that given
func (o *ExportFollowersNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the export followers not acceptable response
func (o *ExportFollowersNotAcceptable) Code() int {
	return 406
}

func (o *ExportFollowersNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersNotAcceptable", 406)
}

func (o *ExportFollowersNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersNotAcceptable", 406)
}

func (o *ExportFollowersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFollowersInternalServerError creates a ExportFollowersInternalServerError with default headers values
func NewExportFollowersInternalServerError() *ExportFollowersInternalServerError {
	return &ExportFollowersInternalServerError{}
}

/*
ExportFollowersInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ExportFollowersInternalServerError struct {
}

// IsSuccess returns true when this export followers internal server error response has a 2xx status code
func (o *ExportFollowersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export followers internal server error response has a 3xx status code
func (o *ExportFollowersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export followers internal server error response has a 4xx status code
func (o *ExportFollowersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this export followers internal server error response has a 5xx status code
func (o *ExportFollowersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this export followers internal server error response a status code equal to that given
func (o *ExportFollowersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the export followers internal server error response
func (o *ExportFollowersInternalServerError) Code() int {
	return 500
}

func (o *ExportFollowersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersInternalServerError", 500)
}

func (o *ExportFollowersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/exports/followers.csv][%d] exportFollowersInternalServerError", 500)
}

func (o *ExportFollowersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}