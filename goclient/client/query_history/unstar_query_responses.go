// Code generated by go-swagger; DO NOT EDIT.

package query_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UnstarQueryReader is a Reader for the UnstarQuery structure.
type UnstarQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnstarQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnstarQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUnstarQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnstarQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnstarQueryOK creates a UnstarQueryOK with default headers values
func NewUnstarQueryOK() *UnstarQueryOK {
	return &UnstarQueryOK{}
}

/*
UnstarQueryOK describes a response with status code 200, with default header values.

(empty)
*/
type UnstarQueryOK struct {
	Payload *models.QueryHistoryResponse
}

func (o *UnstarQueryOK) Error() string {
	return fmt.Sprintf("[DELETE /query-history/star/{query_history_uid}][%d] unstarQueryOK  %+v", 200, o.Payload)
}
func (o *UnstarQueryOK) GetPayload() *models.QueryHistoryResponse {
	return o.Payload
}

func (o *UnstarQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnstarQueryUnauthorized creates a UnstarQueryUnauthorized with default headers values
func NewUnstarQueryUnauthorized() *UnstarQueryUnauthorized {
	return &UnstarQueryUnauthorized{}
}

/*
UnstarQueryUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UnstarQueryUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UnstarQueryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /query-history/star/{query_history_uid}][%d] unstarQueryUnauthorized  %+v", 401, o.Payload)
}
func (o *UnstarQueryUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UnstarQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnstarQueryInternalServerError creates a UnstarQueryInternalServerError with default headers values
func NewUnstarQueryInternalServerError() *UnstarQueryInternalServerError {
	return &UnstarQueryInternalServerError{}
}

/*
UnstarQueryInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UnstarQueryInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UnstarQueryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /query-history/star/{query_history_uid}][%d] unstarQueryInternalServerError  %+v", 500, o.Payload)
}
func (o *UnstarQueryInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UnstarQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
