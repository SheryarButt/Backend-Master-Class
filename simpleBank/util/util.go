package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// init is used to seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// RandomOwner generates a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money.
func RandomMoney() int64 {
	return int64(RandomInt(1, 100))
}

// RandomCurrency generates a random currency name.
func RandomCurrency() string {
	currency := []string{"USD", "CAD", "EUR", "GBP", "JPY", "AUD", "NZD"}
	return currency[rand.Intn(len(currency))]
}
