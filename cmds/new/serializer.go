package new

import (
	"fmt"

	"github.com/one-hole/gonrails-cli/helper"
)

func TouchSerializer(moduleName string) {
	helper.CreateFile(
		fmt.Sprintf("%s/%s/serializers/book_serializer.go", helper.Pwd, moduleName),
		fmt.Sprintf("%s/templates/serializers/book_serializer.go", helper.ProjectPath),
		nil,
	)
}
