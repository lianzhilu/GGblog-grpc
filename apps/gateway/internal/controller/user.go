package controller

import (
	"ggblog-grpc/apps/gateway/rpc"
	userPb "ggblog-grpc/idl/pb/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserIsExist(ctx *gin.Context) {

}

// 添加用户
func AddUser(ctx *gin.Context) {
	var userReq userPb.UserRequest
	ctx.ShouldBindJSON(&userReq)
	resp, _ := rpc.UserRegister(ctx, &userReq)
	ctx.JSON(http.StatusOK, gin.H{
		"status": resp.Code,
	})
}

// 查询用户列表
func GetUsers(ctx *gin.Context) {

}

// 通过关键字查询用户
func SearchUser(ctx *gin.Context) {

}

// 编辑用户
func EditUser(ctx *gin.Context) {

}

// 删除用户
func DeleteUser(ctx *gin.Context) {

}
