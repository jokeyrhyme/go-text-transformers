package transformers

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/text/transform"
)

func TestInitialsTransform(t *testing.T) {
	t.Parallel()

	type Case struct {
		input string
		want  string
	}
	cases := []Case{
	// {
	// 	input: "hello",
	// 	want:  "H",
	// },
	// {
	// 	input: strings.Repeat("hello, world! ", 10),
	// 	want:  strings.Repeat("H, W! ", 10),
	// },
	// {
	// 	input: strings.Repeat("hello,world!", 10),
	// 	want:  strings.Repeat("H,W!", 10),
	// },
	// {
	// 	input: strings.Repeat("hello, world! ", 1000),
	// 	want:  strings.Repeat("H, W! ", 1000),
	// },
	// {
	// 	// too long for single Transform(), 1 word
	// 	input: strings.Repeat("abcdefghij", 500),
	// 	want:  "A",
	// },
	// {
	// 	input: strings.Repeat("hello, world! ", 1000000),
	// 	want:  strings.Repeat("H, W! ", 1000000),
	// },
	}
	for i, c := range cases {
		func(i int, c Case) {
			t.Run(fmt.Sprintf("cases[%d]", i), func(t *testing.T) {
				t.Parallel()

				transformer := InitialsTransformer{}
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
					t.Errorf("got=%s want=%s", got, c.want)
				}
			})
		}(i, c)
	}
}
