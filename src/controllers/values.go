package controllers

import (
	"time"
)

const (
	// wait
	deletionThread  time.Duration = 10 // default is one minute, in production is one hour
	deletionCaptcha time.Duration = 1  // how much time you need to write a captcha this is stupid
)

var (

	// checking existence
	threads   = make(map[string][]Post)
	captchaDB = make(map[string]string)

	// if you are in the black list then you will not be able to access any kind of content
	// banned things
	BlackList   = make(map[string]bool) // add in the future
	BannedWords = make(map[string]bool) // this will be added in the future

	//postTemplate  = "posts.html"
)

type PostApi struct {
	Title     string `json:"title"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Content   string `json:"content"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captchaID"`
}
type Post struct {
	Title   string `json:"title"`
	Name    string `json:"name"`
	Date    string `json:"date"`
	Content string `json:"content"`
	IP      string // this will be used for the blacklisting
}
type CaptchaApi struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}
