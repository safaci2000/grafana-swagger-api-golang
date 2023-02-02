// Code generated by go-swagger; DO NOT EDIT.

package library_elements

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

// NewGetLibraryElementsParams creates a new GetLibraryElementsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLibraryElementsParams() *GetLibraryElementsParams {
	return &GetLibraryElementsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLibraryElementsParamsWithTimeout creates a new GetLibraryElementsParams object
// with the ability to set a timeout on a request.
func NewGetLibraryElementsParamsWithTimeout(timeout time.Duration) *GetLibraryElementsParams {
	return &GetLibraryElementsParams{
		timeout: timeout,
	}
}

// NewGetLibraryElementsParamsWithContext creates a new GetLibraryElementsParams object
// with the ability to set a context for a request.
func NewGetLibraryElementsParamsWithContext(ctx context.Context) *GetLibraryElementsParams {
	return &GetLibraryElementsParams{
		Context: ctx,
	}
}

// NewGetLibraryElementsParamsWithHTTPClient creates a new GetLibraryElementsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLibraryElementsParamsWithHTTPClient(client *http.Client) *GetLibraryElementsParams {
	return &GetLibraryElementsParams{
		HTTPClient: client,
	}
}

/*
GetLibraryElementsParams contains all the parameters to send to the API endpoint

	for the get library elements operation.

	Typically these are written to a http.Request.
*/
type GetLibraryElementsParams struct {

	/* ExcludeUID.

	   Element UID to exclude from search results.
	*/
	ExcludeUID *string

	/* FolderFilter.

	   A comma separated list of folder ID(s) to filter the elements by.
	*/
	FolderFilter *string

	/* Kind.

	   Kind of element to search for.

	   Format: int64
	*/
	Kind *int64

	/* Page.

	   The page for a set of records, given that only perPage records are returned at a time. Numbering starts at 1.

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PerPage.

	   The number of results per page.

	   Format: int64
	   Default: 100
	*/
	PerPage *int64

	/* SearchString.

	   Part of the name or description searched for.
	*/
	SearchString *string

	/* SortDirection.

	   Sort order of elements.
	*/
	SortDirection *string

	/* TypeFilter.

	   A comma separated list of types to filter the elements by
	*/
	TypeFilter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get library elements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLibraryElementsParams) WithDefaults() *GetLibraryElementsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get library elements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLibraryElementsParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		perPageDefault = int64(100)
	)

	val := GetLibraryElementsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get library elements params
func (o *GetLibraryElementsParams) WithTimeout(timeout time.Duration) *GetLibraryElementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get library elements params
func (o *GetLibraryElementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get library elements params
func (o *GetLibraryElementsParams) WithContext(ctx context.Context) *GetLibraryElementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get library elements params
func (o *GetLibraryElementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get library elements params
func (o *GetLibraryElementsParams) WithHTTPClient(client *http.Client) *GetLibraryElementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get library elements params
func (o *GetLibraryElementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeUID adds the excludeUID to the get library elements params
func (o *GetLibraryElementsParams) WithExcludeUID(excludeUID *string) *GetLibraryElementsParams {
	o.SetExcludeUID(excludeUID)
	return o
}

// SetExcludeUID adds the excludeUid to the get library elements params
func (o *GetLibraryElementsParams) SetExcludeUID(excludeUID *string) {
	o.ExcludeUID = excludeUID
}

// WithFolderFilter adds the folderFilter to the get library elements params
func (o *GetLibraryElementsParams) WithFolderFilter(folderFilter *string) *GetLibraryElementsParams {
	o.SetFolderFilter(folderFilter)
	return o
}

// SetFolderFilter adds the folderFilter to the get library elements params
func (o *GetLibraryElementsParams) SetFolderFilter(folderFilter *string) {
	o.FolderFilter = folderFilter
}

// WithKind adds the kind to the get library elements params
func (o *GetLibraryElementsParams) WithKind(kind *int64) *GetLibraryElementsParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get library elements params
func (o *GetLibraryElementsParams) SetKind(kind *int64) {
	o.Kind = kind
}

// WithPage adds the page to the get library elements params
func (o *GetLibraryElementsParams) WithPage(page *int64) *GetLibraryElementsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get library elements params
func (o *GetLibraryElementsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get library elements params
func (o *GetLibraryElementsParams) WithPerPage(perPage *int64) *GetLibraryElementsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get library elements params
func (o *GetLibraryElementsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSearchString adds the searchString to the get library elements params
func (o *GetLibraryElementsParams) WithSearchString(searchString *string) *GetLibraryElementsParams {
	o.SetSearchString(searchString)
	return o
}

// SetSearchString adds the searchString to the get library elements params
func (o *GetLibraryElementsParams) SetSearchString(searchString *string) {
	o.SearchString = searchString
}

// WithSortDirection adds the sortDirection to the get library elements params
func (o *GetLibraryElementsParams) WithSortDirection(sortDirection *string) *GetLibraryElementsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the get library elements params
func (o *GetLibraryElementsParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithTypeFilter adds the typeFilter to the get library elements params
func (o *GetLibraryElementsParams) WithTypeFilter(typeFilter *string) *GetLibraryElementsParams {
	o.SetTypeFilter(typeFilter)
	return o
}

// SetTypeFilter adds the typeFilter to the get library elements params
func (o *GetLibraryElementsParams) SetTypeFilter(typeFilter *string) {
	o.TypeFilter = typeFilter
}

// WriteToRequest writes these params to a swagger request
func (o *GetLibraryElementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeUID != nil {

		// query param excludeUid
		var qrExcludeUID string

		if o.ExcludeUID != nil {
			qrExcludeUID = *o.ExcludeUID
		}
		qExcludeUID := qrExcludeUID
		if qExcludeUID != "" {

			if err := r.SetQueryParam("excludeUid", qExcludeUID); err != nil {
				return err
			}
		}
	}

	if o.FolderFilter != nil {

		// query param folderFilter
		var qrFolderFilter string

		if o.FolderFilter != nil {
			qrFolderFilter = *o.FolderFilter
		}
		qFolderFilter := qrFolderFilter
		if qFolderFilter != "" {

			if err := r.SetQueryParam("folderFilter", qFolderFilter); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind int64

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := swag.FormatInt64(qrKind)
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.SearchString != nil {

		// query param searchString
		var qrSearchString string

		if o.SearchString != nil {
			qrSearchString = *o.SearchString
		}
		qSearchString := qrSearchString
		if qSearchString != "" {

			if err := r.SetQueryParam("searchString", qSearchString); err != nil {
				return err
			}
		}
	}

	if o.SortDirection != nil {

		// query param sortDirection
		var qrSortDirection string

		if o.SortDirection != nil {
			qrSortDirection = *o.SortDirection
		}
		qSortDirection := qrSortDirection
		if qSortDirection != "" {

			if err := r.SetQueryParam("sortDirection", qSortDirection); err != nil {
				return err
			}
		}
	}

	if o.TypeFilter != nil {

		// query param typeFilter
		var qrTypeFilter string

		if o.TypeFilter != nil {
			qrTypeFilter = *o.TypeFilter
		}
		qTypeFilter := qrTypeFilter
		if qTypeFilter != "" {

			if err := r.SetQueryParam("typeFilter", qTypeFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
