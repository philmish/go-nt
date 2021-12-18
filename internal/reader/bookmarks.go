package reader

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Bookmark struct {
	Page      int
	Index     string
	Timestamp int64
	Comment   string
}

type Bookmarks struct {
	Data []Bookmark `json:"data"`
}

func (b *Bookmarks) sortBookmarks() {
	sort.Slice(b.Data, func(i, j int) bool { return b.Data[i].Page < b.Data[j].Page })
}

func (b *Bookmarks) Add(p int, i string, c string) *Bookmark {
	nB := Bookmark{p, i, time.Now().Unix(), c}
	b.Data = append(b.Data, nB)
	return &nB
}

func (b *Bookmarks) ByIndex(index string) *Bookmarks {
	res := []Bookmark{}
	for _, bm := range b.Data {
		if bm.Index == index {
			res = append(res, bm)
		}
	}
	return &Bookmarks{res}
}

func (b *Bookmarks) Remove(id int) error {
	b.sortBookmarks()
	if id < 0 || id > len(b.Data) {
		return errors.New(fmt.Sprintf("%d is out of bounce", id))
	}
	b.Data = append(b.Data[:id-1], b.Data[:id]...)
	return nil
}
