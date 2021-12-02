package helpers

import "math/rand"

const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const length = 8
const maxInt = 99999

func GenerateString() string {
	buf := make([]byte, length)
	for i := 1; i < length; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}

func GenerateInteger() int {
	return rand.Intn(maxInt)
}
