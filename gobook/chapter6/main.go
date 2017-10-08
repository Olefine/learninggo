package main

import (
	"fmt"
)

type Post struct {
	Title string
	Body  string
}

// note receiver is not link
func (p Post) String() string {
	return fmt.Sprintf("%s - %s", p.Title, p.Body)
}

func (p *Post) UpdateTitle(newTitle string) {
	p.Title = newTitle
}

func main() {
	post := &Post{Title: "Hello world", Body: "Hello, my name is Egor, I'm 24 years old."}

	fmt.Println(post)

	post.UpdateTitle("Egor")

	fmt.Println(post)
}
