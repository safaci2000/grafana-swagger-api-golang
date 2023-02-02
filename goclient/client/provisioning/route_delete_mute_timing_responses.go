// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RouteDeleteMuteTimingReader is a Reader for the RouteDeleteMuteTiming structure.
type RouteDeleteMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteDeleteMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRouteDeleteMuteTimingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteDeleteMuteTimingNoContent creates a RouteDeleteMuteTimingNoContent with default headers values
func NewRouteDeleteMuteTimingNoContent() *RouteDeleteMuteTimingNoContent {
	return &RouteDeleteMuteTimingNoContent{}
}

/*
RouteDeleteMuteTimingNoContent describes a response with status code 204, with default header values.

	The mute timing was deleted successfully.
*/
type RouteDeleteMuteTimingNoContent struct {
}

func (o *RouteDeleteMuteTimingNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/provisioning/mute-timings/{name}][%d] routeDeleteMuteTimingNoContent ", 204)
}

func (o *RouteDeleteMuteTimingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
