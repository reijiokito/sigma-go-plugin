package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	js, _ := nc.JetStream()

	streamA, _ := js.AddStream(&nats.StreamConfig{
		Name:     "plugin_a",
		Subjects: []string{"plugin_a.*"},
	})

	streamB, _ := js.AddStream(&nats.StreamConfig{
		Name:     "plugin_b",
		Subjects: []string{"plugin_b.*"},
	})

	fmt.Println(streamA)
	fmt.Println(streamB)
}
