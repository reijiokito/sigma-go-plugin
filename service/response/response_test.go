package response

import (
	"testing"

	"github.com/reijiokito/sigma-go-plugin/bridge"
	"github.com/stretchr/testify/assert"
)

var response Response
var ch chan interface{}

func init() {
	ch = make(chan interface{})
	response = New(ch)
}

func getBack(f func()) interface{} {
	go f()
	d := <-ch
	ch <- nil

	return d
}

func TestGetStatus(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.service.response.get_status"}, getBack(func() { response.GetStatus() }))
}

func TestGetHeader(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.service.response.get_header", Args: []interface{}{"foo"}}, getBack(func() { response.GetHeader("foo") }))
}

func TestGetHeaders(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.service.response.get_headers", Args: []interface{}{1}}, getBack(func() { response.GetHeaders(1) }))
}

func TestGetRawBody(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.service.response.get_raw_body"}, getBack(func() { response.GetRawBody() }))
}
