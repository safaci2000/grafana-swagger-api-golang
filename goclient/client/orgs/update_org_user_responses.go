// Code generated by go-swagger; DO NOT EDIT.

package orgs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// UpdateOrgUserReader is a Reader for the UpdateOrgUser structure.
type UpdateOrgUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrgUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrgUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrgUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrgUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateOrgUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrgUserOK creates a UpdateOrgUserOK with default headers values
func NewUpdateOrgUserOK() *UpdateOrgUserOK {
	return &UpdateOrgUserOK{}
}

/*
UpdateOrgUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateOrgUserOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateOrgUserOK) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org_id}/users/{user_id}][%d] updateOrgUserOK  %+v", 200, o.Payload)
}
func (o *UpdateOrgUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserBadRequest creates a UpdateOrgUserBadRequest with default headers values
func NewUpdateOrgUserBadRequest() *UpdateOrgUserBadRequest {
	return &UpdateOrgUserBadRequest{}
}

/*
UpdateOrgUserBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateOrgUserBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org_id}/users/{user_id}][%d] updateOrgUserBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateOrgUserBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserUnauthorized creates a UpdateOrgUserUnauthorized with default headers values
func NewUpdateOrgUserUnauthorized() *UpdateOrgUserUnauthorized {
	return &UpdateOrgUserUnauthorized{}
}

/*
UpdateOrgUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateOrgUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org_id}/users/{user_id}][%d] updateOrgUserUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateOrgUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserForbidden creates a UpdateOrgUserForbidden with default headers values
func NewUpdateOrgUserForbidden() *UpdateOrgUserForbidden {
	return &UpdateOrgUserForbidden{}
}

/*
UpdateOrgUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateOrgUserForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org_id}/users/{user_id}][%d] updateOrgUserForbidden  %+v", 403, o.Payload)
}
func (o *UpdateOrgUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgUserInternalServerError creates a UpdateOrgUserInternalServerError with default headers values
func NewUpdateOrgUserInternalServerError() *UpdateOrgUserInternalServerError {
	return &UpdateOrgUserInternalServerError{}
}

/*
UpdateOrgUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateOrgUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateOrgUserInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org_id}/users/{user_id}][%d] updateOrgUserInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateOrgUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateOrgUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
