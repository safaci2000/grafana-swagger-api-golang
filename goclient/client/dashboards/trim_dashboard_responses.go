// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// TrimDashboardReader is a Reader for the TrimDashboard structure.
type TrimDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TrimDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTrimDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTrimDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTrimDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTrimDashboardOK creates a TrimDashboardOK with default headers values
func NewTrimDashboardOK() *TrimDashboardOK {
	return &TrimDashboardOK{}
}

/*
TrimDashboardOK describes a response with status code 200, with default header values.

(empty)
*/
type TrimDashboardOK struct {
	Payload *models.TrimDashboardFullWithMeta
}

func (o *TrimDashboardOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/trim][%d] trimDashboardOK  %+v", 200, o.Payload)
}
func (o *TrimDashboardOK) GetPayload() *models.TrimDashboardFullWithMeta {
	return o.Payload
}

func (o *TrimDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrimDashboardFullWithMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrimDashboardUnauthorized creates a TrimDashboardUnauthorized with default headers values
func NewTrimDashboardUnauthorized() *TrimDashboardUnauthorized {
	return &TrimDashboardUnauthorized{}
}

/*
TrimDashboardUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type TrimDashboardUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *TrimDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dashboards/trim][%d] trimDashboardUnauthorized  %+v", 401, o.Payload)
}
func (o *TrimDashboardUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *TrimDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTrimDashboardInternalServerError creates a TrimDashboardInternalServerError with default headers values
func NewTrimDashboardInternalServerError() *TrimDashboardInternalServerError {
	return &TrimDashboardInternalServerError{}
}

/*
TrimDashboardInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type TrimDashboardInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *TrimDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dashboards/trim][%d] trimDashboardInternalServerError  %+v", 500, o.Payload)
}
func (o *TrimDashboardInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *TrimDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
