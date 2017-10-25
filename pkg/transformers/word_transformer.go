package transformers

import (
	"regexp"

	"golang.org/x/text/transform"
)

var nonWordRegexp = regexp.MustCompile("\\W")

// WordTransformer processes up to the next word boundary, one "word" at a time
type WordTransformer struct {
	transform.NopResetter
}

func (t *WordTransformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	// TODO: probably should anticipate that a word crosses the src boundary,
	// requiring multiple reads
	var (
		err  error
		word []byte
	)
	nonWordIndex := nonWordRegexp.FindIndex(src)
	if len(nonWordIndex) == 2 {
		word = make([]byte, nonWordIndex[1])
		copy(word, src[:nonWordIndex[1]])
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

	length := copy(dst, word)
	return length, len(word), err
}
