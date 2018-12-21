package util

import (
	"bytes"
	"math/rand"
)

var (
	s = "abcdefghijklmnopqrstuvwxyz0987654321-"
	l = len(s)
)

func RandStr() string {
	var buf bytes.Buffer
	for i := 0; i < 16; i++ {
		buf.WriteByte(s[rand.Intn(l)])
	}
	return buf.String()
}

func Rand() string {
	var buf bytes.Buffer
	for i := 0; i < 17; i++ {
		buf.WriteByte(s[rand.Intn(l)])
	}
	return buf.String()
}
