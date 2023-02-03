// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// RestoreDashboardVersionByIDReader is a Reader for the RestoreDashboardVersionByID structure.
type RestoreDashboardVersionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreDashboardVersionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreDashboardVersionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestoreDashboardVersionByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestoreDashboardVersionByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreDashboardVersionByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestoreDashboardVersionByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestoreDashboardVersionByIDOK creates a RestoreDashboardVersionByIDOK with default headers values
func NewRestoreDashboardVersionByIDOK() *RestoreDashboardVersionByIDOK {
	return &RestoreDashboardVersionByIDOK{}
}

/*
RestoreDashboardVersionByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type RestoreDashboardVersionByIDOK struct {
	Payload *models.RestoreDashboardVersionByIDOKBody
}

func (o *RestoreDashboardVersionByIDOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/restore][%d] restoreDashboardVersionByIdOK  %+v", 200, o.Payload)
}
func (o *RestoreDashboardVersionByIDOK) GetPayload() *models.RestoreDashboardVersionByIDOKBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreDashboardVersionByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByIDUnauthorized creates a RestoreDashboardVersionByIDUnauthorized with default headers values
func NewRestoreDashboardVersionByIDUnauthorized() *RestoreDashboardVersionByIDUnauthorized {
	return &RestoreDashboardVersionByIDUnauthorized{}
}

/*
RestoreDashboardVersionByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RestoreDashboardVersionByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/restore][%d] restoreDashboardVersionByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *RestoreDashboardVersionByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByIDForbidden creates a RestoreDashboardVersionByIDForbidden with default headers values
func NewRestoreDashboardVersionByIDForbidden() *RestoreDashboardVersionByIDForbidden {
	return &RestoreDashboardVersionByIDForbidden{}
}

/*
RestoreDashboardVersionByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RestoreDashboardVersionByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByIDForbidden) Error() string {
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/restore][%d] restoreDashboardVersionByIdForbidden  %+v", 403, o.Payload)
}
func (o *RestoreDashboardVersionByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByIDNotFound creates a RestoreDashboardVersionByIDNotFound with default headers values
func NewRestoreDashboardVersionByIDNotFound() *RestoreDashboardVersionByIDNotFound {
	return &RestoreDashboardVersionByIDNotFound{}
}

/*
RestoreDashboardVersionByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RestoreDashboardVersionByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByIDNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/restore][%d] restoreDashboardVersionByIdNotFound  %+v", 404, o.Payload)
}
func (o *RestoreDashboardVersionByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByIDInternalServerError creates a RestoreDashboardVersionByIDInternalServerError with default headers values
func NewRestoreDashboardVersionByIDInternalServerError() *RestoreDashboardVersionByIDInternalServerError {
	return &RestoreDashboardVersionByIDInternalServerError{}
}

/*
RestoreDashboardVersionByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RestoreDashboardVersionByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dashboards/id/{DashboardID}/restore][%d] restoreDashboardVersionByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *RestoreDashboardVersionByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
