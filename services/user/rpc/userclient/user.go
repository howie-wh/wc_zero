// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"wczero/services/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdminLoginRequest     = user.AdminLoginRequest
	AdminLoginResponse    = user.AdminLoginResponse
	AdminRegisterRequest  = user.AdminRegisterRequest
	AdminRegisterResponse = user.AdminRegisterResponse
	AdminUserInfoRequest  = user.AdminUserInfoRequest
	AdminUserInfoResponse = user.AdminUserInfoResponse
	WeChatLoginRequest    = user.WeChatLoginRequest
	WeChatLoginResponse   = user.WeChatLoginResponse

	User interface {
		WeChatLogin(ctx context.Context, in *WeChatLoginRequest, opts ...grpc.CallOption) (*WeChatLoginResponse, error)
		AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
		AdminRegister(ctx context.Context, in *AdminRegisterRequest, opts ...grpc.CallOption) (*AdminRegisterResponse, error)
		AdminUserInfo(ctx context.Context, in *AdminUserInfoRequest, opts ...grpc.CallOption) (*AdminUserInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) WeChatLogin(ctx context.Context, in *WeChatLoginRequest, opts ...grpc.CallOption) (*WeChatLoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.WeChatLogin(ctx, in, opts...)
}

func (m *defaultUser) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AdminLogin(ctx, in, opts...)
}

func (m *defaultUser) AdminRegister(ctx context.Context, in *AdminRegisterRequest, opts ...grpc.CallOption) (*AdminRegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AdminRegister(ctx, in, opts...)
}

func (m *defaultUser) AdminUserInfo(ctx context.Context, in *AdminUserInfoRequest, opts ...grpc.CallOption) (*AdminUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AdminUserInfo(ctx, in, opts...)
}
