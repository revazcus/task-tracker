package commonTesting

import (
	"crypto/rand"
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	"math/big"
	"strconv"
)

const (
	defaultMinNumber = int64(1000000)
	defaultMaxNumber = int64(9999999)

	floatNumberFrom0To1Base = 1000000
)

var (
	lettersForRandom = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandomByteSls() []byte {
	bytesCount := RandomNumber(5, 25)
	bytes := make([]byte, bytesCount)
	for i := 0; i < bytesCount; i++ {
		randomByte := byte(RandomNumber(1, 100) % 2)
		bytes[i] = randomByte
	}
	return bytes
}

func RandomDefaultNumberString() string {
	randomNumber := RandomDefaultNumber()
	return strconv.Itoa(randomNumber)
}

func RandomDefaultString() string {
	return RandomString(15)
}

func RandomString(length int) string {
	randomLetters := make([]rune, length)
	for i := 0; i < length; i++ {
		randomLetterIndex := RandomNumber(0, int64(len(lettersForRandom)-1))
		randomLetters[i] = lettersForRandom[randomLetterIndex]
	}
	return string(randomLetters)
}

func RandomDefaultNumber() int {
	return RandomNumber(defaultMinNumber, defaultMaxNumber)
}

func RandomNumber(minNumber, maxNumber int64) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(minNumber-maxNumber))
	if err != nil {
		panic(err)
	}
	return int(nBig.Int64() + minNumber)
}

func RandomFloatNumber(minFloatNumber, maxFloatNumber float64) float64 {
	return RandomFloatNumberFrom0To1()*(maxFloatNumber-minFloatNumber) + minFloatNumber
}

func RandomFloatNumberFrom0To1() float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(floatNumberFrom0To1Base))
	if err != nil {
		panic(err)
	}
	return float64(nBig.Int64()) / floatNumberFrom0To1Base
}

func RandomError() *errors.Error {
	randomErrorCode := errors.ErrorCode(RandomDefaultNumberString())
	randomErrorText := RandomDefaultString()
	return errors.NewError(randomErrorCode, randomErrorText)
}
