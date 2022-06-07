package services

import (
	"math/rand"
	"time"
)

const (
	a = 97
	z = 122
)

func GenerateRandomRequest() string {
	var requestName string
	for i := 0; i < 2; i++ {
		randomLetter := Random(a, z)
		requestName += string(byte(randomLetter))
	}

	return requestName
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}