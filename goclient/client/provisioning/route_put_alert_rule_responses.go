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

// RoutePutAlertRuleReader is a Reader for the RoutePutAlertRule structure.
type RoutePutAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePutAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoutePutAlertRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePutAlertRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePutAlertRuleOK creates a RoutePutAlertRuleOK with default headers values
func NewRoutePutAlertRuleOK() *RoutePutAlertRuleOK {
	return &RoutePutAlertRuleOK{}
}

/*
RoutePutAlertRuleOK describes a response with status code 200, with default header values.

ProvisionedAlertRule
*/
type RoutePutAlertRuleOK struct {
	Payload *models.ProvisionedAlertRule
}

func (o *RoutePutAlertRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/alert-rules/{UID}][%d] routePutAlertRuleOK  %+v", 200, o.Payload)
}
func (o *RoutePutAlertRuleOK) GetPayload() *models.ProvisionedAlertRule {
	return o.Payload
}

func (o *RoutePutAlertRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvisionedAlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePutAlertRuleBadRequest creates a RoutePutAlertRuleBadRequest with default headers values
func NewRoutePutAlertRuleBadRequest() *RoutePutAlertRuleBadRequest {
	return &RoutePutAlertRuleBadRequest{}
}

/*
RoutePutAlertRuleBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePutAlertRuleBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePutAlertRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/alert-rules/{UID}][%d] routePutAlertRuleBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePutAlertRuleBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePutAlertRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
