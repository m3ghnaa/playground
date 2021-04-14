package main

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("Hello, a sample space")}
	p1.save()
}
