// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// SetHelpFlagReader is a Reader for the SetHelpFlag structure.
type SetHelpFlagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetHelpFlagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetHelpFlagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetHelpFlagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetHelpFlagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetHelpFlagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetHelpFlagOK creates a SetHelpFlagOK with default headers values
func NewSetHelpFlagOK() *SetHelpFlagOK {
	return &SetHelpFlagOK{}
}

/*
SetHelpFlagOK describes a response with status code 200, with default header values.

(empty)
*/
type SetHelpFlagOK struct {
	Payload *models.SetHelpFlagOKBody
}

func (o *SetHelpFlagOK) Error() string {
	return fmt.Sprintf("[PUT /user/helpflags/{flag_id}][%d] setHelpFlagOK  %+v", 200, o.Payload)
}
func (o *SetHelpFlagOK) GetPayload() *models.SetHelpFlagOKBody {
	return o.Payload
}

func (o *SetHelpFlagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetHelpFlagOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetHelpFlagUnauthorized creates a SetHelpFlagUnauthorized with default headers values
func NewSetHelpFlagUnauthorized() *SetHelpFlagUnauthorized {
	return &SetHelpFlagUnauthorized{}
}

/*
SetHelpFlagUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SetHelpFlagUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *SetHelpFlagUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user/helpflags/{flag_id}][%d] setHelpFlagUnauthorized  %+v", 401, o.Payload)
}
func (o *SetHelpFlagUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetHelpFlagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetHelpFlagForbidden creates a SetHelpFlagForbidden with default headers values
func NewSetHelpFlagForbidden() *SetHelpFlagForbidden {
	return &SetHelpFlagForbidden{}
}

/*
SetHelpFlagForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SetHelpFlagForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *SetHelpFlagForbidden) Error() string {
	return fmt.Sprintf("[PUT /user/helpflags/{flag_id}][%d] setHelpFlagForbidden  %+v", 403, o.Payload)
}
func (o *SetHelpFlagForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetHelpFlagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetHelpFlagInternalServerError creates a SetHelpFlagInternalServerError with default headers values
func NewSetHelpFlagInternalServerError() *SetHelpFlagInternalServerError {
	return &SetHelpFlagInternalServerError{}
}

/*
SetHelpFlagInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SetHelpFlagInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SetHelpFlagInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user/helpflags/{flag_id}][%d] setHelpFlagInternalServerError  %+v", 500, o.Payload)
}
func (o *SetHelpFlagInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetHelpFlagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
