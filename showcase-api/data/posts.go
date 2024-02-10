package data

import (
	"encoding/json"
	"io"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	Deleted bool   `json:"deleted"`
}

type Posts []*Post

func GetBlogs() Posts {
	return blogDummyData
}

func (p *Posts) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Posts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var blogDummyData = Posts{
	{
		ID:      1,
		Title:   "The first blog",
		Snippet: "this is my first blog hope you enjoy it blah blah blah",
		Deleted: false,
	},
	{
		ID:      2,
		Title:   "The second blog",
		Snippet: "this is my second blog hope you enjoy it blah blah blah",
		Deleted: false,
	},
}
