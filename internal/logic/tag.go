package logic

import (
	"JuneBlog/internal/db"
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/message"
	"JuneBlog/patch/logger"
	"context"
)

func TagListReq(ctx context.Context, req message.None) (resp *message.TagListResp, err error) {
	tags, err := db.Db().GetTags(ctx)
	if err != nil {
		logger.Error("Tag: TagList fail", "err", err)
		return nil, err
	}
	resp = &message.TagListResp{
		RespHeader: message.RespOK,
		Tags:       make([]message.Tag, 0, len(tags)),
	}
	for _, tag := range tags {
		resp.Tags = append(resp.Tags, message.Tag{
			Id:   int32(tag.Id),
			Name: tag.Name,
		})
	}
	return resp, nil
}

func NewTagReq(ctx context.Context, req message.Tag) (resp *message.RespHeader, err error) {
	logger.Debug("Tag: NewTagReq", req)
	err = db.Db().NewTag(ctx, &module.Tag{Name: req.Name})
	if err != nil {
		logger.Error("Tag: NewTag fail", "err", err, "id", req.Id)
		return nil, err
	}
	return &message.RespOK, nil
}

func TagDeleteReq(
	ctx context.Context, id int, req message.None,
) (resp *message.RespHeader, err error) {
	err = db.Db().DelTag(ctx, int64(id))
	if err != nil {
		logger.Error("Tag: DeleteTag fail", "err", err, "id", id)
		return nil, err
	}
	return &message.RespOK, nil
}
