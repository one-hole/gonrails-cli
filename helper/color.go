package helper

import (
	"fmt"
	"io"
)

// 这里提供终端的ASCII输出

type colorWriter struct {
	w io.Writer
}

// Write 实现了 io.Writer 的 Write 接口
func (cw *colorWriter) Write(p []byte) (int, error) {
	return cw.w.Write(p)
}

func bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

func red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}
