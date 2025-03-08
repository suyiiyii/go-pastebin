package post

import (
	"context"

	"pastebin/biz/dal/query"
	post "pastebin/hertz_gen/post"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type GetPostService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetPostService(Context context.Context, RequestContext *app.RequestContext) *GetPostService {
	return &GetPostService{RequestContext: RequestContext, Context: Context}
}

func (h *GetPostService) Run(req *post.GetPostReq) (resp *post.GetPostResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	p := query.Post
	po, err := p.WithContext(h.Context).Where(p.ID.Eq(uint(req.Id))).First()
	if err != nil {
		return nil, err
	}
	resp = &post.GetPostResp{
		Post: &post.Post{
			Id:      int64(po.ID),
			Title:   po.Title,
			Content: po.Content,
		},
	}

	return
}
