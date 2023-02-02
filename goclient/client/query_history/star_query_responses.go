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

// StarQueryReader is a Reader for the StarQuery structure.
type StarQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StarQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStarQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStarQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStarQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStarQueryOK creates a StarQueryOK with default headers values
func NewStarQueryOK() *StarQueryOK {
	return &StarQueryOK{}
}

/*
StarQueryOK describes a response with status code 200, with default header values.

(empty)
*/
type StarQueryOK struct {
	Payload *models.QueryHistoryResponse
}

func (o *StarQueryOK) Error() string {
	return fmt.Sprintf("[POST /query-history/star/{query_history_uid}][%d] starQueryOK  %+v", 200, o.Payload)
}
func (o *StarQueryOK) GetPayload() *models.QueryHistoryResponse {
	return o.Payload
}

func (o *StarQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarQueryUnauthorized creates a StarQueryUnauthorized with default headers values
func NewStarQueryUnauthorized() *StarQueryUnauthorized {
	return &StarQueryUnauthorized{}
}

/*
StarQueryUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type StarQueryUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *StarQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /query-history/star/{query_history_uid}][%d] starQueryUnauthorized  %+v", 401, o.Payload)
}
func (o *StarQueryUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarQueryInternalServerError creates a StarQueryInternalServerError with default headers values
func NewStarQueryInternalServerError() *StarQueryInternalServerError {
	return &StarQueryInternalServerError{}
}

/*
StarQueryInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type StarQueryInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *StarQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /query-history/star/{query_history_uid}][%d] starQueryInternalServerError  %+v", 500, o.Payload)
}
func (o *StarQueryInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
