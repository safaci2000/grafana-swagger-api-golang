// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// NewUpdateDataSourceByUIDParams creates a new UpdateDataSourceByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDataSourceByUIDParams() *UpdateDataSourceByUIDParams {
	return &UpdateDataSourceByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDataSourceByUIDParamsWithTimeout creates a new UpdateDataSourceByUIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDataSourceByUIDParamsWithTimeout(timeout time.Duration) *UpdateDataSourceByUIDParams {
	return &UpdateDataSourceByUIDParams{
		timeout: timeout,
	}
}

// NewUpdateDataSourceByUIDParamsWithContext creates a new UpdateDataSourceByUIDParams object
// with the ability to set a context for a request.
func NewUpdateDataSourceByUIDParamsWithContext(ctx context.Context) *UpdateDataSourceByUIDParams {
	return &UpdateDataSourceByUIDParams{
		Context: ctx,
	}
}

// NewUpdateDataSourceByUIDParamsWithHTTPClient creates a new UpdateDataSourceByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDataSourceByUIDParamsWithHTTPClient(client *http.Client) *UpdateDataSourceByUIDParams {
	return &UpdateDataSourceByUIDParams{
		HTTPClient: client,
	}
}

/*
UpdateDataSourceByUIDParams contains all the parameters to send to the API endpoint

	for the update data source by UID operation.

	Typically these are written to a http.Request.
*/
type UpdateDataSourceByUIDParams struct {

	// Body.
	Body *models.UpdateDataSourceCommand

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update data source by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataSourceByUIDParams) WithDefaults() *UpdateDataSourceByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update data source by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataSourceByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) WithTimeout(timeout time.Duration) *UpdateDataSourceByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) WithContext(ctx context.Context) *UpdateDataSourceByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) WithHTTPClient(client *http.Client) *UpdateDataSourceByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) WithBody(body *models.UpdateDataSourceCommand) *UpdateDataSourceByUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) SetBody(body *models.UpdateDataSourceCommand) {
	o.Body = body
}

// WithUID adds the uid to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) WithUID(uid string) *UpdateDataSourceByUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the update data source by UID params
func (o *UpdateDataSourceByUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDataSourceByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
