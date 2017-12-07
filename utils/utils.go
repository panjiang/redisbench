package utils

import (
	"log"
	"math/rand"
	"time"
)

// FatalErr : If the error is not nil fatal it
func FatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandSeq : Create a string sequence with random chars
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// NowTs : Return timestamp now (second level)
func NowTs() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

// NowMilliTs : Return timestamp now (millisecond level)
func NowMilliTs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
