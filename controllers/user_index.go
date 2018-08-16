package controllers

import (
  "github.com/astaxie/beego"
)

// Base Controller Actions

type IndexController struct {
  BaseController
}


func (this *IndexController) Get() {
  beego.Debug("In IndexController.Get() - Started")
  this.LayoutSections["Footer"] = "indexcontroller/get.tpl"
}
