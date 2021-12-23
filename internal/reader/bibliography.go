package reader

import (
        "fmt"
)


type Quote struct {
        Page int
        Note string
}

func (q *Quote)qView() string {
        return fmt.Sprintf("[%d] %s", q.Page, q.Note)
}

type InternalQ struct {
        Index string
        Content Quote
}

func (iq *InternalQ)iqView() string {
        return fmt.Sprintf("%s\n%s", iq.Index, iq.Content.qView())
}

type ExternalQ struct {
        Author string
        PageEnd int
        Title string
        Publisher string
        Year int
        Content Quote
}

func (eq *ExternalQ)eqView() string {
        header := fmt.Sprintf("%s (%d)\n%s\n%s\n", eq.Title, eq.Year, eq.Author, eq.Publisher)
        body := fmt.Sprintf("[%d - %d]\n%s", eq.Content.Page, eq.PageEnd, eq.Content.Note)
        return fmt.Sprintf("%s%s", header, body)
}

type ReaderQuotes struct {
        Internal []InternalQ `json:"internal"`
        External []ExternalQ `json:"external"`
}

func (q *ReaderQuotes)AddInternal(p int, note,index string) *InternalQ {
        quote := Quote{p, note}
        iquote := InternalQ{index, quote}
        q.Internal = append(q.Internal, iquote)
        return &iquote
}

func (q *ReaderQuotes)AddExternal(p, ep, y int, note, author, title, publ string) *ExternalQ {
        nq := Quote{p, note}
        equote := ExternalQ{author, ep, title, publ, y, nq}
        q.External = append(q.External, equote)
        return &equote
}


