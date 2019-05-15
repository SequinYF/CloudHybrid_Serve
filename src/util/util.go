package util

type UserProfile struct {
	Name   string
	Passwd string
	Dir    string
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

//error
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_NAME     = 10001
	ERROR_NOT_EXIST_NAME = 10002
	ERROR_NOT_EXIST_DIR  = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_NAME:               "已存在该名称",
	ERROR_NOT_EXIST_NAME:           "该名称不存在",
	ERROR_NOT_EXIST_DIR:            "该目录不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
