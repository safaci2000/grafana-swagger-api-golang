// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// StarDashboardByUIDReader is a Reader for the StarDashboardByUID structure.
type StarDashboardByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StarDashboardByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStarDashboardByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStarDashboardByUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStarDashboardByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStarDashboardByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStarDashboardByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStarDashboardByUIDOK creates a StarDashboardByUIDOK with default headers values
func NewStarDashboardByUIDOK() *StarDashboardByUIDOK {
	return &StarDashboardByUIDOK{}
}

/*
StarDashboardByUIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type StarDashboardByUIDOK struct {
	Payload *models.SuccessResponseBody
}

func (o *StarDashboardByUIDOK) Error() string {
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidOK  %+v", 200, o.Payload)
}
func (o *StarDashboardByUIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDBadRequest creates a StarDashboardByUIDBadRequest with default headers values
func NewStarDashboardByUIDBadRequest() *StarDashboardByUIDBadRequest {
	return &StarDashboardByUIDBadRequest{}
}

/*
StarDashboardByUIDBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type StarDashboardByUIDBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *StarDashboardByUIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidBadRequest  %+v", 400, o.Payload)
}
func (o *StarDashboardByUIDBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDUnauthorized creates a StarDashboardByUIDUnauthorized with default headers values
func NewStarDashboardByUIDUnauthorized() *StarDashboardByUIDUnauthorized {
	return &StarDashboardByUIDUnauthorized{}
}

/*
StarDashboardByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type StarDashboardByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *StarDashboardByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidUnauthorized  %+v", 401, o.Payload)
}
func (o *StarDashboardByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDForbidden creates a StarDashboardByUIDForbidden with default headers values
func NewStarDashboardByUIDForbidden() *StarDashboardByUIDForbidden {
	return &StarDashboardByUIDForbidden{}
}

/*
StarDashboardByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type StarDashboardByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *StarDashboardByUIDForbidden) Error() string {
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidForbidden  %+v", 403, o.Payload)
}
func (o *StarDashboardByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStarDashboardByUIDInternalServerError creates a StarDashboardByUIDInternalServerError with default headers values
func NewStarDashboardByUIDInternalServerError() *StarDashboardByUIDInternalServerError {
	return &StarDashboardByUIDInternalServerError{}
}

/*
StarDashboardByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type StarDashboardByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *StarDashboardByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/stars/dashboard/uid/{dashboard_uid}][%d] starDashboardByUidInternalServerError  %+v", 500, o.Payload)
}
func (o *StarDashboardByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *StarDashboardByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
