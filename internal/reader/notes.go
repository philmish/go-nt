package reader

import (
        "time"
        "fmt"
)

type Note struct {
        Id int
        Index string
        Content string
        Timestamp int64
}

func (n *Note)Read() {
        fmt.Printf("[%d]%s (%s)\n\n%s", n.Id, n.Index, time.Unix(n.Timestamp, 0), n.Content)
}

type Notes struct {
        Content []Note
}

func createId(n *Notes) int {
        switch len(n.Content){
                case 0:
                        return 1
                default:
                        return n.Content[len(n.Content)-1].Id + 1
                }
}

func (n *Notes)AddNote(index, content string) *Note {
        id := createId(n)
        nn := Note{id, index, content, time.Now().Unix()}
        n.Content = append(n.Content, nn)
        return &nn
}

func (n *Notes)Get_by_index(index string) *Notes {
        res := []Note{}
        for _, nt := range n.Content {
                if nt.Index == index {
                        res = append(res, nt)
                }
        }
        return &Notes{res}
}
