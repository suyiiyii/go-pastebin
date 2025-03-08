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
	var req post.GetPostRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp := &post.GetPostResponse{}
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
	var req post.ListPostsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp := &post.ListPostsResponse{}
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
	var req post.CreatePostRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusBadRequest, err)
		return
	}

	resp := &post.CreatePostResponse{}
	resp, err = service.NewCreatePostService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
