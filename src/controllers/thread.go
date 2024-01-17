package controllers

import (
	"fmt"
	"net/http"
)

var waitThreads = make(chan string)

// this will handle everything
func Thread(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ReplyThread(w, r)
	case "GET":
		GetThread(w, r)
	default:
		return
	}
}

// this is for loading the thread
func GetThread(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {

		ErrorT(w, "the id is empty", 400)
		return
	}
	v, e := threads[id]
	if !e {
		ErrorT(w, "thread not found", 400)
		return
	}
	templates.ExecuteTemplate(w, postT, v)
}
func ReplyThread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey look at me")
	id := r.URL.Query().Get("id")
	_, err := CheckAndAdd(w, r, id)
	fmt.Println(id, err)

	if err != nil {
		return
	}

}

// i start a new thread with this
func NewThread(w http.ResponseWriter, r *http.Request) {

	id, err := CheckAndAdd(w, r, "")
	if err != nil {
		return
	}
	// this will delete the thread after an hour
	fmt.Println(id)
	go AddToQuery(id, waitThreads)

	w.Write([]byte(id))

}

// i delete the thread after some time
func ThreadDeletion() {
	DeletionMap(threads, waitThreads, deletionThread)

}
