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

// SearchUsersWithPagingReader is a Reader for the SearchUsersWithPaging structure.
type SearchUsersWithPagingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUsersWithPagingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUsersWithPagingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchUsersWithPagingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchUsersWithPagingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchUsersWithPagingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUsersWithPagingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchUsersWithPagingOK creates a SearchUsersWithPagingOK with default headers values
func NewSearchUsersWithPagingOK() *SearchUsersWithPagingOK {
	return &SearchUsersWithPagingOK{}
}

/*
SearchUsersWithPagingOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchUsersWithPagingOK struct {
	Payload *models.SearchUserQueryResult
}

func (o *SearchUsersWithPagingOK) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingOK  %+v", 200, o.Payload)
}
func (o *SearchUsersWithPagingOK) GetPayload() *models.SearchUserQueryResult {
	return o.Payload
}

func (o *SearchUsersWithPagingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchUserQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingUnauthorized creates a SearchUsersWithPagingUnauthorized with default headers values
func NewSearchUsersWithPagingUnauthorized() *SearchUsersWithPagingUnauthorized {
	return &SearchUsersWithPagingUnauthorized{}
}

/*
SearchUsersWithPagingUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SearchUsersWithPagingUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchUsersWithPagingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchUsersWithPagingUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingForbidden creates a SearchUsersWithPagingForbidden with default headers values
func NewSearchUsersWithPagingForbidden() *SearchUsersWithPagingForbidden {
	return &SearchUsersWithPagingForbidden{}
}

/*
SearchUsersWithPagingForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SearchUsersWithPagingForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchUsersWithPagingForbidden) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingForbidden  %+v", 403, o.Payload)
}
func (o *SearchUsersWithPagingForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingNotFound creates a SearchUsersWithPagingNotFound with default headers values
func NewSearchUsersWithPagingNotFound() *SearchUsersWithPagingNotFound {
	return &SearchUsersWithPagingNotFound{}
}

/*
SearchUsersWithPagingNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SearchUsersWithPagingNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchUsersWithPagingNotFound) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingNotFound  %+v", 404, o.Payload)
}
func (o *SearchUsersWithPagingNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingInternalServerError creates a SearchUsersWithPagingInternalServerError with default headers values
func NewSearchUsersWithPagingInternalServerError() *SearchUsersWithPagingInternalServerError {
	return &SearchUsersWithPagingInternalServerError{}
}

/*
SearchUsersWithPagingInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchUsersWithPagingInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchUsersWithPagingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchUsersWithPagingInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
