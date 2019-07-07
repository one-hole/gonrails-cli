package cmds

import (
	"os"
)

var (
	dirs   = []string{"config", "routes", "models", "controllers", "serializers", "helper", "workers", "logs"}
	pwd, _ = os.Getwd()
)
