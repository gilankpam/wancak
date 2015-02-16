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
	goji.Get(regexp.MustCompile(`^/post/(?P<id>[0-9]+)$`), getSinglePost)
	goji.Get(regexp.MustCompile(`^/random$`), getRandomPost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	flag.Set("bind", ":"+port)
	log.Printf("Running on port %s", port)
	goji.Serve()
}

func sectionHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	section := c.URLParams["section"]
	id := c.URLParams["id"]

	posts, err := wancak.GetPosts(section, id)
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

func getRandomPost(c web.C, w http.ResponseWriter, r *http.Request) {
	post, err := wancak.GetPostId("")
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"status": "error", "message": "Can't get posts"})
		return
	}
	render.JSON(w, http.StatusOK, post)
}

func getSinglePost(c web.C, w http.ResponseWriter, r *http.Request) {
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
