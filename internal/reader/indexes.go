package reader

import (
        "errors"
        "fmt"
        "sort"
)

type Index struct {
        Start int
        End int
        Name string
}

func (i *Index)inPageRange(n int) bool {
        return (i.Start <= n && i.End >= n)
}

type Indices struct {
        Data []Index
}

func  (i *Indices)nameUnique(name string) bool {
        for _, index := range i.Data {
                if index.Name == name {
                        return false
                }
        }
        return true
}

func (in *Indices)sortIndices(){
        sort.Slice(in.Data, func(i, j int) bool {return in.Data[i].Start < in.Data[j].Start})
}

func (in *Indices)View() {
        in.sortIndices()
        for i, index := range in.Data {
                fmt.Printf("[%d] %s (%d - %d)", i + 1, index.Name, index.Start, index.End)
        }
}

func (i *Indices)AddIndex(name string, start, end int) (*Index, error) {
        for _, index := range i.Data {
                if ! (start > index.End || end < index.Start) {
                        return nil, errors.New(fmt.Sprintf("The provided Index would overlap with %s", index.Name))
                }
        }
        if ! i.nameUnique(name) {
                return nil, errors.New(fmt.Sprintf("%s is already an index name", name))
        }
        nIndex := Index{start, end, name}
        i.Data = append(i.Data, nIndex)
        return &nIndex, nil
}

func (i *Indices)Remove(id int) (error) {
        if (id < 0 || id > len(i.Data)) {
                return errors.New(fmt.Sprintf("%d is out of bounce.", id))
        }
        i.sortIndices()
        i.Data = append(i.Data[:id-1], i.Data[:id]...)
        return nil
}
