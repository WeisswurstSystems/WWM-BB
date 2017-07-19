package util

import (
	"github.com/FabianWilms/GoReadableID/readableId"
	"math/rand"
)

func GetReadableUID() string {
	return readableId.GetRandomID()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GetUID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}