module github.com/ii64/go-ovo/example/mock

go 1.16

replace github.com/ii64/go-ovo => ../../

require (
	github.com/gogo/gateway v1.1.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ii64/go-ovo v0.0.0-00010101000000-000000000000
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/rakyll/statik v0.1.7
	google.golang.org/grpc v1.39.0
)
