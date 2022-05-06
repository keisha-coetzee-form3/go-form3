// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package metrics_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new metrics api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for metrics api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get metrics prometheus API v1 query API
*/
func (a *GetMetricsPrometheusAPIV1QueryRequest) Do() (*GetMetricsPrometheusAPIV1QueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetricsPrometheusAPIV1Query",
		Method:             "GET",
		PathPattern:        "/metrics/prometheus/api/v1/query",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetMetricsPrometheusAPIV1QueryReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetricsPrometheusAPIV1QueryOK), nil

}

func (a *GetMetricsPrometheusAPIV1QueryRequest) MustDo() *GetMetricsPrometheusAPIV1QueryOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get metrics prometheus API v1 query range API
*/
func (a *GetMetricsPrometheusAPIV1QueryRangeRequest) Do() (*GetMetricsPrometheusAPIV1QueryRangeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetricsPrometheusAPIV1QueryRange",
		Method:             "GET",
		PathPattern:        "/metrics/prometheus/api/v1/query_range",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetMetricsPrometheusAPIV1QueryRangeReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetricsPrometheusAPIV1QueryRangeOK), nil

}

func (a *GetMetricsPrometheusAPIV1QueryRangeRequest) MustDo() *GetMetricsPrometheusAPIV1QueryRangeOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get metrics prometheus federate API
*/
func (a *GetMetricsPrometheusFederateRequest) Do() (*GetMetricsPrometheusFederateOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetricsPrometheusFederate",
		Method:             "GET",
		PathPattern:        "/metrics/prometheus/federate",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetMetricsPrometheusFederateReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetricsPrometheusFederateOK), nil

}

func (a *GetMetricsPrometheusFederateRequest) MustDo() *GetMetricsPrometheusFederateOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post metrics prometheus API v1 query API
*/
func (a *PostMetricsPrometheusAPIV1QueryRequest) Do() (*PostMetricsPrometheusAPIV1QueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostMetricsPrometheusAPIV1Query",
		Method:             "POST",
		PathPattern:        "/metrics/prometheus/api/v1/query",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostMetricsPrometheusAPIV1QueryReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostMetricsPrometheusAPIV1QueryOK), nil

}

func (a *PostMetricsPrometheusAPIV1QueryRequest) MustDo() *PostMetricsPrometheusAPIV1QueryOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post metrics prometheus API v1 query range API
*/
func (a *PostMetricsPrometheusAPIV1QueryRangeRequest) Do() (*PostMetricsPrometheusAPIV1QueryRangeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostMetricsPrometheusAPIV1QueryRange",
		Method:             "POST",
		PathPattern:        "/metrics/prometheus/api/v1/query_range",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostMetricsPrometheusAPIV1QueryRangeReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostMetricsPrometheusAPIV1QueryRangeOK), nil

}

func (a *PostMetricsPrometheusAPIV1QueryRangeRequest) MustDo() *PostMetricsPrometheusAPIV1QueryRangeOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
