package errors

import (
	"go/types"
	"net/http"

	xhttp "github.com/zeromicro/x/http"
)

func HttpErrorHandler(err error) (int, any) {
	switch e := err.(type) {
	case *UserAuthError:
		httpCode, ok := userAuthHttpErrorMap[e.Code]
		if !ok {
			httpCode = http.StatusBadRequest
		}
		return httpCode, xhttp.BaseResponse[types.Nil]{
			Code: int(e.Code),
			Msg:  e.Message,
		}
	case *GeneralError:
		httpCode, ok := generalHttpErrorMap[e.Code]
		if !ok {
			httpCode = http.StatusBadRequest
		}
		return httpCode, xhttp.BaseResponse[types.Nil]{
			Code: int(e.Code),
			Msg:  e.Message,
		}
	default:
		return http.StatusInternalServerError, nil
	}
}
