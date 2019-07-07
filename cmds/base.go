package cmds

var (
	dirs          = []string{"config", "routes", "models", "controllers", "serializers", "helper", "workers", "logs"}
	AvailableCMDS = []string{"new", "help", "generate"}
)

// 后期可能打算把 CMD 封装成一个 Struct 。比如有shortcut 、alias 、params 比较方便来处理
