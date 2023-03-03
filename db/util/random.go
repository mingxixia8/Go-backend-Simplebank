package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) //make sure the generate value will be different for every time

}

// RandomInt to generate a random intger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min -> max

}

// RandomString to generate a random string with length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] //get random int between 0 to length - 1
		sb.WriteByte(c)             //put random char to sb
	}
	return sb.String()
}

// RandomOwner to generate a random owner name with 6 chars
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney to generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency to generate a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
