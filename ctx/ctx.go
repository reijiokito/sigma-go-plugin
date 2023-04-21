/*
Current request context data.

`sigma.ctx.shared`:
A table that has the lifetime of the current request and is shared between
all plugins. It can be used to share data between several plugins in a given
request.

Values inserted in this table by a plugin will be visible by all other
plugins.  One must use caution when interacting with its values, as a naming
conflict could result in the overwrite of data.

Usage:
// Two plugins A and B, if plugin A has a higher priority than B's
// (it executes before B), and plugin A is a Go plugin:

// plugin A PluginA.go

	func (conf Config) Access(sigma *pdk.PDK) {
		err := sigma.Ctx.SetShared("hello world")
		if err != nil {
			sigma.Log.Err(err)
			return
		}
	}

// plugin B handler.lua
function plugin_b_handler:access(conf)

	sigma.log(sigma.ctx.shared.foo) // "hello world"

end
*/
package ctx

import (
	"github.com/Sigma/go-pdk/bridge"
)

// Holds this module's functions.  Accessible as `sigma.Ctx`
type Ctx struct {
	bridge.PdkBridge
}

// Called by the plugin server at initialization.
func New(ch chan interface{}) Ctx {
	return Ctx{bridge.New(ch)}
}

// sigma.Ctx.SetShared() sets a value in the `sigma.ctx.shared` request context table.
func (c Ctx) SetShared(k string, value interface{}) error {
	_, err := c.Ask(`sigma.ctx.shared.set`, k, value)
	return err
}

// sigma.Ctx.GetSharedAny() returns a value from the `sigma.ctx.shared` request context table.
func (c Ctx) GetSharedAny(k string) (interface{}, error) {
	return c.Ask(`sigma.ctx.shared.get`, k)
}

// sigma.Ctx.GetSharedString() returns a string value from the `sigma.ctx.shared` request context table.
func (c Ctx) GetSharedString(k string) (string, error) {
	return c.AskString(`sigma.ctx.shared.get`, k)
}

// sigma.Ctx.GetSharedFloat() returns a float value from the `sigma.ctx.shared` request context table.
func (c Ctx) GetSharedFloat(k string) (float64, error) {
	return c.AskFloat(`sigma.ctx.shared.get`, k)
}

// sigma.Ctx.GetSharedInt() returns an integer value from the `sigma.ctx.shared` request context table.
func (c Ctx) GetSharedInt(k string) (int, error) {
	return c.AskInt(`sigma.ctx.shared.get`, k)
}
