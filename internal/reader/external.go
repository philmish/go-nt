package reader

import (
        "errors"
        "time"
)

type Link struct {
        Url string
        Note string
        Timestamp int64
}

type ExtLinks struct {
        Links []Link
}

func (el *ExtLinks)urls() []string {
        urls := []string{}
        for _, link := range el.Links {
                urls = append(urls, link.Url)
        }
        return urls
}

func (el *ExtLinks)urlUnique(url string) bool {
        urls := el.urls()
        for _, u := range urls {
                if url == u {
                        return false
                }
        }
        return true
}

func (el *ExtLinks)Add(url, note string) error {
        if ! el.urlUnique(url) {
                return errors.New("There already exists a Link with that URL")
        }
        nLink := Link{url, note, time.Now().Unix()}
        el.Links = append(el.Links, nLink)
        return nil
}

func (el *ExtLinks)Remove(id int) error {
        if (id < 0 || id > len(el.Links)) {
                return errors.New("Link ID out of bounce")
        }
        el.Links = append(el.Links[:id-1], el.Links[:id]...)
        return nil
}
