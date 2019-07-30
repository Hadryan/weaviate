//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new things API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for things API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
WeaviateThingsCreate creates a new thing based on a thing template

Registers a new Thing. Given meta-data and schema values are validated.
*/
func (a *Client) WeaviateThingsCreate(params *WeaviateThingsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.create",
		Method:             "POST",
		PathPattern:        "/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsCreateOK), nil

}

/*
WeaviateThingsDelete deletes a thing based on its UUID

Deletes a Thing from the system. All Actions pointing to this Thing, where the Thing is the object of the Action, are also being deleted.
*/
func (a *Client) WeaviateThingsDelete(params *WeaviateThingsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.delete",
		Method:             "DELETE",
		PathPattern:        "/things/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsDeleteNoContent), nil

}

/*
WeaviateThingsGet gets a thing based on its UUID

Returns a particular Thing data.
*/
func (a *Client) WeaviateThingsGet(params *WeaviateThingsGetParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.get",
		Method:             "GET",
		PathPattern:        "/things/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsGetOK), nil

}

/*
WeaviateThingsList gets a list of things

Lists all Things in reverse order of creation, owned by the user that belongs to the used token.
*/
func (a *Client) WeaviateThingsList(params *WeaviateThingsListParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.list",
		Method:             "GET",
		PathPattern:        "/things",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsListOK), nil

}

/*
WeaviateThingsPatch updates a thing based on its UUID using patch semantics

Updates a Thing's data. This method supports patch semantics. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) WeaviateThingsPatch(params *WeaviateThingsPatchParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.patch",
		Method:             "PATCH",
		PathPattern:        "/things/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsPatchOK), nil

}

/*
WeaviateThingsReferencesCreate adds a single reference to a class property when cardinality is set to has many

Add a single reference to a class-property when cardinality is set to 'hasMany'.
*/
func (a *Client) WeaviateThingsReferencesCreate(params *WeaviateThingsReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsReferencesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.references.create",
		Method:             "POST",
		PathPattern:        "/things/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsReferencesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsReferencesCreateOK), nil

}

/*
WeaviateThingsReferencesDelete deletes the single reference that is given in the body from the list of references that this property has

Delete the single reference that is given in the body from the list of references that this property has.
*/
func (a *Client) WeaviateThingsReferencesDelete(params *WeaviateThingsReferencesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsReferencesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsReferencesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.references.delete",
		Method:             "DELETE",
		PathPattern:        "/things/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsReferencesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsReferencesDeleteNoContent), nil

}

/*
WeaviateThingsReferencesUpdate replaces all references to a class property

Replace all references to a class-property.
*/
func (a *Client) WeaviateThingsReferencesUpdate(params *WeaviateThingsReferencesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsReferencesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsReferencesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.references.update",
		Method:             "PUT",
		PathPattern:        "/things/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsReferencesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsReferencesUpdateOK), nil

}

/*
WeaviateThingsUpdate updates a thing based on its UUID

Updates a Thing's data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) WeaviateThingsUpdate(params *WeaviateThingsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.update",
		Method:             "PUT",
		PathPattern:        "/things/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsUpdateOK), nil

}

/*
WeaviateThingsValidate validates things schema

Validate a Thing's schema and meta-data. It has to be based on a schema, which is related to the given Thing to be accepted by this validation.
*/
func (a *Client) WeaviateThingsValidate(params *WeaviateThingsValidateParams, authInfo runtime.ClientAuthInfoWriter) (*WeaviateThingsValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWeaviateThingsValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "weaviate.things.validate",
		Method:             "POST",
		PathPattern:        "/things/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WeaviateThingsValidateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WeaviateThingsValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
