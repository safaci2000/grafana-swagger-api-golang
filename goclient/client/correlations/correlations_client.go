// Code generated by go-swagger; DO NOT EDIT.

package correlations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new correlations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for correlations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCorrelation(params *CreateCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCorrelationOK, error)

	DeleteCorrelation(params *DeleteCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCorrelationOK, error)

	GetCorrelation(params *GetCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorrelationOK, error)

	GetCorrelations(params *GetCorrelationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorrelationsOK, error)

	GetCorrelationsBySourceUID(params *GetCorrelationsBySourceUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorrelationsBySourceUIDOK, error)

	UpdateCorrelation(params *UpdateCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCorrelationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCorrelation adds correlation
*/
func (a *Client) CreateCorrelation(params *CreateCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCorrelationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCorrelationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCorrelation",
		Method:             "POST",
		PathPattern:        "/datasources/uid/{sourceUID}/correlations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCorrelationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCorrelationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCorrelation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCorrelation deletes a correlation
*/
func (a *Client) DeleteCorrelation(params *DeleteCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCorrelationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCorrelationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCorrelation",
		Method:             "DELETE",
		PathPattern:        "/datasources/uid/{uid}/correlations/{correlationUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCorrelationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCorrelationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCorrelation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCorrelation gets a correlation
*/
func (a *Client) GetCorrelation(params *GetCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorrelationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorrelationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCorrelation",
		Method:             "GET",
		PathPattern:        "/datasources/uid/{sourceUID}/correlations/{correlationUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCorrelationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCorrelationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCorrelation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCorrelations gets all correlations
*/
func (a *Client) GetCorrelations(params *GetCorrelationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorrelationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorrelationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCorrelations",
		Method:             "GET",
		PathPattern:        "/datasources/correlations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCorrelationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCorrelationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCorrelations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCorrelationsBySourceUID gets all correlations originating from the given data source
*/
func (a *Client) GetCorrelationsBySourceUID(params *GetCorrelationsBySourceUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorrelationsBySourceUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorrelationsBySourceUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCorrelationsBySourceUID",
		Method:             "GET",
		PathPattern:        "/datasources/uid/{sourceUID}/correlations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCorrelationsBySourceUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCorrelationsBySourceUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCorrelationsBySourceUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCorrelation updates a correlation
*/
func (a *Client) UpdateCorrelation(params *UpdateCorrelationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCorrelationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCorrelationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCorrelation",
		Method:             "PATCH",
		PathPattern:        "/datasources/uid/{sourceUID}/correlations/{correlationUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateCorrelationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCorrelationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCorrelation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
