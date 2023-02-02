// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// SearchDashboardSnapshotsReader is a Reader for the SearchDashboardSnapshots structure.
type SearchDashboardSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchDashboardSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchDashboardSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSearchDashboardSnapshotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchDashboardSnapshotsOK creates a SearchDashboardSnapshotsOK with default headers values
func NewSearchDashboardSnapshotsOK() *SearchDashboardSnapshotsOK {
	return &SearchDashboardSnapshotsOK{}
}

/*
SearchDashboardSnapshotsOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchDashboardSnapshotsOK struct {
	Payload []*models.DashboardSnapshotDTO
}

func (o *SearchDashboardSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/snapshots][%d] searchDashboardSnapshotsOK  %+v", 200, o.Payload)
}
func (o *SearchDashboardSnapshotsOK) GetPayload() []*models.DashboardSnapshotDTO {
	return o.Payload
}

func (o *SearchDashboardSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchDashboardSnapshotsInternalServerError creates a SearchDashboardSnapshotsInternalServerError with default headers values
func NewSearchDashboardSnapshotsInternalServerError() *SearchDashboardSnapshotsInternalServerError {
	return &SearchDashboardSnapshotsInternalServerError{}
}

/*
SearchDashboardSnapshotsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchDashboardSnapshotsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchDashboardSnapshotsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dashboard/snapshots][%d] searchDashboardSnapshotsInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchDashboardSnapshotsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchDashboardSnapshotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
