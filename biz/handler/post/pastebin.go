package post

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	service "pastebin/biz/service/post"
	"pastebin/biz/utils"
	post "pastebin/hertz_gen/post"
)

// GetPost .
// @router /posts/{id} [GET]
func GetPost(ctx context.Context, c *app.RequestContext) {
	var err error
	var req post.GetPostReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp := &post.GetPostResp{}
	resp, err = service.NewGetPostService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListPosts .
// @router /posts [GET]
func ListPosts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req post.ListPostsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp := &post.ListPostsResp{}
	resp, err = service.NewListPostsService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreatePost .
// @router /posts [POST]
func CreatePost(ctx context.Context, c *app.RequestContext) {
	var err error
	var req post.CreatePostReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp := &post.CreatePostResp{}
	resp, err = service.NewCreatePostService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
