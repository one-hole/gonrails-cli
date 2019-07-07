package helper

import (
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
