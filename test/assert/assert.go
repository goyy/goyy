// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package assert

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

// httpResp is a helper that returns HTTP of the response.
func httpResp(handle xhttp.Handle, method, url string, values url.Values) (*httptest.ResponseRecorder, error) {
	w := httptest.NewRecorder()
	r, err := xhttp.NewRequest(method, url, values)
	if err != nil {
		return nil, err
	}
	c := xhttp.NewContext(w, r)
	handle(c)
	return w, nil
}

// httpCode is a helper that returns HTTP code of the response. It returns -1
// if building a new request fails.
func httpCode(handle xhttp.Handle, method, url string, values url.Values) int {
	w, err := httpResp(handle, method, url, values)
	if err != nil {
		return -1
	}
	return w.Code
}

// HTTPSuccess asserts that a specified handle returns a success status code.
//
//  assert.HTTPSuccess(t, handle, "POST", "/a/b/c", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccess(t *testing.T, handle xhttp.Handle, method, url string, values url.Values) bool {
	code := httpCode(handle, method, url, values)
	if code == -1 {
		return false
	}
	return code >= http.StatusOK && code <= http.StatusPartialContent
}

// HTTPRedirect asserts that a specified handle returns a redirect status code.
//
//  assert.HTTPRedirect(t, handle, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirect(t *testing.T, handle xhttp.Handle, method, url string, values url.Values) bool {
	code := httpCode(handle, method, url, values)
	if code == -1 {
		return false
	}
	return code >= http.StatusMultipleChoices && code <= http.StatusTemporaryRedirect
}

// HTTPError asserts that a specified handle returns an error status code.
//
//  assert.HTTPError(t, handle, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPError(t *testing.T, handle xhttp.Handle, method, url string, values url.Values) bool {
	code := httpCode(handle, method, url, values)
	if code == -1 {
		return false
	}
	return code >= http.StatusBadRequest
}

// HTTPBody is a helper that returns HTTP body of the response. It returns
// empty string if building a new request fails.
func HTTPBody(handle xhttp.Handle, method, url string, values url.Values) string {
	w, err := httpResp(handle, method, url, values)
	if err != nil {
		return ""
	}
	return w.Body.String()
}

// HTTPBodyContains asserts that a specified handle returns a
// body that contains a string.
//
//  assert.HTTPBodyContains(t, handle, "/a/b/c", nil, "ok")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContains(t *testing.T, handle xhttp.Handle, method, url string, values url.Values, str interface{}) bool {
	body := HTTPBody(handle, method, url, values)
	return strings.Contains(body, fmt.Sprint(str))
}

// HTTPBodyNotContains asserts that a specified handle returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContains(t, handle, "/a/b/c", nil, "no")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContains(t *testing.T, handle xhttp.Handle, method, url string, values url.Values, str interface{}) bool {
	body := HTTPBody(handle, method, url, values)
	contains := strings.Contains(body, fmt.Sprint(str))
	return !contains
}
