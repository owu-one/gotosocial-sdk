// Code generated by go-swagger; DO NOT EDIT.

package tokens

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

// TokensInfoGetReader is a Reader for the TokensInfoGet structure.
type TokensInfoGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokensInfoGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokensInfoGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTokensInfoGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTokensInfoGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/tokens] tokensInfoGet", response, response.Code())
	}
}

// NewTokensInfoGetOK creates a TokensInfoGetOK with default headers values
func NewTokensInfoGetOK() *TokensInfoGetOK {
	return &TokensInfoGetOK{}
}

/*
TokensInfoGetOK describes a response with status code 200, with default header values.

Array of token info entries.
*/
type TokensInfoGetOK struct {

	/* Links to the next and previous queries.
	 */
	Link string

	Payload []*models.TokenInfo
}

// IsSuccess returns true when this tokens info get o k response has a 2xx status code
func (o *TokensInfoGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tokens info get o k response has a 3xx status code
func (o *TokensInfoGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens info get o k response has a 4xx status code
func (o *TokensInfoGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tokens info get o k response has a 5xx status code
func (o *TokensInfoGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens info get o k response a status code equal to that given
func (o *TokensInfoGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tokens info get o k response
func (o *TokensInfoGetOK) Code() int {
	return 200
}

func (o *TokensInfoGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokensInfoGetOK %s", 200, payload)
}

func (o *TokensInfoGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokensInfoGetOK %s", 200, payload)
}

func (o *TokensInfoGetOK) GetPayload() []*models.TokenInfo {
	return o.Payload
}

func (o *TokensInfoGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTokensInfoGetBadRequest creates a TokensInfoGetBadRequest with default headers values
func NewTokensInfoGetBadRequest() *TokensInfoGetBadRequest {
	return &TokensInfoGetBadRequest{}
}

/*
TokensInfoGetBadRequest describes a response with status code 400, with default header values.

bad request
*/
type TokensInfoGetBadRequest struct {
}

// IsSuccess returns true when this tokens info get bad request response has a 2xx status code
func (o *TokensInfoGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens info get bad request response has a 3xx status code
func (o *TokensInfoGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens info get bad request response has a 4xx status code
func (o *TokensInfoGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens info get bad request response has a 5xx status code
func (o *TokensInfoGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens info get bad request response a status code equal to that given
func (o *TokensInfoGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the tokens info get bad request response
func (o *TokensInfoGetBadRequest) Code() int {
	return 400
}

func (o *TokensInfoGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokensInfoGetBadRequest", 400)
}

func (o *TokensInfoGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokensInfoGetBadRequest", 400)
}

func (o *TokensInfoGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTokensInfoGetUnauthorized creates a TokensInfoGetUnauthorized with default headers values
func NewTokensInfoGetUnauthorized() *TokensInfoGetUnauthorized {
	return &TokensInfoGetUnauthorized{}
}

/*
TokensInfoGetUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type TokensInfoGetUnauthorized struct {
}

// IsSuccess returns true when this tokens info get unauthorized response has a 2xx status code
func (o *TokensInfoGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tokens info get unauthorized response has a 3xx status code
func (o *TokensInfoGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tokens info get unauthorized response has a 4xx status code
func (o *TokensInfoGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this tokens info get unauthorized response has a 5xx status code
func (o *TokensInfoGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this tokens info get unauthorized response a status code equal to that given
func (o *TokensInfoGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the tokens info get unauthorized response
func (o *TokensInfoGetUnauthorized) Code() int {
	return 401
}

func (o *TokensInfoGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokensInfoGetUnauthorized", 401)
}

func (o *TokensInfoGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/tokens][%d] tokensInfoGetUnauthorized", 401)
}

func (o *TokensInfoGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
