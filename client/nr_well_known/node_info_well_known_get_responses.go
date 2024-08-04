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

// NodeInfoWellKnownGetReader is a Reader for the NodeInfoWellKnownGet structure.
type NodeInfoWellKnownGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeInfoWellKnownGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeInfoWellKnownGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /.well-known/nodeinfo] nodeInfoWellKnownGet", response, response.Code())
	}
}

// NewNodeInfoWellKnownGetOK creates a NodeInfoWellKnownGetOK with default headers values
func NewNodeInfoWellKnownGetOK() *NodeInfoWellKnownGetOK {
	return &NodeInfoWellKnownGetOK{}
}

/*
NodeInfoWellKnownGetOK describes a response with status code 200, with default header values.

NodeInfoWellKnownGetOK node info well known get o k
*/
type NodeInfoWellKnownGetOK struct {
	Payload *models.WellKnownResponse
}

// IsSuccess returns true when this node info well known get o k response has a 2xx status code
func (o *NodeInfoWellKnownGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node info well known get o k response has a 3xx status code
func (o *NodeInfoWellKnownGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node info well known get o k response has a 4xx status code
func (o *NodeInfoWellKnownGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node info well known get o k response has a 5xx status code
func (o *NodeInfoWellKnownGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node info well known get o k response a status code equal to that given
func (o *NodeInfoWellKnownGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node info well known get o k response
func (o *NodeInfoWellKnownGetOK) Code() int {
	return 200
}

func (o *NodeInfoWellKnownGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /.well-known/nodeinfo][%d] nodeInfoWellKnownGetOK %s", 200, payload)
}

func (o *NodeInfoWellKnownGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /.well-known/nodeinfo][%d] nodeInfoWellKnownGetOK %s", 200, payload)
}

func (o *NodeInfoWellKnownGetOK) GetPayload() *models.WellKnownResponse {
	return o.Payload
}

func (o *NodeInfoWellKnownGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WellKnownResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}