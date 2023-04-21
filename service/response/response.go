/*
Manipulation of the response from the Service.
*/
package response

import (
	"github.com/reijiokito/sigma-go-plugin/bridge"
)

// Holds this module's functions.  Accessible as `sigma.ServiceResponse`
type Response struct {
	bridge.PdkBridge
}

// Called by the plugin server at initialization.
func New(ch chan interface{}) Response {
	return Response{bridge.New(ch)}
}

// sigma.ServiceResponse.GetStatus() returns the HTTP status code
// of the response from the Service as an integer.
func (r Response) GetStatus() (i int, err error) {
	return r.AskInt(`sigma.service.response.get_status`)
}

// sigma.ServiceResponse.GetHeaders() returns a map holding the headers
// from the response from the Service. Keys are header names.
// Values are either a string with the header value, or an array of strings
// if a header was sent multiple times. Header names in this table are
// case-insensitive and dashes (-) can be written as underscores (_);
// that is, the header X-Custom-Header can also be retrieved as x_custom_header.
//
// Unlike sigma.Response.GetHeaders(), this function will only return headers
// that were present in the response from the Service (ignoring headers added
// by sigma itself). If the request was not proxied to a Service
// (e.g. an authentication plugin rejected a request and produced an HTTP 401 response),
// then the returned headers value might be nil, since no response
// from the Service has been received.
//
// The max_args argument specifies the maximum number of returned headers.
// Must be greater than 1 and not greater than 1000, or -1 to specify the
// default limit of 100 arguments.
func (r Response) GetHeaders(max_headers int) (map[string][]string, error) {
	if max_headers == -1 {
		return r.AskMap(`sigma.service.response.get_headers`)
	}

	return r.AskMap(`sigma.service.response.get_headers`, max_headers)
}

// sigma.ServiceResponse.GetHeader() returns the value of the specified response header.
//
// Unlike sigma.Response.GetHeader(), this function will only return a header
// if it was present in the response from the Service
// (ignoring headers added by sigma itself).
func (r Response) GetHeader(name string) (string, error) {
	return r.AskString(`sigma.service.response.get_header`, name)
}

// sigma.ServiceResponse.GetRawBody() returns the raw body
// of the response from the Service.
func (r Response) GetRawBody() (string, error) {
	return r.AskString(`sigma.service.response.get_raw_body`)
}
