package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/jacksonopp/blogposts"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i walways fail")
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("it should get the title", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: Tag 1
---
hello
world`
			secondBody = `Title: Post 2
Description: Description 2
Tags: Tag 2
---
content 2 content 2
content 2 content 2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFs(fs)

		if err != nil {
			t.Fatal(err)
		}

		// rest of test code cut for brevity
		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        "Tag 1",
			Body: `hello
world`,
		})
	})

	t.Run("With a failing file system", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFs(StubFailingFS{})

		if err == nil {
			t.Errorf("Expected an error, but didn't get one")
		}
	})
}
