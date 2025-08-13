package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/madhusudhan-reddy-oneture/gotbd/blogrenderer"
)

var (
	aPost = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a Post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	bPost = blogrenderer.Post{
		Title:       "concurrency in rust",
		Body:        "Let's dive deep in concurrency in Rust!",
		Description: "This article is easy to understand for begineer rust developers too",
		Tags:        []string{"rust", "concurrency"},
	}
)

func TestRender(t *testing.T) {

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

		// got := buf.String()
		// want := "<h1>hello world</h1>\n<p>This is a description</p>\nTags: <ul><li>go</li><li>tdd</li></ul>"

		// if got != want {
		// 	t.Errorf("got '%s' want '%s'", got, want)
		// }
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{aPost, bPost}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">hello world</a></li><li><a href="/post/concurrency-in-rust">concurrency in rust</a></li></ol>`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
