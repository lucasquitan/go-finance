package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(number int) string {
	var sb strings.Builder

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < number; i++ {
		c := alphabet[r.Intn(10)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
