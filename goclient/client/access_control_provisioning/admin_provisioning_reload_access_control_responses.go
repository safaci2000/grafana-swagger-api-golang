// Code generated by go-swagger; DO NOT EDIT.

package access_control_provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// AdminProvisioningReloadAccessControlReader is a Reader for the AdminProvisioningReloadAccessControl structure.
type AdminProvisioningReloadAccessControlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminProvisioningReloadAccessControlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAdminProvisioningReloadAccessControlAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminProvisioningReloadAccessControlUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminProvisioningReloadAccessControlForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminProvisioningReloadAccessControlAccepted creates a AdminProvisioningReloadAccessControlAccepted with default headers values
func NewAdminProvisioningReloadAccessControlAccepted() *AdminProvisioningReloadAccessControlAccepted {
	return &AdminProvisioningReloadAccessControlAccepted{}
}

/*
AdminProvisioningReloadAccessControlAccepted describes a response with status code 202, with default header values.

AcceptedResponse
*/
type AdminProvisioningReloadAccessControlAccepted struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminProvisioningReloadAccessControlAccepted) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/access-control/reload][%d] adminProvisioningReloadAccessControlAccepted  %+v", 202, o.Payload)
}
func (o *AdminProvisioningReloadAccessControlAccepted) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadAccessControlAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProvisioningReloadAccessControlUnauthorized creates a AdminProvisioningReloadAccessControlUnauthorized with default headers values
func NewAdminProvisioningReloadAccessControlUnauthorized() *AdminProvisioningReloadAccessControlUnauthorized {
	return &AdminProvisioningReloadAccessControlUnauthorized{}
}

/*
AdminProvisioningReloadAccessControlUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminProvisioningReloadAccessControlUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminProvisioningReloadAccessControlUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/access-control/reload][%d] adminProvisioningReloadAccessControlUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminProvisioningReloadAccessControlUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadAccessControlUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProvisioningReloadAccessControlForbidden creates a AdminProvisioningReloadAccessControlForbidden with default headers values
func NewAdminProvisioningReloadAccessControlForbidden() *AdminProvisioningReloadAccessControlForbidden {
	return &AdminProvisioningReloadAccessControlForbidden{}
}

/*
AdminProvisioningReloadAccessControlForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminProvisioningReloadAccessControlForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminProvisioningReloadAccessControlForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/provisioning/access-control/reload][%d] adminProvisioningReloadAccessControlForbidden  %+v", 403, o.Payload)
}
func (o *AdminProvisioningReloadAccessControlForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminProvisioningReloadAccessControlForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
