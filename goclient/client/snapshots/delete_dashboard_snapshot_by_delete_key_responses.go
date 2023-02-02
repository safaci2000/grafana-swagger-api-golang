// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// DeleteDashboardSnapshotByDeleteKeyReader is a Reader for the DeleteDashboardSnapshotByDeleteKey structure.
type DeleteDashboardSnapshotByDeleteKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardSnapshotByDeleteKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDashboardSnapshotByDeleteKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDashboardSnapshotByDeleteKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDashboardSnapshotByDeleteKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDashboardSnapshotByDeleteKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDashboardSnapshotByDeleteKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDashboardSnapshotByDeleteKeyOK creates a DeleteDashboardSnapshotByDeleteKeyOK with default headers values
func NewDeleteDashboardSnapshotByDeleteKeyOK() *DeleteDashboardSnapshotByDeleteKeyOK {
	return &DeleteDashboardSnapshotByDeleteKeyOK{}
}

/*
DeleteDashboardSnapshotByDeleteKeyOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteDashboardSnapshotByDeleteKeyOK struct {
	Payload *models.SuccessResponseBody
}

func (o *DeleteDashboardSnapshotByDeleteKeyOK) Error() string {
	return fmt.Sprintf("[GET /snapshots-delete/{deleteKey}][%d] deleteDashboardSnapshotByDeleteKeyOK  %+v", 200, o.Payload)
}
func (o *DeleteDashboardSnapshotByDeleteKeyOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotByDeleteKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotByDeleteKeyUnauthorized creates a DeleteDashboardSnapshotByDeleteKeyUnauthorized with default headers values
func NewDeleteDashboardSnapshotByDeleteKeyUnauthorized() *DeleteDashboardSnapshotByDeleteKeyUnauthorized {
	return &DeleteDashboardSnapshotByDeleteKeyUnauthorized{}
}

/*
DeleteDashboardSnapshotByDeleteKeyUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteDashboardSnapshotByDeleteKeyUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteDashboardSnapshotByDeleteKeyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /snapshots-delete/{deleteKey}][%d] deleteDashboardSnapshotByDeleteKeyUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDashboardSnapshotByDeleteKeyUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotByDeleteKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotByDeleteKeyForbidden creates a DeleteDashboardSnapshotByDeleteKeyForbidden with default headers values
func NewDeleteDashboardSnapshotByDeleteKeyForbidden() *DeleteDashboardSnapshotByDeleteKeyForbidden {
	return &DeleteDashboardSnapshotByDeleteKeyForbidden{}
}

/*
DeleteDashboardSnapshotByDeleteKeyForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteDashboardSnapshotByDeleteKeyForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteDashboardSnapshotByDeleteKeyForbidden) Error() string {
	return fmt.Sprintf("[GET /snapshots-delete/{deleteKey}][%d] deleteDashboardSnapshotByDeleteKeyForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDashboardSnapshotByDeleteKeyForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotByDeleteKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotByDeleteKeyNotFound creates a DeleteDashboardSnapshotByDeleteKeyNotFound with default headers values
func NewDeleteDashboardSnapshotByDeleteKeyNotFound() *DeleteDashboardSnapshotByDeleteKeyNotFound {
	return &DeleteDashboardSnapshotByDeleteKeyNotFound{}
}

/*
DeleteDashboardSnapshotByDeleteKeyNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteDashboardSnapshotByDeleteKeyNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteDashboardSnapshotByDeleteKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /snapshots-delete/{deleteKey}][%d] deleteDashboardSnapshotByDeleteKeyNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDashboardSnapshotByDeleteKeyNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotByDeleteKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardSnapshotByDeleteKeyInternalServerError creates a DeleteDashboardSnapshotByDeleteKeyInternalServerError with default headers values
func NewDeleteDashboardSnapshotByDeleteKeyInternalServerError() *DeleteDashboardSnapshotByDeleteKeyInternalServerError {
	return &DeleteDashboardSnapshotByDeleteKeyInternalServerError{}
}

/*
DeleteDashboardSnapshotByDeleteKeyInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteDashboardSnapshotByDeleteKeyInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteDashboardSnapshotByDeleteKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /snapshots-delete/{deleteKey}][%d] deleteDashboardSnapshotByDeleteKeyInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDashboardSnapshotByDeleteKeyInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteDashboardSnapshotByDeleteKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
