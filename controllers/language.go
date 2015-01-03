package controllers

import (
	"github.com/astaxie/beego"
)

type LangController struct {
	beego.Controller
}

func (this *LangController) Post() {
	langs := map[string](int64){"eng": 1, "ru": 2, "chn": 3}
	lang := this.GetString(":lg")
	if _, exist := langs[lang]; exist {
		beego.Info(exist)
		this.SetSession("userlang", lang)
	}
	this.Redirect("/", 302)
}
