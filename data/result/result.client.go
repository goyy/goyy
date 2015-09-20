// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	Body       []byte
	Status     string
	StatusCode int
	Header     http.Header
	Cookies    []*http.Cookie
}

func (me *Client) ParseEntity(out *Entity) error {
	return out.ParseJSON(string(me.Body))
}

func (me *Client) ParseEntities(out *Entities) error {
	return out.ParseJSON(string(me.Body))
}

func (me *Client) ParsePage(out *Page) error {
	return out.ParseJSON(string(me.Body))
}

func (me *Client) ParseResult(out *Result) error {
	dec := json.NewDecoder(strings.NewReader(string(me.Body)))
	if err := dec.Decode(out); err == io.EOF {
		return nil
	} else if err != nil {
		log.Fatal(err)
	}
	return nil
}
