// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

type dialects interface {
	DropTable(t *table) string
	CreateTable(t *table) string
	CreateIndex(t *table) string
}
