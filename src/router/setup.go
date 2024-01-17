package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ranon-rat/secretos24/src/controllers"
)

func Setup() error {

	fs := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", controllers.Index)
	http.Handle("/static/", http.StripPrefix("/static", fs)) // basic files
	http.HandleFunc("/thread", controllers.Thread)           // you can post or get
	// api stuff
	http.HandleFunc("/new-thread", controllers.NewThread) // this is just for adding a new thread
	http.HandleFunc("/get-captcha", controllers.GetCaptcha)
	// supporting functions
	controllers.SetupTemplate()
	go controllers.ThreadDeletion()
	go controllers.CaptchaDeletion()
	//basic stuff
	port, e := os.LookupEnv("PORT")
	if !e {
		port = "8080"
	}
	fmt.Println("starting server at:", port)
	return http.ListenAndServe(":"+port, nil)
}
