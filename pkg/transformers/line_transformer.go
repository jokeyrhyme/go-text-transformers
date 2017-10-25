package transformers

import (
	"bytes"

	"golang.org/x/text/transform"
)

var (
	endOfLine = []byte("\n")
)

// LineTransformer processes one line at a time
type LineTransformer struct {
	wholeLine []byte
}

/*
things to be aware of:
-   atEOF?
-   EOL present?
-   line started in previous src?

scenarios (current src is between pipes):
 1.     |.........EOF|
 2.     |............|
 3.  ...|.........EOF|
 4.  ...|............|
 5.     |...EOL...EOF|
 6.     |...EOL......|
 7.  ...|...EOL...EOF|
 8.  ...|...EOL......|
*/

func (t *LineTransformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	eolIndex := bytes.Index(src, endOfLine)

	if eolIndex == -1 { // 1, 2, 3, 4
		t.wholeLine = append(t.wholeLine, src...)
		if atEOF {
			length := copy(dst, t.wholeLine)
			t.wholeLine = nil
			return length, len(src), nil // 1, 3
		}

		return 0, len(src), transform.ErrShortSrc // 2, 4
	}

	// 5, 6, 7, 8
	t.wholeLine = append(t.wholeLine, src[:eolIndex]...)
	if atEOF {
		length := copy(dst, t.wholeLine)
		t.wholeLine = nil
		return length, len(src), nil // 5, 7
	}

	length := copy(dst, t.wholeLine)
	t.wholeLine = src[eolIndex:]

	return length, eolIndex, transform.ErrShortSrc // 6, 8
}

func (t *LineTransformer) Reset() {
	t.wholeLine = nil
}
