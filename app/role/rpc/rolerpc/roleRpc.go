// Code generated by goctl. DO NOT EDIT.
// Source: role.proto

package rolerpc

import (
	"context"

	"xihu/app/role/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq  = pb.GenerateTokenReq
	GenerateTokenResp = pb.GenerateTokenResp
	GetUserInfoReq    = pb.GetUserInfoReq
	GetUserInfoResp   = pb.GetUserInfoResp
	GetVerCodeReq     = pb.GetVerCodeReq
	GetVerCodeResp    = pb.GetVerCodeResp
	LoginReq          = pb.LoginReq
	LoginResp         = pb.LoginResp
	RegisterReq       = pb.RegisterReq
	RegisterResp      = pb.RegisterResp
	UserAuth          = pb.UserAuth
	UserInform        = pb.UserInform

	RoleRpc interface {
		GetVerCode(ctx context.Context, in *GetVerCodeReq, opts ...grpc.CallOption) (*GetVerCodeResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	}

	defaultRoleRpc struct {
		cli zrpc.Client
	}
)

func NewRoleRpc(cli zrpc.Client) RoleRpc {
	return &defaultRoleRpc{
		cli: cli,
	}
}

func (m *defaultRoleRpc) GetVerCode(ctx context.Context, in *GetVerCodeReq, opts ...grpc.CallOption) (*GetVerCodeResp, error) {
	client := pb.NewRoleRpcClient(m.cli.Conn())
	return client.GetVerCode(ctx, in, opts...)
}

func (m *defaultRoleRpc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewRoleRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultRoleRpc) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewRoleRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultRoleRpc) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewRoleRpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultRoleRpc) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewRoleRpcClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}
