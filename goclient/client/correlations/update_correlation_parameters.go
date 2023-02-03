// Code generated by go-swagger; DO NOT EDIT.

package correlations

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

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// NewUpdateCorrelationParams creates a new UpdateCorrelationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCorrelationParams() *UpdateCorrelationParams {
	return &UpdateCorrelationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCorrelationParamsWithTimeout creates a new UpdateCorrelationParams object
// with the ability to set a timeout on a request.
func NewUpdateCorrelationParamsWithTimeout(timeout time.Duration) *UpdateCorrelationParams {
	return &UpdateCorrelationParams{
		timeout: timeout,
	}
}

// NewUpdateCorrelationParamsWithContext creates a new UpdateCorrelationParams object
// with the ability to set a context for a request.
func NewUpdateCorrelationParamsWithContext(ctx context.Context) *UpdateCorrelationParams {
	return &UpdateCorrelationParams{
		Context: ctx,
	}
}

// NewUpdateCorrelationParamsWithHTTPClient creates a new UpdateCorrelationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCorrelationParamsWithHTTPClient(client *http.Client) *UpdateCorrelationParams {
	return &UpdateCorrelationParams{
		HTTPClient: client,
	}
}

/*
UpdateCorrelationParams contains all the parameters to send to the API endpoint

	for the update correlation operation.

	Typically these are written to a http.Request.
*/
type UpdateCorrelationParams struct {

	// Body.
	Body *models.UpdateCorrelationCommand

	// CorrelationUID.
	CorrelationUID string

	// SourceUID.
	SourceUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update correlation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCorrelationParams) WithDefaults() *UpdateCorrelationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update correlation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCorrelationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update correlation params
func (o *UpdateCorrelationParams) WithTimeout(timeout time.Duration) *UpdateCorrelationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update correlation params
func (o *UpdateCorrelationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update correlation params
func (o *UpdateCorrelationParams) WithContext(ctx context.Context) *UpdateCorrelationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update correlation params
func (o *UpdateCorrelationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update correlation params
func (o *UpdateCorrelationParams) WithHTTPClient(client *http.Client) *UpdateCorrelationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update correlation params
func (o *UpdateCorrelationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update correlation params
func (o *UpdateCorrelationParams) WithBody(body *models.UpdateCorrelationCommand) *UpdateCorrelationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update correlation params
func (o *UpdateCorrelationParams) SetBody(body *models.UpdateCorrelationCommand) {
	o.Body = body
}

// WithCorrelationUID adds the correlationUID to the update correlation params
func (o *UpdateCorrelationParams) WithCorrelationUID(correlationUID string) *UpdateCorrelationParams {
	o.SetCorrelationUID(correlationUID)
	return o
}

// SetCorrelationUID adds the correlationUid to the update correlation params
func (o *UpdateCorrelationParams) SetCorrelationUID(correlationUID string) {
	o.CorrelationUID = correlationUID
}

// WithSourceUID adds the sourceUID to the update correlation params
func (o *UpdateCorrelationParams) WithSourceUID(sourceUID string) *UpdateCorrelationParams {
	o.SetSourceUID(sourceUID)
	return o
}

// SetSourceUID adds the sourceUid to the update correlation params
func (o *UpdateCorrelationParams) SetSourceUID(sourceUID string) {
	o.SourceUID = sourceUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCorrelationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param correlationUID
	if err := r.SetPathParam("correlationUID", o.CorrelationUID); err != nil {
		return err
	}

	// path param sourceUID
	if err := r.SetPathParam("sourceUID", o.SourceUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
