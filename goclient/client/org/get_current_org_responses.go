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

// GetCurrentOrgReader is a Reader for the GetCurrentOrg structure.
type GetCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentOrgOK creates a GetCurrentOrgOK with default headers values
func NewGetCurrentOrgOK() *GetCurrentOrgOK {
	return &GetCurrentOrgOK{}
}

/*
GetCurrentOrgOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCurrentOrgOK struct {
	Payload *models.OrgDetailsDTO
}

func (o *GetCurrentOrgOK) Error() string {
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgOK  %+v", 200, o.Payload)
}
func (o *GetCurrentOrgOK) GetPayload() *models.OrgDetailsDTO {
	return o.Payload
}

func (o *GetCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrgDetailsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgUnauthorized creates a GetCurrentOrgUnauthorized with default headers values
func NewGetCurrentOrgUnauthorized() *GetCurrentOrgUnauthorized {
	return &GetCurrentOrgUnauthorized{}
}

/*
GetCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgForbidden creates a GetCurrentOrgForbidden with default headers values
func NewGetCurrentOrgForbidden() *GetCurrentOrgForbidden {
	return &GetCurrentOrgForbidden{}
}

/*
GetCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgForbidden  %+v", 403, o.Payload)
}
func (o *GetCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgInternalServerError creates a GetCurrentOrgInternalServerError with default headers values
func NewGetCurrentOrgInternalServerError() *GetCurrentOrgInternalServerError {
	return &GetCurrentOrgInternalServerError{}
}

/*
GetCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[GET /org][%d] getCurrentOrgInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
