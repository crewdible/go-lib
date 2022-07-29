package random

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandString(n int) string {
	var (
		letterBytes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b           = strconv.Itoa(int(time.Now().Unix()))
	)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < n; i++ {
		b += string(letterBytes[rand.Int63()%int64(len(letterBytes))])
	}
	return b
}
