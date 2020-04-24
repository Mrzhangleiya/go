package main

import (
	_ "code.text.com/Work/beegoDemo/controllers/utils"
	_"code.text.com/Work/beegoDemo/models"
	// 下划线'_'表示只执行routers的init()函数
	_"code.text.com/Work/beegoDemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	//映射函数的执行必须在run之前被调用
	_ = beego.AddFuncMap("ShowPrePage", HandlePrePage)
	_ = beego.AddFuncMap("ShowNextPage", HandleNextPage)
	beego.Run()
}

func HandlePrePage(pageIndex int) int {
	pageIndex -= 1
	return pageIndex
}

func HandleNextPage(pageIndex int) int {
	pageIndex += 1
	return pageIndex
}
