package rpc

import (
	"context"
	userPb "ggblog-grpc/idl/pb/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	resp, err = UserClient.UserRegister(ctx, req)
	return
}
func UserIsExist(ctx context.Context, req *userPb.UserRequest) (*userPb.UserCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserIsExist not implemented")
}
func UserList(ctx context.Context, req *userPb.UserSearchRequest) (*userPb.UserSearchRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func UserSearchByKeyword(ctx context.Context, req *userPb.UserSearchRequest) (*userPb.UserSearchRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSearchByKeyword not implemented")
}
func UserEdit(ctx context.Context, req *userPb.UserRequest) (*userPb.UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserEdit not implemented")
}
func UserDelete(ctx context.Context, req *userPb.UserRequest) (*userPb.UserCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func UserGetInfoByID(ctx context.Context, req *userPb.UserRequest) (*userPb.UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetInfoByID not implemented")
}
