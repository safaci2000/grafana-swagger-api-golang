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
)

// NewRevokeInviteParams creates a new RevokeInviteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeInviteParams() *RevokeInviteParams {
	return &RevokeInviteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeInviteParamsWithTimeout creates a new RevokeInviteParams object
// with the ability to set a timeout on a request.
func NewRevokeInviteParamsWithTimeout(timeout time.Duration) *RevokeInviteParams {
	return &RevokeInviteParams{
		timeout: timeout,
	}
}

// NewRevokeInviteParamsWithContext creates a new RevokeInviteParams object
// with the ability to set a context for a request.
func NewRevokeInviteParamsWithContext(ctx context.Context) *RevokeInviteParams {
	return &RevokeInviteParams{
		Context: ctx,
	}
}

// NewRevokeInviteParamsWithHTTPClient creates a new RevokeInviteParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeInviteParamsWithHTTPClient(client *http.Client) *RevokeInviteParams {
	return &RevokeInviteParams{
		HTTPClient: client,
	}
}

/*
RevokeInviteParams contains all the parameters to send to the API endpoint

	for the revoke invite operation.

	Typically these are written to a http.Request.
*/
type RevokeInviteParams struct {

	// InvitationCode.
	InvitationCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeInviteParams) WithDefaults() *RevokeInviteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeInviteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke invite params
func (o *RevokeInviteParams) WithTimeout(timeout time.Duration) *RevokeInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke invite params
func (o *RevokeInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke invite params
func (o *RevokeInviteParams) WithContext(ctx context.Context) *RevokeInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke invite params
func (o *RevokeInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke invite params
func (o *RevokeInviteParams) WithHTTPClient(client *http.Client) *RevokeInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke invite params
func (o *RevokeInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitationCode adds the invitationCode to the revoke invite params
func (o *RevokeInviteParams) WithInvitationCode(invitationCode string) *RevokeInviteParams {
	o.SetInvitationCode(invitationCode)
	return o
}

// SetInvitationCode adds the invitationCode to the revoke invite params
func (o *RevokeInviteParams) SetInvitationCode(invitationCode string) {
	o.InvitationCode = invitationCode
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitation_code
	if err := r.SetPathParam("invitation_code", o.InvitationCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
