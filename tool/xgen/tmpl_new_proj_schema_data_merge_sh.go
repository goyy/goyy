// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjSchemaDataMergeSh = `#!/bin/sh

echo [INFO] Merged SQL Files.

export I18N_LOCALE=zh_CN
xtab -regexp=^insert.[\S]+.sql$ -newfile=init.sql -merge
`
