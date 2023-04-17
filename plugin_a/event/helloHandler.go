package event

import (
	"fmt"
	"github.com/reijiokito/plugin-manager/sigma"
	"sigma-plugin-a/proto"

	"log"
)

func HelloHandler(ctx *sigma.Context, message *proto.AccountCreated) {
	log.Println(fmt.Sprintf("Plugin A: Receive event Hello: %v", message))
}
