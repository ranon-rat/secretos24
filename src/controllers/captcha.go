package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

var waitCaptcha = make(chan string)

func VerifyCaptcha(id, answer string) bool {
	a, e := captchaDB[id]
	if !e {
		return false
	}
	delete(captchaDB, id)
	return a == answer
}
func GetCaptcha(w http.ResponseWriter, r *http.Request) {
	// genero el captcha
	driver := base64Captcha.NewDriverDigit(100, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, answer, _ := captcha.Generate()
	// guardo el captcha

	captchaDB[id] = answer
	// lo agrego al query
	go AddToQuery(id, waitCaptcha)
	// lo elimino
	json.NewEncoder(w).Encode(CaptchaApi{Image: b64s, ID: id})

}
func CaptchaDeletion() {
	DeletionMap(captchaDB, waitCaptcha, deletionCaptcha)

}
