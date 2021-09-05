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

func TestNewBlogPosts(t *testing.T) {
	t.Run("it should get the title", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := blogposts.NewPostsFromFs(fs)

		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("With a failing file system", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFs(StubFailingFS{})

		if err == nil {
			t.Errorf("Expected an error, but didn't get one")
		}
	})
}
