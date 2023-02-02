// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// GetReportsReader is a Reader for the GetReports structure.
type GetReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportsOK creates a GetReportsOK with default headers values
func NewGetReportsOK() *GetReportsOK {
	return &GetReportsOK{}
}

/*
GetReportsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetReportsOK struct {
	Payload []*models.ConfigDTO
}

func (o *GetReportsOK) Error() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsOK  %+v", 200, o.Payload)
}
func (o *GetReportsOK) GetPayload() []*models.ConfigDTO {
	return o.Payload
}

func (o *GetReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsUnauthorized creates a GetReportsUnauthorized with default headers values
func NewGetReportsUnauthorized() *GetReportsUnauthorized {
	return &GetReportsUnauthorized{}
}

/*
GetReportsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetReportsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReportsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsForbidden creates a GetReportsForbidden with default headers values
func NewGetReportsForbidden() *GetReportsForbidden {
	return &GetReportsForbidden{}
}

/*
GetReportsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetReportsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportsForbidden) Error() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsForbidden  %+v", 403, o.Payload)
}
func (o *GetReportsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsInternalServerError creates a GetReportsInternalServerError with default headers values
func NewGetReportsInternalServerError() *GetReportsInternalServerError {
	return &GetReportsInternalServerError{}
}

/*
GetReportsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetReportsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports][%d] getReportsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReportsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
