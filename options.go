package ovo

import (
	"github.com/ii64/go-ovo/service"
)

type Options struct {
	*App
	*service.ServiceEndpoints
}

func NewDefaultOptions() *Options {
	return &Options{
		App:              NewDefaultApp(),
		ServiceEndpoints: service.NewDefaultServiceEndpoints(),
	}
}
