package router

import (
	"ggblog-grpc/apps/gateway/internal/controller"
	"ggblog-grpc/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.AppConf.Mode)
	r := gin.Default()

	// 需要jwt验证的接口
	authV1 := r.Group("api/v1")
	{
		// user模块
		// 编辑用户
		authV1.PUT("user/:id", controller.EditUser)

		// 删除用户
		authV1.DELETE("user/:id", controller.DeleteUser)

		// article 模块
		// 添加文章
		authV1.POST("article/add", controller.AddArticle)

		// 编辑文章
		authV1.PUT("article/:id", controller.EditArticle)

		// 删除文章
		authV1.DELETE("article/:id", controller.DeleteArticle)
	}

	// 不需要jwt验证的接口
	routerV1 := r.Group("/api/v1")
	{
		// 测试联通
		routerV1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})

		// 登录
		routerV1.POST("/login", controller.Login)

		// user模块
		// 添加用户
		routerV1.POST("user/add", controller.AddUser)

		// 获取用户列表
		routerV1.GET("user/users", controller.GetUsers)

		// 通过关键字查询用户
		routerV1.GET("user/search", controller.SearchUser)

		// article模块
		// 查询某个文章的详细信息
		routerV1.GET("article/:id", controller.GetArticleByID)

		// 根据关键字搜索文章
		routerV1.GET("article/search", controller.SearchArticle)

		// 查询某个用户发布的所有文章
		routerV1.GET("article/search/:userid", controller.SearchUserArticle)

		// 查询文章列表
		routerV1.GET("article/articles", controller.GetArticles)

	}

	r.Run(config.AppConf.Port)
}
