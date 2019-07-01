` gonrails-cli new yourproject-name `

1. 生成相应的项目（目录和文件）
2. 生成 Go Mod 文件
3. 执行 `go mod tidy`


```
mkdir config
mkdir routes
mkdir controllers
mkdir models

exec `go mod init`
```

` gonrails-cli help`

` gonrails-cli generate model your-model-name `

` gonrails-cli generate controller your-controller-name`

` gonrails-cli generate help`


---
目前已经可以使用的功能:

1. __gonrails-cli new__ (unfinished)

    `gonrails-cli new yourproject`
    
    `cd yourproject`
    
    `go mon init yourproject`
    
    `GO_ENV=debug go run main.go`

---


一些其他的思考：

1. 如何把 `Gonrails` 项目里面的部分东西组件化抽象，这样就可以避免生成新的琐碎的代码
2. 如何使用大仓库整合 `Gonrails` 和 `gonrails-cli` 两个项目