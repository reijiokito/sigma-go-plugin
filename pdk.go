/*
Package sigma/go-pdk implements sigma's Plugin Development Kit for Go.

It directly parallels the existing sigma PDK for Lua plugins.

sigma plugins written in Go implement event handlers as methods on the Plugin's
structure, with the given signature:

	func (conf *MyConfig) Access (sigma *pdk.PDK) {
		...
	}

The `sigma` argument of type `*pdk.PDK` is the entrypoint for all PDK functions.
For example, to get the client's IP address, you'd use `sigma.Client.GetIp()`.
*/
package pdk

import (
	"github.com/reijiokito/sigma-go-plugin/client"
	"github.com/reijiokito/sigma-go-plugin/ctx"
	"github.com/reijiokito/sigma-go-plugin/log"
	"github.com/reijiokito/sigma-go-plugin/request"
	"github.com/reijiokito/sigma-go-plugin/response"
	"github.com/reijiokito/sigma-go-plugin/service"
	service_request "github.com/reijiokito/sigma-go-plugin/service/request"
	service_response "github.com/reijiokito/sigma-go-plugin/service/response"
)

// PDK go pdk module
type PDK struct {
	Client          client.Client
	Ctx             ctx.Ctx
	Log             log.Log
	Request         request.Request
	Response        response.Response
	Service         service.Service
	ServiceRequest  service_request.Request
	ServiceResponse service_response.Response
}

// Init initialize go pdk.  Called by the pluginserver at initialization.
func Init(ch chan interface{}) *PDK {
	return &PDK{
		Client:          client.New(ch),
		Ctx:             ctx.New(ch),
		Log:             log.New(ch),
		Request:         request.New(ch),
		Response:        response.New(ch),
		Service:         service.New(ch),
		ServiceRequest:  service_request.New(ch),
		ServiceResponse: service_response.New(ch),
	}
}
