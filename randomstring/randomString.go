package randomstring

import (
	"math/rand"
	"time"
)

func Generate(length int) string {

	rand.Seed(time.Now().Unix())

	ran_str := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = byte(65 + rand.Intn(25))
	}

	// Displaying the random string
	return string(ran_str)
}
