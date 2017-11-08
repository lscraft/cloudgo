package controllers

import (
	"github.com/astaxie/beego"
)
//handle"/"
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world!\n")
}
//handle"/user/:id"
type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	uid := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("hello " + uid + "!\n")
}
