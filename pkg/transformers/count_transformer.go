package transformers

import (
	"fmt"
)

// CountTransformer keeps a running total in state and outputs the total number of bytes
type CountTransformer struct {
	count int
}

func (t *CountTransformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	t.count = t.count + len(src)
	if !atEOF {
		return 0, len(src), nil
	}
	length := copy(dst, []byte(fmt.Sprintf("%d", t.count)))
	return length, len(src), nil
}

func (t *CountTransformer) Reset() {
	t.count = 0
}
