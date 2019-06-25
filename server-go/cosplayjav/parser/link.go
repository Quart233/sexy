package parser

import (
    "github.com/PuerkitoBio/goquery"
    "sexy/engine"
    "sync"
)

func ParseLink(doc *goquery.Document, guard *sync.Mutex, post *Post) engine.ParseResult {
    var link, _ = doc.Find(`.btn.btn-primary.btn-download`).Attr("href")
    guard.Lock()
    post.Parts = append(post.Parts, link)
    guard.Unlock()

    return engine.ParseResult{}
}