// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dialect

var Domains domains = domains{}

type domains struct {
}

func (me *domains) String(length int) Domain {
	return Domain{Type: DTString, Length: length}
}

func (me *domains) Bool() Domain {
	return Domain{Type: DTBool}
}

func (me *domains) Int(length int) Domain {
	return Domain{Type: DTInt, Length: length}
}

func (me *domains) Float(length, precision int) Domain {
	return Domain{Type: DTFloat, Length: length, Precision: precision}
}

func (me *domains) Time() Domain {
	return Domain{Type: DTTime}
}
