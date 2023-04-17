package sigma

type InitPlugin interface {
	Access(pdk *PDK)
}

type PDK struct {
	Client        string
	Configuration Configuration
}

func Access(pdk *PDK) {
	//Function add plugin
}
