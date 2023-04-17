package event

import (
	"fmt"
	"github.com/reijiokito/plugin-manager/sigma"
	"sigma-plugin-a/proto"

	"log"
)

func HelloHandler(ctx *sigma.Context, message *proto.Hello) {
	log.Println(fmt.Sprintf("Plugin B: Receive event Hello: %v", message))

}
