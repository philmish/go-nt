package reader_test

import (
        "github.com/philmish/go-nt/internal/reader"
        "testing"
)

type tLink struct {
        url string
        note string
        throws bool
}

type tRemove struct {
        id int
        throws bool
}

type tCase struct {
        tl []tLink
        tr []tRemove
}

func (c *tCase)run(t *testing.T) {
        nEl := reader.ExtLinks{[]reader.Link{}}
        for _, tL := range c.tl {
                err := nEl.Add(tL.url, tL.note)
                parseErr(tL.throws, err, t)
        }
        for _, r := range c.tr {
                err := nEl.Remove(r.id)
                parseErr(r.throws, err, t)
        }
}

func TestExtLinks(t *testing.T) {
        // Definition of test cases
        //
        // Test Case 1
        tc1L1 := tLink{"https://example.com/", "Testing", false}
        tc1L2 := tLink{"https://example.com/", "this throws", true}
        tc1L3 := tLink{"https://github.com/philmish/go-nt", "Repository", false}
        tc1L4 := tLink{"https://google.com", "still testing", false}
        tc1L := []tLink{tc1L1, tc1L2, tc1L3, tc1L4}

        tc1R1 := tRemove{-1, true}
        tc1R2 := tRemove{6, true}
        tc1R3 := tRemove{2, false}
        tc1R := []tRemove{tc1R1, tc1R2, tc1R3}

        tc1 := tCase{tc1L, tc1R}
        // ...
        tCases := []tCase{tc1}
        for _, test := range tCases {
                test.run(t)
        }
}
