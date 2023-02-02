// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AddDataSourceReader is a Reader for the AddDataSource structure.
type AddDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddDataSourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddDataSourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddDataSourceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDataSourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddDataSourceOK creates a AddDataSourceOK with default headers values
func NewAddDataSourceOK() *AddDataSourceOK {
	return &AddDataSourceOK{}
}

/*
AddDataSourceOK describes a response with status code 200, with default header values.

(empty)
*/
type AddDataSourceOK struct {
	Payload *models.AddDataSourceOKBody
}

func (o *AddDataSourceOK) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceOK  %+v", 200, o.Payload)
}
func (o *AddDataSourceOK) GetPayload() *models.AddDataSourceOKBody {
	return o.Payload
}

func (o *AddDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddDataSourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceUnauthorized creates a AddDataSourceUnauthorized with default headers values
func NewAddDataSourceUnauthorized() *AddDataSourceUnauthorized {
	return &AddDataSourceUnauthorized{}
}

/*
AddDataSourceUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddDataSourceUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AddDataSourceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceUnauthorized  %+v", 401, o.Payload)
}
func (o *AddDataSourceUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceForbidden creates a AddDataSourceForbidden with default headers values
func NewAddDataSourceForbidden() *AddDataSourceForbidden {
	return &AddDataSourceForbidden{}
}

/*
AddDataSourceForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddDataSourceForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AddDataSourceForbidden) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceForbidden  %+v", 403, o.Payload)
}
func (o *AddDataSourceForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceConflict creates a AddDataSourceConflict with default headers values
func NewAddDataSourceConflict() *AddDataSourceConflict {
	return &AddDataSourceConflict{}
}

/*
AddDataSourceConflict describes a response with status code 409, with default header values.

ConflictError
*/
type AddDataSourceConflict struct {
	Payload *models.ErrorResponseBody
}

func (o *AddDataSourceConflict) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceConflict  %+v", 409, o.Payload)
}
func (o *AddDataSourceConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDataSourceInternalServerError creates a AddDataSourceInternalServerError with default headers values
func NewAddDataSourceInternalServerError() *AddDataSourceInternalServerError {
	return &AddDataSourceInternalServerError{}
}

/*
AddDataSourceInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddDataSourceInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AddDataSourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasources][%d] addDataSourceInternalServerError  %+v", 500, o.Payload)
}
func (o *AddDataSourceInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddDataSourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
