package util

import (
	"github.com/FabianWilms/GoReadableID/readableId"
	"math/rand"
	"time"
)

func GetReadableUID() string {
	return readableId.GetRandomID()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")


var zufall = rand.New(rand.NewSource(time.Now().Unix()))

func GetUID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[zufall.Intn(len(letters))]
	}
	return string(b)
}