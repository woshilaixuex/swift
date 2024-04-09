package UserController

import (
	"context"

	"xihu/app/role/api/internal/svc"
	"xihu/app/role/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneVerCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneVerCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneVerCodeLogic {
	return &PhoneVerCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneVerCodeLogic) PhoneVerCode(req *types.PhoneVerCodeReq) (resp *types.PhoneVerCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
