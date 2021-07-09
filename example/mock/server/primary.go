package server

import (
	"context"

	gen "github.com/ii64/go-ovo/gen"
)

func (b *Backend) GetBudget(ctx context.Context, req *gen.GetBudgetRequest) (res *gen.GetBudgetResponse, err error) {
	res = &gen.GetBudgetResponse{}
	return
}

func (b *Backend) GetFront(ctx context.Context, req *gen.GetFrontRequest) (res *gen.GetFrontResponse, err error) {
	res = &gen.GetFrontResponse{}
	return
}

func (b *Backend) Login2FA(ctx context.Context, req *gen.Login2FARequest) (res *gen.Login2FAResponse, err error) {
	res = &gen.Login2FAResponse{}
	return
}

func (b *Backend) Login2FAVerify(ctx context.Context, req *gen.Login2FAVerifyRequest) (res *gen.Login2FAVerifyResponse, err error) {
	res = &gen.Login2FAVerifyResponse{}
	return
}

func (b *Backend) LoginSecurityCode(ctx context.Context, req *gen.LoginSecurityCodeRequest) (res *gen.LoginSecurityCodeResponse, err error) {
	res = &gen.LoginSecurityCodeResponse{}
	return
}
