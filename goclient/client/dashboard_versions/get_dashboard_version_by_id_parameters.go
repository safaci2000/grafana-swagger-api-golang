// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDashboardVersionByIDParams creates a new GetDashboardVersionByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardVersionByIDParams() *GetDashboardVersionByIDParams {
	return &GetDashboardVersionByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardVersionByIDParamsWithTimeout creates a new GetDashboardVersionByIDParams object
// with the ability to set a timeout on a request.
func NewGetDashboardVersionByIDParamsWithTimeout(timeout time.Duration) *GetDashboardVersionByIDParams {
	return &GetDashboardVersionByIDParams{
		timeout: timeout,
	}
}

// NewGetDashboardVersionByIDParamsWithContext creates a new GetDashboardVersionByIDParams object
// with the ability to set a context for a request.
func NewGetDashboardVersionByIDParamsWithContext(ctx context.Context) *GetDashboardVersionByIDParams {
	return &GetDashboardVersionByIDParams{
		Context: ctx,
	}
}

// NewGetDashboardVersionByIDParamsWithHTTPClient creates a new GetDashboardVersionByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardVersionByIDParamsWithHTTPClient(client *http.Client) *GetDashboardVersionByIDParams {
	return &GetDashboardVersionByIDParams{
		HTTPClient: client,
	}
}

/*
GetDashboardVersionByIDParams contains all the parameters to send to the API endpoint

	for the get dashboard version by ID operation.

	Typically these are written to a http.Request.
*/
type GetDashboardVersionByIDParams struct {

	// DashboardID.
	//
	// Format: int64
	DashboardID int64

	// DashboardVersionID.
	//
	// Format: int64
	DashboardVersionID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard version by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardVersionByIDParams) WithDefaults() *GetDashboardVersionByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard version by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardVersionByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) WithTimeout(timeout time.Duration) *GetDashboardVersionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) WithContext(ctx context.Context) *GetDashboardVersionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) WithHTTPClient(client *http.Client) *GetDashboardVersionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) WithDashboardID(dashboardID int64) *GetDashboardVersionByIDParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) SetDashboardID(dashboardID int64) {
	o.DashboardID = dashboardID
}

// WithDashboardVersionID adds the dashboardVersionID to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) WithDashboardVersionID(dashboardVersionID int64) *GetDashboardVersionByIDParams {
	o.SetDashboardVersionID(dashboardVersionID)
	return o
}

// SetDashboardVersionID adds the dashboardVersionId to the get dashboard version by ID params
func (o *GetDashboardVersionByIDParams) SetDashboardVersionID(dashboardVersionID int64) {
	o.DashboardVersionID = dashboardVersionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardVersionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param DashboardID
	if err := r.SetPathParam("DashboardID", swag.FormatInt64(o.DashboardID)); err != nil {
		return err
	}

	// path param DashboardVersionID
	if err := r.SetPathParam("DashboardVersionID", swag.FormatInt64(o.DashboardVersionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
