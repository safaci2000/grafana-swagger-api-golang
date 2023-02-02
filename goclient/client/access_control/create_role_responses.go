// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// CreateRoleReader is a Reader for the CreateRole structure.
type CreateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRoleCreated creates a CreateRoleCreated with default headers values
func NewCreateRoleCreated() *CreateRoleCreated {
	return &CreateRoleCreated{}
}

/*
CreateRoleCreated describes a response with status code 201, with default header values.

(empty)
*/
type CreateRoleCreated struct {
	Payload *models.RoleDTO
}

func (o *CreateRoleCreated) Error() string {
	return fmt.Sprintf("[POST /access-control/roles][%d] createRoleCreated  %+v", 201, o.Payload)
}
func (o *CreateRoleCreated) GetPayload() *models.RoleDTO {
	return o.Payload
}

func (o *CreateRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleBadRequest creates a CreateRoleBadRequest with default headers values
func NewCreateRoleBadRequest() *CreateRoleBadRequest {
	return &CreateRoleBadRequest{}
}

/*
CreateRoleBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateRoleBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /access-control/roles][%d] createRoleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRoleBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleForbidden creates a CreateRoleForbidden with default headers values
func NewCreateRoleForbidden() *CreateRoleForbidden {
	return &CreateRoleForbidden{}
}

/*
CreateRoleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateRoleForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /access-control/roles][%d] createRoleForbidden  %+v", 403, o.Payload)
}
func (o *CreateRoleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoleInternalServerError creates a CreateRoleInternalServerError with default headers values
func NewCreateRoleInternalServerError() *CreateRoleInternalServerError {
	return &CreateRoleInternalServerError{}
}

/*
CreateRoleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateRoleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access-control/roles][%d] createRoleInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRoleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
