package event

import (
	"fmt"
	"github.com/reijiokito/plugin-manager/sigma"
	"sigma-plugin-a/proto"

	"log"
)

func HandShakeHandler(ctx *sigma.Context, message *proto.AccountCreated) {
	log.Println(fmt.Sprintf("Receive event Hello: %v", message))
}
