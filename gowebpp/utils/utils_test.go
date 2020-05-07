package utils

import (
	"path"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMD5(t *testing.T) {
	s1 := MD5("kkzz")
	s2 := MD5("kkzz")
	assert.Equal(t, s1, s2)
}

func TestPath(t *testing.T) {
	str := `/Users/medea/Pictures/golang-developer-roadmap-zh-CN.png`
	dir, name := path.Split(str)
	assert.Equal(t, name, dir)
}
