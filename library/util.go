package library

import (
	"math/rand"
	"strings"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// StrIndex 返回 str 中 第 index 个 substr 的位置
// EXP s := "0:00:000:0000:00000"
// index(s, ":", 1) = 1
// index(s, ":", 2) = 4
// index(s, ":", 3) = 8
// index(s, ":", 100) = -1
func StrIndex(str string, substr string, index int) int {
	sum := 0
	for i := 0; i < index; i++ {
		p := strings.Index(str, substr)
		if p == -1 {
			return -1
		}
		sum += p
		str = str[p+1:]
	}
	return sum + index - 1
}
