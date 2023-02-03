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

// RestoreDashboardVersionByUIDReader is a Reader for the RestoreDashboardVersionByUID structure.
type RestoreDashboardVersionByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreDashboardVersionByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreDashboardVersionByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestoreDashboardVersionByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestoreDashboardVersionByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreDashboardVersionByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestoreDashboardVersionByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestoreDashboardVersionByUIDOK creates a RestoreDashboardVersionByUIDOK with default headers values
func NewRestoreDashboardVersionByUIDOK() *RestoreDashboardVersionByUIDOK {
	return &RestoreDashboardVersionByUIDOK{}
}

/*
RestoreDashboardVersionByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type RestoreDashboardVersionByUIDOK struct {
	Payload *models.RestoreDashboardVersionByUIDOKBody
}

func (o *RestoreDashboardVersionByUIDOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/uid/{uid}/restore][%d] restoreDashboardVersionByUidOK  %+v", 200, o.Payload)
}
func (o *RestoreDashboardVersionByUIDOK) GetPayload() *models.RestoreDashboardVersionByUIDOKBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreDashboardVersionByUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByUIDUnauthorized creates a RestoreDashboardVersionByUIDUnauthorized with default headers values
func NewRestoreDashboardVersionByUIDUnauthorized() *RestoreDashboardVersionByUIDUnauthorized {
	return &RestoreDashboardVersionByUIDUnauthorized{}
}

/*
RestoreDashboardVersionByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RestoreDashboardVersionByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dashboards/uid/{uid}/restore][%d] restoreDashboardVersionByUidUnauthorized  %+v", 401, o.Payload)
}
func (o *RestoreDashboardVersionByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByUIDForbidden creates a RestoreDashboardVersionByUIDForbidden with default headers values
func NewRestoreDashboardVersionByUIDForbidden() *RestoreDashboardVersionByUIDForbidden {
	return &RestoreDashboardVersionByUIDForbidden{}
}

/*
RestoreDashboardVersionByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RestoreDashboardVersionByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByUIDForbidden) Error() string {
	return fmt.Sprintf("[POST /dashboards/uid/{uid}/restore][%d] restoreDashboardVersionByUidForbidden  %+v", 403, o.Payload)
}
func (o *RestoreDashboardVersionByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByUIDNotFound creates a RestoreDashboardVersionByUIDNotFound with default headers values
func NewRestoreDashboardVersionByUIDNotFound() *RestoreDashboardVersionByUIDNotFound {
	return &RestoreDashboardVersionByUIDNotFound{}
}

/*
RestoreDashboardVersionByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RestoreDashboardVersionByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByUIDNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards/uid/{uid}/restore][%d] restoreDashboardVersionByUidNotFound  %+v", 404, o.Payload)
}
func (o *RestoreDashboardVersionByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDashboardVersionByUIDInternalServerError creates a RestoreDashboardVersionByUIDInternalServerError with default headers values
func NewRestoreDashboardVersionByUIDInternalServerError() *RestoreDashboardVersionByUIDInternalServerError {
	return &RestoreDashboardVersionByUIDInternalServerError{}
}

/*
RestoreDashboardVersionByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RestoreDashboardVersionByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RestoreDashboardVersionByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dashboards/uid/{uid}/restore][%d] restoreDashboardVersionByUidInternalServerError  %+v", 500, o.Payload)
}
func (o *RestoreDashboardVersionByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RestoreDashboardVersionByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
