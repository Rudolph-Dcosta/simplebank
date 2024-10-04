package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxuz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// random int
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// random Owner
func RandomOwner() string {
	return RandomString(6)
}

// random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "INR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
