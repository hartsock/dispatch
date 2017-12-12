///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new subscriptions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscriptions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddSubscription adds a new subscription
*/
func (a *Client) AddSubscription(params *AddSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*AddSubscriptionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSubscription",
		Method:             "POST",
		PathPattern:        "/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddSubscriptionCreated), nil

}

/*
DeleteSubscription deletes a subscription
*/
func (a *Client) DeleteSubscription(params *DeleteSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSubscription",
		Method:             "DELETE",
		PathPattern:        "/subscriptions/{subscriptionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSubscriptionOK), nil

}

/*
GetSubscription finds subscription by name

Returns a single subscription
*/
func (a *Client) GetSubscription(params *GetSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscription",
		Method:             "GET",
		PathPattern:        "/subscriptions/{subscriptionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionOK), nil

}

/*
GetSubscriptions lists all existing subscriptions
*/
func (a *Client) GetSubscriptions(params *GetSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSubscriptions",
		Method:             "GET",
		PathPattern:        "/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSubscriptionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}