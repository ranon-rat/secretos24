package controllers

import (
	"time"
)

// i check that nothing is empty
func CheckParameters(thread, title, name, content string) bool {
	return thread == "" || title == "" || name == "" || content == ""

}

// i reduce the size of the content
func ReduceSize(content string, limit int) string {
	if len(content) > limit {
		return content[limit:]
	}
	return content
}

// i add a new post
func AddPost(thread, title, name, content string) {
	v, e := threads[thread]
	if !e {
		v = []Post{}
	}

	title = ReduceSize(title, 50)
	name = ReduceSize(name, 50)
	thread = ReduceSize(thread, 5000)
	threads[thread] = append(v, Post{
		Title:   title,
		Name:    name,
		Content: content,
		Date:    time.Now().String(),
	})

}
