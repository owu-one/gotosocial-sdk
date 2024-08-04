// Code generated by go-swagger; DO NOT EDIT.

package timelines

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

// HomeTimelineReader is a Reader for the HomeTimeline structure.
type HomeTimelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HomeTimelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHomeTimelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHomeTimelineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHomeTimelineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/timelines/home] homeTimeline", response, response.Code())
	}
}

// NewHomeTimelineOK creates a HomeTimelineOK with default headers values
func NewHomeTimelineOK() *HomeTimelineOK {
	return &HomeTimelineOK{}
}

/*
HomeTimelineOK describes a response with status code 200, with default header values.

Array of statuses.
*/
type HomeTimelineOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.Status
}

// IsSuccess returns true when this home timeline o k response has a 2xx status code
func (o *HomeTimelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this home timeline o k response has a 3xx status code
func (o *HomeTimelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this home timeline o k response has a 4xx status code
func (o *HomeTimelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this home timeline o k response has a 5xx status code
func (o *HomeTimelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this home timeline o k response a status code equal to that given
func (o *HomeTimelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the home timeline o k response
func (o *HomeTimelineOK) Code() int {
	return 200
}

func (o *HomeTimelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/timelines/home][%d] homeTimelineOK %s", 200, payload)
}

func (o *HomeTimelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/timelines/home][%d] homeTimelineOK %s", 200, payload)
}

func (o *HomeTimelineOK) GetPayload() []*models.Status {
	return o.Payload
}

func (o *HomeTimelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewHomeTimelineBadRequest creates a HomeTimelineBadRequest with default headers values
func NewHomeTimelineBadRequest() *HomeTimelineBadRequest {
	return &HomeTimelineBadRequest{}
}

/*
HomeTimelineBadRequest describes a response with status code 400, with default header values.

bad request
*/
type HomeTimelineBadRequest struct {
}

// IsSuccess returns true when this home timeline bad request response has a 2xx status code
func (o *HomeTimelineBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this home timeline bad request response has a 3xx status code
func (o *HomeTimelineBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this home timeline bad request response has a 4xx status code
func (o *HomeTimelineBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this home timeline bad request response has a 5xx status code
func (o *HomeTimelineBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this home timeline bad request response a status code equal to that given
func (o *HomeTimelineBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the home timeline bad request response
func (o *HomeTimelineBadRequest) Code() int {
	return 400
}

func (o *HomeTimelineBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/timelines/home][%d] homeTimelineBadRequest", 400)
}

func (o *HomeTimelineBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/timelines/home][%d] homeTimelineBadRequest", 400)
}

func (o *HomeTimelineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHomeTimelineUnauthorized creates a HomeTimelineUnauthorized with default headers values
func NewHomeTimelineUnauthorized() *HomeTimelineUnauthorized {
	return &HomeTimelineUnauthorized{}
}

/*
HomeTimelineUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type HomeTimelineUnauthorized struct {
}

// IsSuccess returns true when this home timeline unauthorized response has a 2xx status code
func (o *HomeTimelineUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this home timeline unauthorized response has a 3xx status code
func (o *HomeTimelineUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this home timeline unauthorized response has a 4xx status code
func (o *HomeTimelineUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this home timeline unauthorized response has a 5xx status code
func (o *HomeTimelineUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this home timeline unauthorized response a status code equal to that given
func (o *HomeTimelineUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the home timeline unauthorized response
func (o *HomeTimelineUnauthorized) Code() int {
	return 401
}

func (o *HomeTimelineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/timelines/home][%d] homeTimelineUnauthorized", 401)
}

func (o *HomeTimelineUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/timelines/home][%d] homeTimelineUnauthorized", 401)
}

func (o *HomeTimelineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
