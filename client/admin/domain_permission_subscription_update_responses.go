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

// DomainPermissionSubscriptionUpdateReader is a Reader for the DomainPermissionSubscriptionUpdate structure.
type DomainPermissionSubscriptionUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPermissionSubscriptionUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPermissionSubscriptionUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDomainPermissionSubscriptionUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDomainPermissionSubscriptionUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDomainPermissionSubscriptionUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDomainPermissionSubscriptionUpdateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDomainPermissionSubscriptionUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDomainPermissionSubscriptionUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}] domainPermissionSubscriptionUpdate", response, response.Code())
	}
}

// NewDomainPermissionSubscriptionUpdateOK creates a DomainPermissionSubscriptionUpdateOK with default headers values
func NewDomainPermissionSubscriptionUpdateOK() *DomainPermissionSubscriptionUpdateOK {
	return &DomainPermissionSubscriptionUpdateOK{}
}

/*
DomainPermissionSubscriptionUpdateOK describes a response with status code 200, with default header values.

The updated domain permission subscription.
*/
type DomainPermissionSubscriptionUpdateOK struct {
	Payload *models.DomainPermissionSubscription
}

// IsSuccess returns true when this domain permission subscription update o k response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain permission subscription update o k response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update o k response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscription update o k response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription update o k response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain permission subscription update o k response
func (o *DomainPermissionSubscriptionUpdateOK) Code() int {
	return 200
}

func (o *DomainPermissionSubscriptionUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateOK %s", 200, payload)
}

func (o *DomainPermissionSubscriptionUpdateOK) GetPayload() *models.DomainPermissionSubscription {
	return o.Payload
}

func (o *DomainPermissionSubscriptionUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainPermissionSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDomainPermissionSubscriptionUpdateBadRequest creates a DomainPermissionSubscriptionUpdateBadRequest with default headers values
func NewDomainPermissionSubscriptionUpdateBadRequest() *DomainPermissionSubscriptionUpdateBadRequest {
	return &DomainPermissionSubscriptionUpdateBadRequest{}
}

/*
DomainPermissionSubscriptionUpdateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DomainPermissionSubscriptionUpdateBadRequest struct {
}

// IsSuccess returns true when this domain permission subscription update bad request response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription update bad request response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update bad request response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription update bad request response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription update bad request response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the domain permission subscription update bad request response
func (o *DomainPermissionSubscriptionUpdateBadRequest) Code() int {
	return 400
}

func (o *DomainPermissionSubscriptionUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateBadRequest", 400)
}

func (o *DomainPermissionSubscriptionUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateBadRequest", 400)
}

func (o *DomainPermissionSubscriptionUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionUpdateUnauthorized creates a DomainPermissionSubscriptionUpdateUnauthorized with default headers values
func NewDomainPermissionSubscriptionUpdateUnauthorized() *DomainPermissionSubscriptionUpdateUnauthorized {
	return &DomainPermissionSubscriptionUpdateUnauthorized{}
}

/*
DomainPermissionSubscriptionUpdateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DomainPermissionSubscriptionUpdateUnauthorized struct {
}

// IsSuccess returns true when this domain permission subscription update unauthorized response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription update unauthorized response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update unauthorized response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription update unauthorized response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription update unauthorized response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the domain permission subscription update unauthorized response
func (o *DomainPermissionSubscriptionUpdateUnauthorized) Code() int {
	return 401
}

func (o *DomainPermissionSubscriptionUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateUnauthorized", 401)
}

func (o *DomainPermissionSubscriptionUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionUpdateForbidden creates a DomainPermissionSubscriptionUpdateForbidden with default headers values
func NewDomainPermissionSubscriptionUpdateForbidden() *DomainPermissionSubscriptionUpdateForbidden {
	return &DomainPermissionSubscriptionUpdateForbidden{}
}

/*
DomainPermissionSubscriptionUpdateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DomainPermissionSubscriptionUpdateForbidden struct {
}

// IsSuccess returns true when this domain permission subscription update forbidden response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription update forbidden response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update forbidden response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription update forbidden response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription update forbidden response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the domain permission subscription update forbidden response
func (o *DomainPermissionSubscriptionUpdateForbidden) Code() int {
	return 403
}

func (o *DomainPermissionSubscriptionUpdateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateForbidden", 403)
}

func (o *DomainPermissionSubscriptionUpdateForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateForbidden", 403)
}

func (o *DomainPermissionSubscriptionUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionUpdateNotAcceptable creates a DomainPermissionSubscriptionUpdateNotAcceptable with default headers values
func NewDomainPermissionSubscriptionUpdateNotAcceptable() *DomainPermissionSubscriptionUpdateNotAcceptable {
	return &DomainPermissionSubscriptionUpdateNotAcceptable{}
}

/*
DomainPermissionSubscriptionUpdateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type DomainPermissionSubscriptionUpdateNotAcceptable struct {
}

// IsSuccess returns true when this domain permission subscription update not acceptable response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription update not acceptable response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update not acceptable response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription update not acceptable response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription update not acceptable response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the domain permission subscription update not acceptable response
func (o *DomainPermissionSubscriptionUpdateNotAcceptable) Code() int {
	return 406
}

func (o *DomainPermissionSubscriptionUpdateNotAcceptable) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionUpdateNotAcceptable) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateNotAcceptable", 406)
}

func (o *DomainPermissionSubscriptionUpdateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionUpdateConflict creates a DomainPermissionSubscriptionUpdateConflict with default headers values
func NewDomainPermissionSubscriptionUpdateConflict() *DomainPermissionSubscriptionUpdateConflict {
	return &DomainPermissionSubscriptionUpdateConflict{}
}

/*
DomainPermissionSubscriptionUpdateConflict describes a response with status code 409, with default header values.

conflict
*/
type DomainPermissionSubscriptionUpdateConflict struct {
}

// IsSuccess returns true when this domain permission subscription update conflict response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription update conflict response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update conflict response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this domain permission subscription update conflict response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this domain permission subscription update conflict response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the domain permission subscription update conflict response
func (o *DomainPermissionSubscriptionUpdateConflict) Code() int {
	return 409
}

func (o *DomainPermissionSubscriptionUpdateConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateConflict", 409)
}

func (o *DomainPermissionSubscriptionUpdateConflict) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateConflict", 409)
}

func (o *DomainPermissionSubscriptionUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDomainPermissionSubscriptionUpdateInternalServerError creates a DomainPermissionSubscriptionUpdateInternalServerError with default headers values
func NewDomainPermissionSubscriptionUpdateInternalServerError() *DomainPermissionSubscriptionUpdateInternalServerError {
	return &DomainPermissionSubscriptionUpdateInternalServerError{}
}

/*
DomainPermissionSubscriptionUpdateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type DomainPermissionSubscriptionUpdateInternalServerError struct {
}

// IsSuccess returns true when this domain permission subscription update internal server error response has a 2xx status code
func (o *DomainPermissionSubscriptionUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this domain permission subscription update internal server error response has a 3xx status code
func (o *DomainPermissionSubscriptionUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain permission subscription update internal server error response has a 4xx status code
func (o *DomainPermissionSubscriptionUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain permission subscription update internal server error response has a 5xx status code
func (o *DomainPermissionSubscriptionUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this domain permission subscription update internal server error response a status code equal to that given
func (o *DomainPermissionSubscriptionUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the domain permission subscription update internal server error response
func (o *DomainPermissionSubscriptionUpdateInternalServerError) Code() int {
	return 500
}

func (o *DomainPermissionSubscriptionUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/domain_permission_subscriptions/${id}][%d] domainPermissionSubscriptionUpdateInternalServerError", 500)
}

func (o *DomainPermissionSubscriptionUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}