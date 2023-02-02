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

// RouteResetPolicyTreeReader is a Reader for the RouteResetPolicyTree structure.
type RouteResetPolicyTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteResetPolicyTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRouteResetPolicyTreeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteResetPolicyTreeAccepted creates a RouteResetPolicyTreeAccepted with default headers values
func NewRouteResetPolicyTreeAccepted() *RouteResetPolicyTreeAccepted {
	return &RouteResetPolicyTreeAccepted{}
}

/*
RouteResetPolicyTreeAccepted describes a response with status code 202, with default header values.

Ack
*/
type RouteResetPolicyTreeAccepted struct {
	Payload models.Ack
}

func (o *RouteResetPolicyTreeAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/provisioning/policies][%d] routeResetPolicyTreeAccepted  %+v", 202, o.Payload)
}
func (o *RouteResetPolicyTreeAccepted) GetPayload() models.Ack {
	return o.Payload
}

func (o *RouteResetPolicyTreeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
