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

// DatasourceProxyPOSTByUIDcallsReader is a Reader for the DatasourceProxyPOSTByUIDcalls structure.
type DatasourceProxyPOSTByUIDcallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatasourceProxyPOSTByUIDcallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDatasourceProxyPOSTByUIDcallsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDatasourceProxyPOSTByUIDcallsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDatasourceProxyPOSTByUIDcallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDatasourceProxyPOSTByUIDcallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatasourceProxyPOSTByUIDcallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDatasourceProxyPOSTByUIDcallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatasourceProxyPOSTByUIDcallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDatasourceProxyPOSTByUIDcallsCreated creates a DatasourceProxyPOSTByUIDcallsCreated with default headers values
func NewDatasourceProxyPOSTByUIDcallsCreated() *DatasourceProxyPOSTByUIDcallsCreated {
	return &DatasourceProxyPOSTByUIDcallsCreated{}
}

/*
DatasourceProxyPOSTByUIDcallsCreated describes a response with status code 201, with default header values.

(empty)
*/
type DatasourceProxyPOSTByUIDcallsCreated struct {
}

func (o *DatasourceProxyPOSTByUIDcallsCreated) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsCreated ", 201)
}

func (o *DatasourceProxyPOSTByUIDcallsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsAccepted creates a DatasourceProxyPOSTByUIDcallsAccepted with default headers values
func NewDatasourceProxyPOSTByUIDcallsAccepted() *DatasourceProxyPOSTByUIDcallsAccepted {
	return &DatasourceProxyPOSTByUIDcallsAccepted{}
}

/*
DatasourceProxyPOSTByUIDcallsAccepted describes a response with status code 202, with default header values.

(empty)
*/
type DatasourceProxyPOSTByUIDcallsAccepted struct {
}

func (o *DatasourceProxyPOSTByUIDcallsAccepted) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsAccepted ", 202)
}

func (o *DatasourceProxyPOSTByUIDcallsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsBadRequest creates a DatasourceProxyPOSTByUIDcallsBadRequest with default headers values
func NewDatasourceProxyPOSTByUIDcallsBadRequest() *DatasourceProxyPOSTByUIDcallsBadRequest {
	return &DatasourceProxyPOSTByUIDcallsBadRequest{}
}

/*
DatasourceProxyPOSTByUIDcallsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DatasourceProxyPOSTByUIDcallsBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *DatasourceProxyPOSTByUIDcallsBadRequest) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsBadRequest  %+v", 400, o.Payload)
}
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsUnauthorized creates a DatasourceProxyPOSTByUIDcallsUnauthorized with default headers values
func NewDatasourceProxyPOSTByUIDcallsUnauthorized() *DatasourceProxyPOSTByUIDcallsUnauthorized {
	return &DatasourceProxyPOSTByUIDcallsUnauthorized{}
}

/*
DatasourceProxyPOSTByUIDcallsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DatasourceProxyPOSTByUIDcallsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsUnauthorized  %+v", 401, o.Payload)
}
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsForbidden creates a DatasourceProxyPOSTByUIDcallsForbidden with default headers values
func NewDatasourceProxyPOSTByUIDcallsForbidden() *DatasourceProxyPOSTByUIDcallsForbidden {
	return &DatasourceProxyPOSTByUIDcallsForbidden{}
}

/*
DatasourceProxyPOSTByUIDcallsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DatasourceProxyPOSTByUIDcallsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *DatasourceProxyPOSTByUIDcallsForbidden) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsForbidden  %+v", 403, o.Payload)
}
func (o *DatasourceProxyPOSTByUIDcallsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsNotFound creates a DatasourceProxyPOSTByUIDcallsNotFound with default headers values
func NewDatasourceProxyPOSTByUIDcallsNotFound() *DatasourceProxyPOSTByUIDcallsNotFound {
	return &DatasourceProxyPOSTByUIDcallsNotFound{}
}

/*
DatasourceProxyPOSTByUIDcallsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DatasourceProxyPOSTByUIDcallsNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *DatasourceProxyPOSTByUIDcallsNotFound) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsNotFound  %+v", 404, o.Payload)
}
func (o *DatasourceProxyPOSTByUIDcallsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsInternalServerError creates a DatasourceProxyPOSTByUIDcallsInternalServerError with default headers values
func NewDatasourceProxyPOSTByUIDcallsInternalServerError() *DatasourceProxyPOSTByUIDcallsInternalServerError {
	return &DatasourceProxyPOSTByUIDcallsInternalServerError{}
}

/*
DatasourceProxyPOSTByUIDcallsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DatasourceProxyPOSTByUIDcallsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsInternalServerError  %+v", 500, o.Payload)
}
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
