package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozero-mall/common/cryptx"
	"gozero-mall/service/user/model"

	"gozero-mall/service/user/rpc/internal/svc"
	"gozero-mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

// 用户注册流程，先判断注册手机号是否已经被注册，手机号未被注册，将用户信息写入数据库，用户密码需要进行加密存储。

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil {
		return nil, status.Error(100, "手机号已经注册")
	}

	if err == model.ErrNotFound {
		// 注册
		newUser := model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(&newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:     newUser.Id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
