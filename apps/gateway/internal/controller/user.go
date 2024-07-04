package controller

import (
	"fmt"
	"ggblog-grpc/apps/gateway/rpc"
	userPb "ggblog-grpc/idl/pb/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserIsExist(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {
	var userReq userPb.UserRequest
	ctx.ShouldBindJSON(&userReq)
	resp, err := rpc.UserLogin(ctx, &userReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Status.Code,
		"msg":  resp.Status.Msg,
	})
}

// 添加用户
func AddUser(ctx *gin.Context) {
	var userReq userPb.UserRequest
	ctx.ShouldBindJSON(&userReq)
	resp, err := rpc.UserRegister(ctx, &userReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Status.Code,
		"msg":  resp.Status.Msg,
	})
}

// 查询用户列表
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	var userReq userPb.UserSearchRequest
	userReq.Pagesize = int32(pageSize)
	userReq.Pagenum = int32(pageNum)

	resp, err := rpc.UserList(ctx, &userReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  resp.Status.Code,
		"msg":   resp.Status.Msg,
		"users": resp.Users,
	})
}

// 通过关键字查询用户
func SearchUser(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	var userReq userPb.UserSearchRequest
	userReq.Keyword = keyword
	userReq.Pagesize = int32(pageSize)
	userReq.Pagenum = int32(pageNum)

	resp, err := rpc.UserSearchByKeyword(ctx, &userReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  resp.Status.Code,
		"msg":   resp.Status.Msg,
		"users": resp.Users,
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	var userReq userPb.UserTargetRequest
	id, _ := strconv.Atoi(ctx.Param("id"))
	userReq.Id = int32(id)
	ctx.ShouldBindJSON(&userReq.User)

	resp, err := rpc.UserEdit(ctx, &userReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Status.Code,
		"msg":  resp.Status.Msg,
	})
}

// 删除用户
func DeleteUser(ctx *gin.Context) {
	var userReq userPb.UserTargetRequest
	id, _ := strconv.Atoi(ctx.Param("id"))
	userReq.Id = int32(id)

	resp, err := rpc.UserDelete(ctx, &userReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Status.Code,
		"msg":  resp.Status.Msg,
	})
}
