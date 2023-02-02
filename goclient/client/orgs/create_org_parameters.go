// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// NewCreateOrgParams creates a new CreateOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrgParams() *CreateOrgParams {
	return &CreateOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgParamsWithTimeout creates a new CreateOrgParams object
// with the ability to set a timeout on a request.
func NewCreateOrgParamsWithTimeout(timeout time.Duration) *CreateOrgParams {
	return &CreateOrgParams{
		timeout: timeout,
	}
}

// NewCreateOrgParamsWithContext creates a new CreateOrgParams object
// with the ability to set a context for a request.
func NewCreateOrgParamsWithContext(ctx context.Context) *CreateOrgParams {
	return &CreateOrgParams{
		Context: ctx,
	}
}

// NewCreateOrgParamsWithHTTPClient creates a new CreateOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrgParamsWithHTTPClient(client *http.Client) *CreateOrgParams {
	return &CreateOrgParams{
		HTTPClient: client,
	}
}

/*
CreateOrgParams contains all the parameters to send to the API endpoint

	for the create org operation.

	Typically these are written to a http.Request.
*/
type CreateOrgParams struct {

	// Body.
	Body *models.CreateOrgCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgParams) WithDefaults() *CreateOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create org params
func (o *CreateOrgParams) WithTimeout(timeout time.Duration) *CreateOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org params
func (o *CreateOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org params
func (o *CreateOrgParams) WithContext(ctx context.Context) *CreateOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org params
func (o *CreateOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org params
func (o *CreateOrgParams) WithHTTPClient(client *http.Client) *CreateOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org params
func (o *CreateOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create org params
func (o *CreateOrgParams) WithBody(body *models.CreateOrgCommand) *CreateOrgParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create org params
func (o *CreateOrgParams) SetBody(body *models.CreateOrgCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
