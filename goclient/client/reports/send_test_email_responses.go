// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// SendTestEmailReader is a Reader for the SendTestEmail structure.
type SendTestEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendTestEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendTestEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendTestEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSendTestEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSendTestEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSendTestEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSendTestEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSendTestEmailOK creates a SendTestEmailOK with default headers values
func NewSendTestEmailOK() *SendTestEmailOK {
	return &SendTestEmailOK{}
}

/*
SendTestEmailOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type SendTestEmailOK struct {
	Payload *models.SuccessResponseBody
}

func (o *SendTestEmailOK) Error() string {
	return fmt.Sprintf("[POST /reports/test-email][%d] sendTestEmailOK  %+v", 200, o.Payload)
}
func (o *SendTestEmailOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *SendTestEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestEmailBadRequest creates a SendTestEmailBadRequest with default headers values
func NewSendTestEmailBadRequest() *SendTestEmailBadRequest {
	return &SendTestEmailBadRequest{}
}

/*
SendTestEmailBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type SendTestEmailBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *SendTestEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /reports/test-email][%d] sendTestEmailBadRequest  %+v", 400, o.Payload)
}
func (o *SendTestEmailBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SendTestEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestEmailUnauthorized creates a SendTestEmailUnauthorized with default headers values
func NewSendTestEmailUnauthorized() *SendTestEmailUnauthorized {
	return &SendTestEmailUnauthorized{}
}

/*
SendTestEmailUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SendTestEmailUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *SendTestEmailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reports/test-email][%d] sendTestEmailUnauthorized  %+v", 401, o.Payload)
}
func (o *SendTestEmailUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SendTestEmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestEmailForbidden creates a SendTestEmailForbidden with default headers values
func NewSendTestEmailForbidden() *SendTestEmailForbidden {
	return &SendTestEmailForbidden{}
}

/*
SendTestEmailForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SendTestEmailForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *SendTestEmailForbidden) Error() string {
	return fmt.Sprintf("[POST /reports/test-email][%d] sendTestEmailForbidden  %+v", 403, o.Payload)
}
func (o *SendTestEmailForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SendTestEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestEmailNotFound creates a SendTestEmailNotFound with default headers values
func NewSendTestEmailNotFound() *SendTestEmailNotFound {
	return &SendTestEmailNotFound{}
}

/*
SendTestEmailNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SendTestEmailNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *SendTestEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /reports/test-email][%d] sendTestEmailNotFound  %+v", 404, o.Payload)
}
func (o *SendTestEmailNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SendTestEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendTestEmailInternalServerError creates a SendTestEmailInternalServerError with default headers values
func NewSendTestEmailInternalServerError() *SendTestEmailInternalServerError {
	return &SendTestEmailInternalServerError{}
}

/*
SendTestEmailInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SendTestEmailInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SendTestEmailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /reports/test-email][%d] sendTestEmailInternalServerError  %+v", 500, o.Payload)
}
func (o *SendTestEmailInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SendTestEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
