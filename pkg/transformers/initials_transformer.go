package transformers

import (
	"bytes"

	"golang.org/x/text/transform"
)

const (
	initialsLength = 1
)

var (
	emptyBytes = make([]byte, 0)
)

// InitialsTransformer replaces each word with its uppercased initial
type InitialsTransformer struct{}

func (t *InitialsTransformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	var (
		err       error
		processed int
		word      []byte
	)
	nonWordIndex := nonWordRegexp.FindIndex(src)
	if len(nonWordIndex) == 2 {
		word = make([]byte, nonWordIndex[1])
		if nonWordIndex[0] != 0 {
			copy(word, src[0:1])
			processed = len(word)
		} else {
			processed = 1
		}
		if !atEOF {
			// we have more, so go ask for more
			err = transform.ErrShortSrc
		} else {
			// we've run out, so cause a flush instead
			err = transform.ErrShortDst
		}
	} else {
		word = make([]byte, len(src))
		copy(word, src)
	}

	if len(word) > cap(dst) {
		return 0, 0, transform.ErrShortDst
	}

	length := copy(dst, bytes.ToUpper(word))
	return length, processed, err
}

func (t *InitialsTransformer) Reset() {}
