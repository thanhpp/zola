package application

import (
	"math/rand"
	"strings"
	"time"
)

const source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(length int) string {
	seedRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var strB = new(strings.Builder)
	strB.Grow(length)
	for i := 0; i < length; i++ {
		strB.WriteByte(source[seedRand.Intn(len(source))])
	}

	return strB.String()
}
