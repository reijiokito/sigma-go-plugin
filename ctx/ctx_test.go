package ctx

import (
	"testing"

	"github.com/reijiokito/sigma-go-plugin/bridge"
	"github.com/stretchr/testify/assert"
)

var ctx Ctx
var ch chan interface{}

func init() {
	ch = make(chan interface{})
	ctx = New(ch)
}

func getBack(f func()) interface{} {
	go f()
	d := <-ch
	ch <- nil

	return d
}

func TestSetShared(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.ctx.shared.set", Args: []interface{}{"key", "value"}}, getBack(func() { ctx.SetShared("key", "value") }))
}

func TestGetSharedAny(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.ctx.shared.get", Args: []interface{}{"key"}}, getBack(func() { ctx.GetSharedAny("key") }))
}

func TestGetSharedString(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.ctx.shared.get", Args: []interface{}{"key"}}, getBack(func() { ctx.GetSharedString("key") }))
}

func TestGetSharedFloat(t *testing.T) {
	assert.Equal(t, bridge.StepData{Method: "sigma.ctx.shared.get", Args: []interface{}{"key"}}, getBack(func() { ctx.GetSharedFloat("key") }))
}
