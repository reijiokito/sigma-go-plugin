/*
Write to log file.
*/
package log

import (
	"github.com/Sigma/go-pdk/bridge"
)

// Holds this module's functions.  Accessible as `sigma.Log`
type Log struct {
	bridge.PdkBridge
}

// Called by the plugin server at initialization.
func New(ch chan interface{}) Log {
	return Log{bridge.New(ch)}
}

func (r Log) Alert(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.alert`, args...)
	return err
}

func (r Log) Crit(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.crit`, args...)
	return err
}

func (r Log) Err(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.err`, args...)
	return err
}

func (r Log) Warn(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.warn`, args...)
	return err
}

func (r Log) Notice(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.notice`, args...)
	return err
}

func (r Log) Info(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.info`, args...)
	return err
}

func (r Log) Debug(args ...interface{}) error {
	_, err := r.Ask(`sigma.log.debug`, args...)
	return err
}

var (
	modeSet     = map[string]string{"mode": "set"}
	modeAdd     = map[string]string{"mode": "add"}
	modeReplace = map[string]string{"mode": "replace"}
)

func (r Log) SetSerializeValue(key string, v interface{}) error {
	_, err := r.Ask(`sigma.log.set_serialize_value`, key, v, modeSet)
	return err
}

func (r Log) SetSerializeValueAdd(key string, v interface{}) error {
	_, err := r.Ask(`sigma.log.set_serialize_value`, key, v, modeAdd)
	return err
}

func (r Log) SetSerializeValueReplace(key string, v interface{}) error {
	_, err := r.Ask(`sigma.log.set_serialize_value`, key, v, modeReplace)
	return err
}

func (r Log) Serialize() (s string, err error) {
	return r.AskString(`sigma.log.serialize`)
}
