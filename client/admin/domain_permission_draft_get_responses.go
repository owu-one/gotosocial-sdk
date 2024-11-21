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

// DomainPermissionDraftGetReader is a Reader for the DomainPermissionDraftGet structure.
type DomainPermissionDraftGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionDraftGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionDraftGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDomainPermissionDraftGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionDraftGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDomainPermissionDraftGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionDraftGetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionDraftGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/admin/domain_permission_drafts/{id}] domainPermissionDraftGet", response, response.Code())
	}
}

// NewDomainPermissionDraftGetOK creates a DomainPermissionDraftGetOK with default headers values
func NewDomainPermissionDraftGetOK() *DomainPermissionDraftGetOK {
	return &DomainPermissionDraftGetOK{}
}

/*
DomainPermissionDraftGetOK describes a response with status code 200, with default header values.

Domain permission draft.
*/
type DomainPermissionDraftGetOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain permission draft get o k response has a 2xx status code
func (o *DomainPermissionDraftGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission draft get o k response has a 3xx status code
func (o *DomainPermissionDraftGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft get o k response has a 4xx status code
func (o *DomainPermissionDraftGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission draft get o k response has a 5xx status code
func (o *DomainPermissionDraftGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft get o k response a status code equal to that given
func (o *DomainPermissionDraftGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission draft get o k response
func (o *DomainPermissionDraftGetOK) Code() int {
	return 200
}

func (o *DomainPermissionDraftGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetOK %s", 200, payload)
}

func (o *DomainPermissionDraftGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetOK %s", 200, payload)
}

func (o *DomainPermissionDraftGetOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainPermissionDraftGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainPermissionDraftGetUnauthorized creates a DomainPermissionDraftGetUnauthorized with default headers values
func NewDomainPermissionDraftGetUnauthorized() *DomainPermissionDraftGetUnauthorized {
	return &DomainPermissionDraftGetUnauthorized{}
}

/*
DomainPermissionDraftGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionDraftGetUnauthorized struct {
}

// IsSuccess returns true when this domain permission draft get unauthorized response has a 2xx status code
func (o *DomainPermissionDraftGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft get unauthorized response has a 3xx status code
func (o *DomainPermissionDraftGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft get unauthorized response has a 4xx status code
func (o *DomainPermissionDraftGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft get unauthorized response has a 5xx status code
func (o *DomainPermissionDraftGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft get unauthorized response a status code equal to that given
func (o *DomainPermissionDraftGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission draft get unauthorized response
func (o *DomainPermissionDraftGetUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionDraftGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetUnauthorized", 401)
}

func (o *DomainPermissionDraftGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetUnauthorized", 401)
}

func (o *DomainPermissionDraftGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftGetForbidden creates a DomainPermissionDraftGetForbidden with default headers values
func NewDomainPermissionDraftGetForbidden() *DomainPermissionDraftGetForbidden {
	return &DomainPermissionDraftGetForbidden{}
}

/*
DomainPermissionDraftGetForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionDraftGetForbidden struct {
}

// IsSuccess returns true when this domain permission draft get forbidden response has a 2xx status code
func (o *DomainPermissionDraftGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft get forbidden response has a 3xx status code
func (o *DomainPermissionDraftGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft get forbidden response has a 4xx status code
func (o *DomainPermissionDraftGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft get forbidden response has a 5xx status code
func (o *DomainPermissionDraftGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft get forbidden response a status code equal to that given
func (o *DomainPermissionDraftGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission draft get forbidden response
func (o *DomainPermissionDraftGetForbidden) Code() int {
	return 403
}

func (o *DomainPermissionDraftGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetForbidden", 403)
}

func (o *DomainPermissionDraftGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetForbidden", 403)
}

func (o *DomainPermissionDraftGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftGetNotFound creates a DomainPermissionDraftGetNotFound with default headers values
func NewDomainPermissionDraftGetNotFound() *DomainPermissionDraftGetNotFound {
	return &DomainPermissionDraftGetNotFound{}
}

/*
DomainPermissionDraftGetNotFound describes a response with status code 404, with default header values.

not found
*/
type DomainPermissionDraftGetNotFound struct {
}

// IsSuccess returns true when this domain permission draft get not found response has a 2xx status code
func (o *DomainPermissionDraftGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft get not found response has a 3xx status code
func (o *DomainPermissionDraftGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft get not found response has a 4xx status code
func (o *DomainPermissionDraftGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft get not found response has a 5xx status code
func (o *DomainPermissionDraftGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft get not found response a status code equal to that given
func (o *DomainPermissionDraftGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the domain permission draft get not found response
func (o *DomainPermissionDraftGetNotFound) Code() int {
	return 404
}

func (o *DomainPermissionDraftGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetNotFound", 404)
}

func (o *DomainPermissionDraftGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetNotFound", 404)
}

func (o *DomainPermissionDraftGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftGetNotAcceptable creates a DomainPermissionDraftGetNotAcceptable with default headers values
func NewDomainPermissionDraftGetNotAcceptable() *DomainPermissionDraftGetNotAcceptable {
	return &DomainPermissionDraftGetNotAcceptable{}
}

/*
DomainPermissionDraftGetNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionDraftGetNotAcceptable struct {
}

// IsSuccess returns true when this domain permission draft get not acceptable response has a 2xx status code
func (o *DomainPermissionDraftGetNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft get not acceptable response has a 3xx status code
func (o *DomainPermissionDraftGetNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft get not acceptable response has a 4xx status code
func (o *DomainPermissionDraftGetNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft get not acceptable response has a 5xx status code
func (o *DomainPermissionDraftGetNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft get not acceptable response a status code equal to that given
func (o *DomainPermissionDraftGetNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission draft get not acceptable response
func (o *DomainPermissionDraftGetNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionDraftGetNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetNotAcceptable", 406)
}

func (o *DomainPermissionDraftGetNotAcceptable) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetNotAcceptable", 406)
}

func (o *DomainPermissionDraftGetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftGetInternalServerError creates a DomainPermissionDraftGetInternalServerError with default headers values
func NewDomainPermissionDraftGetInternalServerError() *DomainPermissionDraftGetInternalServerError {
	return &DomainPermissionDraftGetInternalServerError{}
}

/*
DomainPermissionDraftGetInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionDraftGetInternalServerError struct {
}

// IsSuccess returns true when this domain permission draft get internal server error response has a 2xx status code
func (o *DomainPermissionDraftGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft get internal server error response has a 3xx status code
func (o *DomainPermissionDraftGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft get internal server error response has a 4xx status code
func (o *DomainPermissionDraftGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission draft get internal server error response has a 5xx status code
func (o *DomainPermissionDraftGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission draft get internal server error response a status code equal to that given
func (o *DomainPermissionDraftGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission draft get internal server error response
func (o *DomainPermissionDraftGetInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionDraftGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetInternalServerError", 500)
}

func (o *DomainPermissionDraftGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/admin/domain_permission_drafts/{id}][%d] domainPermissionDraftGetInternalServerError", 500)
}

func (o *DomainPermissionDraftGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}