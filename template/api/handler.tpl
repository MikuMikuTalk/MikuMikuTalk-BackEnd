package handler

import (
    "net/http"
    "im_server/common/response"
    {{.ImportPackages}}
    {{if .HasRequest}}
    "github.com/zeromicro/go-zero/rest/httpx"
    {{end}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := httpx.Parse(r, &req); err != nil {
            			//解析请求体失败
			//❌就响应错误信息，使用common包中的response进行数据封装
             response.Response(r, w, nil, err)
            return
        }{{end}}

        l := logic.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        // 这里如果正常，err就是nil,响应的包装好的json数据里的code就是0,如果Open_login这个逻辑在调用中发生了错误，那么会把错误信息和响应包装在响应的json数据中
        {{if .HasResp}}response.Response(r, w, resp, err){{else}}response.Response(r,w, nil, err){{end}}

    }
}