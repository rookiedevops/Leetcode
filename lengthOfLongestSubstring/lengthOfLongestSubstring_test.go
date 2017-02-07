package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	s    string
	want int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	c := []Case{
		Case{
			s:    "abcdefcd",
			want: 6,
		},
		Case{
			s:    "ddddd",
			want: 1,
		},
		Case{
			s:    "a",
			want: 1,
		},
		Case{
			s:    "",
			want: 0,
		},
		Case{
			s:    "au",
			want: 2,
		},
	}
	for _, i := range c {
		w := lengthOfLongestSubstring(i.s)
		assert.Equal(t, i.want, w)
	}
}
