package service

import (
	"content_manage/api/operate"
	"content_manage/internal/biz"
	"context"
)

func (a *AppService) FindContent(ctx context.Context,
	req *operate.FindContentReq) (*operate.FindContentRsp, error) {
	//构建请求参数结构
	findParams := &biz.FindParams{
		ID:       req.GetId(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     req.Page,
		PageSize: req.GetPageSize(),
	}
	uc := a.uc
	results, total, err := uc.FindContent(ctx, findParams) //调用biz层的实现
	if err != nil {
		return nil, err
	}
	//构建返回的Rsp。FindContentRsp
	var contents []*operate.Content //构建内容
	//用一个for循环，遍历从数据库中拿出来的结果
	for _, r := range results {
		contents = append(contents, &operate.Content{
			Id:             r.ID,
			Title:          r.Title,
			VideoUrl:       r.VideoURL,
			Author:         r.Author,
			Description:    r.Description,
			Thumbnail:      r.Thumbnail,
			Category:       r.Category,
			Duration:       r.Duration.Milliseconds(),
			Resolution:     r.Resolution,
			FileSize:       r.FileSize,
			Format:         r.Format,
			Quality:        r.Quality,
			ApprovalStatus: r.ApprovalStatus,
		})
	}
	rsp := &operate.FindContentRsp{ //api的app.proto里面，定义了FindContentRsp结构需要两个内容，total和Contents
		Total:    total,
		Contents: contents,
	}
	return rsp, nil
}
