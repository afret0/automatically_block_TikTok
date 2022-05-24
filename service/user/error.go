package user

import "errors"

var CheckVerificationCodeError error

func init() {
	CheckVerificationCodeError = errors.New("check verification code error")
}
