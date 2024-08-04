// Code generated by go-swagger; DO NOT EDIT.

package nr_well_known

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

// WebfingerGetReader is a Reader for the WebfingerGet structure.
type WebfingerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebfingerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebfingerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /.well-known/webfinger] webfingerGet", response, response.Code())
	}
}

// NewWebfingerGetOK creates a WebfingerGetOK with default headers values
func NewWebfingerGetOK() *WebfingerGetOK {
	return &WebfingerGetOK{}
}

/*
WebfingerGetOK describes a response with status code 200, with default header values.

WebfingerGetOK webfinger get o k
*/
type WebfingerGetOK struct {
	Payload *models.WellKnownResponse
}

// IsSuccess returns true when this webfinger get o k response has a 2xx status code
func (o *WebfingerGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webfinger get o k response has a 3xx status code
func (o *WebfingerGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webfinger get o k response has a 4xx status code
func (o *WebfingerGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webfinger get o k response has a 5xx status code
func (o *WebfingerGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webfinger get o k response a status code equal to that given
func (o *WebfingerGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webfinger get o k response
func (o *WebfingerGetOK) Code() int {
	return 200
}

func (o *WebfingerGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /.well-known/webfinger][%d] webfingerGetOK %s", 200, payload)
}

func (o *WebfingerGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /.well-known/webfinger][%d] webfingerGetOK %s", 200, payload)
}

func (o *WebfingerGetOK) GetPayload() *models.WellKnownResponse {
	return o.Payload
}

func (o *WebfingerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WellKnownResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}