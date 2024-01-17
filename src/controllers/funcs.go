package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func GenHash(content string) string {
	msgHash := sha256.New()
	msgHash.Write([]byte(content))
	return hex.EncodeToString(msgHash.Sum(nil))
}

func GenHashID(content, title, name string) string {
	return GenHash(content +
		name +
		title +
		strconv.Itoa(int(time.Now().UnixNano())) +
		strconv.Itoa(rand.Intn(77)))
}

// this will wait a certain time and then delete the element in the captcha
func DeletionMap[K any](m map[string]K, c chan string, d time.Duration) {
	for {
		id := <-c

		go func() {
			time.Sleep(d * time.Minute)
			delete(m, id)

		}()

	}
}
func AddToQuery(id string, c chan string) {
	c <- id
}
