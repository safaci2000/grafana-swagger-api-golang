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

// RouteGetPolicyTreeReader is a Reader for the RouteGetPolicyTree structure.
type RouteGetPolicyTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetPolicyTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetPolicyTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetPolicyTreeOK creates a RouteGetPolicyTreeOK with default headers values
func NewRouteGetPolicyTreeOK() *RouteGetPolicyTreeOK {
	return &RouteGetPolicyTreeOK{}
}

/*
RouteGetPolicyTreeOK describes a response with status code 200, with default header values.

Route
*/
type RouteGetPolicyTreeOK struct {
	Payload *models.Route
}

func (o *RouteGetPolicyTreeOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/policies][%d] routeGetPolicyTreeOK  %+v", 200, o.Payload)
}
func (o *RouteGetPolicyTreeOK) GetPayload() *models.Route {
	return o.Payload
}

func (o *RouteGetPolicyTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Route)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
