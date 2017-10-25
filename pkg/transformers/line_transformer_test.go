package transformers

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"testing"

	"golang.org/x/text/transform"
)

func assertEqualStrings(t *testing.T, got, want string) {
	if len(got) != len(want) {
		t.Errorf("len(got)=%d len(want)=%d", len(got), len(want))
	}
	if got != want {
		t.Errorf("got=%s want=%s", got, want)
	}
}

func testLineTransformCase(t *testing.T, s string) {
	transformer := LineTransformer{}
	reader := transform.NewReader(strings.NewReader(s), &transformer)
	buffer := bytes.Buffer{}
	_, err := buffer.ReadFrom(reader)
	if err != nil {
		t.Fatal(err)
	}
	got := buffer.String()
	assertEqualStrings(t, got, s)
}

func TestLineTransform(t *testing.T) {
	t.Parallel()

	for i := 0; i < 2; i++ {
		func(n int) {
			ts := strings.Repeat("no EOL, trailing space, ", n)
			t.Run(fmt.Sprintf("len(ts)=%d, trailing space", len(ts)), func(t *testing.T) {
				t.Parallel()
				testLineTransformCase(t, ts)
			})

			nts := strings.Repeat("no EOL, no trailing space, ", n) + "done!"
			t.Run(fmt.Sprintf("len(nts)=%d, no trailing space", len(nts)), func(t *testing.T) {
				t.Parallel()
				testLineTransformCase(t, nts)
			})

			eolts := strings.Repeat("EOL, trailing space\n ", n)
			t.Run(fmt.Sprintf("len(eolts)=%d, trailing space", len(eolts)), func(t *testing.T) {
				t.Parallel()
				testLineTransformCase(t, eolts)
			})

			eolnts := strings.Repeat("EOL, no trailing space\n ", n) + "done!"
			t.Run(fmt.Sprintf("len(eolnts)=%d, no trailing space", len(eolnts)), func(t *testing.T) {
				t.Parallel()
				testLineTransformCase(t, eolnts)
			})
		}(int(math.Pow10(i)))
	}
}
