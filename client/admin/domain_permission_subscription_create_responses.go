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

// DomainPermissionSubscriptionCreateReader is a Reader for the DomainPermissionSubscriptionCreate structure.
type DomainPermissionSubscriptionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionSubscriptionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionSubscriptionCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainPermissionSubscriptionCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainPermissionSubscriptionCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionSubscriptionCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionSubscriptionCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainPermissionSubscriptionCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionSubscriptionCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/domain_permission_subscriptions] domainPermissionSubscriptionCreate", response, response.Code())
	}
}

// NewDomainPermissionSubscriptionCreateOK creates a DomainPermissionSubscriptionCreateOK with default headers values
func NewDomainPermissionSubscriptionCreateOK() *DomainPermissionSubscriptionCreateOK {
	return &DomainPermissionSubscriptionCreateOK{}
}

/*
DomainPermissionSubscriptionCreateOK describes a response with status code 200, with default header values.

The newly created domain permission subscription.
*/
type DomainPermissionSubscriptionCreateOK struct {
	Payload *models.DomainPermissionSubscription
}

// IsSuccess returns true when this domain permission subscription create o k response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission subscription create o k response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create o k response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscription create o k response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription create o k response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission subscription create o k response
func (o *DomainPermissionSubscriptionCreateOK) Code() int {
	return 200
}

func (o *DomainPermissionSubscriptionCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionCreateOK) GetPayload() *models.DomainPermissionSubscription {
	return o.Payload
}

func (o *DomainPermissionSubscriptionCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermissionSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainPermissionSubscriptionCreateBadRequest creates a DomainPermissionSubscriptionCreateBadRequest with default headers values
func NewDomainPermissionSubscriptionCreateBadRequest() *DomainPermissionSubscriptionCreateBadRequest {
	return &DomainPermissionSubscriptionCreateBadRequest{}
}

/*
DomainPermissionSubscriptionCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainPermissionSubscriptionCreateBadRequest struct {
}

// IsSuccess returns true when this domain permission subscription create bad request response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription create bad request response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create bad request response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription create bad request response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription create bad request response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain permission subscription create bad request response
func (o *DomainPermissionSubscriptionCreateBadRequest) Code() int {
	return 400
}

func (o *DomainPermissionSubscriptionCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateBadRequest", 400)
}

func (o *DomainPermissionSubscriptionCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateBadRequest", 400)
}

func (o *DomainPermissionSubscriptionCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionCreateUnauthorized creates a DomainPermissionSubscriptionCreateUnauthorized with default headers values
func NewDomainPermissionSubscriptionCreateUnauthorized() *DomainPermissionSubscriptionCreateUnauthorized {
	return &DomainPermissionSubscriptionCreateUnauthorized{}
}

/*
DomainPermissionSubscriptionCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionSubscriptionCreateUnauthorized struct {
}

// IsSuccess returns true when this domain permission subscription create unauthorized response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription create unauthorized response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create unauthorized response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription create unauthorized response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription create unauthorized response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission subscription create unauthorized response
func (o *DomainPermissionSubscriptionCreateUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionSubscriptionCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionCreateForbidden creates a DomainPermissionSubscriptionCreateForbidden with default headers values
func NewDomainPermissionSubscriptionCreateForbidden() *DomainPermissionSubscriptionCreateForbidden {
	return &DomainPermissionSubscriptionCreateForbidden{}
}

/*
DomainPermissionSubscriptionCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionSubscriptionCreateForbidden struct {
}

// IsSuccess returns true when this domain permission subscription create forbidden response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription create forbidden response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create forbidden response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription create forbidden response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription create forbidden response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission subscription create forbidden response
func (o *DomainPermissionSubscriptionCreateForbidden) Code() int {
	return 403
}

func (o *DomainPermissionSubscriptionCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateForbidden", 403)
}

func (o *DomainPermissionSubscriptionCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateForbidden", 403)
}

func (o *DomainPermissionSubscriptionCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionCreateNotAcceptable creates a DomainPermissionSubscriptionCreateNotAcceptable with default headers values
func NewDomainPermissionSubscriptionCreateNotAcceptable() *DomainPermissionSubscriptionCreateNotAcceptable {
	return &DomainPermissionSubscriptionCreateNotAcceptable{}
}

/*
DomainPermissionSubscriptionCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionSubscriptionCreateNotAcceptable struct {
}

// IsSuccess returns true when this domain permission subscription create not acceptable response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription create not acceptable response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create not acceptable response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription create not acceptable response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription create not acceptable response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission subscription create not acceptable response
func (o *DomainPermissionSubscriptionCreateNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionSubscriptionCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionCreateConflict creates a DomainPermissionSubscriptionCreateConflict with default headers values
func NewDomainPermissionSubscriptionCreateConflict() *DomainPermissionSubscriptionCreateConflict {
	return &DomainPermissionSubscriptionCreateConflict{}
}

/*
DomainPermissionSubscriptionCreateConflict describes a response with status code 409, with default header values.

conflict
*/
type DomainPermissionSubscriptionCreateConflict struct {
}

// IsSuccess returns true when this domain permission subscription create conflict response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription create conflict response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create conflict response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription create conflict response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription create conflict response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain permission subscription create conflict response
func (o *DomainPermissionSubscriptionCreateConflict) Code() int {
	return 409
}

func (o *DomainPermissionSubscriptionCreateConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateConflict", 409)
}

func (o *DomainPermissionSubscriptionCreateConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateConflict", 409)
}

func (o *DomainPermissionSubscriptionCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionCreateInternalServerError creates a DomainPermissionSubscriptionCreateInternalServerError with default headers values
func NewDomainPermissionSubscriptionCreateInternalServerError() *DomainPermissionSubscriptionCreateInternalServerError {
	return &DomainPermissionSubscriptionCreateInternalServerError{}
}

/*
DomainPermissionSubscriptionCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionSubscriptionCreateInternalServerError struct {
}

// IsSuccess returns true when this domain permission subscription create internal server error response has a 2xx status code
func (o *DomainPermissionSubscriptionCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription create internal server error response has a 3xx status code
func (o *DomainPermissionSubscriptionCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription create internal server error response has a 4xx status code
func (o *DomainPermissionSubscriptionCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscription create internal server error response has a 5xx status code
func (o *DomainPermissionSubscriptionCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission subscription create internal server error response a status code equal to that given
func (o *DomainPermissionSubscriptionCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission subscription create internal server error response
func (o *DomainPermissionSubscriptionCreateInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionSubscriptionCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions][%d] domainPermissionSubscriptionCreateInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
