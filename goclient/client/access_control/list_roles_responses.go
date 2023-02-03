// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// ListRolesReader is a Reader for the ListRoles structure.
type ListRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRolesOK creates a ListRolesOK with default headers values
func NewListRolesOK() *ListRolesOK {
	return &ListRolesOK{}
}

/*
ListRolesOK describes a response with status code 200, with default header values.

(empty)
*/
type ListRolesOK struct {
	Payload []*models.RoleDTO
}

func (o *ListRolesOK) Error() string {
	return fmt.Sprintf("[GET /access-control/roles][%d] listRolesOK  %+v", 200, o.Payload)
}
func (o *ListRolesOK) GetPayload() []*models.RoleDTO {
	return o.Payload
}

func (o *ListRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRolesForbidden creates a ListRolesForbidden with default headers values
func NewListRolesForbidden() *ListRolesForbidden {
	return &ListRolesForbidden{}
}

/*
ListRolesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListRolesForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *ListRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /access-control/roles][%d] listRolesForbidden  %+v", 403, o.Payload)
}
func (o *ListRolesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRolesInternalServerError creates a ListRolesInternalServerError with default headers values
func NewListRolesInternalServerError() *ListRolesInternalServerError {
	return &ListRolesInternalServerError{}
}

/*
ListRolesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ListRolesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ListRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access-control/roles][%d] listRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListRolesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
