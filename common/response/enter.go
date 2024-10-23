package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
Response 函数 接受三个参数，分别为请求io,写入io,要响应的内容resp和错误内容
如果❌错误存在，那么就返回错误信息，如果错误不存在，那就返回正确内容
*/
func Response(r *http.Request, w http.ResponseWriter, resp any, err error) {
	if err == nil {
		//如果没有错误
		r := &Body{
			Code: 0,
			Msg:  "成功",
			Data: resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	// 错误返回
	errCode := uint32(7)
	httpx.WriteJson(w, http.StatusOK, &Body{
		Code: errCode,
		Msg:  err.Error(),
		Data: nil,
	})
}
