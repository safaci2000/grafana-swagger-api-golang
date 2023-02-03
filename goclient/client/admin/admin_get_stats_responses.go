// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// AdminGetStatsReader is a Reader for the AdminGetStats structure.
type AdminGetStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGetStatsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminGetStatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminGetStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminGetStatsOK creates a AdminGetStatsOK with default headers values
func NewAdminGetStatsOK() *AdminGetStatsOK {
	return &AdminGetStatsOK{}
}

/*
AdminGetStatsOK describes a response with status code 200, with default header values.

(empty)
*/
type AdminGetStatsOK struct {
	Payload *models.AdminStats
}

func (o *AdminGetStatsOK) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsOK  %+v", 200, o.Payload)
}
func (o *AdminGetStatsOK) GetPayload() *models.AdminStats {
	return o.Payload
}

func (o *AdminGetStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetStatsUnauthorized creates a AdminGetStatsUnauthorized with default headers values
func NewAdminGetStatsUnauthorized() *AdminGetStatsUnauthorized {
	return &AdminGetStatsUnauthorized{}
}

/*
AdminGetStatsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminGetStatsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminGetStatsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminGetStatsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetStatsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetStatsForbidden creates a AdminGetStatsForbidden with default headers values
func NewAdminGetStatsForbidden() *AdminGetStatsForbidden {
	return &AdminGetStatsForbidden{}
}

/*
AdminGetStatsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminGetStatsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminGetStatsForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsForbidden  %+v", 403, o.Payload)
}
func (o *AdminGetStatsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetStatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetStatsInternalServerError creates a AdminGetStatsInternalServerError with default headers values
func NewAdminGetStatsInternalServerError() *AdminGetStatsInternalServerError {
	return &AdminGetStatsInternalServerError{}
}

/*
AdminGetStatsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminGetStatsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminGetStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/stats][%d] adminGetStatsInternalServerError  %+v", 500, o.Payload)
}
func (o *AdminGetStatsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
