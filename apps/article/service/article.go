package service

import (
	"context"
	"errors"
	"ggblog-grpc/apps/article/dao/mysql"
	"ggblog-grpc/apps/article/dao/redis"
	articlePb "ggblog-grpc/idl/pb/article"
	"ggblog-grpc/pkg/errmsg"
)

type ArticleServer struct {
	articlePb.UnimplementedArticleServiceServer
}

func newArticleResponse(article mysql.Article) *articlePb.ArticleResponse {
	resp := new(articlePb.ArticleResponse)
	resp.Id = int32(article.ID)
	resp.Title = article.Title
	resp.Description = article.Description
	resp.Content = article.Content
	resp.Userid = int32(article.UserID)
	return resp
}

func (as *ArticleServer) ArticleCreate(ctx context.Context, req *articlePb.ArticleRequest) (resp *articlePb.ArticleCommonResponse, err error) {
	article := new(mysql.Article)
	article.Title = req.Title
	article.Description = req.Description
	article.Content = req.Content
	article.UserID = uint(req.Userid)

	resp = new(articlePb.ArticleCommonResponse)
	resp.Status = new(articlePb.ArticleStatusResponse)

	code := mysql.NewArticleModel(ctx).CreateArticle(article)
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}

	return
}
func (as *ArticleServer) ArticleDetail(ctx context.Context, req *articlePb.ArticleTargetRequest) (resp *articlePb.ArticleDetailResponse, err error) {
	resp = new(articlePb.ArticleDetailResponse)
	resp.Status = new(articlePb.ArticleStatusResponse)

	article, code := mysql.NewArticleModel(ctx).GetArticleByID(int(req.Id))
	resp.Article = newArticleResponse(article)
	resp.Article.View = int32(redis.AddClicks(ctx, int(req.Id)))
	resp.Status.Code = int32(code)
	resp.Status.Msg = errmsg.GetErrorMessage(code)
	if code != errmsg.SUCCESS {
		err = errors.New(errmsg.GetErrorMessage(code))
	}

	return
}
func (as *ArticleServer) ArticleSearchBykeyword(ctx context.Context, req *articlePb.ArticleSearchRequest) (resp *articlePb.ArticleSearchResponse, err error) {
	return
}
func (as *ArticleServer) ArticleSearchByUser(ctx context.Context, req *articlePb.ArticleSearchRequest) (resp *articlePb.ArticleSearchResponse, err error) {
	return
}
func (as *ArticleServer) ArticleList(ctx context.Context, req *articlePb.ArticleSearchRequest) (resp *articlePb.ArticleSearchResponse, err error) {
	return
}
func (as *ArticleServer) ArticleEdit(ctx context.Context, req *articlePb.ArticleTargetRequest) (resp *articlePb.ArticleDetailResponse, err error) {
	return
}
func (as *ArticleServer) ArticleDelete(ctx context.Context, req *articlePb.ArticleTargetRequest) (resp *articlePb.ArticleCommonResponse, err error) {
	return
}
