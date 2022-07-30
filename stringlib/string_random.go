package stringlib

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
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

// src : https://gist.github.com/asessa/3aaec43d93044fc42b7c6d5f728cb039
func StrPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}
