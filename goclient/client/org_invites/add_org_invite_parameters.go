// Code generated by go-swagger; DO NOT EDIT.

package org_invites

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

// NewAddOrgInviteParams creates a new AddOrgInviteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddOrgInviteParams() *AddOrgInviteParams {
	return &AddOrgInviteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrgInviteParamsWithTimeout creates a new AddOrgInviteParams object
// with the ability to set a timeout on a request.
func NewAddOrgInviteParamsWithTimeout(timeout time.Duration) *AddOrgInviteParams {
	return &AddOrgInviteParams{
		timeout: timeout,
	}
}

// NewAddOrgInviteParamsWithContext creates a new AddOrgInviteParams object
// with the ability to set a context for a request.
func NewAddOrgInviteParamsWithContext(ctx context.Context) *AddOrgInviteParams {
	return &AddOrgInviteParams{
		Context: ctx,
	}
}

// NewAddOrgInviteParamsWithHTTPClient creates a new AddOrgInviteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddOrgInviteParamsWithHTTPClient(client *http.Client) *AddOrgInviteParams {
	return &AddOrgInviteParams{
		HTTPClient: client,
	}
}

/*
AddOrgInviteParams contains all the parameters to send to the API endpoint

	for the add org invite operation.

	Typically these are written to a http.Request.
*/
type AddOrgInviteParams struct {

	// Body.
	Body *models.AddInviteForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add org invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOrgInviteParams) WithDefaults() *AddOrgInviteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add org invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOrgInviteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add org invite params
func (o *AddOrgInviteParams) WithTimeout(timeout time.Duration) *AddOrgInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add org invite params
func (o *AddOrgInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add org invite params
func (o *AddOrgInviteParams) WithContext(ctx context.Context) *AddOrgInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add org invite params
func (o *AddOrgInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add org invite params
func (o *AddOrgInviteParams) WithHTTPClient(client *http.Client) *AddOrgInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add org invite params
func (o *AddOrgInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add org invite params
func (o *AddOrgInviteParams) WithBody(body *models.AddInviteForm) *AddOrgInviteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add org invite params
func (o *AddOrgInviteParams) SetBody(body *models.AddInviteForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrgInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
