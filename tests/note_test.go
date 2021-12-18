package reader_test

import (
	"github.com/philmish/go-nt/internal/reader"
	"log"
	"testing"
)

func checkErr(e error, t *testing.T) {
	if e != nil {
		t.Errorf("%v", e)
	}
}

func parseErr(th bool, e error, t *testing.T) {
	if th && e != nil {
		log.Println("Error Test success, error thrown.")
	} else if !th && e == nil {
		log.Println("Error Test success, no error detected")
	} else {
		t.Error("Error Test failed.")
	}
}

type tNVals struct {
	index   string
	content string
	idrm    int
	throws  bool
}

func TestNotes(t *testing.T) {
	th1 := tNVals{"testing", "This is my test note", 3, true}
	th2 := tNVals{"testing", "Still testing", -1, true}
	s := tNVals{"testing", "My Note", 1, false}
	v := []tNVals{th1, th2, s}
	for _, i := range v {
		nN := reader.Notes{[]reader.Note{}}
		_ = nN.Add(i.index, i.content)
		err := nN.Remove(i.idrm)
		parseErr(i.throws, err, t)
	}
}
