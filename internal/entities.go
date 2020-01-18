package internal

import "strings"

type FullName string

func (name FullName) Owner() string {
	return string(name[:strings.Index(string(name), "/")])
}

func (name FullName) Name() string {
	return string(name[strings.Index(string(name), "/")+1:])
}

type Repository struct {
	ID       int64
	FullName FullName
	Labels   []Label
}

type Label struct {
	ID          int64
	Name        string
	Color       string
	Description string
}
