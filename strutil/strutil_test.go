package strutil

import (
	"io/ioutil"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func Test_ReplaceInvalidRune(t *testing.T) {
	assert.Equal(t, ReplaceInvalidRune(""), "")
	assert.Equal(t, ReplaceInvalidRune("1"), "1")
	assert.Equal(t, ReplaceInvalidRune("123456"), "123456")
	assert.Equal(t, ReplaceInvalidRune("\x80123456"), "*123456")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0123456"), "**123456")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0123\xfe456"), "**123*456")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0123\xff\xfe456"), "**123**456")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0123\xff\xfe\x88\x89456"), "**123****456")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0123\xff\xfe\x88\x89456\x99"), "**123****456*")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0123\xff\xfe\x88\x89456\x99\x9a"), "**123****456**")
	assert.Equal(t, ReplaceInvalidRune("\x88"), "*")
	assert.Equal(t, ReplaceInvalidRune("\x89\x88"), "**")
	assert.Equal(t, ReplaceInvalidRune("\x89\x88\x89\x88"), "****")
	assert.Equal(t, ReplaceInvalidRune("赵"), "赵")
	assert.Equal(t, ReplaceInvalidRune("赵钱孙李周武"), "赵钱孙李周武")
	assert.Equal(t, ReplaceInvalidRune("\x80赵钱孙李周武"), "*赵钱孙李周武")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0赵钱孙李周武"), "**赵钱孙李周武")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0赵钱孙\xfe李周武"), "**赵钱孙*李周武")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0赵钱孙\xff\xfe李周武"), "**赵钱孙**李周武")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0赵钱孙\xff\xfe\x88\x89李周武"), "**赵钱孙****李周武")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0赵钱孙\xff\xfe\x88\x89李周武\x99"), "**赵钱孙****李周武*")
	assert.Equal(t, ReplaceInvalidRune("\x80\xf0赵钱孙\xff\xfe\x88\x89李周武\x99\x9a"), "**赵钱孙****李周武**")
}

func Test_ReplaceInvalidRune2(t *testing.T) {
	assert.Equal(t, ReplaceInvalidRune2(""), "")
	assert.Equal(t, ReplaceInvalidRune2("1"), "1")
	assert.Equal(t, ReplaceInvalidRune2("123456"), "123456")
	assert.Equal(t, ReplaceInvalidRune2("\x80123456"), "*123456")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0123456"), "**123456")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0123\xfe456"), "**123*456")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0123\xff\xfe456"), "**123**456")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0123\xff\xfe\x88\x89456"), "**123****456")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0123\xff\xfe\x88\x89456\x99"), "**123****456*")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0123\xff\xfe\x88\x89456\x99\x9a"), "**123****456**")
	assert.Equal(t, ReplaceInvalidRune2("\x88"), "*")
	assert.Equal(t, ReplaceInvalidRune2("\x89\x88"), "**")
	assert.Equal(t, ReplaceInvalidRune2("\x89\x88\x89\x88"), "****")
	assert.Equal(t, ReplaceInvalidRune2("赵"), "赵")
	assert.Equal(t, ReplaceInvalidRune2("赵钱孙李周武"), "赵钱孙李周武")
	assert.Equal(t, ReplaceInvalidRune2("\x80赵钱孙李周武"), "*赵钱孙李周武")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0赵钱孙李周武"), "**赵钱孙李周武")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0赵钱孙\xfe李周武"), "**赵钱孙*李周武")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0赵钱孙\xff\xfe李周武"), "**赵钱孙**李周武")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0赵钱孙\xff\xfe\x88\x89李周武"), "**赵钱孙****李周武")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0赵钱孙\xff\xfe\x88\x89李周武\x99"), "**赵钱孙****李周武*")
	assert.Equal(t, ReplaceInvalidRune2("\x80\xf0赵钱孙\xff\xfe\x88\x89李周武\x99\x9a"), "**赵钱孙****李周武**")
}

func Benchmark_ReplaceInvalidRune(b *testing.B) {
	dat, err := ioutil.ReadFile("testdata/invalid.dat")
	assert.Nil(b, err)
	str := string(dat)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceInvalidRune(str)
	}
}

func Benchmark_ReplaceInvalidRune2(b *testing.B) {
	dat, err := ioutil.ReadFile("testdata/invalid.dat")
	assert.Nil(b, err)
	str := string(dat)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceInvalidRune2(str)
	}
}

func Benchmark_replaceInvalidRune3(b *testing.B) {
	dat, err := ioutil.ReadFile("testdata/invalid.dat")
	assert.Nil(b, err)
	str := string(dat)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceInvalidRune3(str)
	}
}

func Benchmark_replaceInvalidRune4(b *testing.B) {
	dat, err := ioutil.ReadFile("testdata/invalid.dat")
	assert.Nil(b, err)
	str := string(dat)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceInvalidRune4(str)
	}
}

func replaceInvalidRune3(s string) string {
	if len(s) == 0 {
		return s
	}
	var res string
	var v rune
	for _, v = range s {
		if v != rune(0xfffd) {
			res = res + string(v)
			continue
		}
		res = res + "*"
	}
	return res
}

// replaceInvalidRune4 replaces invalid runes to '*'.
// Use bytes.Buffer
func replaceInvalidRune4(s string) string {
	if len(s) == 0 {
		return s
	}
	var builder strings.Builder
	startIdx := 0
	idx := 0
	for _, v := range s {
		if v != rune(0xfffd) {
			idx += utf8.RuneLen(v)
			continue
		}
		if startIdx < idx {
			builder.WriteString(s[startIdx:idx])
		}
		builder.WriteString("*")
		idx += 1
		startIdx = idx
	}
	if startIdx < idx {
		builder.WriteString(s[startIdx:idx])
	}
	return builder.String()
}
