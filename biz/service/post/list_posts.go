package post

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	post "pastebin/hertz_gen/post"
)

type ListPostsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListPostsService(Context context.Context, RequestContext *app.RequestContext) *ListPostsService {
	return &ListPostsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListPostsService) Run(req *post.ListPostsRequest) (resp *post.ListPostsResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
