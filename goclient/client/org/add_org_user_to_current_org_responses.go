// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AddOrgUserToCurrentOrgReader is a Reader for the AddOrgUserToCurrentOrg structure.
type AddOrgUserToCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrgUserToCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOrgUserToCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddOrgUserToCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddOrgUserToCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrgUserToCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrgUserToCurrentOrgOK creates a AddOrgUserToCurrentOrgOK with default headers values
func NewAddOrgUserToCurrentOrgOK() *AddOrgUserToCurrentOrgOK {
	return &AddOrgUserToCurrentOrgOK{}
}

/*
AddOrgUserToCurrentOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddOrgUserToCurrentOrgOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AddOrgUserToCurrentOrgOK) Error() string {
	return fmt.Sprintf("[POST /org/users][%d] addOrgUserToCurrentOrgOK  %+v", 200, o.Payload)
}
func (o *AddOrgUserToCurrentOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddOrgUserToCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgUserToCurrentOrgUnauthorized creates a AddOrgUserToCurrentOrgUnauthorized with default headers values
func NewAddOrgUserToCurrentOrgUnauthorized() *AddOrgUserToCurrentOrgUnauthorized {
	return &AddOrgUserToCurrentOrgUnauthorized{}
}

/*
AddOrgUserToCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddOrgUserToCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgUserToCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[POST /org/users][%d] addOrgUserToCurrentOrgUnauthorized  %+v", 401, o.Payload)
}
func (o *AddOrgUserToCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgUserToCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgUserToCurrentOrgForbidden creates a AddOrgUserToCurrentOrgForbidden with default headers values
func NewAddOrgUserToCurrentOrgForbidden() *AddOrgUserToCurrentOrgForbidden {
	return &AddOrgUserToCurrentOrgForbidden{}
}

/*
AddOrgUserToCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddOrgUserToCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgUserToCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[POST /org/users][%d] addOrgUserToCurrentOrgForbidden  %+v", 403, o.Payload)
}
func (o *AddOrgUserToCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgUserToCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrgUserToCurrentOrgInternalServerError creates a AddOrgUserToCurrentOrgInternalServerError with default headers values
func NewAddOrgUserToCurrentOrgInternalServerError() *AddOrgUserToCurrentOrgInternalServerError {
	return &AddOrgUserToCurrentOrgInternalServerError{}
}

/*
AddOrgUserToCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddOrgUserToCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AddOrgUserToCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[POST /org/users][%d] addOrgUserToCurrentOrgInternalServerError  %+v", 500, o.Payload)
}
func (o *AddOrgUserToCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddOrgUserToCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
