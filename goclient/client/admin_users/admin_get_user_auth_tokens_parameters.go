// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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

// NewAdminGetUserAuthTokensParams creates a new AdminGetUserAuthTokensParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminGetUserAuthTokensParams() *AdminGetUserAuthTokensParams {
	return &AdminGetUserAuthTokensParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetUserAuthTokensParamsWithTimeout creates a new AdminGetUserAuthTokensParams object
// with the ability to set a timeout on a request.
func NewAdminGetUserAuthTokensParamsWithTimeout(timeout time.Duration) *AdminGetUserAuthTokensParams {
	return &AdminGetUserAuthTokensParams{
		timeout: timeout,
	}
}

// NewAdminGetUserAuthTokensParamsWithContext creates a new AdminGetUserAuthTokensParams object
// with the ability to set a context for a request.
func NewAdminGetUserAuthTokensParamsWithContext(ctx context.Context) *AdminGetUserAuthTokensParams {
	return &AdminGetUserAuthTokensParams{
		Context: ctx,
	}
}

// NewAdminGetUserAuthTokensParamsWithHTTPClient creates a new AdminGetUserAuthTokensParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminGetUserAuthTokensParamsWithHTTPClient(client *http.Client) *AdminGetUserAuthTokensParams {
	return &AdminGetUserAuthTokensParams{
		HTTPClient: client,
	}
}

/*
AdminGetUserAuthTokensParams contains all the parameters to send to the API endpoint

	for the admin get user auth tokens operation.

	Typically these are written to a http.Request.
*/
type AdminGetUserAuthTokensParams struct {

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin get user auth tokens params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGetUserAuthTokensParams) WithDefaults() *AdminGetUserAuthTokensParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin get user auth tokens params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGetUserAuthTokensParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) WithTimeout(timeout time.Duration) *AdminGetUserAuthTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) WithContext(ctx context.Context) *AdminGetUserAuthTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) WithHTTPClient(client *http.Client) *AdminGetUserAuthTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) WithUserID(userID int64) *AdminGetUserAuthTokensParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin get user auth tokens params
func (o *AdminGetUserAuthTokensParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetUserAuthTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
