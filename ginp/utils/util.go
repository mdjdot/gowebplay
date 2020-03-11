package utils

import (
	"math/rand"
	"time"
)

// RandName 生成随机名字
func RandName(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	newName := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		newName[i] = letters[rand.Intn(len(letters))]
	}
	return string(newName)
}
