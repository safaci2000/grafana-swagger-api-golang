// Code generated by go-swagger; DO NOT EDIT.

package org_invites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AddOrgInviteReader is a Reader for the AddOrgInvite structure.
type AddOrgInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrgInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOrgInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOrgInviteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddOrgInviteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddOrgInviteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAddOrgInvitePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrgInviteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrgInviteOK creates a AddOrgInviteOK with default headers values
func NewAddOrgInviteOK() *AddOrgInviteOK {
	return &AddOrgInviteOK{}
}

/*
AddOrgInviteOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddOrgInviteOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AddOrgInviteOK) Error() string {
	return fmt.Sprintf("[POST /org/invites][%d] addOrgInviteOK  %+v", 200, o.Payload)
}
func (o *AddOrgInviteOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddOrgInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgInviteBadRequest creates a AddOrgInviteBadRequest with default headers values
func NewAddOrgInviteBadRequest() *AddOrgInviteBadRequest {
	return &AddOrgInviteBadRequest{}
}

/*
AddOrgInviteBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AddOrgInviteBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgInviteBadRequest) Error() string {
	return fmt.Sprintf("[POST /org/invites][%d] addOrgInviteBadRequest  %+v", 400, o.Payload)
}
func (o *AddOrgInviteBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgInviteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgInviteUnauthorized creates a AddOrgInviteUnauthorized with default headers values
func NewAddOrgInviteUnauthorized() *AddOrgInviteUnauthorized {
	return &AddOrgInviteUnauthorized{}
}

/*
AddOrgInviteUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddOrgInviteUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgInviteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /org/invites][%d] addOrgInviteUnauthorized  %+v", 401, o.Payload)
}
func (o *AddOrgInviteUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgInviteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgInviteForbidden creates a AddOrgInviteForbidden with default headers values
func NewAddOrgInviteForbidden() *AddOrgInviteForbidden {
	return &AddOrgInviteForbidden{}
}

/*
AddOrgInviteForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddOrgInviteForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgInviteForbidden) Error() string {
	return fmt.Sprintf("[POST /org/invites][%d] addOrgInviteForbidden  %+v", 403, o.Payload)
}
func (o *AddOrgInviteForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgInviteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgInvitePreconditionFailed creates a AddOrgInvitePreconditionFailed with default headers values
func NewAddOrgInvitePreconditionFailed() *AddOrgInvitePreconditionFailed {
	return &AddOrgInvitePreconditionFailed{}
}

/*
AddOrgInvitePreconditionFailed describes a response with status code 412, with default header values.

(empty)
*/
type AddOrgInvitePreconditionFailed struct {
}

func (o *AddOrgInvitePreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /org/invites][%d] addOrgInvitePreconditionFailed ", 412)
}

func (o *AddOrgInvitePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrgInviteInternalServerError creates a AddOrgInviteInternalServerError with default headers values
func NewAddOrgInviteInternalServerError() *AddOrgInviteInternalServerError {
	return &AddOrgInviteInternalServerError{}
}

/*
AddOrgInviteInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddOrgInviteInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgInviteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /org/invites][%d] addOrgInviteInternalServerError  %+v", 500, o.Payload)
}
func (o *AddOrgInviteInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgInviteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
