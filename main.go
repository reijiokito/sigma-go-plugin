package main

import (
	"flag"
	"github.com/reijiokito/plugin-manager/sigma"
	"sigma-plugin-a/event"
)

const PLUGIN_CODE = "PLUGIN_A"

func main() {
	natsUrl := flag.String("nats_url", "127.0.0.1", "Nats URL")
	natsUsername := flag.String("nats_username", "", "Nats Username")
	natsPassword := flag.String("nats_password", "", "Nats Password")

	flag.Parse()

	config := sigma.Configuration{
		NatsUrl:      *natsUrl,
		NatsUsername: *natsUsername,
		NatsPassword: *natsPassword,
	}
	flag.Parse()

	sigma.Init(PLUGIN_CODE, &config)
	defer sigma.Release()

	sigma.RegisterEvent("PLUGIN_B", "hello", event.HelloHandler)
	
	sigma.Start()
}
