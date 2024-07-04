package controller

import (
	"fmt"
	"ggblog-grpc/apps/gateway/rpc"
	articlePb "ggblog-grpc/idl/pb/article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(ctx *gin.Context) {

}

// 查询单个文章的详细信息
func GetArticleByID(ctx *gin.Context) {
	var articleReq articlePb.ArticleTargetRequest
	id, _ := strconv.Atoi(ctx.Param("id"))
	articleReq.Id = int32(id)

	resp, err := rpc.ArticleDetail(ctx, &articleReq)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    resp.Status.Code,
		"msg":     resp.Status.Msg,
		"article": resp.Article,
	})
}

// 根据关键字搜索文章
func SearchArticle(ctx *gin.Context) {

}

// 查询某个用户发布的所有文章
func SearchUserArticle(ctx *gin.Context) {

}

// 查询文章列表
func GetArticles(ctx *gin.Context) {

}

// 编辑文章
func EditArticle(ctx *gin.Context) {

}

// 删除文章
func DeleteArticle(ctx *gin.Context) {

}
