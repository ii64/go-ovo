package ovo

import (
	gateway "github.com/ii64/go-ovo/service/gateway"
	primary "github.com/ii64/go-ovo/service/primary"
)

type Ovo struct {
	options *Options

	*primary.PrimaryService
	*gateway.GatewayService
}

func New(opt *Options) (r *Ovo) {
	if opt == nil {
		opt = NewDefaultOptions()
	}
	r = &Ovo{
		options: opt,
	}
	return
}
