package auth

import (
	"context"
	"fmt"

	kerrors "kapibara-apigateway-gozero/internal/errors"
	"kapibara-apigateway-gozero/restful/auth/internal/models/mysql"
	"kapibara-apigateway-gozero/restful/auth/internal/svc"
	"kapibara-apigateway-gozero/restful/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// TODO: validator
	if req.Account == "" || req.Password == "" || req.Username == "" {
		return nil, &kerrors.GeneralError{
			Code:    kerrors.InvalidParamsError,
			Message: fmt.Sprintf("Invalid input for [%s]", req.Account),
		}
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		l.Errorf("[GenerateFromPassword][traceid: %s][%v]", req.TraceId, err)
		return nil, err
	}
	_, err = l.svcCtx.UsersModel.Insert(
		l.ctx,
		&mysql.Users{
			Account:  req.Account,
			PwdHash:  string(pwdHash),
			UserName: req.Username,
		},
	)
	if err != nil {
		// TODO: need enhancement
		l.Errorf("[InsertMySQL][traceid: %s][%v]", req.TraceId, err)
	}

	l.Infof("[Register]Account[%s], TraceId[%s]", req.Account, req.TraceId)
	return &types.RegisterResponse{
		Message: "Register success!",
	}, nil
}
