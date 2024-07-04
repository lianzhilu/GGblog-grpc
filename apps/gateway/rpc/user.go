package rpc

import (
	"context"
	userPb "ggblog-grpc/idl/pb/user"
)

func UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	resp, err = UserClient.UserLogin(ctx, req)
	return
}
func UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	resp, err = UserClient.UserRegister(ctx, req)
	return
}
func UserIsExist(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	resp, err = UserClient.UserIsExist(ctx, req)
	return
}
func UserList(ctx context.Context, req *userPb.UserSearchRequest) (resp *userPb.UserSearchResponse, err error) {
	resp, err = UserClient.UserList(ctx, req)
	return
}
func UserSearchByKeyword(ctx context.Context, req *userPb.UserSearchRequest) (resp *userPb.UserSearchResponse, err error) {
	resp, err = UserClient.UserSearchByKeyword(ctx, req)
	return
}
func UserEdit(ctx context.Context, req *userPb.UserTargetRequest) (resp *userPb.UserCommonResponse, err error) {
	resp, err = UserClient.UserEdit(ctx, req)
	return
}
func UserDelete(ctx context.Context, req *userPb.UserTargetRequest) (resp *userPb.UserCommonResponse, err error) {
	resp, err = UserClient.UserDelete(ctx, req)
	return
}
func UserGetInfoByID(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	resp, err = UserClient.UserGetInfoByID(ctx, req)
	return
}
