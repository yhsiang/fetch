package request

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func getAttr(tkn *html.Tokenizer, key string) string {
	for {
		attrKey, attrValue, moreAttr := tkn.TagAttr()
		if string(attrKey) == key {
			return string(attrValue)
		}

		if !moreAttr {
			return ""
		}
	}
}

func GetMetadata(content string) ([]string, []string, []string, []string, error) {
	tkn := html.NewTokenizer(strings.NewReader(content))

	links := []string{}
	images := []string{}
	scripts := []string{}
	stylesheets := []string{}

	for {

		tt := tkn.Next()
		if tt == html.ErrorToken {
			err := tkn.Err()
			if err != io.EOF {
				return nil, nil, nil, nil, err
			}
			return links, images, scripts, stylesheets, nil
		}

		tag, _ := tkn.TagName()

		// handle <a>
		if tt == html.StartTagToken && string(tag) == "a" {
			links = append(links, getAttr(tkn, "href"))
		}

		// handle <script>
		if tt == html.StartTagToken && string(tag) == "script" {
			scripts = append(scripts, getAttr(tkn, "src"))
		}

		// handle <img>
		if tt == html.StartTagToken && string(tag) == "img" {
			images = append(images, getAttr(tkn, "src"))
		}

		if tt == html.SelfClosingTagToken && string(tag) == "img" {
			images = append(images, getAttr(tkn, "src"))
		}

		// handle <link>
		if tt == html.StartTagToken && string(tag) == "link" {
			stylesheets = append(stylesheets, getAttr(tkn, "href"))
		}

		if tt == html.SelfClosingTagToken && string(tag) == "link" {
			stylesheets = append(stylesheets, getAttr(tkn, "href"))
		}
	}
}
