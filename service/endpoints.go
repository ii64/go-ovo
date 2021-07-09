package service

import (
	"github.com/ii64/go-ovo/service/gateway"
	"github.com/ii64/go-ovo/service/primary"
)

type ServiceType int

const (
	ServiceType_Primary ServiceType = iota
	ServiceType_Gateway ServiceType = iota
)

type ServiceEndpoints struct {
	Primary *primary.Endpoints
	Gateway *gateway.Endpoints
}

func NewDefaultServiceEndpoints() *ServiceEndpoints {
	return &ServiceEndpoints{
		Primary: primary.NewDefaultEndpoints(),
		Gateway: gateway.NewDefaultEndpoints(),
	}
}
