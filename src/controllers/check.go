package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// this will add a new post or start a new thread depending if you add the id
func CheckAndAdd(w http.ResponseWriter, r *http.Request, id string) (string, error) {
	var post PostApi
	json.NewDecoder(r.Body).Decode(&post)
	// I get the info
	content := strings.TrimSpace(post.Content)
	name := strings.TrimSpace(post.Name)
	title := strings.TrimSpace(post.Title)

	captchaID := post.CaptchaID
	captchaAnswer := post.Captcha
	// i need to check if the id actually exists, and this means that is a post to /thread
	if id != "" {
		_, e := threads[id]
		if !e {
			return "", errors.New("404 thread not found")
		}
	}
	// if the id is empty then it means that is the /new-thread
	if id == "" {
		id = GenHashID(content, title, name)
	}

	// i just verify that the captcha is working correctly
	if !VerifyCaptcha(captchaID, captchaAnswer) {
		ErrorT(w, "bad captcha", 400)

		return "", errors.New("bad captcha")
	}
	// I check that everything is okay
	if CheckParameters(id, title, name, content) {
		ErrorT(w, "not everything is fulfilled", 400)
		return "", errors.New("invalid values")
	}
	AddPost(id, title, name, content) // esto es para poder mostrarlo
	return id, nil
}
