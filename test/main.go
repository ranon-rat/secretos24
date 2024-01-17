package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/mojocn/base64Captcha"
)

func init() {
	//init rand seed
	rand.Seed(time.Now().UnixNano())

}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		driver := base64Captcha.NewDriverDigit(100, 240, 4, 0.7, 80)
		captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
		id, b64s, answer, _ := captcha.Generate()
		fmt.Println(id, answer)
		w.Write([]byte(`
		<!DOCTYPE html>
	<html lang="en">
	<body>	
	<img src="` + b64s + `">
		</body></html>`))

	})
	fmt.Println("en 8080")
	http.ListenAndServe(":8080", nil)
}
