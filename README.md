## __gonrails-cli__ is command line tools for project __[gonrails](https://github.com/one-hole/gonrails)__


### 1. Install

`go get -u  github.com/one-hole/gonrails-cli`

### 2. New

`gonrails-cli new yourproject-name`


 * so if I want to create a project named __kalista__ , I just run command __`gonrails-cli new kalista`__
 * the docs of the created project was in __[here](https://github.com/one-hole/gonrails)__


### 3. Generate

##### 3.1

 The command is `gonrails-cli generate controller yourcontroller action list`

##### 3.2



---
### N. Others

` gonrails-cli new yourproject-name `
    
    1. 生成相应的项目（目录和文件）
    2. 生成 Go Mod 文件
    3. 执行 `go mod tidy`

` gonrails-cli help [command]`

` gonrails-cli generate model your-model-name `

` gonrails-cli generate controller your-controller-name action list`

` gonrails-cli generate help`


* useage .
* help for each command (useage for each cmd)
* colorful output
* shorcut


---
目前已经可以使用的功能:

1. __gonrails-cli new__ 

---


一些其他的思考：

1. 如何把 `Gonrails` 项目里面的部分东西组件化抽象，这样就可以避免生成新的琐碎的代码****
2. 如何使用大仓库整合 `Gonrails` 和 `gonrails-cli` 两个项目