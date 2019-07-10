package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/one-hole/gonrails-cli/helper"
)

var validActions = map[string]bool{
	"index":  true,
	"create": true,
	"delete": true,
	"update": true,
	"show":   true,
}

var basePath string

type ventory struct {
	ModuleName string
	ActionName string
}

func init() {
	if "true" == os.Getenv("DEV") {
		basePath = fmt.Sprintf("%s/watermelon", helper.Pwd)
	} else {
		basePath, _ = os.Getwd() // 非开发环境、那么这里应该是项目的根目录
	}
}

func Exec(args []string) {
	if len(args) == 2 {
		log.Fatal("Missing controller name")
	}

	mkdir(args[2])
	touchActions(args)
}

func mkdir(path string) {
	var dirPath string
	if os.Getenv("DEV") == "true" {
		dirPath = fmt.Sprintf("%s%s/%s", helper.Pwd, "/watermelon/controllers", path)
	} else {
		path, _ = os.Getwd() // 这个命令我希望是在项目的目录里面执行
		dirPath = fmt.Sprintf("%s%s/%s", path, "/controllers", path)
	}

	err := helper.FindOrCreateDir(dirPath)
	helper.PanicError(err)
}

func touchActions(args []string) {
	actionArgs := args[3:]

	for _, action := range actionArgs {
		touchAction(action, args[2])
	}
}

func touchAction(action, path string) {
	if _, ok := validActions[action]; ok {
		helper.CreateFile(
			fmt.Sprintf("%s/controllers/%s/%s.go", basePath, path, action),
			fmt.Sprintf("%s/templates/controllers/action.go.template", helper.ProjectPath),
			ventory{
				ModuleName: path,
				ActionName: action,
			},
		)
	}
}

/*
	1. nested
  2. namespace
	3. TODO 自动大写转小写
*/

/*
	gonrails-cli generate controller user
	gonrails-cli generate controller admin/user

查找或者创建目录
创建相关的 Action 文件
*/

/*
	需要判断是否是 gonrails 项目 & 命令需要在
*/
