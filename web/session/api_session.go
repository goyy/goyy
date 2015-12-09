// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package session

// Options stores configuration for a session or session store.
//
// Fields are a subset of http.Cookie fields.
type Options struct {
	Path   string
	Domain string
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

// Interface stores the values and optional configuration for a session.
type Interface interface {
	// Get returns the session value associated to the given key.
	Get(key string) (string, error)
	// Set sets the session value associated to the given key.
	Set(key string, val string) error
	// Delete removes the session value associated to the given key.
	Delete(key string) error
	// Clear deletes all values in the session.
	Clear() error
	// AddFlash adds a flash message to the session.
	// A single variadic argument is accepted, and it is optional: it defines the flash key.
	// If not defined "_flash" is used by default.
	AddFlash(value string, vars ...string)
	// Flashes returns a slice of flash messages from the session.
	// A single variadic argument is accepted, and it is optional: it defines the flash key.
	// If not defined "_flash" is used by default.
	Flashes(vars ...string) []string
	// Options sets confuguration for a session.
	Options(*Options)
	// To determine whether it has been logged in
	IsLogin() bool
	// Get principal from session
	Principal() (Principal, error)
	// Set principal to session
	SetPrincipal(value Principal) error
	// Set principal to session
	ResetPrincipal() error
}
