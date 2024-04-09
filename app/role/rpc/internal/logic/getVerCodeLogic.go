package logic

import (
	"context"

	"xihu/app/role/rpc/internal/svc"
	"xihu/app/role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVerCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVerCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerCodeLogic {
	return &GetVerCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVerCodeLogic) GetVerCode(in *pb.GetVerCodeReq) (*pb.GetVerCodeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVerCodeResp{}, nil
}
