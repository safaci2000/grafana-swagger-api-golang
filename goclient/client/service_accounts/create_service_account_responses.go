// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// CreateServiceAccountReader is a Reader for the CreateServiceAccount structure.
type CreateServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServiceAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateServiceAccountCreated creates a CreateServiceAccountCreated with default headers values
func NewCreateServiceAccountCreated() *CreateServiceAccountCreated {
	return &CreateServiceAccountCreated{}
}

/*
CreateServiceAccountCreated describes a response with status code 201, with default header values.

(empty)
*/
type CreateServiceAccountCreated struct {
	Payload *models.ServiceAccountDTO
}

func (o *CreateServiceAccountCreated) Error() string {
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountCreated  %+v", 201, o.Payload)
}
func (o *CreateServiceAccountCreated) GetPayload() *models.ServiceAccountDTO {
	return o.Payload
}

func (o *CreateServiceAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccountDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountBadRequest creates a CreateServiceAccountBadRequest with default headers values
func NewCreateServiceAccountBadRequest() *CreateServiceAccountBadRequest {
	return &CreateServiceAccountBadRequest{}
}

/*
CreateServiceAccountBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateServiceAccountBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountBadRequest  %+v", 400, o.Payload)
}
func (o *CreateServiceAccountBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountUnauthorized creates a CreateServiceAccountUnauthorized with default headers values
func NewCreateServiceAccountUnauthorized() *CreateServiceAccountUnauthorized {
	return &CreateServiceAccountUnauthorized{}
}

/*
CreateServiceAccountUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateServiceAccountUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateServiceAccountUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountForbidden creates a CreateServiceAccountForbidden with default headers values
func NewCreateServiceAccountForbidden() *CreateServiceAccountForbidden {
	return &CreateServiceAccountForbidden{}
}

/*
CreateServiceAccountForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateServiceAccountForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountForbidden  %+v", 403, o.Payload)
}
func (o *CreateServiceAccountForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountInternalServerError creates a CreateServiceAccountInternalServerError with default headers values
func NewCreateServiceAccountInternalServerError() *CreateServiceAccountInternalServerError {
	return &CreateServiceAccountInternalServerError{}
}

/*
CreateServiceAccountInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateServiceAccountInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateServiceAccountInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
