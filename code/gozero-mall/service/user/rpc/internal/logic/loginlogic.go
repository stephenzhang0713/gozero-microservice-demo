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

// 用户登录流程，通过手机号查询判断用户是否是注册用户，如果是注册用户，需要将用户输入的密码进行加密与数据库中用户加密密码进行对比验证。

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Errorf(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != in.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.LoginResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
