package sigma_go_plugin

import "github.com/reijiokito/sigma-go-plugin/proto/generate"

type InitPlugin interface {
	Access(pdk *PDK)
}

type PDK struct {
	Client        string
	Configuration generate.Configuration
}

func Access(pdk *PDK) {
	//Function add plugin
}
