package errors

import "errors"

var ErrRepeatLogin = errors.New("重复登录")

var ErrNotOnlineLogin = errors.New("用户未登录")
