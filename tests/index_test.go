package reader_test

import (
        "github.com/philmish/go-nt/internal/reader"
        "testing"
        "log"
)



func TestIndices(t *testing.T) {
        nIndices := reader.Indices{[]reader.Index{}}
        log.Println("Indices created successfully")
        _, err := nIndices.AddIndex("testing", 0, 50)
        if err != nil {
                t.Errorf("Failed to add index\n%v", err)
        }
        err = nIndices.Remove(1)
        if err != nil {
                t.Errorf("Failed to remove index\n%v", err)
        }
}
