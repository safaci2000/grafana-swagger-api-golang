// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// TestAlertReader is a Reader for the TestAlert structure.
type TestAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestAlertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTestAlertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTestAlertUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTestAlertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTestAlertOK creates a TestAlertOK with default headers values
func NewTestAlertOK() *TestAlertOK {
	return &TestAlertOK{}
}

/*
TestAlertOK describes a response with status code 200, with default header values.

(empty)
*/
type TestAlertOK struct {
	Payload *models.AlertTestResult
}

func (o *TestAlertOK) Error() string {
	return fmt.Sprintf("[POST /alerts/test][%d] testAlertOK  %+v", 200, o.Payload)
}
func (o *TestAlertOK) GetPayload() *models.AlertTestResult {
	return o.Payload
}

func (o *TestAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAlertBadRequest creates a TestAlertBadRequest with default headers values
func NewTestAlertBadRequest() *TestAlertBadRequest {
	return &TestAlertBadRequest{}
}

/*
TestAlertBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type TestAlertBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *TestAlertBadRequest) Error() string {
	return fmt.Sprintf("[POST /alerts/test][%d] testAlertBadRequest  %+v", 400, o.Payload)
}
func (o *TestAlertBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *TestAlertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAlertForbidden creates a TestAlertForbidden with default headers values
func NewTestAlertForbidden() *TestAlertForbidden {
	return &TestAlertForbidden{}
}

/*
TestAlertForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type TestAlertForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *TestAlertForbidden) Error() string {
	return fmt.Sprintf("[POST /alerts/test][%d] testAlertForbidden  %+v", 403, o.Payload)
}
func (o *TestAlertForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *TestAlertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAlertUnprocessableEntity creates a TestAlertUnprocessableEntity with default headers values
func NewTestAlertUnprocessableEntity() *TestAlertUnprocessableEntity {
	return &TestAlertUnprocessableEntity{}
}

/*
TestAlertUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntityError
*/
type TestAlertUnprocessableEntity struct {
	Payload *models.ErrorResponseBody
}

func (o *TestAlertUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /alerts/test][%d] testAlertUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *TestAlertUnprocessableEntity) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *TestAlertUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestAlertInternalServerError creates a TestAlertInternalServerError with default headers values
func NewTestAlertInternalServerError() *TestAlertInternalServerError {
	return &TestAlertInternalServerError{}
}

/*
TestAlertInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type TestAlertInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *TestAlertInternalServerError) Error() string {
	return fmt.Sprintf("[POST /alerts/test][%d] testAlertInternalServerError  %+v", 500, o.Payload)
}
func (o *TestAlertInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *TestAlertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
