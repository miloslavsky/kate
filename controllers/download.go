package controllers

import (
	"Kate/features"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
)

type DownloadController struct {
	beego.Controller
}

func (this *DownloadController) Get() {
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)

	if sess_username != "" {

		this.Data["LogStr"], this.Data["LogURL"] = features.Strings("", sess_userlang)
		this.Data["Main"] = features.Translate("Главная", sess_userlang)
		this.Data["Lang"] = features.Translate("Язык", sess_userlang)
		this.Data["Login"] = features.Translate("Логин", sess_userlang)
		this.Data["Password"] = features.Translate("Пароль", sess_userlang)
		this.Data["Enter"] = features.Translate("Войти", sess_userlang)
		this.TplNames = "download.tpl"

	} else {
		this.Redirect("/", 302)
	}
}

func (this *DownloadController) Download() {
	file, _, _ := this.GetFile("file")
	name := this.GetString("name")
	sess_id, _ := this.GetSession("userid").(int64)
	out, err := os.Create(fmt.Sprintf("../Downloads/%d/%s", sess_id, name))
	beego.Info(err)
	_, err = io.Copy(out, file)
	if err != nil {
		this.Redirect(`/download`, 302)
	} else {
		this.Redirect(`/`, 302)
	}
}
