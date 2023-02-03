// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetTeamGroupsAPIReader is a Reader for the GetTeamGroupsAPI structure.
type GetTeamGroupsAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamGroupsAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamGroupsAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTeamGroupsAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTeamGroupsAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTeamGroupsAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamGroupsAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamGroupsAPIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamGroupsAPIOK creates a GetTeamGroupsAPIOK with default headers values
func NewGetTeamGroupsAPIOK() *GetTeamGroupsAPIOK {
	return &GetTeamGroupsAPIOK{}
}

/*
GetTeamGroupsAPIOK describes a response with status code 200, with default header values.

(empty)
*/
type GetTeamGroupsAPIOK struct {
	Payload []*models.TeamGroupDTO
}

func (o *GetTeamGroupsAPIOK) Error() string {
	return fmt.Sprintf("[GET /teams/{teamId}/groups][%d] getTeamGroupsApiOK  %+v", 200, o.Payload)
}
func (o *GetTeamGroupsAPIOK) GetPayload() []*models.TeamGroupDTO {
	return o.Payload
}

func (o *GetTeamGroupsAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamGroupsAPIBadRequest creates a GetTeamGroupsAPIBadRequest with default headers values
func NewGetTeamGroupsAPIBadRequest() *GetTeamGroupsAPIBadRequest {
	return &GetTeamGroupsAPIBadRequest{}
}

/*
GetTeamGroupsAPIBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetTeamGroupsAPIBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *GetTeamGroupsAPIBadRequest) Error() string {
	return fmt.Sprintf("[GET /teams/{teamId}/groups][%d] getTeamGroupsApiBadRequest  %+v", 400, o.Payload)
}
func (o *GetTeamGroupsAPIBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamGroupsAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamGroupsAPIUnauthorized creates a GetTeamGroupsAPIUnauthorized with default headers values
func NewGetTeamGroupsAPIUnauthorized() *GetTeamGroupsAPIUnauthorized {
	return &GetTeamGroupsAPIUnauthorized{}
}

/*
GetTeamGroupsAPIUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetTeamGroupsAPIUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetTeamGroupsAPIUnauthorized) Error() string {
	return fmt.Sprintf("[GET /teams/{teamId}/groups][%d] getTeamGroupsApiUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTeamGroupsAPIUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamGroupsAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamGroupsAPIForbidden creates a GetTeamGroupsAPIForbidden with default headers values
func NewGetTeamGroupsAPIForbidden() *GetTeamGroupsAPIForbidden {
	return &GetTeamGroupsAPIForbidden{}
}

/*
GetTeamGroupsAPIForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetTeamGroupsAPIForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetTeamGroupsAPIForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{teamId}/groups][%d] getTeamGroupsApiForbidden  %+v", 403, o.Payload)
}
func (o *GetTeamGroupsAPIForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamGroupsAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamGroupsAPINotFound creates a GetTeamGroupsAPINotFound with default headers values
func NewGetTeamGroupsAPINotFound() *GetTeamGroupsAPINotFound {
	return &GetTeamGroupsAPINotFound{}
}

/*
GetTeamGroupsAPINotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetTeamGroupsAPINotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetTeamGroupsAPINotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{teamId}/groups][%d] getTeamGroupsApiNotFound  %+v", 404, o.Payload)
}
func (o *GetTeamGroupsAPINotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamGroupsAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamGroupsAPIInternalServerError creates a GetTeamGroupsAPIInternalServerError with default headers values
func NewGetTeamGroupsAPIInternalServerError() *GetTeamGroupsAPIInternalServerError {
	return &GetTeamGroupsAPIInternalServerError{}
}

/*
GetTeamGroupsAPIInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetTeamGroupsAPIInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetTeamGroupsAPIInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams/{teamId}/groups][%d] getTeamGroupsApiInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTeamGroupsAPIInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamGroupsAPIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
