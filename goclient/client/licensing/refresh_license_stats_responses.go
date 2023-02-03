// Code generated by go-swagger; DO NOT EDIT.

package licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// RefreshLicenseStatsReader is a Reader for the RefreshLicenseStats structure.
type RefreshLicenseStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshLicenseStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshLicenseStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRefreshLicenseStatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshLicenseStatsOK creates a RefreshLicenseStatsOK with default headers values
func NewRefreshLicenseStatsOK() *RefreshLicenseStatsOK {
	return &RefreshLicenseStatsOK{}
}

/*
RefreshLicenseStatsOK describes a response with status code 200, with default header values.

(empty)
*/
type RefreshLicenseStatsOK struct {
	Payload *models.ActiveUserStats
}

func (o *RefreshLicenseStatsOK) Error() string {
	return fmt.Sprintf("[GET /licensing/refresh-stats][%d] refreshLicenseStatsOK  %+v", 200, o.Payload)
}
func (o *RefreshLicenseStatsOK) GetPayload() *models.ActiveUserStats {
	return o.Payload
}

func (o *RefreshLicenseStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveUserStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshLicenseStatsInternalServerError creates a RefreshLicenseStatsInternalServerError with default headers values
func NewRefreshLicenseStatsInternalServerError() *RefreshLicenseStatsInternalServerError {
	return &RefreshLicenseStatsInternalServerError{}
}

/*
RefreshLicenseStatsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RefreshLicenseStatsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RefreshLicenseStatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /licensing/refresh-stats][%d] refreshLicenseStatsInternalServerError  %+v", 500, o.Payload)
}
func (o *RefreshLicenseStatsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RefreshLicenseStatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
