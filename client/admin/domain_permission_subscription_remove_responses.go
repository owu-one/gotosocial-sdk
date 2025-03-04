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

// DomainPermissionSubscriptionRemoveReader is a Reader for the DomainPermissionSubscriptionRemove structure.
type DomainPermissionSubscriptionRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionSubscriptionRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionSubscriptionRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainPermissionSubscriptionRemoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainPermissionSubscriptionRemoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionSubscriptionRemoveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionSubscriptionRemoveNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainPermissionSubscriptionRemoveConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionSubscriptionRemoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove] domainPermissionSubscriptionRemove", response, response.Code())
	}
}

// NewDomainPermissionSubscriptionRemoveOK creates a DomainPermissionSubscriptionRemoveOK with default headers values
func NewDomainPermissionSubscriptionRemoveOK() *DomainPermissionSubscriptionRemoveOK {
	return &DomainPermissionSubscriptionRemoveOK{}
}

/*
DomainPermissionSubscriptionRemoveOK describes a response with status code 200, with default header values.

The removed domain permission subscription.
*/
type DomainPermissionSubscriptionRemoveOK struct {
	Payload *models.DomainPermissionSubscription
}

// IsSuccess returns true when this domain permission subscription remove o k response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission subscription remove o k response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove o k response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscription remove o k response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription remove o k response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission subscription remove o k response
func (o *DomainPermissionSubscriptionRemoveOK) Code() int {
	return 200
}

func (o *DomainPermissionSubscriptionRemoveOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionRemoveOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionRemoveOK) GetPayload() *models.DomainPermissionSubscription {
	return o.Payload
}

func (o *DomainPermissionSubscriptionRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermissionSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainPermissionSubscriptionRemoveBadRequest creates a DomainPermissionSubscriptionRemoveBadRequest with default headers values
func NewDomainPermissionSubscriptionRemoveBadRequest() *DomainPermissionSubscriptionRemoveBadRequest {
	return &DomainPermissionSubscriptionRemoveBadRequest{}
}

/*
DomainPermissionSubscriptionRemoveBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainPermissionSubscriptionRemoveBadRequest struct {
}

// IsSuccess returns true when this domain permission subscription remove bad request response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription remove bad request response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove bad request response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription remove bad request response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription remove bad request response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain permission subscription remove bad request response
func (o *DomainPermissionSubscriptionRemoveBadRequest) Code() int {
	return 400
}

func (o *DomainPermissionSubscriptionRemoveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveBadRequest", 400)
}

func (o *DomainPermissionSubscriptionRemoveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveBadRequest", 400)
}

func (o *DomainPermissionSubscriptionRemoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionRemoveUnauthorized creates a DomainPermissionSubscriptionRemoveUnauthorized with default headers values
func NewDomainPermissionSubscriptionRemoveUnauthorized() *DomainPermissionSubscriptionRemoveUnauthorized {
	return &DomainPermissionSubscriptionRemoveUnauthorized{}
}

/*
DomainPermissionSubscriptionRemoveUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionSubscriptionRemoveUnauthorized struct {
}

// IsSuccess returns true when this domain permission subscription remove unauthorized response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription remove unauthorized response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove unauthorized response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription remove unauthorized response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription remove unauthorized response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission subscription remove unauthorized response
func (o *DomainPermissionSubscriptionRemoveUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionSubscriptionRemoveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionRemoveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionRemoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionRemoveForbidden creates a DomainPermissionSubscriptionRemoveForbidden with default headers values
func NewDomainPermissionSubscriptionRemoveForbidden() *DomainPermissionSubscriptionRemoveForbidden {
	return &DomainPermissionSubscriptionRemoveForbidden{}
}

/*
DomainPermissionSubscriptionRemoveForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionSubscriptionRemoveForbidden struct {
}

// IsSuccess returns true when this domain permission subscription remove forbidden response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription remove forbidden response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove forbidden response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription remove forbidden response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription remove forbidden response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission subscription remove forbidden response
func (o *DomainPermissionSubscriptionRemoveForbidden) Code() int {
	return 403
}

func (o *DomainPermissionSubscriptionRemoveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveForbidden", 403)
}

func (o *DomainPermissionSubscriptionRemoveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveForbidden", 403)
}

func (o *DomainPermissionSubscriptionRemoveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionRemoveNotAcceptable creates a DomainPermissionSubscriptionRemoveNotAcceptable with default headers values
func NewDomainPermissionSubscriptionRemoveNotAcceptable() *DomainPermissionSubscriptionRemoveNotAcceptable {
	return &DomainPermissionSubscriptionRemoveNotAcceptable{}
}

/*
DomainPermissionSubscriptionRemoveNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionSubscriptionRemoveNotAcceptable struct {
}

// IsSuccess returns true when this domain permission subscription remove not acceptable response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription remove not acceptable response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove not acceptable response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription remove not acceptable response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription remove not acceptable response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission subscription remove not acceptable response
func (o *DomainPermissionSubscriptionRemoveNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionSubscriptionRemoveNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionRemoveNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionRemoveNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionRemoveConflict creates a DomainPermissionSubscriptionRemoveConflict with default headers values
func NewDomainPermissionSubscriptionRemoveConflict() *DomainPermissionSubscriptionRemoveConflict {
	return &DomainPermissionSubscriptionRemoveConflict{}
}

/*
DomainPermissionSubscriptionRemoveConflict describes a response with status code 409, with default header values.

conflict
*/
type DomainPermissionSubscriptionRemoveConflict struct {
}

// IsSuccess returns true when this domain permission subscription remove conflict response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription remove conflict response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove conflict response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription remove conflict response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription remove conflict response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain permission subscription remove conflict response
func (o *DomainPermissionSubscriptionRemoveConflict) Code() int {
	return 409
}

func (o *DomainPermissionSubscriptionRemoveConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveConflict", 409)
}

func (o *DomainPermissionSubscriptionRemoveConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveConflict", 409)
}

func (o *DomainPermissionSubscriptionRemoveConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionRemoveInternalServerError creates a DomainPermissionSubscriptionRemoveInternalServerError with default headers values
func NewDomainPermissionSubscriptionRemoveInternalServerError() *DomainPermissionSubscriptionRemoveInternalServerError {
	return &DomainPermissionSubscriptionRemoveInternalServerError{}
}

/*
DomainPermissionSubscriptionRemoveInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionSubscriptionRemoveInternalServerError struct {
}

// IsSuccess returns true when this domain permission subscription remove internal server error response has a 2xx status code
func (o *DomainPermissionSubscriptionRemoveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription remove internal server error response has a 3xx status code
func (o *DomainPermissionSubscriptionRemoveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription remove internal server error response has a 4xx status code
func (o *DomainPermissionSubscriptionRemoveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscription remove internal server error response has a 5xx status code
func (o *DomainPermissionSubscriptionRemoveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission subscription remove internal server error response a status code equal to that given
func (o *DomainPermissionSubscriptionRemoveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission subscription remove internal server error response
func (o *DomainPermissionSubscriptionRemoveInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionSubscriptionRemoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionRemoveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/domain_permission_subscriptions/{id}/remove][%d] domainPermissionSubscriptionRemoveInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionRemoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
