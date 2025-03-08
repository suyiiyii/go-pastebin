package post

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	post "pastebin/hertz_gen/post"
)

type GetPostService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetPostService(Context context.Context, RequestContext *app.RequestContext) *GetPostService {
	return &GetPostService{RequestContext: RequestContext, Context: Context}
}

func (h *GetPostService) Run(req *post.GetPostRequest) (resp *post.GetPostResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
