package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozero-mall/service/product/model"

	"gozero-mall/service/product/rpc/internal/svc"
	"gozero-mall/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Errorf(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Name != "" {
		res.Name = in.Name
	}
	if in.Desc != "" {
		res.Desc = in.Desc
	}
	if in.Stock != 0 {
		res.Stock = in.Stock
	}
	if in.Amount != 0 {
		res.Amount = in.Amount
	}
	if in.Status != 0 {
		res.Status = in.Status
	}
	err = l.svcCtx.ProductModel.Update(res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &product.UpdateResponse{}, nil
}
