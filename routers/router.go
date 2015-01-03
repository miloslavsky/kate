package routers

import (
	"Kate/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.MainController{})
	beego.Router("/post", &controllers.MainController{})
	beego.Router("/delete", &controllers.MainController{})
	beego.Router("/delete/posts", &controllers.MainController{})
	beego.Router("/registration", &controllers.RegController{})
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{})
	beego.Router("/users/:id([0-9]+)", &controllers.UserController{})
	beego.Router("/login", &controllers.SessionController{})
	beego.Router("/logout", &controllers.SessionController{}, "*:Delete")
	beego.Router("/add", &controllers.BlogPostController{}, "get:Add")
	beego.Router("/add", &controllers.BlogPostController{}, "post:Post")
	beego.Router("/post/:id([0-9]+)", &controllers.BlogPostController{}, "get:Get")
	beego.Router("/posts/:id([0-9]+)", &controllers.BlogPostController{}, "get:Get")
	beego.Router("/delete/post/:id([0-9]+)", &controllers.BlogPostController{}, "*:Delete")
	beego.Router("/edit/post/:id([0-9]+)", &controllers.BlogPostController{}, "get:Edit")
	beego.Router("/edit/post/:id([0-9]+)", &controllers.BlogPostController{}, "post:Edition")
	beego.Router("/follow/:id([0-9]+)", &controllers.FollowController{}, "*:Follow")
	beego.Router("/followers/:id([0-9]+)", &controllers.FollowController{}, "get:Get")
	beego.Router("/following/:id([0-9]+)", &controllers.FollowController{}, "get:Following")
	beego.Router("/lang/:lg([a-z]+)", &controllers.LangController{}, "*:Post")
	beego.Router("/post/:id([0-9]+)", &controllers.CommentController{}, "post:Post")
	beego.Router("/posts/:id([0-9]+)", &controllers.CommentController{}, "post:Post")
	beego.Router("/delete/comment/:id([0-9]+)", &controllers.CommentController{}, "*:Delete")
	beego.Router("/download", &controllers.DownloadController{}, "get:Get")
	beego.Router("/download", &controllers.DownloadController{}, "post:Download")
}
