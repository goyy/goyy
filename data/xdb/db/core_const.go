// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

// Operator
const (
	o_eq string = "Eq" // equal( = )
	o_ne string = "Ne" // not equal( <> )
	o_gt string = "Gt" // greater than( > )
	o_lt string = "Lt" // less than( < )
	o_ge string = "Ge" // greater than or equal( >= )
	o_le string = "Le" // less than or equal( <= )
	o_li string = "Li" // like ?
	o_lk string = "Lk" // like %?%
	o_ll string = "Ll" // like %?
	o_lr string = "Lr" // like ?%
	o_be string = "Be" // between
	o_in string = "In" // in
	o_nu string = "Nu" // is null
	o_nn string = "Nn" // is not null
)
