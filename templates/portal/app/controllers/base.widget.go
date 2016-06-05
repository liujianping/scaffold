package controllers

import (
	"fmt"
	goio "io"
	"os"
	"path"
	"time"

	cf "github.com/qiniu/api.v6/conf"
	io "github.com/qiniu/api.v6/io"
	rs "github.com/qiniu/api.v6/rs"
	"github.com/revel/revel"
)

type WidgetResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type WidgetController struct {
	*revel.Controller
}

func (widget WidgetController) Index() revel.Result {
	revel.TRACE.Printf("GET >> widget.index ...")
	return widget.RenderTemplate("widget/index.html")
}

func (widget WidgetController) Editor() revel.Result {
	revel.TRACE.Printf("GET >> widget.editor ...")
	return widget.RenderTemplate("widget/editor.html")
}

func (widget WidgetController) EditorPost() revel.Result {
	revel.TRACE.Printf("POST >> widget.editor ...")

	var callback string
	widget.Params.Bind(&callback, "CKEditorFuncNum")

	revel.INFO.Printf("POST >> widget.editor ...callback(%s)", callback)

	widget.Validation.Required(widget.Params.Files["upload"])
	if widget.Validation.HasErrors() {
		return widget.RenderJson(WidgetResponse{Code: 401, Message: "files absent"})
	}

	file := widget.Params.Files["upload"][0]
	fname := file.Filename
	ftype := file.Header.Get("Content-Type")
	flength := file.Header.Get("Content-Length")

	revel.INFO.Printf("POST >> widget.upload file(%s:%s:%s) ...", fname, ftype, flength)

	fd, err := file.Open()
	if err != nil {
		html := fmt.Sprintf("<html><body><script type=\"text/javascript\">alert('%s')</script></body></html>", err.Error())
		return widget.RenderHtml(html)
	}
	defer fd.Close()

	var url string
	if !QiniuEnable {
		sub := fmt.Sprintf("%d", time.Now().Unix()%31)
		dir := path.Join(revel.BasePath, "upload", sub)
		os.MkdirAll(dir, os.ModePerm)

		wr, err := os.Create(path.Join(dir, fname))
		if err != nil {
			return widget.RenderJson(WidgetResponse{Code: 402, Message: err.Error()})
		}
		defer wr.Close()

		if _, err := goio.Copy(wr, fd); err != nil {
			return widget.RenderJson(WidgetResponse{Code: 403, Message: err.Error()})
		}

		url = fmt.Sprintf("/upload/%s/%s", sub, fname)
	} else {
		var ret io.PutRet
		var extra = &io.PutExtra{}

		uptoken := QiniuToken(QiniuBucket)
		err = io.PutWithoutKey(nil, &ret, uptoken, fd, extra)
		if err != nil {
			return widget.RenderJson(WidgetResponse{Code: 405, Message: err.Error()})
		}
		url = fmt.Sprintf("http://%s/%s", QiniuDomain, ret.Key)
	}

	html := fmt.Sprintf("<html><body><script type=\"text/javascript\">window.parent.CKEDITOR.tools.callFunction(%s, '%s')</script></body></html>",
		callback, url)
	return widget.RenderHtml(html)
}

func (widget WidgetController) Upload() revel.Result {
	revel.TRACE.Printf("GET >> widget.upload ...")
	return widget.RenderTemplate("widget/upload.html")
}

func (widget WidgetController) UploadPost() revel.Result {
	revel.TRACE.Printf("POST >> widget.upload ...")

	widget.Validation.Required(widget.Params.Files["upload"])
	if widget.Validation.HasErrors() {
		return widget.RenderJson(WidgetResponse{Code: 401, Message: "files absent"})
	}

	file := widget.Params.Files["upload"][0]
	fname := file.Filename
	ftype := file.Header.Get("Content-Type")
	flength := file.Header.Get("Content-Length")

	revel.INFO.Printf("POST >> widget.upload file(%s:%s:%s) ...", fname, ftype, flength)

	fd, err := file.Open()
	if err != nil {
		return widget.RenderJson(WidgetResponse{Code: 402, Message: err.Error()})
	}
	defer fd.Close()

	var url string
	if !QiniuEnable {
		sub := fmt.Sprintf("%d", time.Now().Unix()%31)
		dir := path.Join(revel.BasePath, "upload", sub)
		os.MkdirAll(dir, os.ModePerm)

		wr, err := os.Create(path.Join(dir, fname))
		if err != nil {
			return widget.RenderJson(WidgetResponse{Code: 402, Message: err.Error()})
		}
		defer wr.Close()

		if _, err := goio.Copy(wr, fd); err != nil {
			return widget.RenderJson(WidgetResponse{Code: 403, Message: err.Error()})
		}

		url = fmt.Sprintf("/upload/%s/%s", sub, fname)
	} else {
		var ret io.PutRet
		var extra = &io.PutExtra{}

		uptoken := QiniuToken(QiniuBucket)
		err = io.PutWithoutKey(nil, &ret, uptoken, fd, extra)
		if err != nil {
			return widget.RenderJson(WidgetResponse{Code: 405, Message: err.Error()})
		}
		url = fmt.Sprintf("http://%s/%s", QiniuDomain, ret.Key)
	}

	return widget.RenderJson(WidgetResponse{Code: 0, Data: map[string]interface{}{
		"url": url,
	}})
}

func QiniuToken(bucket string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucket,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody:  returnBody,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	return putPolicy.Token(nil)
}

var (
	QiniuEnable bool
	QiniuAccess string
	QiniuSecret string
	QiniuBucket string
	QiniuDomain string
)

func InitWidget() {
	QiniuEnable = revel.Config.BoolDefault("qiniu.enable", false)
	QiniuAccess = revel.Config.StringDefault("qiniu.access", "")
	QiniuSecret = revel.Config.StringDefault("qiniu.secret", "")
	QiniuBucket = revel.Config.StringDefault("qiniu.bucket", "")
	QiniuDomain = revel.Config.StringDefault("qiniu.domain", "")

	cf.ACCESS_KEY = QiniuAccess
	cf.SECRET_KEY = QiniuSecret
}

func init() {
	revel.OnAppStart(InitWidget)
}
