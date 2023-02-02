// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// GetUserOrgListReader is a Reader for the GetUserOrgList structure.
type GetUserOrgListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserOrgListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOrgListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserOrgListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserOrgListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserOrgListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserOrgListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserOrgListOK creates a GetUserOrgListOK with default headers values
func NewGetUserOrgListOK() *GetUserOrgListOK {
	return &GetUserOrgListOK{}
}

/*
GetUserOrgListOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserOrgListOK struct {
	Payload []*models.UserOrgDTO
}

func (o *GetUserOrgListOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/orgs][%d] getUserOrgListOK  %+v", 200, o.Payload)
}
func (o *GetUserOrgListOK) GetPayload() []*models.UserOrgDTO {
	return o.Payload
}

func (o *GetUserOrgListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserOrgListUnauthorized creates a GetUserOrgListUnauthorized with default headers values
func NewGetUserOrgListUnauthorized() *GetUserOrgListUnauthorized {
	return &GetUserOrgListUnauthorized{}
}

/*
GetUserOrgListUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserOrgListUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserOrgListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/orgs][%d] getUserOrgListUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserOrgListUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserOrgListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserOrgListForbidden creates a GetUserOrgListForbidden with default headers values
func NewGetUserOrgListForbidden() *GetUserOrgListForbidden {
	return &GetUserOrgListForbidden{}
}

/*
GetUserOrgListForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetUserOrgListForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserOrgListForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/orgs][%d] getUserOrgListForbidden  %+v", 403, o.Payload)
}
func (o *GetUserOrgListForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserOrgListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserOrgListNotFound creates a GetUserOrgListNotFound with default headers values
func NewGetUserOrgListNotFound() *GetUserOrgListNotFound {
	return &GetUserOrgListNotFound{}
}

/*
GetUserOrgListNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetUserOrgListNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserOrgListNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/orgs][%d] getUserOrgListNotFound  %+v", 404, o.Payload)
}
func (o *GetUserOrgListNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserOrgListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserOrgListInternalServerError creates a GetUserOrgListInternalServerError with default headers values
func NewGetUserOrgListInternalServerError() *GetUserOrgListInternalServerError {
	return &GetUserOrgListInternalServerError{}
}

/*
GetUserOrgListInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserOrgListInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserOrgListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/orgs][%d] getUserOrgListInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserOrgListInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserOrgListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
