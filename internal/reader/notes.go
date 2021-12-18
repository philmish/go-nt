package reader

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Note struct {
	Index     string `json:"index"`
	Content   string `json:"content"`
	Timestamp int64  `json:"created"`
}

func (n *Note) View() {
	fmt.Printf("[%s] %s\n\n%s", n.Index, time.Unix(n.Timestamp, 0), n.Content)
}

type Notes struct {
	Data []Note
}

func (n *Notes) sortNotes() {
	sort.Slice(n.Data, func(i, j int) bool { return n.Data[i].Timestamp > n.Data[j].Timestamp })
}

func (n *Notes) View() {
	n.sortNotes()
	for i, note := range n.Data {
		seperator := strings.Repeat("#", 10)
		fmt.Printf("[%d] %s(%s)\n%s\n%s\n", i+1, note.Index, time.Unix(note.Timestamp, 0), note.Content, seperator)
	}
}

func (n *Notes) Add(index, content string) *Note {
	nn := Note{index, content, time.Now().Unix()}
	n.Data = append(n.Data, nn)
	return &nn
}

func (n *Notes) ByIndex(index string) *Notes {
	res := []Note{}
	for _, nt := range n.Data {
		if nt.Index == index {
			res = append(res, nt)
		}
	}
	return &Notes{res}
}

func (n *Notes) Remove(id int) error {
	n.sortNotes()
	if id < 0 || id > len(n.Data) {
		return errors.New(fmt.Sprintf("%d is out of bounce", id))
	}
	n.Data = append(n.Data[:id-1], n.Data[:id]...)
	return nil
}
