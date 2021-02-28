package log

import "io"

type nilWriter struct{}

func (w nilWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

var _ io.Writer = nilWriter{}
