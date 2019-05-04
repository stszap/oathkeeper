// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new health API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for health API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
IsInstanceAlive checks the alive status

This endpoint returns a 200 status code when the HTTP server is up running.
This status does currently not include checks whether the database connection is working.
This endpoint does not require the `X-Forwarded-Proto` header when TLS termination is set.

Be aware that if you are running multiple nodes of ORY Oathkeeper, the health status will never refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceAlive(params *IsInstanceAliveParams) (*IsInstanceAliveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceAliveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "isInstanceAlive",
		Method:             "GET",
		PathPattern:        "/health/alive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceAliveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IsInstanceAliveOK), nil

}

/*
IsInstanceReady checks the readiness status

This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g.
the database) are responsive as well.

This status does currently not include checks whether the database connection is working.
This endpoint does not require the `X-Forwarded-Proto` header when TLS termination is set.

Be aware that if you are running multiple nodes of ORY Oathkeeper, the health status will never refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceReady(params *IsInstanceReadyParams) (*IsInstanceReadyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceReadyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "isInstanceReady",
		Method:             "GET",
		PathPattern:        "/health/ready",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceReadyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IsInstanceReadyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
