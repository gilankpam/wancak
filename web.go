package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	rndr "github.com/unrolled/render"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"log"
	"os"
)

const (
	Url1cak string = "http://1cak.com"
)

var (
	httpClient *http.Client
	render     *rndr.Render
)

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Img   string `json:"img"`
	Votes string `json:"votes"`
}

type Posts struct {
	Page struct {
		Next string `json:"next"`
	} `json:"page"`
	Posts []Post `json:"posts"`
}

func main() {
	httpClient = &http.Client{}
	render = rndr.New(rndr.Options{})
	api := mux.NewRouter()
	api.StrictSlash(true)

	defaultHandler := handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(ApiHandlerIndex))

	api.Handle("/{section:lol|trend|recent|legendary}", defaultHandler).Methods("GET")
	api.Handle("/{section:lol|trend|recent|legendary}/{id:[0-9]+}", defaultHandler).Methods("GET")

	log.Printf("Running on port %s", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), api)
}

func ApiHandlerIndex(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	section := vars["section"]
	if id, ok := vars["id"]; ok {
		url = fmt.Sprintf("%s/%s-%s", Url1cak, section, id)
	} else {
		url = fmt.Sprintf("%s/%s", Url1cak, section)
	}
	posts, err := GetPosts(url)
	if err != nil {
		render.JSON(w, http.StatusInternalServerError, map[string]string{"msg": "error"})
	} else {
		render.JSON(w, http.StatusOK, posts)
	}
}

func GetPosts(url string) (*Posts, error) {
	var posts []Post
	// fmt.Println("get: " + url)
	res, err := makeRequest(url)
	if err != nil {
		return nil, err
	}
	doc, _ := goquery.NewDocumentFromResponse(res)
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if id, _ := s.Find(".upperSpan").Attr("rel"); id != "" {
			title := s.Find("h3").Text()
			//If relative link, add Url1Cak
			img, _ := s.Find("img").Attr("src")
			if !strings.HasPrefix(img, "http://") {
				img = Url1cak + img
			}
			votes := s.Find("#span_vote_" + id).Text()
			url := fmt.Sprintf("%s/%s", Url1cak, id)
			post := Post{
				Id:    id,
				Title: title,
				Url:   url,
				Img:   img,
				Votes: votes,
			}
			posts = append(posts, post)
		}
	})
	next, _ := doc.Find("#next_page_link").Attr("href")
	nextSplit := strings.Split(next, "-")

	if len(nextSplit) < 2 {
		return nil, fmt.Errorf("index out of range")
	}
	nextId := nextSplit[1]
	p := &Posts{
		Page: struct {
			Next string `json:"next"`
		}{Next: nextId},
		Posts: posts,
	}
	return p, nil
}

func makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
