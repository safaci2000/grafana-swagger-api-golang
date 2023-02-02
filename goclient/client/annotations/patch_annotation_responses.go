// Code generated by go-swagger; DO NOT EDIT.

package annotations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// PatchAnnotationReader is a Reader for the PatchAnnotation structure.
type PatchAnnotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAnnotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAnnotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAnnotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAnnotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAnnotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAnnotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAnnotationOK creates a PatchAnnotationOK with default headers values
func NewPatchAnnotationOK() *PatchAnnotationOK {
	return &PatchAnnotationOK{}
}

/*
PatchAnnotationOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type PatchAnnotationOK struct {
	Payload *models.SuccessResponseBody
}

func (o *PatchAnnotationOK) Error() string {
	return fmt.Sprintf("[PATCH /annotations/{annotation_id}][%d] patchAnnotationOK  %+v", 200, o.Payload)
}
func (o *PatchAnnotationOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *PatchAnnotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnnotationUnauthorized creates a PatchAnnotationUnauthorized with default headers values
func NewPatchAnnotationUnauthorized() *PatchAnnotationUnauthorized {
	return &PatchAnnotationUnauthorized{}
}

/*
PatchAnnotationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PatchAnnotationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchAnnotationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /annotations/{annotation_id}][%d] patchAnnotationUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchAnnotationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchAnnotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnnotationForbidden creates a PatchAnnotationForbidden with default headers values
func NewPatchAnnotationForbidden() *PatchAnnotationForbidden {
	return &PatchAnnotationForbidden{}
}

/*
PatchAnnotationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type PatchAnnotationForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchAnnotationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /annotations/{annotation_id}][%d] patchAnnotationForbidden  %+v", 403, o.Payload)
}
func (o *PatchAnnotationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchAnnotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnnotationNotFound creates a PatchAnnotationNotFound with default headers values
func NewPatchAnnotationNotFound() *PatchAnnotationNotFound {
	return &PatchAnnotationNotFound{}
}

/*
PatchAnnotationNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type PatchAnnotationNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchAnnotationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /annotations/{annotation_id}][%d] patchAnnotationNotFound  %+v", 404, o.Payload)
}
func (o *PatchAnnotationNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchAnnotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnnotationInternalServerError creates a PatchAnnotationInternalServerError with default headers values
func NewPatchAnnotationInternalServerError() *PatchAnnotationInternalServerError {
	return &PatchAnnotationInternalServerError{}
}

/*
PatchAnnotationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PatchAnnotationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchAnnotationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /annotations/{annotation_id}][%d] patchAnnotationInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchAnnotationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchAnnotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
