// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

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

// NewChangeUserPasswordParams creates a new ChangeUserPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeUserPasswordParams() *ChangeUserPasswordParams {
	return &ChangeUserPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeUserPasswordParamsWithTimeout creates a new ChangeUserPasswordParams object
// with the ability to set a timeout on a request.
func NewChangeUserPasswordParamsWithTimeout(timeout time.Duration) *ChangeUserPasswordParams {
	return &ChangeUserPasswordParams{
		timeout: timeout,
	}
}

// NewChangeUserPasswordParamsWithContext creates a new ChangeUserPasswordParams object
// with the ability to set a context for a request.
func NewChangeUserPasswordParamsWithContext(ctx context.Context) *ChangeUserPasswordParams {
	return &ChangeUserPasswordParams{
		Context: ctx,
	}
}

// NewChangeUserPasswordParamsWithHTTPClient creates a new ChangeUserPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeUserPasswordParamsWithHTTPClient(client *http.Client) *ChangeUserPasswordParams {
	return &ChangeUserPasswordParams{
		HTTPClient: client,
	}
}

/*
ChangeUserPasswordParams contains all the parameters to send to the API endpoint

	for the change user password operation.

	Typically these are written to a http.Request.
*/
type ChangeUserPasswordParams struct {

	/* Body.

	   To change the email, name, login, theme, provide another one.
	*/
	Body *models.ChangeUserPasswordCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change user password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeUserPasswordParams) WithDefaults() *ChangeUserPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change user password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeUserPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change user password params
func (o *ChangeUserPasswordParams) WithTimeout(timeout time.Duration) *ChangeUserPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change user password params
func (o *ChangeUserPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change user password params
func (o *ChangeUserPasswordParams) WithContext(ctx context.Context) *ChangeUserPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change user password params
func (o *ChangeUserPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change user password params
func (o *ChangeUserPasswordParams) WithHTTPClient(client *http.Client) *ChangeUserPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change user password params
func (o *ChangeUserPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change user password params
func (o *ChangeUserPasswordParams) WithBody(body *models.ChangeUserPasswordCommand) *ChangeUserPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change user password params
func (o *ChangeUserPasswordParams) SetBody(body *models.ChangeUserPasswordCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeUserPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
