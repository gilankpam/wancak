package main

import (
	"flag"
	"github.com/gilankpam/wancak"
	rndr "github.com/unrolled/render"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"os"
	"regexp"
)

var render *rndr.Render

func main() {
	render = rndr.New(rndr.Options{})

	goji.Get(regexp.MustCompile(`^/(?P<section>lol|trend|recent|legendary)$`), sectionHandler)
	goji.Get(regexp.MustCompile(`^/(?P<section>lol|trend|recent|legendary)/(?P<id>[0-9]+)$`), sectionHandler)
	goji.Get(regexp.MustCompile(`^/post/(?P<id>[0-9]+)$`), singlePostHandler)
	goji.Get("/random", randomPostHandler)
	goji.Get("/search", searchHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	flag.Set("bind", ":"+port)
	log.Printf("Running on port %s", port)
	goji.Serve()
}

func sectionHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	var posts *wancak.Posts
	var err error

	section := c.URLParams["section"]
	if id := c.URLParams["id"]; id == "" {
		posts, err = wancak.GetSectionPosts(section)
	} else {
		posts, err = wancak.GetSectionPosts(section, id)
	}

	if err == wancak.NotFoundErr {
		render.JSON(w, http.StatusNotFound, map[string]string{"status": "error", "message": "Posts not found"})
		return
	}
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": "Can't get posts"})
		return
	}
	render.JSON(w, http.StatusOK, posts)
}

func randomPostHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	post, err := wancak.GetPostId("")
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": "Can't get posts"})
		return
	}
	render.JSON(w, http.StatusOK, post)
}

func singlePostHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	postId := c.URLParams["id"]
	post, err := wancak.GetPostId(postId)
	if err == wancak.NotFoundErr {
		render.JSON(w, http.StatusNotFound, map[string]string{"status": "error", "message": "Post not found"})
		return
	}
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": "Can't get post"})
		return
	}
	render.JSON(w, http.StatusOK, post)
}

func searchHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	var posts *wancak.Posts
	var err error

	q := r.URL.Query().Get("q")
	if pageId := r.URL.Query().Get("pageid"); pageId == "" {
		posts, err = wancak.Search(q)
	} else {
		posts, err = wancak.Search(q, pageId)
	}

	if err == wancak.NotFoundErr {
		render.JSON(w, http.StatusNotFound, map[string]string{"status": "error", "message": "Posts not found"})
		return
	}
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": "Can't get posts"})
		return
	}

	render.JSON(w, http.StatusOK, posts)
}
