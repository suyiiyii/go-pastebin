package post

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	post "pastebin/hertz_gen/post"
)

type CreatePostService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreatePostService(Context context.Context, RequestContext *app.RequestContext) *CreatePostService {
	return &CreatePostService{RequestContext: RequestContext, Context: Context}
}

func (h *CreatePostService) Run(req *post.CreatePostRequest) (resp *post.CreatePostResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
