// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AdminGetSettingsReader is a Reader for the AdminGetSettings structure.
type AdminGetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGetSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminGetSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminGetSettingsOK creates a AdminGetSettingsOK with default headers values
func NewAdminGetSettingsOK() *AdminGetSettingsOK {
	return &AdminGetSettingsOK{}
}

/*
AdminGetSettingsOK describes a response with status code 200, with default header values.

(empty)
*/
type AdminGetSettingsOK struct {
	Payload models.SettingsBag
}

func (o *AdminGetSettingsOK) Error() string {
	return fmt.Sprintf("[GET /admin/settings][%d] adminGetSettingsOK  %+v", 200, o.Payload)
}
func (o *AdminGetSettingsOK) GetPayload() models.SettingsBag {
	return o.Payload
}

func (o *AdminGetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetSettingsUnauthorized creates a AdminGetSettingsUnauthorized with default headers values
func NewAdminGetSettingsUnauthorized() *AdminGetSettingsUnauthorized {
	return &AdminGetSettingsUnauthorized{}
}

/*
AdminGetSettingsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminGetSettingsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminGetSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /admin/settings][%d] adminGetSettingsUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminGetSettingsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetSettingsForbidden creates a AdminGetSettingsForbidden with default headers values
func NewAdminGetSettingsForbidden() *AdminGetSettingsForbidden {
	return &AdminGetSettingsForbidden{}
}

/*
AdminGetSettingsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminGetSettingsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminGetSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/settings][%d] adminGetSettingsForbidden  %+v", 403, o.Payload)
}
func (o *AdminGetSettingsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminGetSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
