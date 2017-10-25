package transformers

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/text/transform"
)

func TestNoopTransform(t *testing.T) {
	t.Parallel()

	type Case struct {
		input string
		want  string
	}
	cases := []Case{
		{
			input: "hello",
			want:  "hello",
		},
		{
			input: strings.Repeat("hello,world!", 10),
			want:  strings.Repeat("hello,world!", 10),
		},
		{
			input: strings.Repeat("hello, world! ", 10),
			want:  strings.Repeat("hello, world! ", 10),
		},
		{
			input: strings.Repeat("hello, world! ", 1000),
			want:  strings.Repeat("hello, world! ", 1000),
		},
		{
			input: strings.Repeat("hello, world! ", 1000000),
			want:  strings.Repeat("hello, world! ", 1000000),
		},
	}
	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				transformer := NoopTransformer{}
				reader := transform.NewReader(strings.NewReader(c.input), &transformer)
				buffer := bytes.Buffer{}
				_, err := buffer.ReadFrom(reader)
				if err != nil {
					t.Error(err)
				}
				got := buffer.String()
				if len(got) != len(c.want) {
					t.Errorf("len(got)=%d len(want)=%d", len(got), len(c.want))
				}
				if got != c.want {
					t.Error("got != want")
				}
			})
		}(i, c)
	}
}
