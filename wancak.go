package wancak

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

const (
	Url1cak string = "http://1cak.com"
)

var (
	NotFoundErr       error = errors.New("Post not found")
	InvalidSectionErr error = errors.New("Invalid section")
)

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Img   string `json:"img"`
	Votes string `json:"votes"`
	NSFW  bool   `json:"nsfw"`
}

type Posts struct {
	Page struct {
		Next string `json:"next"`
	} `json:"page"`
	Posts []Post `json:"posts"`
}

func GetPosts(section, pageId string) (*Posts, error) {
	var posts []Post
	var url string

	if !isValidSection(section) {
		return nil, InvalidSectionErr
	}

	if pageId == "" {
		url = Url1cak + "/" + section
	} else {
		url = fmt.Sprintf("%s/%s-%s", Url1cak, section, pageId)
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Printf("Error getting web pages: %v", err)
		return nil, err
	}
	if isNotFound(doc) {
		return nil, NotFoundErr
	}
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if id, _ := s.Find(".upperSpan").Attr("rel"); id != "" {
			title := s.Find("h3").Text()
			//If relative link, add Url1Cak
			img, _ := s.Find("img").Attr("src")
			//NSFW post
			nsfw := false
			if !strings.HasPrefix(img, "http://") {
				img = Url1cak + img
				nsfw = true
			}
			votes := s.Find("#span_vote_" + id).Text()
			url := fmt.Sprintf("%s/%s", Url1cak, id)
			post := Post{
				Id:    id,
				Title: title,
				Url:   url,
				Img:   img,
				Votes: votes,
				NSFW:  nsfw,
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

//Get sigle post by id, if id is empty then get suffle post :p
func GetPostId(id string) (*Post, error) {
	var url string
	post := new(Post)
	if id == "" {
		url = Url1cak + "/shuffle"
	} else {
		url = Url1cak + "/" + id
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	if isNotFound(doc) {
		return nil, NotFoundErr
	}
	url, _ = doc.Find(".fb-comments").Attr("data-href")
	post.Id = strings.Split(url, "/")[3]
	post.Title = doc.Find("h3").Text()
	post.NSFW = false
	post.Img, _ = doc.Find("img[title]").Attr("src")
	if !strings.HasPrefix(post.Img, "http://") {
		post.Img = Url1cak + "/images/unsave.jpg"
		post.NSFW = true
	}
	post.Url = url
	post.Votes = doc.Find("#span_vote_" + post.Id).Text()
	return post, nil
}

func isNotFound(doc *goquery.Document) bool {
	return doc.Has(`img[src="templates/v1/img/error.png"]`).Size() > 0
}

func isValidSection(section string) bool {
	sections := []string{"lol", "trend", "recent", "legendary"}
	for _, s := range sections {
		if section == s {
			return true
		}
	}
	return false
}
