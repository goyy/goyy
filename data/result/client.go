// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Client result.Client.
type Client struct {
	Body       []byte
	Status     string
	StatusCode int
	Header     http.Header
	Cookies    []*http.Cookie
}

// ParseEntity result.Client.Body to result.Entity.
func (me *Client) ParseEntity(out *Entity) error {
	return out.ParseJSON(string(me.Body))
}

// ParseEntities result.Client.Body to result.Entities.
func (me *Client) ParseEntities(out *Entities) error {
	return out.ParseJSON(string(me.Body))
}

// ParsePage result.Client.Body to result.Page.
func (me *Client) ParsePage(out *Page) error {
	return out.ParseJSON(string(me.Body))
}

// ParseResult result.Client.Body to result.Result.
func (me *Client) ParseResult(out *Result) error {
	body := string(me.Body)
	dec := json.NewDecoder(strings.NewReader(body))
	if err := dec.Decode(out); err == io.EOF {
		return nil
	} else if err != nil {
		logger.Error(err.Error())
		logger.Println(body)
		return err
	}
	return nil
}
