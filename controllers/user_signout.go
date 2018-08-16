package controllers

import (
  "github.com/astaxie/beego"
  _"github.com/astaxie/beego/cache/redis"
)


type SignoutController struct {
  beego.Controller
}


func (this *SignoutController) Get() {
  beego.Debug("In SignoutController:Get - Start")

  this.DelSession("session")

  beego.Debug("In SignoutController:Get - Deleted user session and redirecting back to /user/signin")

  this.Redirect("/user/signin", 303)
}
