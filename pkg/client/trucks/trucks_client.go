// Code generated by go-swagger; DO NOT EDIT.

package trucks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new trucks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trucks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddTruck(params *AddTruckParams, opts ...ClientOption) (*AddTruckCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddTruck adds truck

Add new truck
*/
func (a *Client) AddTruck(params *AddTruckParams, opts ...ClientOption) (*AddTruckCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTruckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addTruck",
		Method:             "POST",
		PathPattern:        "/api/v1/trucks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddTruckReader{formats: a.formats},
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
	success, ok := result.(*AddTruckCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddTruckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
