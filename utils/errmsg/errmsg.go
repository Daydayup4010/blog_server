package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000xxx 用户模块的错误
	ERROR_USERNAME_USED    = 10001
	ERROR_PASSWORD_WRONG   = 10002
	ERROR_USER_NOT_EXIST   = 10003
	ERROR_TOKEN_EXIST      = 10004
	ERROR_TOKEN_TIMEOUT    = 10005
	ERROR_TOKEN_WRONG      = 10006
	ERROR_TOKEN_TYPE_WRONG = 10007

	// code = 2000xxx 分类模块的错误
	ERROR_CATEGORYNAME_USED  = 20001
	ERROR_CATEGORY_NOT_EXIST = 20002

	// code = 3000xxx 分类模块的错误
)

var codeMsg = map[int]string{
	SUCCESS:                  "success",
	ERROR:                    "error",
	ERROR_USERNAME_USED:      "用户名重复",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_TOKEN_EXIST:        "TOKEN不存在",
	ERROR_TOKEN_TIMEOUT:      "TOKEN已过期",
	ERROR_TOKEN_WRONG:        "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误",
	ERROR_CATEGORYNAME_USED:  "分类名重复",
	ERROR_CATEGORY_NOT_EXIST: "此分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
