// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

// Validation validate.Validation.
type Validation struct {
	Errors []*Error
}

// Clear clean all Validation Errors.
func (me *Validation) Clear() {
	me.Errors = nil
}

// HasErrors has Validation Errors nor not.
func (me *Validation) HasErrors() bool {
	return me.Errors != nil && len(me.Errors) > 0
}

// Set Validation Errors.
func (me *Validation) err(field, typ string, err error) *Error {
	if err != nil {
		e := &Error{field: field, typ: typ, message: err.Error()}
		if me.Errors == nil {
			me.Errors = make([]*Error, 0)
		}
		me.Errors = append(me.Errors, e)
		return e
	}
	return nil
}

// Required returns error if the provided input is empty, nil otherwise.
func (me *Validation) Required(field, input string) *Error {
	return me.err(field, typRequired, Required(input))
}

// Min returns error if the provided input is greater than or equal to {min},
// nil otherwise.
func (me *Validation) Min(field, input string, min int) *Error {
	return me.err(field, typMin, Min(input, min))
}

// Max returns error if the provided input is less than or equal to {max},
// nil otherwise.
func (me *Validation) Max(field, input string, max int) *Error {
	return me.err(field, typMax, Max(input, max))
}

// Range returns error if the provided input is between {min} and {max},
// nil otherwise.
func (me *Validation) Range(field, input string, min, max int) *Error {
	return me.err(field, typRange, Range(input, min, max))
}

// Minf returns error if the provided input is greater than or equal to {min},
// nil otherwise.
func (me *Validation) Minf(field, input string, min float64) *Error {
	return me.err(field, typMinf, Minf(input, min))
}

// Maxf returns error if the provided input is less than or equal to {max},
// nil otherwise.
func (me *Validation) Maxf(field, input string, max float64) *Error {
	return me.err(field, typMaxf, Maxf(input, max))
}

// Rangef returns error if the provided input is between {min} and {max},
// nil otherwise.
func (me *Validation) Rangef(field, input string, min, max float64) *Error {
	return me.err(field, typRangef, Rangef(input, min, max))
}

// Minlen returns error if the provided input is least {min} characters,
// nil otherwise.
func (me *Validation) Minlen(field, input string, min int) *Error {
	return me.err(field, typMinlen, Minlen(input, min))
}

// Maxlen returns error if the provided input is more than {max} characters,
// nil otherwise.
func (me *Validation) Maxlen(field, input string, max int) *Error {
	return me.err(field, typMaxlen, Maxlen(input, max))
}

// Rangelen returns error if the provided input is between {min} and {max} characters
// long, nil otherwise.
func (me *Validation) Rangelen(field, input string, min, max int) *Error {
	return me.err(field, typRangelen, Rangelen(input, min, max))
}

// Float returns error if the provided input is not a floating point number, nil
// otherwise.
func (me *Validation) Float(field, input string) *Error {
	return me.err(field, typFloat, Float(input))
}

// Integer returns error if the provided input is not an integer value, nil otherwise.
func (me *Validation) Integer(field, input string) *Error {
	return me.err(field, typInteger, Integer(input))
}

// Alpha returns error if the provided input is not an alphabetic (a-zA-Z) string,
// nil otherwise.
func (me *Validation) Alpha(field, input string) *Error {
	return me.err(field, typAlpha, Alpha(input))
}

// Alrod returns error if the provided input is not an alphabetic or rod (a-zA-Z\-_) string,
// nil otherwise.
func (me *Validation) Alrod(field, input string) *Error {
	return me.err(field, typAlrod, Alrod(input))
}

// Alnum returns error if the provided input is not an alphanumeric (a-zA-Z0-9)
// string, nil otherwise.
func (me *Validation) Alnum(field, input string) *Error {
	return me.err(field, typAlnum, Alnum(input))
}

// Alnumrod returns error if the provided input is not an alphanumeric or rod (a-zA-Z0-9\-_)
// string, nil otherwise.
func (me *Validation) Alnumrod(field, input string) *Error {
	return me.err(field, typAlnumrod, Alnumrod(input))
}

// Alnumhan returns error if the provided input is not an alphanumeric or chinese (a-zA-Z0-9\p{Han})
// string, nil otherwise.
func (me *Validation) Alnumhan(field, input string) *Error {
	return me.err(field, typAlnumhan, Alnumhan(input))
}

// Alnumhanrod returns error if the provided input is not an alphanumeric or chinese or rod (a-zA-Z0-9\p{Han}\-_)
// string, nil otherwise.
func (me *Validation) Alnumhanrod(field, input string) *Error {
	return me.err(field, typAlnumhanrod, Alnumhanrod(input))
}

// Alhan returns error if the provided input is not an alphabetic or chinese (a-zA-Z\p{Han})
// string, nil otherwise.
func (me *Validation) Alhan(field, input string) *Error {
	return me.err(field, typAlhan, Alhan(input))
}

// Alhanrod returns error if the provided input is not an alphabetic or chinese or rod (a-zA-Z\p{Han}\-_)
// string, nil otherwise.
func (me *Validation) Alhanrod(field, input string) *Error {
	return me.err(field, typAlhanrod, Alhanrod(input))
}

// Han returns error if the provided input is not an chinese (\p{Han})
// string, nil otherwise.
func (me *Validation) Han(field, input string) *Error {
	return me.err(field, typHan, Han(input))
}

// Hanrod returns error if the provided input is not an chinese or rod (\p{Han}\-_)
// string, nil otherwise.
func (me *Validation) Hanrod(field, input string) *Error {
	return me.err(field, typHanrod, Hanrod(input))
}

// Match returns error if the provided input is not match {regexp} string,
// nil otherwise.
func (me *Validation) Match(field, input, regexps string) *Error {
	return me.err(field, typMatch, Match(input, regexps))
}

// Nomatch returns error if the provided input is match {regexp} string,
// nil otherwise.
func (me *Validation) Nomatch(field, input, regexps string) *Error {
	return me.err(field, typNomatch, Nomatch(input, regexps))
}

// Email returns error if the provided input is not an email, nil otherwise.
func (me *Validation) Email(field, input string) *Error {
	return me.err(field, typEmail, Email(input))
}

// URL returns error if the provided input is not an URL, nil otherwise.
func (me *Validation) URL(field, input string) *Error {
	return me.err(field, typURL, URL(input))
}

// IP returns error if the provided input is not an IP, nil otherwise.
func (me *Validation) IP(field, input string) *Error {
	return me.err(field, typIP, IP(input))
}

// Mobile returns error if the provided input is not an mobile, nil otherwise.
func (me *Validation) Mobile(field, input string) *Error {
	return me.err(field, typMobile, Mobile(input))
}

// Tel returns error if the provided input is not an tel, nil otherwise.
func (me *Validation) Tel(field, input string) *Error {
	return me.err(field, typTel, Tel(input))
}

// Phone returns error if the provided input is not an phone, nil otherwise.
func (me *Validation) Phone(field, input string) *Error {
	return me.err(field, typPhone, Phone(input))
}

// Zipcode returns error if the provided input is not an zipcode, nil otherwise.
func (me *Validation) Zipcode(field, input string) *Error {
	return me.err(field, typZipcode, Zipcode(input))
}
