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

// RuleCreateReader is a Reader for the RuleCreate structure.
type RuleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RuleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRuleCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRuleCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRuleCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRuleCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRuleCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewRuleCreateNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRuleCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/v1/admin/instance/rules] ruleCreate", response, response.Code())
	}
}

// NewRuleCreateOK creates a RuleCreateOK with default headers values
func NewRuleCreateOK() *RuleCreateOK {
	return &RuleCreateOK{}
}

/*
RuleCreateOK describes a response with status code 200, with default header values.

The newly-created instance rule.
*/
type RuleCreateOK struct {
	Payload *models.InstanceRule
}

// IsSuccess returns true when this rule create o k response has a 2xx status code
func (o *RuleCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rule create o k response has a 3xx status code
func (o *RuleCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create o k response has a 4xx status code
func (o *RuleCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rule create o k response has a 5xx status code
func (o *RuleCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rule create o k response a status code equal to that given
func (o *RuleCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rule create o k response
func (o *RuleCreateOK) Code() int {
	return 200
}

func (o *RuleCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateOK %s", 200, payload)
}

func (o *RuleCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateOK %s", 200, payload)
}

func (o *RuleCreateOK) GetPayload() *models.InstanceRule {
	return o.Payload
}

func (o *RuleCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRuleCreateBadRequest creates a RuleCreateBadRequest with default headers values
func NewRuleCreateBadRequest() *RuleCreateBadRequest {
	return &RuleCreateBadRequest{}
}

/*
RuleCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type RuleCreateBadRequest struct {
}

// IsSuccess returns true when this rule create bad request response has a 2xx status code
func (o *RuleCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule create bad request response has a 3xx status code
func (o *RuleCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create bad request response has a 4xx status code
func (o *RuleCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule create bad request response has a 5xx status code
func (o *RuleCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rule create bad request response a status code equal to that given
func (o *RuleCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rule create bad request response
func (o *RuleCreateBadRequest) Code() int {
	return 400
}

func (o *RuleCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateBadRequest", 400)
}

func (o *RuleCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateBadRequest", 400)
}

func (o *RuleCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleCreateUnauthorized creates a RuleCreateUnauthorized with default headers values
func NewRuleCreateUnauthorized() *RuleCreateUnauthorized {
	return &RuleCreateUnauthorized{}
}

/*
RuleCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type RuleCreateUnauthorized struct {
}

// IsSuccess returns true when this rule create unauthorized response has a 2xx status code
func (o *RuleCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule create unauthorized response has a 3xx status code
func (o *RuleCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create unauthorized response has a 4xx status code
func (o *RuleCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule create unauthorized response has a 5xx status code
func (o *RuleCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this rule create unauthorized response a status code equal to that given
func (o *RuleCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the rule create unauthorized response
func (o *RuleCreateUnauthorized) Code() int {
	return 401
}

func (o *RuleCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateUnauthorized", 401)
}

func (o *RuleCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateUnauthorized", 401)
}

func (o *RuleCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleCreateForbidden creates a RuleCreateForbidden with default headers values
func NewRuleCreateForbidden() *RuleCreateForbidden {
	return &RuleCreateForbidden{}
}

/*
RuleCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type RuleCreateForbidden struct {
}

// IsSuccess returns true when this rule create forbidden response has a 2xx status code
func (o *RuleCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule create forbidden response has a 3xx status code
func (o *RuleCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create forbidden response has a 4xx status code
func (o *RuleCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule create forbidden response has a 5xx status code
func (o *RuleCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this rule create forbidden response a status code equal to that given
func (o *RuleCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the rule create forbidden response
func (o *RuleCreateForbidden) Code() int {
	return 403
}

func (o *RuleCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateForbidden", 403)
}

func (o *RuleCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateForbidden", 403)
}

func (o *RuleCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleCreateNotFound creates a RuleCreateNotFound with default headers values
func NewRuleCreateNotFound() *RuleCreateNotFound {
	return &RuleCreateNotFound{}
}

/*
RuleCreateNotFound describes a response with status code 404, with default header values.

not found
*/
type RuleCreateNotFound struct {
}

// IsSuccess returns true when this rule create not found response has a 2xx status code
func (o *RuleCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule create not found response has a 3xx status code
func (o *RuleCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create not found response has a 4xx status code
func (o *RuleCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule create not found response has a 5xx status code
func (o *RuleCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rule create not found response a status code equal to that given
func (o *RuleCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rule create not found response
func (o *RuleCreateNotFound) Code() int {
	return 404
}

func (o *RuleCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateNotFound", 404)
}

func (o *RuleCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateNotFound", 404)
}

func (o *RuleCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleCreateNotAcceptable creates a RuleCreateNotAcceptable with default headers values
func NewRuleCreateNotAcceptable() *RuleCreateNotAcceptable {
	return &RuleCreateNotAcceptable{}
}

/*
RuleCreateNotAcceptable describes a response with status code 406, with default header values.

not acceptable
*/
type RuleCreateNotAcceptable struct {
}

// IsSuccess returns true when this rule create not acceptable response has a 2xx status code
func (o *RuleCreateNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule create not acceptable response has a 3xx status code
func (o *RuleCreateNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create not acceptable response has a 4xx status code
func (o *RuleCreateNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this rule create not acceptable response has a 5xx status code
func (o *RuleCreateNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this rule create not acceptable response a status code equal to that given
func (o *RuleCreateNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the rule create not acceptable response
func (o *RuleCreateNotAcceptable) Code() int {
	return 406
}

func (o *RuleCreateNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateNotAcceptable", 406)
}

func (o *RuleCreateNotAcceptable) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateNotAcceptable", 406)
}

func (o *RuleCreateNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRuleCreateInternalServerError creates a RuleCreateInternalServerError with default headers values
func NewRuleCreateInternalServerError() *RuleCreateInternalServerError {
	return &RuleCreateInternalServerError{}
}

/*
RuleCreateInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type RuleCreateInternalServerError struct {
}

// IsSuccess returns true when this rule create internal server error response has a 2xx status code
func (o *RuleCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rule create internal server error response has a 3xx status code
func (o *RuleCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rule create internal server error response has a 4xx status code
func (o *RuleCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rule create internal server error response has a 5xx status code
func (o *RuleCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rule create internal server error response a status code equal to that given
func (o *RuleCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rule create internal server error response
func (o *RuleCreateInternalServerError) Code() int {
	return 500
}

func (o *RuleCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateInternalServerError", 500)
}

func (o *RuleCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/admin/instance/rules][%d] ruleCreateInternalServerError", 500)
}

func (o *RuleCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}