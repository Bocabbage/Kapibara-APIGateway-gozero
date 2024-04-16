package auth

import (
	"context"
	"fmt"

	kerrors "kapibara-apigateway-gozero/internal/errors"
	"kapibara-apigateway-gozero/internal/utils"
	"kapibara-apigateway-gozero/restful/auth/internal/svc"
	"kapibara-apigateway-gozero/restful/auth/internal/types"

	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// TODO: validator
	if req.Account == "" || req.Password == "" {
		return nil, &kerrors.GeneralError{
			Code:    kerrors.InvalidParamsError,
			Message: fmt.Sprintf("Invalid input for [%s]", req.Account),
		}
	}

	// TODO: set redis cache
	sqlRes, err := l.svcCtx.UsersModel.FindOneByAccount(l.ctx, req.Account)
	if err == sqlx.ErrNotFound {
		return nil, &kerrors.UserAuthError{
			Code:    kerrors.UserNotFoundError,
			Message: fmt.Sprintf("Account: [%s] not exist.", req.Account),
		}
	} else if err != nil {
		l.Errorf("[FindMySQL][traceid: %s][%v]", req.TraceId, err)
		return nil, err
	}

	cmpResult := bcrypt.CompareHashAndPassword([]byte(sqlRes.PwdHash), []byte(req.Password))
	if cmpResult != nil {
		return nil, &kerrors.UserAuthError{
			Code:    kerrors.IncorrectPwdError,
			Message: "Incorrect password",
		}
	}

	jwtToken, err := utils.GenerateJWT(
		sqlRes.RoleBitmap, sqlRes.UserName, sqlRes.Account,
		l.svcCtx.Config.JwtExpired,
		l.svcCtx.Config.MySQLPwdSalt,
		l.svcCtx.Config.JwtSecretKey,
	)
	if err != nil {
		l.Errorf("[GenerateJWT][traceid: %s][%v]", req.TraceId, err)
		return nil, err
	}

	l.Infof("[LoginSuccess]Account[%s], TraceId[%s]", req.Account, req.TraceId)
	return &types.LoginResponse{
		AccessToken: jwtToken,
		UserName:    sqlRes.UserName,
		Uuid:        uuid.NewV4().String(),
	}, nil
}
