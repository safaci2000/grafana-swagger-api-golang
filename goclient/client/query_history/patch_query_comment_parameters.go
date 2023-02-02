// Code generated by go-swagger; DO NOT EDIT.

package query_history

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

// NewPatchQueryCommentParams creates a new PatchQueryCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchQueryCommentParams() *PatchQueryCommentParams {
	return &PatchQueryCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchQueryCommentParamsWithTimeout creates a new PatchQueryCommentParams object
// with the ability to set a timeout on a request.
func NewPatchQueryCommentParamsWithTimeout(timeout time.Duration) *PatchQueryCommentParams {
	return &PatchQueryCommentParams{
		timeout: timeout,
	}
}

// NewPatchQueryCommentParamsWithContext creates a new PatchQueryCommentParams object
// with the ability to set a context for a request.
func NewPatchQueryCommentParamsWithContext(ctx context.Context) *PatchQueryCommentParams {
	return &PatchQueryCommentParams{
		Context: ctx,
	}
}

// NewPatchQueryCommentParamsWithHTTPClient creates a new PatchQueryCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchQueryCommentParamsWithHTTPClient(client *http.Client) *PatchQueryCommentParams {
	return &PatchQueryCommentParams{
		HTTPClient: client,
	}
}

/*
PatchQueryCommentParams contains all the parameters to send to the API endpoint

	for the patch query comment operation.

	Typically these are written to a http.Request.
*/
type PatchQueryCommentParams struct {

	// Body.
	Body *models.PatchQueryCommentInQueryHistoryCommand

	// QueryHistoryUID.
	QueryHistoryUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch query comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueryCommentParams) WithDefaults() *PatchQueryCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch query comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchQueryCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch query comment params
func (o *PatchQueryCommentParams) WithTimeout(timeout time.Duration) *PatchQueryCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch query comment params
func (o *PatchQueryCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch query comment params
func (o *PatchQueryCommentParams) WithContext(ctx context.Context) *PatchQueryCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch query comment params
func (o *PatchQueryCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch query comment params
func (o *PatchQueryCommentParams) WithHTTPClient(client *http.Client) *PatchQueryCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch query comment params
func (o *PatchQueryCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch query comment params
func (o *PatchQueryCommentParams) WithBody(body *models.PatchQueryCommentInQueryHistoryCommand) *PatchQueryCommentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch query comment params
func (o *PatchQueryCommentParams) SetBody(body *models.PatchQueryCommentInQueryHistoryCommand) {
	o.Body = body
}

// WithQueryHistoryUID adds the queryHistoryUID to the patch query comment params
func (o *PatchQueryCommentParams) WithQueryHistoryUID(queryHistoryUID string) *PatchQueryCommentParams {
	o.SetQueryHistoryUID(queryHistoryUID)
	return o
}

// SetQueryHistoryUID adds the queryHistoryUid to the patch query comment params
func (o *PatchQueryCommentParams) SetQueryHistoryUID(queryHistoryUID string) {
	o.QueryHistoryUID = queryHistoryUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchQueryCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param query_history_uid
	if err := r.SetPathParam("query_history_uid", o.QueryHistoryUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
