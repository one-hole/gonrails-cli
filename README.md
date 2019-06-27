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

` gonrails-cli generate model your-model-name `

` gonrails-cli generate controller your-controller-name`