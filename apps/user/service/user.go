package service

import (
	"context"
	"errors"
	"fmt"
	"ggblog-grpc/apps/user/dao/mysql"
	userPb "ggblog-grpc/idl/pb/user"
	"ggblog-grpc/pkg/errmsg"
)

type UserServer struct {
	userPb.UnimplementedUserServiceServer
}

func newUserResponse(user mysql.User) *userPb.UserResponse {
	uresp := new(userPb.UserResponse)
	uresp.Id = int32(user.ID)
	uresp.Username = user.Username
	uresp.Password = user.Password
	uresp.Gender = int32(user.Gender)
	return uresp
}

func newUserResponseSlice(users []mysql.User) []*userPb.UserResponse {
	uresps := make([]*userPb.UserResponse, 0)
	for i := 0; i < len(users); i++ {
		uresp := newUserResponse(users[i])
		uresps = append(uresps, uresp)
	}
	return uresps
}

func (us *UserServer) UserLogin(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	fmt.Println(req.Username)
	fmt.Println(req.Password)
	resp = new(userPb.UserDetailResponse)
	resp.Status = new(userPb.UserStatusResponse)

	user, code := mysql.NewUserModel(ctx).CheckLogin(req.Username, req.Password)
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}

	resp.User = newUserResponse(user)

	return
}
func (us *UserServer) UserRegister(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	user := new(mysql.User)
	user.Username = req.Username
	user.Password = req.Password
	user.Gender = int(req.Gender)
	resp = new(userPb.UserCommonResponse)
	resp.Status = new(userPb.UserStatusResponse)

	code := mysql.NewUserModel(ctx).CreateUser(user)
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}
	return

}
func (us *UserServer) UserIsExist(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserCommonResponse, err error) {
	resp = new(userPb.UserCommonResponse)
	resp.Status = new(userPb.UserStatusResponse)

	code := mysql.NewUserModel(ctx).CheckUser(req.Username)
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}

	return
}
func (us *UserServer) UserList(ctx context.Context, req *userPb.UserSearchRequest) (resp *userPb.UserSearchResponse, err error) {
	resp = new(userPb.UserSearchResponse)
	resp.Status = new(userPb.UserStatusResponse)

	fmt.Println("Pagesize:", req.Pagesize)
	fmt.Println("Pagenum:", req.Pagenum)

	users, _ := mysql.NewUserModel(ctx).GetUsers(int(req.Pagesize), int(req.Pagenum))
	fmt.Println("users: ", users)

	code := errmsg.SUCCESS
	if len(users) <= 0 {
		code = errmsg.ERROR_RESULT_NOT_FOUND
		err = errors.New(errmsg.GetErrorMessage(code))
	} else {
		resp.Users = newUserResponseSlice(users)
	}
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)

	return
}
func (us *UserServer) UserSearchByKeyword(ctx context.Context, req *userPb.UserSearchRequest) (resp *userPb.UserSearchResponse, err error) {
	resp = new(userPb.UserSearchResponse)
	resp.Status = new(userPb.UserStatusResponse)

	users, _ := mysql.NewUserModel(ctx).GetUsersByKeyWord(req.Keyword, int(req.Pagesize), int(req.Pagenum))
	code := errmsg.SUCCESS
	if len(users) <= 0 {
		code = errmsg.ERROR_RESULT_NOT_FOUND
		err = errors.New(errmsg.GetErrorMessage(code))
	} else {
		resp.Users = newUserResponseSlice(users)
	}
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)

	return
}
func (us *UserServer) UserEdit(ctx context.Context, req *userPb.UserTargetRequest) (resp *userPb.UserCommonResponse, err error) {
	user := new(mysql.User)
	user.Username = req.User.Username
	user.Password = req.User.Password
	user.Gender = int(req.User.Gender)
	resp = new(userPb.UserCommonResponse)
	resp.Status = new(userPb.UserStatusResponse)

	code := mysql.NewUserModel(ctx).UpdateUser(int(req.Id), user)
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}

	return
}
func (us *UserServer) UserDelete(ctx context.Context, req *userPb.UserTargetRequest) (resp *userPb.UserCommonResponse, err error) {
	resp = new(userPb.UserCommonResponse)
	resp.Status = new(userPb.UserStatusResponse)

	code := mysql.NewUserModel(ctx).DeleteUser(int(req.Id))
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}

	return
}
func (us *UserServer) UserGetInfoByID(ctx context.Context, req *userPb.UserRequest) (resp *userPb.UserDetailResponse, err error) {
	return
}
