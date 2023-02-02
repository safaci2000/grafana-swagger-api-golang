// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// RoutePostAlertRuleReader is a Reader for the RoutePostAlertRule structure.
type RoutePostAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePostAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRoutePostAlertRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePostAlertRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePostAlertRuleCreated creates a RoutePostAlertRuleCreated with default headers values
func NewRoutePostAlertRuleCreated() *RoutePostAlertRuleCreated {
	return &RoutePostAlertRuleCreated{}
}

/*
RoutePostAlertRuleCreated describes a response with status code 201, with default header values.

ProvisionedAlertRule
*/
type RoutePostAlertRuleCreated struct {
	Payload *models.ProvisionedAlertRule
}

func (o *RoutePostAlertRuleCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/provisioning/alert-rules][%d] routePostAlertRuleCreated  %+v", 201, o.Payload)
}
func (o *RoutePostAlertRuleCreated) GetPayload() *models.ProvisionedAlertRule {
	return o.Payload
}

func (o *RoutePostAlertRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvisionedAlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePostAlertRuleBadRequest creates a RoutePostAlertRuleBadRequest with default headers values
func NewRoutePostAlertRuleBadRequest() *RoutePostAlertRuleBadRequest {
	return &RoutePostAlertRuleBadRequest{}
}

/*
RoutePostAlertRuleBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePostAlertRuleBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePostAlertRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/provisioning/alert-rules][%d] routePostAlertRuleBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePostAlertRuleBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePostAlertRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
