// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// ReportCreateReader is a Reader for the ReportCreate structure.
type ReportCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReportCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReportCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReportCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewReportCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReportCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/reports] reportCreate", response, response.Code())
	}
}

// NewReportCreateOK creates a ReportCreateOK with default headers values
func NewReportCreateOK() *ReportCreateOK {
	return &ReportCreateOK{}
}

/*
ReportCreateOK describes a response with status code 200, with default header values.

The created report.
*/
type ReportCreateOK struct {
	Payload *models.Report
}

// IsSuccess returns true when this report create o k response has a 2xx status code
func (o *ReportCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report create o k response has a 3xx status code
func (o *ReportCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report create o k response has a 4xx status code
func (o *ReportCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report create o k response has a 5xx status code
func (o *ReportCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report create o k response a status code equal to that given
func (o *ReportCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the report create o k response
func (o *ReportCreateOK) Code() int {
	return 200
}

func (o *ReportCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateOK %s", 200, payload)
}

func (o *ReportCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateOK %s", 200, payload)
}

func (o *ReportCreateOK) GetPayload() *models.Report {
	return o.Payload
}

func (o *ReportCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportCreateBadRequest creates a ReportCreateBadRequest with default headers values
func NewReportCreateBadRequest() *ReportCreateBadRequest {
	return &ReportCreateBadRequest{}
}

/*
ReportCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ReportCreateBadRequest struct {
}

// IsSuccess returns true when this report create bad request response has a 2xx status code
func (o *ReportCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report create bad request response has a 3xx status code
func (o *ReportCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report create bad request response has a 4xx status code
func (o *ReportCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this report create bad request response has a 5xx status code
func (o *ReportCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this report create bad request response a status code equal to that given
func (o *ReportCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the report create bad request response
func (o *ReportCreateBadRequest) Code() int {
	return 400
}

func (o *ReportCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateBadRequest", 400)
}

func (o *ReportCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateBadRequest", 400)
}

func (o *ReportCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportCreateUnauthorized creates a ReportCreateUnauthorized with default headers values
func NewReportCreateUnauthorized() *ReportCreateUnauthorized {
	return &ReportCreateUnauthorized{}
}

/*
ReportCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ReportCreateUnauthorized struct {
}

// IsSuccess returns true when this report create unauthorized response has a 2xx status code
func (o *ReportCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report create unauthorized response has a 3xx status code
func (o *ReportCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report create unauthorized response has a 4xx status code
func (o *ReportCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this report create unauthorized response has a 5xx status code
func (o *ReportCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this report create unauthorized response a status code equal to that given
func (o *ReportCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the report create unauthorized response
func (o *ReportCreateUnauthorized) Code() int {
	return 401
}

func (o *ReportCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateUnauthorized", 401)
}

func (o *ReportCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateUnauthorized", 401)
}

func (o *ReportCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportCreateNotFound creates a ReportCreateNotFound with default headers values
func NewReportCreateNotFound() *ReportCreateNotFound {
	return &ReportCreateNotFound{}
}

/*
ReportCreateNotFound describes a response with status code 404, with default header values.

not found
*/
type ReportCreateNotFound struct {
}

// IsSuccess returns true when this report create not found response has a 2xx status code
func (o *ReportCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report create not found response has a 3xx status code
func (o *ReportCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report create not found response has a 4xx status code
func (o *ReportCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this report create not found response has a 5xx status code
func (o *ReportCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this report create not found response a status code equal to that given
func (o *ReportCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the report create not found response
func (o *ReportCreateNotFound) Code() int {
	return 404
}

func (o *ReportCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateNotFound", 404)
}

func (o *ReportCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateNotFound", 404)
}

func (o *ReportCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportCreateNotAcceptable creates a ReportCreateNotAcceptable with default headers values
func NewReportCreateNotAcceptable() *ReportCreateNotAcceptable {
	return &ReportCreateNotAcceptable{}
}

/*
ReportCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type ReportCreateNotAcceptable struct {
}

// IsSuccess returns true when this report create not acceptable response has a 2xx status code
func (o *ReportCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report create not acceptable response has a 3xx status code
func (o *ReportCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report create not acceptable response has a 4xx status code
func (o *ReportCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this report create not acceptable response has a 5xx status code
func (o *ReportCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this report create not acceptable response a status code equal to that given
func (o *ReportCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the report create not acceptable response
func (o *ReportCreateNotAcceptable) Code() int {
	return 406
}

func (o *ReportCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateNotAcceptable", 406)
}

func (o *ReportCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateNotAcceptable", 406)
}

func (o *ReportCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReportCreateInternalServerError creates a ReportCreateInternalServerError with default headers values
func NewReportCreateInternalServerError() *ReportCreateInternalServerError {
	return &ReportCreateInternalServerError{}
}

/*
ReportCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ReportCreateInternalServerError struct {
}

// IsSuccess returns true when this report create internal server error response has a 2xx status code
func (o *ReportCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this report create internal server error response has a 3xx status code
func (o *ReportCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report create internal server error response has a 4xx status code
func (o *ReportCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this report create internal server error response has a 5xx status code
func (o *ReportCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this report create internal server error response a status code equal to that given
func (o *ReportCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the report create internal server error response
func (o *ReportCreateInternalServerError) Code() int {
	return 500
}

func (o *ReportCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateInternalServerError", 500)
}

func (o *ReportCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/reports][%d] reportCreateInternalServerError", 500)
}

func (o *ReportCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
