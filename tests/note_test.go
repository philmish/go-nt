package reader_test

import (
        "github.com/philmish/go-nt/internal/reader"
        "testing"
        "log"
)

func checkErr(e error, t *testing.T) {
        if e != nil {
                t.Errorf("%v", e)
        }
}

func TestStructs(t *testing.T) {
        _ = reader.Note{123, "testing", "this is a test note", 0}
        log.Println("Note created successfully")
        newNotes := reader.Notes{[]reader.Note{}}
        log.Println("New Notes object created")
        newNotes.AddNote("testing", "This is the second test note")
        log.Println("Note added successfully")
        pulled := newNotes.Get_by_index("testing")
        if len(pulled.Content) < 1 {
                t.Errorf("Failed to get all notes with index testing excpected 1 result, got %d", len(pulled.Content))
        }
        log.Println("Pulled all related notes from notes object successfully")
}
