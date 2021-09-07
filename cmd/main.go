package main

import (
	"log"
	"os"

	blogposts "github.com/jacksonopp/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFs(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
