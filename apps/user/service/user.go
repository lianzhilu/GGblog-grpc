package service

import (
	"context"
	"errors"
	"ggblog-grpc/apps/user/model"
	userPb "ggblog-grpc/idl/pb/user"
	"ggblog-grpc/pkg/errmsg"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userPb.UnimplementedUserServiceServer
}

func (us *UserServer) UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (us *UserServer) UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	user := new(model.User)
	user.Username = req.Username
	user.Password = req.Password
	user.Gender = int(req.Gender)
	resp = new(userPb.UserCommonResponse)

	code := model.NewUserModel(ctx).CreateUser(user)
	resp.Code = int32(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
		return
	}
	return

}
func (us *UserServer) UserIsExist(ctx context.Context, req *userPb.UserRequest) (*userPb.UserCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserIsExist not implemented")
}
func (us *UserServer) UserList(ctx context.Context, req *userPb.UserSearchRequest) (*userPb.UserSearchRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (us *UserServer) UserSearchByKeyword(ctx context.Context, req *userPb.UserSearchRequest) (*userPb.UserSearchRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSearchByKeyword not implemented")
}
func (us *UserServer) UserEdit(ctx context.Context, req *userPb.UserRequest) (*userPb.UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserEdit not implemented")
}
func (us *UserServer) UserDelete(ctx context.Context, req *userPb.UserRequest) (*userPb.UserCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (us *UserServer) UserGetInfoByID(ctx context.Context, req *userPb.UserRequest) (*userPb.UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetInfoByID not implemented")
}
