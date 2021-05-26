package exception

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	AUTH_FAIL : "用户验证失败",
	INVALID_PARAMS : "请求参数错误",
	INVALID_AUTH_TOKEN : "token校验错误",
	SERVER_ERROR : "服务器繁忙",
	TAG_ALREADY_EXISTS : "该标签已存在",
	TAG_NOT_EXIST : "该标签不存在",
	ARTICLE_NOT_EXIST : "该文章不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}