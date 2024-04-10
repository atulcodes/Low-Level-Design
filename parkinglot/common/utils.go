package common

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits = "0123456789"
	numberPlateLength = 6
)

func GenerateUniqueID() string {
	return uuid.New().String()
}

func GenerateRandomNumberPlate() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	byteArr := make([]byte, numberPlateLength)
	for idx := range byteArr {
		if idx < 2 {
			byteArr[idx] = alphabet[randomGenerator.Int63n(int64(len(alphabet)))]
		} else {
			byteArr[idx] = digits[randomGenerator.Int63n(int64(len(digits)))]
		}
	}
	return string(byteArr)
}