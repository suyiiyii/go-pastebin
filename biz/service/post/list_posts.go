package post

import (
	"context"

	"pastebin/biz/dal/query"
	post "pastebin/hertz_gen/post"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ListPostsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListPostsService(Context context.Context, RequestContext *app.RequestContext) *ListPostsService {
	return &ListPostsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListPostsService) Run(req *post.ListPostsReq) (resp *post.ListPostsResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	size := int(req.Size)
	page := int(req.Page)
	p := query.Post
	posts, err := p.WithContext(h.Context).Offset((page - 1) * size).Limit(size).Find()
	if err != nil {
		return nil, err
	}
	resp = &post.ListPostsResp{
		Posts: make([]*post.Post, 0, len(posts)),
	}
	for _, po := range posts {
		resp.Posts = append(resp.Posts, &post.Post{
			Id:      int64(po.ID),
			Title:   po.Title,
			Content: po.Content,
		})
	}
	return
}
