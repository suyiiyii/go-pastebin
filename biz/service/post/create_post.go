package post

import (
	"context"
	"math/rand"

	"pastebin/biz/dal/model"
	"pastebin/biz/dal/query"
	post "pastebin/hertz_gen/post"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CreatePostService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreatePostService(Context context.Context, RequestContext *app.RequestContext) *CreatePostService {
	return &CreatePostService{RequestContext: RequestContext, Context: Context}
}

func (h *CreatePostService) Run(req *post.CreatePostReq) (resp *post.CreatePostResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	p := query.Post
	po := &model.Post{
		Title:   req.Post.Title,
		Content: req.Post.Content,
	}
	// generate 10-digit random number
	po.ID = uint(rand.Intn(9000000000) + 1000000000)
	err = p.WithContext(h.Context).Create(po)
	if err != nil {
		return nil, err
	}
	resp = &post.CreatePostResp{
		Post: &post.Post{
			Id:      int64(po.ID),
			Title:   po.Title,
			Content: po.Content,
		},
	}
	return
}
