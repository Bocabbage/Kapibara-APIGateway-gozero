package errors

import (
	"fmt"
	"net/http"
)

// ------- General Errors -------
type GeneralErrorCode int

const (
	InvalidParamsError GeneralErrorCode = 1 + iota
)

var generalHttpErrorMap = map[GeneralErrorCode]int{
	InvalidParamsError: http.StatusBadRequest,
}

type GeneralError struct {
	Code    GeneralErrorCode
	Message string
}

func (e *GeneralError) Error() string {
	return fmt.Sprintf("[GeneralError]%d: %s", int(e.Code), e.Message)
}

// ------- User Authentication Errors -------
type UserAuthErrorCode int

const (
	UserNotFoundError UserAuthErrorCode = 1 + iota
	IncorrectPwdError
	KeyNotExistInMapError
	JwtTokenParseError
)

var userAuthHttpErrorMap = map[UserAuthErrorCode]int{
	UserNotFoundError:     http.StatusNotFound,
	IncorrectPwdError:     http.StatusUnauthorized,
	KeyNotExistInMapError: http.StatusUnauthorized,
	JwtTokenParseError:    http.StatusUnauthorized,
}

type UserAuthError struct {
	Code    UserAuthErrorCode
	Message string
}

func (e *UserAuthError) Error() string {
	return fmt.Sprintf("[UserAuthError]%d: %s", int(e.Code), e.Message)
}
