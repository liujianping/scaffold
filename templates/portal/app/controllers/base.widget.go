package controllers

import (
	"fmt"

	"[[.project]]/app/models"
	"github.com/liujianping/scaffold/symbol"
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

	url, err := Upload(fname, fd)
	if err != nil {
		return widget.RenderJson(WidgetResponse{Code: 405, Message: err.Error()})
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

	url, err := Upload(fname, fd)
	if err != nil {
		return widget.RenderJson(WidgetResponse{Code: 403, Message: err.Error()})
	}
	revel.INFO.Printf("POST >> widget.upload file (url:%s) completed.", url)
	return widget.RenderJson(WidgetResponse{Code: 0, Data: map[string]interface{}{
		"url": url,
	}})
}

func (widget WidgetController) SystemOptionPost(code, relate_code string) revel.Result {
	revel.TRACE.Printf("POST >> widget.system.option ...")

	var options []models.SystemOption
	if err := db.Where("code = ? and relate_code = ?", code, relate_code).Order("id ASC").Find(&options).Error; err != nil {
		return widget.RenderJson(WidgetResponse{Code: 401, Message: err.Error()})
	}

	return widget.RenderJson(WidgetResponse{Code: 0, Data: map[string]interface{}{
		"options": options,
	}})
}

func (widget WidgetController) TableObjectPost(table string, id int64) revel.Result {
	revel.TRACE.Printf("POST >> widget.table.object ...")
	obj := GetTableObject(table, id)
	if obj == nil {
		return widget.RenderJson(WidgetResponse{Code: 404, Message: "object no exist."})
	}
	return widget.RenderJson(WidgetResponse{Code: 0, Data: obj})
}

func (widget WidgetController) ManyCreatePost() revel.Result {
	revel.TRACE.Printf("POST >> widget.many.create ...")

	var relation, belong, many string
	var belong_id int64
	var manies []int64

	widget.Params.Bind(&relation, "relation")
	widget.Params.Bind(&belong, "belong")
	widget.Params.Bind(&many, "many")
	widget.Params.Bind(&belong_id, "belong_id")
	widget.Params.Bind(&manies, "manies")

	revel.INFO.Printf("POST >> widget.many.create ... manies:%v", manies)
	tx := db.Begin()
	if relation != many {
		for _, many_id := range manies {
			sql := fmt.Sprintf("INSERT INTO %s(%s_id, %s_id) VALUES(?, ?)", 
				   relation, symbol.Singular(belong), symbol.Singular(many))
			if err := tx.Exec(sql, belong_id, many_id).Error; err != nil {
				tx.Rollback()
				return widget.RenderJson(WidgetResponse{Code: 401, Message: err.Error()})
			}
		}
	}
	tx.Commit()

	return widget.RenderJson(WidgetResponse{Code: 0, Message: "ok"})
}

func (widget WidgetController) ManyRemovePost() revel.Result {
	revel.TRACE.Printf("POST >> widget.many.remove ...")

	var relation, belong, many string
	var belong_id int64
	var manies []int64

	widget.Params.Bind(&relation, "relation")
	widget.Params.Bind(&belong, "belong")
	widget.Params.Bind(&many, "many")
	widget.Params.Bind(&belong_id, "belong_id")
	widget.Params.Bind(&manies, "manies")

	revel.INFO.Printf("POST >> widget.many.remove ... manies:%v", manies)

	obj := models.DefaultTableObject(relation)
	tx := db.Begin()
	if relation != many {
		conditon := fmt.Sprintf("%s_id = ? and %s_id = ?", symbol.Singular(many), symbol.Singular(belong))	
		for _, many_id := range manies {
			if err := tx.Where(conditon, many_id, belong_id).Delete(obj).Error; err != nil {
				tx.Rollback()
				return widget.RenderJson(WidgetResponse{Code: 401, Message: err.Error()})
			}
		}
	} else {
		if err := tx.Where(manies).Delete(obj).Error; err != nil {
			tx.Rollback()
			return widget.RenderJson(WidgetResponse{Code: 401, Message: err.Error()})
		}
	}
	tx.Commit()

	return widget.RenderJson(WidgetResponse{Code: 0, Message: "ok"})
}

func (widget WidgetController) ManyAppendPost() revel.Result {
	revel.TRACE.Printf("POST >> widget.many.append ...")
	return widget.RenderJson(WidgetResponse{Code: 0, Message: "ok"})
}
