package ptools

import "github.com/gogf/gf/net/ghttp"

type resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 200 // 成功
	ERROR   = 400 // 异常
)

// 成功
func succ(data interface{}) resp {
	return resp{SUCCESS, "success", data}
}

// 失败
func fail(msg string, data interface{}) resp {
	if data != nil {
		return resp{ERROR, msg, data}
	}
	return resp{ERROR, msg, ""}
}

func Success(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(succ(data))
	r.Exit()
}

func Error(r *ghttp.Request, msg string, data ...interface{}) {
	ret := fail(msg, "")
	if len(data) > 0 {
		ret = fail(msg, data[0])
	}
	r.Response.WriteJson(ret)
	r.Exit()
}
