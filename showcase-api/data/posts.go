package data

import (
	"encoding/json"
	"io"
)

type Post struct {
	ID      int    `json:"id"`
	title   string `json:"title"`
	snippet string `json:"snippet"`
	deleted bool   `json:deleted"`
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
		title:   "The first blog",
		snippet: "this is my first blog hope you enjoy it blah blah blah",
		deleted: false,
	},
	{
		ID:      2,
		title:   "The second blog",
		snippet: "this is my second blog hope you enjoy it blah blah blah",
		deleted: false,
	},
}
