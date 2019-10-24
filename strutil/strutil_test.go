package strutil

import (
	"testing"

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
