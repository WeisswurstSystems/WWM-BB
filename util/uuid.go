package util

import (
	"github.com/FabianWilms/GoReadableID/readableId"
)

func GetUID(n int) string {
	return readableId.GetRandomID()
}
