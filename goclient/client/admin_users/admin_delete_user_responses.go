// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AdminDeleteUserReader is a Reader for the AdminDeleteUser structure.
type AdminDeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminDeleteUserOK creates a AdminDeleteUserOK with default headers values
func NewAdminDeleteUserOK() *AdminDeleteUserOK {
	return &AdminDeleteUserOK{}
}

/*
AdminDeleteUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminDeleteUserOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AdminDeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserOK  %+v", 200, o.Payload)
}
func (o *AdminDeleteUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserUnauthorized creates a AdminDeleteUserUnauthorized with default headers values
func NewAdminDeleteUserUnauthorized() *AdminDeleteUserUnauthorized {
	return &AdminDeleteUserUnauthorized{}
}

/*
AdminDeleteUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminDeleteUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminDeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminDeleteUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserForbidden creates a AdminDeleteUserForbidden with default headers values
func NewAdminDeleteUserForbidden() *AdminDeleteUserForbidden {
	return &AdminDeleteUserForbidden{}
}

/*
AdminDeleteUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminDeleteUserForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminDeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserForbidden  %+v", 403, o.Payload)
}
func (o *AdminDeleteUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserNotFound creates a AdminDeleteUserNotFound with default headers values
func NewAdminDeleteUserNotFound() *AdminDeleteUserNotFound {
	return &AdminDeleteUserNotFound{}
}

/*
AdminDeleteUserNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AdminDeleteUserNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminDeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserNotFound  %+v", 404, o.Payload)
}
func (o *AdminDeleteUserNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteUserInternalServerError creates a AdminDeleteUserInternalServerError with default headers values
func NewAdminDeleteUserInternalServerError() *AdminDeleteUserInternalServerError {
	return &AdminDeleteUserInternalServerError{}
}

/*
AdminDeleteUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminDeleteUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminDeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /admin/users/{user_id}][%d] adminDeleteUserInternalServerError  %+v", 500, o.Payload)
}
func (o *AdminDeleteUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminDeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
