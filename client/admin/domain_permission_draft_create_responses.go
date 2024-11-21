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

// DomainPermissionDraftCreateReader is a Reader for the DomainPermissionDraftCreate structure.
type DomainPermissionDraftCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionDraftCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionDraftCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainPermissionDraftCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainPermissionDraftCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionDraftCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionDraftCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainPermissionDraftCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionDraftCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/domain_permission_drafts] domainPermissionDraftCreate", response, response.Code())
	}
}

// NewDomainPermissionDraftCreateOK creates a DomainPermissionDraftCreateOK with default headers values
func NewDomainPermissionDraftCreateOK() *DomainPermissionDraftCreateOK {
	return &DomainPermissionDraftCreateOK{}
}

/*
DomainPermissionDraftCreateOK describes a response with status code 200, with default header values.

The newly created domain permission draft.
*/
type DomainPermissionDraftCreateOK struct {
	Payload *models.DomainPermission
}

// IsSuccess returns true when this domain permission draft create o k response has a 2xx status code
func (o *DomainPermissionDraftCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission draft create o k response has a 3xx status code
func (o *DomainPermissionDraftCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create o k response has a 4xx status code
func (o *DomainPermissionDraftCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission draft create o k response has a 5xx status code
func (o *DomainPermissionDraftCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft create o k response a status code equal to that given
func (o *DomainPermissionDraftCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission draft create o k response
func (o *DomainPermissionDraftCreateOK) Code() int {
	return 200
}

func (o *DomainPermissionDraftCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateOK %s", 200, payload)
}

func (o *DomainPermissionDraftCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateOK %s", 200, payload)
}

func (o *DomainPermissionDraftCreateOK) GetPayload() *models.DomainPermission {
	return o.Payload
}

func (o *DomainPermissionDraftCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainPermissionDraftCreateBadRequest creates a DomainPermissionDraftCreateBadRequest with default headers values
func NewDomainPermissionDraftCreateBadRequest() *DomainPermissionDraftCreateBadRequest {
	return &DomainPermissionDraftCreateBadRequest{}
}

/*
DomainPermissionDraftCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainPermissionDraftCreateBadRequest struct {
}

// IsSuccess returns true when this domain permission draft create bad request response has a 2xx status code
func (o *DomainPermissionDraftCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft create bad request response has a 3xx status code
func (o *DomainPermissionDraftCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create bad request response has a 4xx status code
func (o *DomainPermissionDraftCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft create bad request response has a 5xx status code
func (o *DomainPermissionDraftCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft create bad request response a status code equal to that given
func (o *DomainPermissionDraftCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain permission draft create bad request response
func (o *DomainPermissionDraftCreateBadRequest) Code() int {
	return 400
}

func (o *DomainPermissionDraftCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateBadRequest", 400)
}

func (o *DomainPermissionDraftCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateBadRequest", 400)
}

func (o *DomainPermissionDraftCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftCreateUnauthorized creates a DomainPermissionDraftCreateUnauthorized with default headers values
func NewDomainPermissionDraftCreateUnauthorized() *DomainPermissionDraftCreateUnauthorized {
	return &DomainPermissionDraftCreateUnauthorized{}
}

/*
DomainPermissionDraftCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionDraftCreateUnauthorized struct {
}

// IsSuccess returns true when this domain permission draft create unauthorized response has a 2xx status code
func (o *DomainPermissionDraftCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft create unauthorized response has a 3xx status code
func (o *DomainPermissionDraftCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create unauthorized response has a 4xx status code
func (o *DomainPermissionDraftCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft create unauthorized response has a 5xx status code
func (o *DomainPermissionDraftCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft create unauthorized response a status code equal to that given
func (o *DomainPermissionDraftCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission draft create unauthorized response
func (o *DomainPermissionDraftCreateUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionDraftCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateUnauthorized", 401)
}

func (o *DomainPermissionDraftCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateUnauthorized", 401)
}

func (o *DomainPermissionDraftCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftCreateForbidden creates a DomainPermissionDraftCreateForbidden with default headers values
func NewDomainPermissionDraftCreateForbidden() *DomainPermissionDraftCreateForbidden {
	return &DomainPermissionDraftCreateForbidden{}
}

/*
DomainPermissionDraftCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionDraftCreateForbidden struct {
}

// IsSuccess returns true when this domain permission draft create forbidden response has a 2xx status code
func (o *DomainPermissionDraftCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft create forbidden response has a 3xx status code
func (o *DomainPermissionDraftCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create forbidden response has a 4xx status code
func (o *DomainPermissionDraftCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft create forbidden response has a 5xx status code
func (o *DomainPermissionDraftCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft create forbidden response a status code equal to that given
func (o *DomainPermissionDraftCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission draft create forbidden response
func (o *DomainPermissionDraftCreateForbidden) Code() int {
	return 403
}

func (o *DomainPermissionDraftCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateForbidden", 403)
}

func (o *DomainPermissionDraftCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateForbidden", 403)
}

func (o *DomainPermissionDraftCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftCreateNotAcceptable creates a DomainPermissionDraftCreateNotAcceptable with default headers values
func NewDomainPermissionDraftCreateNotAcceptable() *DomainPermissionDraftCreateNotAcceptable {
	return &DomainPermissionDraftCreateNotAcceptable{}
}

/*
DomainPermissionDraftCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionDraftCreateNotAcceptable struct {
}

// IsSuccess returns true when this domain permission draft create not acceptable response has a 2xx status code
func (o *DomainPermissionDraftCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft create not acceptable response has a 3xx status code
func (o *DomainPermissionDraftCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create not acceptable response has a 4xx status code
func (o *DomainPermissionDraftCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft create not acceptable response has a 5xx status code
func (o *DomainPermissionDraftCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft create not acceptable response a status code equal to that given
func (o *DomainPermissionDraftCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission draft create not acceptable response
func (o *DomainPermissionDraftCreateNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionDraftCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateNotAcceptable", 406)
}

func (o *DomainPermissionDraftCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateNotAcceptable", 406)
}

func (o *DomainPermissionDraftCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftCreateConflict creates a DomainPermissionDraftCreateConflict with default headers values
func NewDomainPermissionDraftCreateConflict() *DomainPermissionDraftCreateConflict {
	return &DomainPermissionDraftCreateConflict{}
}

/*
DomainPermissionDraftCreateConflict describes a response with status code 409, with default header values.

conflict
*/
type DomainPermissionDraftCreateConflict struct {
}

// IsSuccess returns true when this domain permission draft create conflict response has a 2xx status code
func (o *DomainPermissionDraftCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft create conflict response has a 3xx status code
func (o *DomainPermissionDraftCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create conflict response has a 4xx status code
func (o *DomainPermissionDraftCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission draft create conflict response has a 5xx status code
func (o *DomainPermissionDraftCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission draft create conflict response a status code equal to that given
func (o *DomainPermissionDraftCreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain permission draft create conflict response
func (o *DomainPermissionDraftCreateConflict) Code() int {
	return 409
}

func (o *DomainPermissionDraftCreateConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateConflict", 409)
}

func (o *DomainPermissionDraftCreateConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateConflict", 409)
}

func (o *DomainPermissionDraftCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionDraftCreateInternalServerError creates a DomainPermissionDraftCreateInternalServerError with default headers values
func NewDomainPermissionDraftCreateInternalServerError() *DomainPermissionDraftCreateInternalServerError {
	return &DomainPermissionDraftCreateInternalServerError{}
}

/*
DomainPermissionDraftCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionDraftCreateInternalServerError struct {
}

// IsSuccess returns true when this domain permission draft create internal server error response has a 2xx status code
func (o *DomainPermissionDraftCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission draft create internal server error response has a 3xx status code
func (o *DomainPermissionDraftCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission draft create internal server error response has a 4xx status code
func (o *DomainPermissionDraftCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission draft create internal server error response has a 5xx status code
func (o *DomainPermissionDraftCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission draft create internal server error response a status code equal to that given
func (o *DomainPermissionDraftCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission draft create internal server error response
func (o *DomainPermissionDraftCreateInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionDraftCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateInternalServerError", 500)
}

func (o *DomainPermissionDraftCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_drafts][%d] domainPermissionDraftCreateInternalServerError", 500)
}

func (o *DomainPermissionDraftCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
