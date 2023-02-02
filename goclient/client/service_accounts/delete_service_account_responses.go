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

// DeleteServiceAccountReader is a Reader for the DeleteServiceAccount structure.
type DeleteServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceAccountOK creates a DeleteServiceAccountOK with default headers values
func NewDeleteServiceAccountOK() *DeleteServiceAccountOK {
	return &DeleteServiceAccountOK{}
}

/*
DeleteServiceAccountOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteServiceAccountOK struct {
	Payload *models.SuccessResponseBody
}

func (o *DeleteServiceAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}][%d] deleteServiceAccountOK  %+v", 200, o.Payload)
}
func (o *DeleteServiceAccountOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountBadRequest creates a DeleteServiceAccountBadRequest with default headers values
func NewDeleteServiceAccountBadRequest() *DeleteServiceAccountBadRequest {
	return &DeleteServiceAccountBadRequest{}
}

/*
DeleteServiceAccountBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DeleteServiceAccountBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}][%d] deleteServiceAccountBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteServiceAccountBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountUnauthorized creates a DeleteServiceAccountUnauthorized with default headers values
func NewDeleteServiceAccountUnauthorized() *DeleteServiceAccountUnauthorized {
	return &DeleteServiceAccountUnauthorized{}
}

/*
DeleteServiceAccountUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteServiceAccountUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}][%d] deleteServiceAccountUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteServiceAccountUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountForbidden creates a DeleteServiceAccountForbidden with default headers values
func NewDeleteServiceAccountForbidden() *DeleteServiceAccountForbidden {
	return &DeleteServiceAccountForbidden{}
}

/*
DeleteServiceAccountForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteServiceAccountForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}][%d] deleteServiceAccountForbidden  %+v", 403, o.Payload)
}
func (o *DeleteServiceAccountForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountInternalServerError creates a DeleteServiceAccountInternalServerError with default headers values
func NewDeleteServiceAccountInternalServerError() *DeleteServiceAccountInternalServerError {
	return &DeleteServiceAccountInternalServerError{}
}

/*
DeleteServiceAccountInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteServiceAccountInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /serviceaccounts/{serviceAccountId}][%d] deleteServiceAccountInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteServiceAccountInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
