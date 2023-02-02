// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UpdateFolderReader is a Reader for the UpdateFolder structure.
type UpdateFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateFolderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFolderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateFolderConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFolderOK creates a UpdateFolderOK with default headers values
func NewUpdateFolderOK() *UpdateFolderOK {
	return &UpdateFolderOK{}
}

/*
UpdateFolderOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateFolderOK struct {
	Payload *models.Folder
}

func (o *UpdateFolderOK) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderOK  %+v", 200, o.Payload)
}
func (o *UpdateFolderOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *UpdateFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderBadRequest creates a UpdateFolderBadRequest with default headers values
func NewUpdateFolderBadRequest() *UpdateFolderBadRequest {
	return &UpdateFolderBadRequest{}
}

/*
UpdateFolderBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateFolderBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateFolderBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderUnauthorized creates a UpdateFolderUnauthorized with default headers values
func NewUpdateFolderUnauthorized() *UpdateFolderUnauthorized {
	return &UpdateFolderUnauthorized{}
}

/*
UpdateFolderUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateFolderUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateFolderUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderForbidden creates a UpdateFolderForbidden with default headers values
func NewUpdateFolderForbidden() *UpdateFolderForbidden {
	return &UpdateFolderForbidden{}
}

/*
UpdateFolderForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateFolderForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderForbidden) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderForbidden  %+v", 403, o.Payload)
}
func (o *UpdateFolderForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderNotFound creates a UpdateFolderNotFound with default headers values
func NewUpdateFolderNotFound() *UpdateFolderNotFound {
	return &UpdateFolderNotFound{}
}

/*
UpdateFolderNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateFolderNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderNotFound) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderNotFound  %+v", 404, o.Payload)
}
func (o *UpdateFolderNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderConflict creates a UpdateFolderConflict with default headers values
func NewUpdateFolderConflict() *UpdateFolderConflict {
	return &UpdateFolderConflict{}
}

/*
UpdateFolderConflict describes a response with status code 409, with default header values.

ConflictError
*/
type UpdateFolderConflict struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderConflict) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderConflict  %+v", 409, o.Payload)
}
func (o *UpdateFolderConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderInternalServerError creates a UpdateFolderInternalServerError with default headers values
func NewUpdateFolderInternalServerError() *UpdateFolderInternalServerError {
	return &UpdateFolderInternalServerError{}
}

/*
UpdateFolderInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateFolderInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateFolderInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /folders/{folder_uid}][%d] updateFolderInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateFolderInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
