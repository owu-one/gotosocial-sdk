// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// AdminReportResolveReader is a Reader for the AdminReportResolve structure.
type AdminReportResolveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminReportResolveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminReportResolveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminReportResolveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminReportResolveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminReportResolveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAdminReportResolveNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminReportResolveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/reports/{id}/resolve] adminReportResolve", response, response.Code())
	}
}

// NewAdminReportResolveOK creates a AdminReportResolveOK with default headers values
func NewAdminReportResolveOK() *AdminReportResolveOK {
	return &AdminReportResolveOK{}
}

/*
AdminReportResolveOK describes a response with status code 200, with default header values.

The resolved report.
*/
type AdminReportResolveOK struct {
	Payload *models.AdminReport
}

// IsSuccess returns true when this admin report resolve o k response has a 2xx status code
func (o *AdminReportResolveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin report resolve o k response has a 3xx status code
func (o *AdminReportResolveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin report resolve o k response has a 4xx status code
func (o *AdminReportResolveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin report resolve o k response has a 5xx status code
func (o *AdminReportResolveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin report resolve o k response a status code equal to that given
func (o *AdminReportResolveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin report resolve o k response
func (o *AdminReportResolveOK) Code() int {
	return 200
}

func (o *AdminReportResolveOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveOK %s", 200, payload)
}

func (o *AdminReportResolveOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveOK %s", 200, payload)
}

func (o *AdminReportResolveOK) GetPayload() *models.AdminReport {
	return o.Payload
}

func (o *AdminReportResolveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminReportResolveBadRequest creates a AdminReportResolveBadRequest with default headers values
func NewAdminReportResolveBadRequest() *AdminReportResolveBadRequest {
	return &AdminReportResolveBadRequest{}
}

/*
AdminReportResolveBadRequest describes a response with status code 400, with default header values.

bad request
*/
type AdminReportResolveBadRequest struct {
}

// IsSuccess returns true when this admin report resolve bad request response has a 2xx status code
func (o *AdminReportResolveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin report resolve bad request response has a 3xx status code
func (o *AdminReportResolveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin report resolve bad request response has a 4xx status code
func (o *AdminReportResolveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin report resolve bad request response has a 5xx status code
func (o *AdminReportResolveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin report resolve bad request response a status code equal to that given
func (o *AdminReportResolveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin report resolve bad request response
func (o *AdminReportResolveBadRequest) Code() int {
	return 400
}

func (o *AdminReportResolveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveBadRequest", 400)
}

func (o *AdminReportResolveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveBadRequest", 400)
}

func (o *AdminReportResolveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminReportResolveUnauthorized creates a AdminReportResolveUnauthorized with default headers values
func NewAdminReportResolveUnauthorized() *AdminReportResolveUnauthorized {
	return &AdminReportResolveUnauthorized{}
}

/*
AdminReportResolveUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type AdminReportResolveUnauthorized struct {
}

// IsSuccess returns true when this admin report resolve unauthorized response has a 2xx status code
func (o *AdminReportResolveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin report resolve unauthorized response has a 3xx status code
func (o *AdminReportResolveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin report resolve unauthorized response has a 4xx status code
func (o *AdminReportResolveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin report resolve unauthorized response has a 5xx status code
func (o *AdminReportResolveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin report resolve unauthorized response a status code equal to that given
func (o *AdminReportResolveUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin report resolve unauthorized response
func (o *AdminReportResolveUnauthorized) Code() int {
	return 401
}

func (o *AdminReportResolveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveUnauthorized", 401)
}

func (o *AdminReportResolveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveUnauthorized", 401)
}

func (o *AdminReportResolveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminReportResolveNotFound creates a AdminReportResolveNotFound with default headers values
func NewAdminReportResolveNotFound() *AdminReportResolveNotFound {
	return &AdminReportResolveNotFound{}
}

/*
AdminReportResolveNotFound describes a response with status code 404, with default header values.

not found
*/
type AdminReportResolveNotFound struct {
}

// IsSuccess returns true when this admin report resolve not found response has a 2xx status code
func (o *AdminReportResolveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin report resolve not found response has a 3xx status code
func (o *AdminReportResolveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin report resolve not found response has a 4xx status code
func (o *AdminReportResolveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin report resolve not found response has a 5xx status code
func (o *AdminReportResolveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin report resolve not found response a status code equal to that given
func (o *AdminReportResolveNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin report resolve not found response
func (o *AdminReportResolveNotFound) Code() int {
	return 404
}

func (o *AdminReportResolveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveNotFound", 404)
}

func (o *AdminReportResolveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveNotFound", 404)
}

func (o *AdminReportResolveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminReportResolveNotAcceptable creates a AdminReportResolveNotAcceptable with default headers values
func NewAdminReportResolveNotAcceptable() *AdminReportResolveNotAcceptable {
	return &AdminReportResolveNotAcceptable{}
}

/*
AdminReportResolveNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type AdminReportResolveNotAcceptable struct {
}

// IsSuccess returns true when this admin report resolve not acceptable response has a 2xx status code
func (o *AdminReportResolveNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin report resolve not acceptable response has a 3xx status code
func (o *AdminReportResolveNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin report resolve not acceptable response has a 4xx status code
func (o *AdminReportResolveNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin report resolve not acceptable response has a 5xx status code
func (o *AdminReportResolveNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this admin report resolve not acceptable response a status code equal to that given
func (o *AdminReportResolveNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the admin report resolve not acceptable response
func (o *AdminReportResolveNotAcceptable) Code() int {
	return 406
}

func (o *AdminReportResolveNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveNotAcceptable", 406)
}

func (o *AdminReportResolveNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveNotAcceptable", 406)
}

func (o *AdminReportResolveNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminReportResolveInternalServerError creates a AdminReportResolveInternalServerError with default headers values
func NewAdminReportResolveInternalServerError() *AdminReportResolveInternalServerError {
	return &AdminReportResolveInternalServerError{}
}

/*
AdminReportResolveInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AdminReportResolveInternalServerError struct {
}

// IsSuccess returns true when this admin report resolve internal server error response has a 2xx status code
func (o *AdminReportResolveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin report resolve internal server error response has a 3xx status code
func (o *AdminReportResolveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin report resolve internal server error response has a 4xx status code
func (o *AdminReportResolveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin report resolve internal server error response has a 5xx status code
func (o *AdminReportResolveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin report resolve internal server error response a status code equal to that given
func (o *AdminReportResolveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin report resolve internal server error response
func (o *AdminReportResolveInternalServerError) Code() int {
	return 500
}

func (o *AdminReportResolveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveInternalServerError", 500)
}

func (o *AdminReportResolveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/reports/{id}/resolve][%d] adminReportResolveInternalServerError", 500)
}

func (o *AdminReportResolveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
