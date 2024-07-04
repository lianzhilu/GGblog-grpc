package rpc

import (
	"context"
	articlePb "ggblog-grpc/idl/pb/article"
)

func ArticleCreate(ctx context.Context, req *articlePb.ArticleRequest) (resp *articlePb.ArticleCommonResponse, err error) {
	resp, err = ArticleClient.ArticleCreate(ctx, req)
	return
}
func ArticleDetail(ctx context.Context, req *articlePb.ArticleTargetRequest) (resp *articlePb.ArticleDetailResponse, err error) {
	resp, err = ArticleClient.ArticleDetail(ctx, req)
	return
}
func ArticleSearchBykeyword(ctx context.Context, req *articlePb.ArticleSearchRequest) (resp *articlePb.ArticleSearchResponse, err error) {
	resp, err = ArticleClient.ArticleSearchBykeyword(ctx, req)
	return
}
func ArticleSearchByUser(ctx context.Context, req *articlePb.ArticleSearchRequest) (resp *articlePb.ArticleSearchResponse, err error) {
	resp, err = ArticleClient.ArticleSearchByUser(ctx, req)
	return
}
func ArticleList(ctx context.Context, req *articlePb.ArticleSearchRequest) (resp *articlePb.ArticleSearchResponse, err error) {
	resp, err = ArticleClient.ArticleList(ctx, req)
	return
}
func ArticleEdit(ctx context.Context, req *articlePb.ArticleTargetRequest) (resp *articlePb.ArticleDetailResponse, err error) {
	resp, err = ArticleClient.ArticleEdit(ctx, req)
	return
}
func ArticleDelete(ctx context.Context, req *articlePb.ArticleTargetRequest) (resp *articlePb.ArticleCommonResponse, err error) {
	resp, err = ArticleClient.ArticleDelete(ctx, req)
	return
}
