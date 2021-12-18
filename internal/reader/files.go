package reader

import (
	"errors"
	"fmt"
)

type File struct {
	Name    string
	Comment string
	Added   int64
}

type ReaderIndex struct {
	Name    string
	Content []File
}

func (r *ReaderIndex) fileNames() []string {
	data := []string{}
	for _, file := range r.Content {
		data = append(data, file.Name)
	}
	return data
}

type FileStore struct {
	Root    string
	Readers []ReaderIndex
}

func (fs *FileStore) readerNames() []string {
	data := []string{}
	for _, reader := range fs.Readers {
		data = append(data, reader.Name)
	}
	return data
}

func (fs *FileStore) readerByName(name string) (ReaderIndex, error) {
	for _, reader := range fs.Readers {
		if reader.Name == name {
			return reader, nil
		}
	}
	return ReaderIndex{}, errors.New(fmt.Sprintf("No reader with name %s found.", name))
}

func (fs *FileStore) getReaderPath(name string) (string, error) {
	reader, err := fs.readerByName(name)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%s/%s/", fs.Root, reader.Name)
	return result, nil
}
