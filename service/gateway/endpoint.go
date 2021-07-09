package gateway

var (
	EndpointBase = "https://apigw01.aws.ovo.id"
)

type Endpoints struct {
	Base string
}

func NewDefaultEndpoints() *Endpoints {
	return &Endpoints{}
}
