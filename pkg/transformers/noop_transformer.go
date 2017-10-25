package transformers

import (
	"golang.org/x/text/transform"
)

// NoopTransformer passes incoming bytes straight through without modification
type NoopTransformer struct {
	transform.NopResetter
}

func (t *NoopTransformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	length := copy(dst, src)
	return length, len(src), nil
}
