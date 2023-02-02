// Code generated by go-swagger; DO NOT EDIT.

package folder_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UpdateFolderPermissionsReader is a Reader for the UpdateFolderPermissions structure.
type UpdateFolderPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFolderPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFolderPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateFolderPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFolderPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFolderPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateFolderPermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFolderPermissionsOK creates a UpdateFolderPermissionsOK with default headers values
func NewUpdateFolderPermissionsOK() *UpdateFolderPermissionsOK {
	return &UpdateFolderPermissionsOK{}
}

/*
UpdateFolderPermissionsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateFolderPermissionsOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateFolderPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/permissions][%d] updateFolderPermissionsOK  %+v", 200, o.Payload)
}
func (o *UpdateFolderPermissionsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateFolderPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderPermissionsUnauthorized creates a UpdateFolderPermissionsUnauthorized with default headers values
func NewUpdateFolderPermissionsUnauthorized() *UpdateFolderPermissionsUnauthorized {
	return &UpdateFolderPermissionsUnauthorized{}
}

/*
UpdateFolderPermissionsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateFolderPermissionsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/permissions][%d] updateFolderPermissionsUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateFolderPermissionsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderPermissionsForbidden creates a UpdateFolderPermissionsForbidden with default headers values
func NewUpdateFolderPermissionsForbidden() *UpdateFolderPermissionsForbidden {
	return &UpdateFolderPermissionsForbidden{}
}

/*
UpdateFolderPermissionsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateFolderPermissionsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderPermissionsForbidden) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/permissions][%d] updateFolderPermissionsForbidden  %+v", 403, o.Payload)
}
func (o *UpdateFolderPermissionsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderPermissionsNotFound creates a UpdateFolderPermissionsNotFound with default headers values
func NewUpdateFolderPermissionsNotFound() *UpdateFolderPermissionsNotFound {
	return &UpdateFolderPermissionsNotFound{}
}

/*
UpdateFolderPermissionsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateFolderPermissionsNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderPermissionsNotFound) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/permissions][%d] updateFolderPermissionsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateFolderPermissionsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderPermissionsInternalServerError creates a UpdateFolderPermissionsInternalServerError with default headers values
func NewUpdateFolderPermissionsInternalServerError() *UpdateFolderPermissionsInternalServerError {
	return &UpdateFolderPermissionsInternalServerError{}
}

/*
UpdateFolderPermissionsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateFolderPermissionsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderPermissionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/permissions][%d] updateFolderPermissionsInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateFolderPermissionsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderPermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
