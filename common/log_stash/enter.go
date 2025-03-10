package log_stash

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"im_server/common/contexts"
	"io"
	"net/http"
	"strings"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

type Pusher struct {
	LogType    int8   `json:"logType"`
	IP         string `json:"ip"`
	UserID     uint   `json:"userID"`
	Level      string `json:"level"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Service    string `json:"service"`
	client     *kq.Pusher
	items      []string
	ctx        context.Context
	isRequest  bool
	isHeaders  bool
	isResponse bool
	request    string
	headers    string
	count      int
	response   string
}

func (p *Pusher) IsRequest() {
	p.isRequest = true
}

func (p *Pusher) IsHeaders() {
	p.isHeaders = true
}

func (p *Pusher) IsResponse() {
	p.isResponse = true
}
func (p *Pusher) GetResponse() bool {
	return p.isResponse
}

// SetRequest 设置一组入参
func (p *Pusher) SetRequest(r *http.Request) {
	// 请求头
	// 请求体
	// 请求路径，请求方法
	// 关于请求体的问题，拿了之后要还回去
	// 一定要在参数绑定之前调用
	method := r.Method
	path := r.URL.String()
	byteData, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(byteData))
	p.request = fmt.Sprintf(
		`<div class="log_request">
  <div class="log_request_head">
    <span class="log_request_method %s">%s</span>
    <span class="log_request_path">%s</span>
  </div>
  <div class="log_request_body">
    <pre class="log_json_body">%s</pre>
  </div>
</div>`, strings.ToLower(method), method, path, string(byteData))
}

func (p *Pusher) SetHeaders(r *http.Request) {
	byteData, _ := json.Marshal(r.Header)
	p.headers = fmt.Sprintf(
		`<div class="log_request_header">
  <div class="log_request_body">
    <pre class="log_json_body">%s</pre>
  </div>
</div>`, string(byteData))
}

func (p *Pusher) SetResponse(w string) {
	p.response = fmt.Sprintf("<div class=\"log_response\">\n  <pre class=\"log_json_body\">%s</pre>\n</div>", w)
	p.Save(p.ctx)
}

func (p *Pusher) Info(title string) {
	p.Title = title
	p.Level = "info"
}
func (p *Pusher) Warning(title string) {
	p.Title = title
	p.Level = "warning"
}
func (p *Pusher) Err(title string) {
	p.Title = title
	p.Level = "err"
}

func (p *Pusher) setItem(level string, label string, val any) {
	var str string
	switch value := val.(type) {
	case string:
		str = fmt.Sprintf("<div class=\"log_item_label\">%s</div> <div class=\"log_item_content\">%s</div>", label, value)
	case int, uint, uint32, uint64, int32, int8:
		str = fmt.Sprintf("<div class=\"log_item_label\">%s</div> <div class=\"log_item_content\">%d</div>", label, value)
	default:
		byteData, _ := json.Marshal(val)
		str = fmt.Sprintf("<div class=\"log_item_label\">%s</div> <div class=\"log_item_content\">%s</div>", label, string(byteData))

	}
	logItem := fmt.Sprintf("<div class=\"log_item %s\">%s</div>", level, str)
	p.items = append(p.items, logItem)
	if p.LogType == 3 {
		// 如果是运行日志 那就调一下发送
		p.Save(p.ctx)
	}
}
func (p *Pusher) SetItem(label string, val any) {
	p.setItem("info", label, val)
}
func (p *Pusher) SetItemInfo(label, val string) {
	p.setItem("info", label, val)
}

func (p *Pusher) SetItemWarning(label, val string) {
	p.setItem("warning", label, val)
}
func (p *Pusher) SetItemErr(label, val string) {
	p.setItem("err", label, val)
}

func (p *Pusher) SetCtx(ctx context.Context) {
	p.ctx = ctx
}
func (p *Pusher) Save(ctx context.Context) {
	if p.isResponse && p.count == 0 {
		p.count = 1
		return
	}
	if ctx == nil {
		p.ctx = context.Background()
	}
	if p.ctx == nil {
		p.ctx = ctx
	}
	if p.client == nil {
		return
	}

	var items []string
	if p.isRequest {
		items = append(items, p.request)
	}
	if p.isHeaders {
		items = append(items, p.headers)
	}

	items = append(items, p.items...)
	if p.isResponse {
		items = append(items, p.response)
	}

	for _, item := range items {
		p.Content += item
	}

	p.items = []string{}

	userID, ok := p.ctx.Value(contexts.ContextKeyUserID).(uint)
	if !ok {
		logx.Error("userID 不存在")
		return
	}
	p.UserID = userID
	clientIP, ok := p.ctx.Value(contexts.ContextKeyClientIP).(string)
	if !ok {
		logx.Error("clientIP 不存在")
		return
	}
	p.IP = clientIP

	if p.Level == "" {
		p.Level = "info"
	}
	byteData, err := json.Marshal(p)
	if err != nil {
		logx.Error(err)
		return
	}
	err = p.client.Push(p.ctx, string(byteData))
	if err != nil {
		logx.Error(err)
		return
	}
}

func NewPusher(client *kq.Pusher, LogType int8, serviceName string) *Pusher {
	return &Pusher{
		Service: serviceName,
		client:  client,
		LogType: LogType,
	}

}
func NewActionPusher(client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(client, 2, serviceName)
}
func NewRuntimePusher(client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(client, 3, serviceName)
}
