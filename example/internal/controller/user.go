package controller

import (
	"context"

	"github.com/UncleChair/gf-jwt/v3/example/api"
	"github.com/UncleChair/gf-jwt/v3/example/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

type userController struct{}

var User = userController{}

// Info should be authenticated to view.
// It is the get user data handler
func (c *userController) Info(ctx context.Context, req *api.UserGetInfoReq) (res *api.UserGetInfoRes, err error) {
	return &api.UserGetInfoRes{
		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
		IdentityKey: service.Auth().IdentityKey,
		Payload:     service.Auth().GetPayload(ctx),
	}, nil
}
