package primary

var (
	EndpointBase = "https://api.ovo.id"
)

type Endpoints struct {
	Base string
}

func NewDefaultEndpoints() *Endpoints {
	return &Endpoints{}
}
