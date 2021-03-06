package controllers

import (
	"github.com/TruthHun/BookStack/models"
	"github.com/astaxie/beego"
)

type CateController struct {
	BaseController
}

//分类
func (this *CateController) List() {
	this.Data["IsCate"] = true
	if cates, err := new(models.Category).GetCates(-1, 1); err == nil {
		this.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}
	this.GetSeoByPage("cate", map[string]string{
		"title":       "书籍分类",
		"keywords":    "文档托管,在线创作,文档在线管理,在线知识管理,文档托管平台,在线写书,文档在线转换,在线编辑,在线阅读,开发手册,api手册,文档在线学习,技术文档,在线编辑",
		"description": this.Sitename + "专注于文档在线写作、协作、分享、阅读与托管，让每个人更方便地发布、分享和获得知识。",
	})
	this.TplName = "cates/list.html"
}
